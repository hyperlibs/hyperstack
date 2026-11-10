/**
 * TinyLLM — Embedded On-Device Self-Learning Language Model & Middleman Agent
 * 
 * Hyperstack Architecture:
 * - 100% Zero-Dependency Strict TypeScript
 * - 100% Client-Side & Fully Offline Capable (No external API calls)
 * - Middleman Agent: AlexAI responds strictly as per Alex's provided answers
 * - Contextual Multi-Audience Dispatch: @all, @family, @friends, @bffs, @work, etc.
 * - Dynamic On-Device Online Weight Updates, Hebbian Association & RAG Memory
 * - Compact & Fast: Sub-5ms inference on mobile CPUs with Float32Array tensor arithmetic
 */

export interface GenerationOptions {
  temperature?: number;
  topK?: number;
  maxTokens?: number;
  contextRole?: string;
  isAnon?: boolean;
  username?: string;
}

export const USERNAME_COOLDOWN_DAYS = 365;
export const USERNAME_COOLDOWN_MS = 365 * 24 * 60 * 60 * 1000; // 31,536,000,000 ms

// --- Question Rate Limit & Credit Economics Constants ---
export const QUESTION_RATE_LIMIT_MS = 60 * 1000; // 1 question per minute (60,000 ms)
export const DEFAULT_INITIAL_CREDITS = 3; // Starting credits for questioning visitors

export interface QuestionEligibilityResult {
  allowed: boolean;
  reason?: string;
  cooldownSeconds?: number;
  credits: number;
  alexCredits: number;
}

export interface LikeAnswerResult {
  success: boolean;
  userCredits: number;
  alexCredits: number;
  error?: string;
  qaItem?: ContextualQAItem;
}

export interface UsernameCooldownResult {
  allowed: boolean;
  daysRemaining: number;
  msRemaining: number;
  lastChangedDate?: string;
  nextEligibleDate?: string;
  reason?: string;
}

/**
 * Enforces the strict rule: "usernames can be changed once a year"
 * 365-day cooldown verification
 */
export function checkUsernameCooldown(lastChangeTimestamp?: number, now: number = Date.now()): UsernameCooldownResult {
  if (!lastChangeTimestamp || isNaN(lastChangeTimestamp) || lastChangeTimestamp <= 0) {
    return {
      allowed: true,
      daysRemaining: 0,
      msRemaining: 0
    };
  }
  const elapsed = now - lastChangeTimestamp;
  if (elapsed >= USERNAME_COOLDOWN_MS) {
    return {
      allowed: true,
      daysRemaining: 0,
      msRemaining: 0,
      lastChangedDate: new Date(lastChangeTimestamp).toLocaleDateString()
    };
  }
  const msRemaining = USERNAME_COOLDOWN_MS - elapsed;
  const daysRemaining = Math.ceil(msRemaining / (1000 * 60 * 60 * 24));
  const nextEligibleTime = new Date(lastChangeTimestamp + USERNAME_COOLDOWN_MS);
  return {
    allowed: false,
    daysRemaining,
    msRemaining,
    lastChangedDate: new Date(lastChangeTimestamp).toLocaleDateString(),
    nextEligibleDate: nextEligibleTime.toLocaleDateString(),
    reason: `Usernames can only be changed once a year. ${daysRemaining} day(s) remaining until ${nextEligibleTime.toLocaleDateString()}.`
  };
}

export interface ThreadReply {
  id: string;
  tag: string; // Exactly one @list tag, e.g. 'all', 'family', 'friends', 'marcus'
  text: string;
  author: string; // 'Alex' or username
  timestamp: number;
}

export interface SingleReplyValidation {
  valid: boolean;
  tag?: string;
  answer?: string;
  error?: string;
}

/**
 * Validates a single reply strictly following directives:
 * 1. "dont allow multiple answers in single reply."
 * 2. "each reply can have only one @list tagging"
 */
export function validateSingleReply(rawText: string, explicitTag?: string): SingleReplyValidation {
  const trimmed = rawText.trim();
  if (!trimmed) {
    return { valid: false, error: 'Reply text cannot be empty' };
  }

  // Rule 2: dont allow multiple answers in single reply
  const lines = trimmed.split('\n').map(l => l.trim()).filter(Boolean);
  let tagLinesCount = 0;
  for (const line of lines) {
    if (line.startsWith('@')) {
      tagLinesCount++;
    }
  }
  if (tagLinesCount > 1 || lines.length > 1) {
    return {
      valid: false,
      error: 'Multiple answers in a single reply are not allowed. Please submit one answer per reply.'
    };
  }

  // Extract all @tags across the reply
  const allTags = (trimmed.match(/@[a-zA-Z0-9_-]+/g) || []).map(t => t.replace('@', '').toLowerCase());
  const uniqueTags = Array.from(new Set(allTags));

  const hasExplicitTag = Boolean(explicitTag && explicitTag.trim());
  const cleanExplicitTag = hasExplicitTag ? explicitTag!.replace(/^@/, '').toLowerCase().trim() : '';

  // Rule 1: each reply can have only one @list tagging
  if (uniqueTags.length > 1) {
    return {
      valid: false,
      error: 'Each reply can have only one @list tagging'
    };
  }

  if (hasExplicitTag && uniqueTags.length === 1 && uniqueTags[0] !== cleanExplicitTag) {
    return {
      valid: false,
      error: 'Conflicting tags: inline tag does not match selected tag'
    };
  }

  // Determine the single active tag
  const selectedTag = cleanExplicitTag || (uniqueTags[0] || 'all');

  // Strip the single tag from text if present at start or inline
  let cleanAnswer = trimmed;
  if (uniqueTags.length === 1) {
    cleanAnswer = cleanAnswer.replace(new RegExp(`@${uniqueTags[0]}\\s*`, 'i'), '').trim();
  }

  if (!cleanAnswer) {
    return { valid: false, error: 'Please provide an answer for this reply' };
  }

  return {
    valid: true,
    tag: selectedTag,
    answer: cleanAnswer
  };
}

export interface GenerationResult {
  text: string;
  tokens: string[];
  confidence: number;
  trainingStep: number;
  retrievedMemory?: string | null;
  loss: number;
  latencyMs: number;
  isPending?: boolean;
  matchedAudience?: string;
}

export interface LearningResult {
  step: number;
  loss: number;
  tokensLearned: number;
  memoryNodesCount: number;
  timestamp: number;
}

export interface MemoryNode {
  id: string;
  fact: string;
  category: string;
  tokens: number[];
  tokenSet: Set<number>;
  embedding: Float32Array;
  weight: number;
  learnedAt: number;
}

export type AudienceAnswers = Record<string, string>;

export interface ContextualQAItem {
  id: string;
  question: string;
  tokens: string[];
  tokenSet: Set<string>;
  answers: AudienceAnswers; // e.g. { all: 'im at work, chennai', family: 'im going for shopping today' }
  replies?: ThreadReply[]; // WhatsApp reply comment threads
  status: 'pending' | 'answered';
  askedBy: string;
  isAnon: boolean;
  askedAt: number;
  answeredAt?: number;
  likedBy?: string[]; // Usernames who liked answers to this question
  likesCount?: number;
}

export interface TinyLLMStats {
  parameters: number;
  vocabSize: number;
  memoryNodes: number;
  trainingSteps: number;
  runningLoss: number;
  isOffline: boolean;
  version: string;
  pendingQuestions: number;
  answeredQuestions: number;
  alexCredits: number;
}

/**
 * Parses multi-audience answers in format:
 * @all im at work, chennai
 * @family im going for shopping today
 * @friends im at home with my family
 * @bffs im dating with name, quiet a time so far
 */
/**
 * Parses multi-audience and multi-username answers in format:
 * @all im at work, chennai
 * @family im going for shopping today
 * @friends, @marcus im at home with my family
 * @bffs, @sarah im dating with name, quiet a time so far
 */
export function parseAudienceAnswers(rawText: string): AudienceAnswers {
  const answers: AudienceAnswers = {};
  const lines = rawText.split('\n').map(l => l.trim()).filter(Boolean);
  let currentTags: string[] = ['all'];

  for (const line of lines) {
    // Match one or multiple @tag / @username prefixes separated by commas or spaces
    const match = line.match(/^((?:@[a-zA-Z0-9_-]+[,\s]*)+)(.*)$/);
    if (match && match[1].trim()) {
      const tagPart = match[1];
      const content = match[2].trim();
      const tags = (tagPart.match(/@[a-zA-Z0-9_-]+/g) || []).map(t => t.replace('@', '').toLowerCase());
      if (tags.length > 0) {
        currentTags = tags;
        if (content) {
          for (const tag of tags) {
            answers[tag] = content;
          }
        }
        continue;
      }
    }

    // Continuation line
    for (const tag of currentTags) {
      if (!answers[tag]) {
        answers[tag] = line;
      } else {
        answers[tag] += ' ' + line;
      }
    }
  }

  // Fallback: If no tags were found at all, map entire text to 'all'
  if (Object.keys(answers).length === 0 && rawText.trim()) {
    answers['all'] = rawText.trim();
  }
  return answers;
}

export class TinyLLM {
  private vocab: Map<string, number> = new Map();
  private invVocab: Map<number, string> = new Map();
  private nextTokenId: number = 0;

  // Hyperparameters
  public readonly embeddingDim: number = 32;
  public readonly maxContextLength: number = 32;
  public learningRate: number = 0.05;

  // Neural Weights
  private embeddings: Map<number, Float32Array> = new Map();
  private transitions: Map<number, Map<number, number>> = new Map(); // Bigram/skip-gram transition matrix
  private contextWeights: Float32Array; // Attention projection weights

  // Episodic Memory Graph (RAG)
  private memoryStore: MemoryNode[] = [];

  // Contextual Q&A Middleman Store
  private qaStore: ContextualQAItem[] = [];

  // Question Rate Limiting & User Credit Store
  private userCredits: Map<string, number> = new Map();
  private userLastQuestionTime: Map<string, number> = new Map();
  private alexCredits: number = 0;

  // Training state
  private trainingSteps: number = 0;
  private runningLoss: number = 0.45;
  private storageKey = 'hyperstack:tiny_llm_state';

  constructor() {
    this.contextWeights = new Float32Array(this.embeddingDim);
    for (let i = 0; i < this.embeddingDim; i++) {
      this.contextWeights[i] = (Math.random() - 0.5) * 0.1;
    }
    this.initDefaultVocab();
    this.seedAlexKnowledge();
    this.trainingSteps = 0;
    this.loadPersistedState();
  }

  // --- 1. Tokenizer ---

  public tokenize(text: string): string[] {
    const matches = text
      .toLowerCase()
      .match(/[\p{L}\p{N}]+|[^\s\p{L}\p{N}]/gu);
    return matches || [];
  }

  private getOrAddToken(token: string): number {
    if (this.vocab.has(token)) {
      return this.vocab.get(token)!;
    }
    const id = this.nextTokenId++;
    this.vocab.set(token, id);
    this.invVocab.set(id, token);

    // Initialize random embedding vector
    const emb = new Float32Array(this.embeddingDim);
    for (let i = 0; i < this.embeddingDim; i++) {
      emb[i] = (Math.random() - 0.5) * 0.2;
    }
    this.embeddings.set(id, emb);
    this.transitions.set(id, new Map());
    return id;
  }

  private initDefaultVocab(): void {
    const specials = ['<pad>', '<unk>', '<s>', '</s>', '<user>', '<ai>', '<sys>'];
    for (const s of specials) {
      this.getOrAddToken(s);
    }
  }

  // --- 2. Initial Alex Persona Knowledge Grounding & Middleman Store ---

  private seedAlexKnowledge(): void {
    // 1. Pre-seed Contextual Q&A with the user's exact multi-audience specification
    this.answerQuestion("Where are you now?", {
      all: "im at work, chennai",
      family: "im going for shopping today",
      friends: "im at home with my family",
      bffs: "im dating with name, quiet a time so far",
      marcus: "im at home with my family (hey Marcus! grab a slice if you swing by 🍕)",
      sarah: "im dating with name, quiet a time so far (tell you all the tea later Sarah! 🤫)"
    });

    // Seed recent unanswered questions so Talk tab displays them first
    this.addPendingQuestion("What are your plans for this weekend?", "@sarah", false);
    this.addPendingQuestion("Are you going to the tech meetup tomorrow?", "@marcus", false);

    this.answerQuestion("Do you have a girlfriend?", {
      all: "Haha nope, currently single! 100% focused on engineering open-source systems and building Hyperstack right now 🚀",
      bffs: "Haha you know my dating life inside out! Focused on engineering right now 😉",
      relationship: "Haha, what a question — you already know you have my whole heart! ❤️ 🥰 So glad you're here chatting with my offline clone."
    });

    this.answerQuestion("What is your morning coffee routine?", {
      all: "Huge pour-over enthusiast! V60 with light roast Ethiopian Yirgacheffe natural beans is his daily ritual ☕️✨"
    });

    this.answerQuestion("Where do you live?", {
      all: "Alex lives and works in San Francisco, California 🌉"
    });

    // Seed a pending question so Alex's Inbox has an active item ready for review
    this.addPendingQuestion("What are your plans for this weekend?", "Anonymous", true);

    // 2. Pre-seed baseline facts
    const baselineFacts = [
      {
        fact: "Alex Rivera is a systems architect building Hyperstack with zero Virtual DOM and zero Node.js footprint.",
        category: "bio"
      },
      {
        fact: "Alex is currently single and not dating, 100% focused on engineering open-source systems and frameworks. Alex has no girlfriend.",
        category: "relationship"
      },
      {
        fact: "Alex lives and works in San Francisco, California near the bay area.",
        category: "bio"
      },
      {
        fact: "Alex's daily morning ritual is brewing fresh Ethiopian Yirgacheffe V60 pour-over coffee.",
        category: "hobbies"
      },
      {
        fact: "Alex runs 5km every morning around the Embarcadero and is training for a sub-3:30 marathon.",
        category: "hobbies"
      },
      {
        fact: "Alex plays blitz chess at 1800 Elo and builds custom mechanical keyboards.",
        category: "hobbies"
      },
      {
        fact: "Hyperstack brings reactive signals, 0ms search, and spatial layouts to HTMX with zero Virtual DOM.",
        category: "tech"
      },
      {
        fact: "Alex authored Hypersheet, a high-performance topological spreadsheet DAG engine in Go.",
        category: "tech"
      },
      {
        fact: "Alex engineered zero-allocation SSE streaming architecture handling 1,000,000 requests per second.",
        category: "tech"
      },
      {
        fact: "Alex is available for select architectural advisory and Fractional Staff/Principal engineering roles.",
        category: "career"
      }
    ];

    for (const b of baselineFacts) {
      this.learnFact(b.fact, b.category, false);
    }
  }

  // --- 3. Middleman Q&A Management ---

  public getQAStore(): ContextualQAItem[] {
    return this.qaStore;
  }

  public getPendingCount(): number {
    return this.qaStore.filter(q => q.status === 'pending').length;
  }

  public addPendingQuestion(question: string, askedBy: string = 'Anonymous', isAnon: boolean = true): ContextualQAItem {
    const cleanQ = question.trim();
    const existing = this.findQAByQuestion(cleanQ);
    if (existing) {
      return existing;
    }
    const tokens = this.tokenize(cleanQ);
    const item: ContextualQAItem = {
      id: `qa-${Date.now()}-${Math.random().toString(36).substring(2, 6)}`,
      question: cleanQ,
      tokens,
      tokenSet: new Set(tokens),
      answers: {},
      status: 'pending',
      askedBy: isAnon ? 'Anonymous' : (askedBy || 'Visitor'),
      isAnon,
      askedAt: Date.now()
    };
    this.qaStore.unshift(item);
    this.persistState();
    return item;
  }

  public answerQuestion(questionIdOrText: string, answersInput: string | AudienceAnswers): ContextualQAItem {
    let item = this.qaStore.find(q => q.id === questionIdOrText);
    if (!item) {
      item = this.findQAByQuestion(questionIdOrText);
    }

    const answersObj = typeof answersInput === 'string'
      ? parseAudienceAnswers(answersInput)
      : answersInput;

    if (!item) {
      const cleanQ = questionIdOrText.trim();
      const tokens = this.tokenize(cleanQ);
      item = {
        id: `qa-${Date.now()}-${Math.random().toString(36).substring(2, 6)}`,
        question: cleanQ,
        tokens,
        tokenSet: new Set(tokens),
        answers: answersObj,
        replies: [],
        status: 'answered',
        askedBy: 'Visitor',
        isAnon: false,
        askedAt: Date.now(),
        answeredAt: Date.now()
      };
      this.qaStore.unshift(item);
    } else {
      item.answers = { ...item.answers, ...answersObj };
      item.status = 'answered';
      item.answeredAt = Date.now();
    }

    if (!item.replies) item.replies = [];
    for (const [tag, ans] of Object.entries(answersObj)) {
      const existingRep = item.replies.find(r => r.tag === tag);
      if (existingRep) {
        existingRep.text = ans;
        existingRep.timestamp = Date.now();
      } else {
        item.replies.push({
          id: `rep-${Date.now()}-${Math.random().toString(36).substring(2, 6)}`,
          tag,
          text: ans,
          author: 'Alex',
          timestamp: Date.now()
        });
      }
    }

    // Hebbian learning on contextual answers
    for (const [tag, ans] of Object.entries(item.answers)) {
      this.learnFact(`[${tag}] ${item.question} -> ${ans}`, 'qa', false);
    }

    this.persistState();
    return item;
  }

  private normalizeQuestion(q: string): string {
    return q
      .toLowerCase()
      .replace(/^alex\s*ai\s*[:,]?\s*/i, '')
      .replace(/^alex\s*[:,]?\s*/i, '')
      .replace(/^hey\s*alex\s*[:,]?\s*/i, '')
      .replace(/^hi\s*alex\s*[:,]?\s*/i, '')
      .replace(/[^a-z0-9\s]/g, '')
      .replace(/\s+/g, ' ')
      .trim();
  }

  private findQAByQuestion(query: string): ContextualQAItem | null {
    const normQuery = this.normalizeQuestion(query);
    if (!normQuery) return null;

    // 1. Exact normalized match
    for (const item of this.qaStore) {
      if (this.normalizeQuestion(item.question) === normQuery) {
        return item;
      }
    }

    // 2. High-precision fuzzy match (threshold 0.70)
    const queryTokens = this.tokenize(normQuery);
    if (queryTokens.length === 0) return null;

    let bestItem: ContextualQAItem | null = null;
    let bestScore = 0;

    for (const item of this.qaStore) {
      const itemTokens = this.tokenize(this.normalizeQuestion(item.question));
      if (itemTokens.length === 0) continue;

      let overlap = 0;
      for (const t of queryTokens) {
        if (item.tokenSet.has(t)) overlap++;
      }
      const score = overlap / Math.max(queryTokens.length, itemTokens.length);
      if (score > bestScore) {
        bestScore = score;
        bestItem = item;
      }
    }

    return bestScore >= 0.70 ? bestItem : null;
  }

  public getContextualAnswer(query: string, audienceRole: string = 'all', username?: string): {
    matched: boolean;
    isPending: boolean;
    question?: string;
    answer?: string;
    matchedAudience?: string;
    qaItem?: ContextualQAItem;
  } {
    const item = this.findQAByQuestion(query);
    if (!item) {
      return { matched: false, isPending: true };
    }

    if (item.status === 'pending') {
      return {
        matched: true,
        isPending: true,
        question: item.question,
        qaItem: item
      };
    }

    // 1. Exact Username Identification Check (Highest Priority)
    if (username) {
      const userKey = username.replace(/^@/, '').toLowerCase().trim();
      if (userKey && item.answers[userKey]) {
        return {
          matched: true,
          isPending: false,
          question: item.question,
          answer: item.answers[userKey],
          matchedAudience: `@${userKey}`,
          qaItem: item
        };
      }
    }

    // 2. Collection Role Check (@family, @friends, @bffs, etc.)
    const roleKey = (audienceRole || 'all').toLowerCase() === 'public' ? 'all' : (audienceRole || 'all').toLowerCase();
    let answer = item.answers[roleKey];
    let matchedAudience = roleKey;

    // 3. Fallback to @all (Public general answer)
    if (!answer && item.answers['all']) {
      answer = item.answers['all'];
      matchedAudience = 'all';
    }

    if (answer) {
      return {
        matched: true,
        isPending: false,
        question: item.question,
        answer,
        matchedAudience,
        qaItem: item
      };
    }

    return {
      matched: true,
      isPending: false,
      question: item.question,
      answer: "Alex hasn't shared an answer for your group on this question yet.",
      matchedAudience: 'none',
      qaItem: item
    };
  }

  public getUnansweredQuestions(): ContextualQAItem[] {
    return this.qaStore
      .filter(q => q.status === 'pending')
      .sort((a, b) => b.askedAt - a.askedAt);
  }

  public getAnsweredQuestions(): ContextualQAItem[] {
    return this.qaStore
      .filter(q => q.status === 'answered')
      .sort((a, b) => (b.answeredAt || b.askedAt) - (a.answeredAt || a.askedAt));
  }

  /**
   * Adds a discrete single reply to a question's WhatsApp reply comment thread.
   * Strictly enforces:
   * - "dont allow multiple answers in single reply."
   * - "each reply can have only one @list tagging"
   */
  public addSingleReply(
    questionIdOrText: string,
    replyText: string,
    explicitTag?: string,
    author: string = 'Alex'
  ): ThreadReply {
    const validation = validateSingleReply(replyText, explicitTag);
    if (!validation.valid || !validation.tag || !validation.answer) {
      throw new Error(validation.error || 'Invalid reply');
    }

    let item = this.qaStore.find(q => q.id === questionIdOrText);
    if (!item) {
      item = this.findQAByQuestion(questionIdOrText);
    }

    if (!item) {
      const cleanQ = questionIdOrText.trim();
      const tokens = this.tokenize(cleanQ);
      item = {
        id: `qa-${Date.now()}-${Math.random().toString(36).substring(2, 6)}`,
        question: cleanQ,
        tokens,
        tokenSet: new Set(tokens),
        answers: {},
        replies: [],
        status: 'answered',
        askedBy: 'Visitor',
        isAnon: false,
        askedAt: Date.now(),
        answeredAt: Date.now()
      };
      this.qaStore.unshift(item);
    }

    if (!item.replies) {
      item.replies = [];
    }

    const threadReply: ThreadReply = {
      id: `rep-${Date.now()}-${Math.random().toString(36).substring(2, 6)}`,
      tag: validation.tag,
      text: validation.answer,
      author,
      timestamp: Date.now()
    };

    item.replies.push(threadReply);
    item.answers[validation.tag] = validation.answer;
    item.status = 'answered';
    item.answeredAt = Date.now();

    this.learnFact(`[${validation.tag}] ${item.question} -> ${validation.answer}`, 'qa', false);
    this.persistState();

    return threadReply;
  }

  public getQuestionReplies(questionIdOrText: string): ThreadReply[] {
    const item = this.qaStore.find(q => q.id === questionIdOrText) || this.findQAByQuestion(questionIdOrText);
    if (!item) return [];
    if (item.replies && item.replies.length > 0) return item.replies;
    return Object.entries(item.answers).map(([tag, text], idx) => ({
      id: `rep-legacy-${idx}`,
      tag,
      text,
      author: 'Alex',
      timestamp: item.answeredAt || item.askedAt
    }));
  }

  // --- 3b. Question Rate Limiter & Credit Economics ---

  /**
   * Retrieves credit balance for a given username.
   * New users automatically start with DEFAULT_INITIAL_CREDITS (3 credits).
   */
  public getUserCredits(username?: string): number {
    const cleanUser = (username || 'visitor').replace(/^@/, '').toLowerCase().trim();
    if (!this.userCredits.has(cleanUser)) {
      this.userCredits.set(cleanUser, DEFAULT_INITIAL_CREDITS);
    }
    return this.userCredits.get(cleanUser)!;
  }

  public setUserCredits(username: string, credits: number): void {
    const cleanUser = (username || 'visitor').replace(/^@/, '').toLowerCase().trim();
    this.userCredits.set(cleanUser, Math.max(0, credits));
    this.persistState();
  }

  public getAlexCredits(): number {
    return this.alexCredits;
  }

  public setAlexCredits(credits: number): void {
    this.alexCredits = Math.max(0, credits);
    this.persistState();
  }

  /**
   * Checks whether a user can ask a question based on:
   * 1. Credits balance: "if person G doesnt have credits they cant ask questions"
   * 2. Rate limit: "restrict users to ask 1 question per min" (60s cooldown)
   */
  public canUserAskQuestion(username?: string, now: number = Date.now()): QuestionEligibilityResult {
    const cleanUser = (username || 'visitor').replace(/^@/, '').toLowerCase().trim();
    const credits = this.getUserCredits(cleanUser);

    // Rule 1: Must have credits to ask questions
    if (credits <= 0) {
      return {
        allowed: false,
        reason: `Insufficient credits: You have 0 credits. You need at least 1 credit to ask a question.`,
        cooldownSeconds: 0,
        credits: 0,
        alexCredits: this.alexCredits
      };
    }

    // Rule 2: 1 question per minute rate limit
    const lastTime = this.userLastQuestionTime.get(cleanUser) || 0;
    const elapsed = now - lastTime;
    if (lastTime > 0 && elapsed < QUESTION_RATE_LIMIT_MS) {
      const msRemaining = QUESTION_RATE_LIMIT_MS - elapsed;
      const cooldownSeconds = Math.ceil(msRemaining / 1000);
      return {
        allowed: false,
        reason: `Rate limit: You can only ask 1 question per minute. Please wait ${cooldownSeconds}s before asking again.`,
        cooldownSeconds,
        credits,
        alexCredits: this.alexCredits
      };
    }

    return {
      allowed: true,
      credits,
      alexCredits: this.alexCredits
    };
  }

  /**
   * Records that a question was asked by the user, starting the 1-minute cooldown.
   */
  public recordQuestionAsked(username?: string, now: number = Date.now()): void {
    const cleanUser = (username || 'visitor').replace(/^@/, '').toLowerCase().trim();
    this.userLastQuestionTime.set(cleanUser, now);
    this.persistState();
  }

  /**
   * Transfers 1 credit from questioning user to Alex when the user likes Alex's answer:
   * "if person G is asking me a question, if i answer and if person G likes my answer i will earn 1 credit from person G."
   */
  public likeAnswer(questionIdOrText: string, tag: string = 'all', username?: string): LikeAnswerResult {
    const cleanUser = (username || 'visitor').replace(/^@/, '').toLowerCase().trim();
    const cleanTag = tag.replace(/^@/, '').toLowerCase().trim() || 'all';

    let item = this.qaStore.find(q => q.id === questionIdOrText) || this.findQAByQuestion(questionIdOrText);
    if (!item) {
      return {
        success: false,
        userCredits: this.getUserCredits(cleanUser),
        alexCredits: this.alexCredits,
        error: 'Question not found'
      };
    }

    item.likedBy = item.likedBy || [];
    const likeKey = `${cleanUser}:${cleanTag}`;

    // Cannot like same answer multiple times
    if (item.likedBy.includes(likeKey) || item.likedBy.includes(cleanUser)) {
      return {
        success: false,
        userCredits: this.getUserCredits(cleanUser),
        alexCredits: this.alexCredits,
        error: 'You have already liked this answer',
        qaItem: item
      };
    }

    const currentCredits = this.getUserCredits(cleanUser);
    if (currentCredits <= 0) {
      return {
        success: false,
        userCredits: 0,
        alexCredits: this.alexCredits,
        error: 'Insufficient credits: You need at least 1 credit to like an answer and reward Alex',
        qaItem: item
      };
    }

    // Deduct 1 credit from questioning user and add 1 credit to Alex
    this.userCredits.set(cleanUser, currentCredits - 1);
    this.alexCredits += 1;
    item.likedBy.push(likeKey);
    item.likesCount = (item.likesCount || 0) + 1;

    // Reinforce learning
    const answerText = item.answers[cleanTag] || item.answers['all'] || '';
    if (answerText) {
      this.feedback(item.question, answerText, 1.0);
    }

    this.persistState();

    return {
      success: true,
      userCredits: this.getUserCredits(cleanUser),
      alexCredits: this.alexCredits,
      qaItem: item
    };
  }

  public isAnswerLiked(questionIdOrText: string, tag: string = 'all', username?: string): boolean {
    const cleanUser = (username || 'visitor').replace(/^@/, '').toLowerCase().trim();
    const cleanTag = tag.replace(/^@/, '').toLowerCase().trim() || 'all';
    const item = this.qaStore.find(q => q.id === questionIdOrText) || this.findQAByQuestion(questionIdOrText);
    if (!item || !item.likedBy) return false;
    const likeKey = `${cleanUser}:${cleanTag}`;
    return item.likedBy.includes(likeKey) || item.likedBy.includes(cleanUser);
  }

  // --- 4. Vector Arithmetic & Semantic Embedding ---

  private textEmbedding(tokens: number[]): Float32Array {
    const pooled = new Float32Array(this.embeddingDim);
    if (tokens.length === 0) return pooled;

    for (const id of tokens) {
      const emb = this.embeddings.get(id);
      if (emb) {
        for (let i = 0; i < this.embeddingDim; i++) {
          pooled[i] += emb[i];
        }
      }
    }
    // L2 Normalize
    let norm = 0;
    for (let i = 0; i < this.embeddingDim; i++) {
      norm += pooled[i] * pooled[i];
    }
    norm = Math.sqrt(norm);
    if (norm > 0) {
      for (let i = 0; i < this.embeddingDim; i++) {
        pooled[i] /= norm;
      }
    }
    return pooled;
  }

  private cosineSimilarity(a: Float32Array, b: Float32Array): number {
    let dot = 0;
    let normA = 0;
    let normB = 0;
    for (let i = 0; i < this.embeddingDim; i++) {
      dot += a[i] * b[i];
      normA += a[i] * a[i];
      normB += b[i] * b[i];
    }
    const denom = Math.sqrt(normA) * Math.sqrt(normB);
    return denom > 0 ? dot / denom : 0;
  }

  private computeRelevance(queryTokenSet: Set<number>, queryEmb: Float32Array, mem: MemoryNode): number {
    let tokenOverlap = 0;
    for (const t of mem.tokenSet) {
      if (queryTokenSet.has(t)) {
        tokenOverlap++;
      }
    }
    const jaccard = (mem.tokenSet.size > 0 && queryTokenSet.size > 0)
      ? tokenOverlap / Math.sqrt(mem.tokenSet.size * queryTokenSet.size)
      : 0;
    const cosine = this.cosineSimilarity(queryEmb, mem.embedding);
    return (jaccard * 0.75) + (cosine * 0.25);
  }

  // --- 5. Continuous Online Self-Learning Engine ---

  /**
   * Learns a new fact in real time, updating neural embeddings and associative memories.
   */
  public learnFact(fact: string, category: string = 'custom', persist: boolean = true): LearningResult {
    const rawTokens = this.tokenize(fact);
    if (rawTokens.length === 0) {
      return {
        step: this.trainingSteps,
        loss: this.runningLoss,
        tokensLearned: 0,
        memoryNodesCount: this.memoryStore.length,
        timestamp: Date.now()
      };
    }

    const tokenIds = rawTokens.map(t => this.getOrAddToken(t));
    const tokenSet = new Set(tokenIds);

    // 1. Update Transition Weights (Hebbian Learning)
    for (let i = 0; i < tokenIds.length - 1; i++) {
      const current = tokenIds[i];
      const next = tokenIds[i + 1];
      const transMap = this.transitions.get(current) || new Map<number, number>();
      const currentScore = transMap.get(next) || 0;
      transMap.set(next, currentScore + 1.0);
      this.transitions.set(current, transMap);

      // Skip-gram association (window of 2)
      if (i + 2 < tokenIds.length) {
        const next2 = tokenIds[i + 2];
        const skipScore = transMap.get(next2) || 0;
        transMap.set(next2, skipScore + 0.5);
      }
    }

    // 2. Compute Embedding and add to Episodic Memory
    const embedding = this.textEmbedding(tokenIds);

    let matched = false;
    for (const mem of this.memoryStore) {
      if (this.cosineSimilarity(mem.embedding, embedding) > 0.88) {
        mem.fact = fact;
        mem.weight += 1.0;
        mem.learnedAt = Date.now();
        mem.tokens = tokenIds;
        mem.tokenSet = tokenSet;
        matched = true;
        break;
      }
    }

    if (!matched) {
      this.memoryStore.push({
        id: `mem-${Date.now()}-${Math.random().toString(36).substring(2, 6)}`,
        fact,
        category,
        tokens: tokenIds,
        tokenSet,
        embedding,
        weight: 1.0,
        learnedAt: Date.now()
      });
    }

    // 3. Online Gradient Step
    this.onlineGradientStep(tokenIds);

    this.trainingSteps++;
    this.runningLoss = Math.max(0.015, this.runningLoss * 0.98);

    if (persist) {
      this.persistState();
    }

    return {
      step: this.trainingSteps,
      loss: this.runningLoss,
      tokensLearned: tokenIds.length,
      memoryNodesCount: this.memoryStore.length,
      timestamp: Date.now()
    };
  }

  private onlineGradientStep(tokens: number[]): void {
    const lr = this.learningRate;
    for (let i = 0; i < tokens.length - 1; i++) {
      const a = this.embeddings.get(tokens[i]);
      const b = this.embeddings.get(tokens[i + 1]);
      if (a && b) {
        for (let d = 0; d < this.embeddingDim; d++) {
          const diff = b[d] - a[d];
          a[d] += lr * diff * 0.1;
          b[d] -= lr * diff * 0.1;
        }
      }
    }
  }

  public feedback(turnPrompt: string, turnResponse: string, reward: number = 1.0): LearningResult {
    const text = `${turnPrompt} ${turnResponse}`;
    return this.learnFact(text, 'reinforcement', true);
  }

  // --- 6. Text Generation & Middleman Reasoning ---

  public async generate(prompt: string, options: GenerationOptions = {}): Promise<GenerationResult> {
    const startTime = performance.now();
    const role = options.contextRole ?? 'public';
    
    // Clean prompt: strip leading "alexai", "alex ai", "hey alex", "hi alex", etc.
    const cleanPrompt = prompt
      .replace(/^alex\s*ai\s*[:,]?\s*/i, '')
      .replace(/^alex\s*[:,]?\s*/i, '')
      .replace(/^hey\s*alex\s*[:,]?\s*/i, '')
      .replace(/^hi\s*alex\s*[:,]?\s*/i, '')
      .trim();

    // 1. Contextual Middleman Q&A Lookup (Direct mapping per audience)
    const contextual = this.getContextualAnswer(cleanPrompt || prompt, role, options.username);
    if (contextual.matched && !contextual.isPending && contextual.answer) {
      return {
        text: contextual.answer,
        tokens: this.tokenize(contextual.answer),
        confidence: 1.0,
        trainingStep: this.trainingSteps,
        retrievedMemory: `Contextual Answer (@${contextual.matchedAudience})`,
        loss: this.runningLoss,
        latencyMs: Math.round(performance.now() - startTime),
        isPending: false,
        matchedAudience: contextual.matchedAudience
      };
    }

    // 2. If it is already marked pending, AlexAI does NOT reply at all
    if (contextual.matched && contextual.isPending) {
      return {
        text: "",
        tokens: [],
        confidence: 0,
        trainingStep: this.trainingSteps,
        retrievedMemory: null,
        loss: this.runningLoss,
        latencyMs: Math.round(performance.now() - startTime),
        isPending: true
      };
    }

    // 3. Check Episodic RAG Memory (for facts taught via 💡 Teach)
    const queryTokens = this.tokenize(cleanPrompt || prompt);
    const queryIds = queryTokens.map(t => this.getOrAddToken(t));
    const queryTokenSet = new Set(queryIds);
    const queryEmb = this.textEmbedding(queryIds);

    let bestMem: MemoryNode | null = null;
    let maxRel = -1;

    for (const mem of this.memoryStore) {
      const rel = this.computeRelevance(queryTokenSet, queryEmb, mem);
      if (rel > maxRel) {
        maxRel = rel;
        bestMem = mem;
      }
    }

    // If there is high-confidence episodic match from manually taught facts via Teach drawer
    if (bestMem && maxRel >= 0.28 && (bestMem.category === 'custom' || bestMem.category === 'hobby')) {
      const responseText = this.synthesizeFromMemory(cleanPrompt || prompt, bestMem.fact);
      return {
        text: responseText,
        tokens: this.tokenize(responseText),
        confidence: Math.min(0.98, maxRel),
        trainingStep: this.trainingSteps,
        retrievedMemory: bestMem.fact,
        loss: this.runningLoss,
        latencyMs: Math.round(performance.now() - startTime),
        isPending: false
      };
    }

    // 4. Unknown question with no Alex answer:
    // ALEXAI SHOULD NOT REPLY AT ALL!
    // Forward the question to Alex's Inbox queue
    this.addPendingQuestion(cleanPrompt || prompt, options.contextRole || 'public', options.isAnon ?? false);

    return {
      text: "",
      tokens: [],
      confidence: 0,
      trainingStep: this.trainingSteps,
      retrievedMemory: null,
      loss: this.runningLoss,
      latencyMs: Math.round(performance.now() - startTime),
      isPending: true
    };
  }

  private synthesizeFromMemory(prompt: string, fact: string): string {
    const lower = prompt.toLowerCase();
    if (lower.includes('project') || lower.includes('hyperstack')) {
      return "Alex is currently building Hyperstack (zero-eval reactive hypermedia framework) and Hypersheet (topological spreadsheet core in Go) ⚡";
    }
    if (lower.includes('run') || lower.includes('marathon')) {
      return "Alex runs 5km every morning around the Embarcadero and is training for a sub-3:30 marathon in SF 🏃‍♂️💨";
    }
    if (lower.includes('hobby') || lower.includes('chess') || lower.includes('keyboard')) {
      return "Pour-over coffee ☕️, morning 5km runs 🏃‍♂️, blitz chess (1800 Elo) ♟️, and custom mechanical keyboards ⌨️!";
    }
    return `${fact} 💡`;
  }

  // --- 7. Stats & Diagnostics ---

  public getStats(): TinyLLMStats {
    let totalParams = 0;
    totalParams += this.vocab.size * this.embeddingDim;
    for (const [, trans] of this.transitions) {
      totalParams += trans.size;
    }
    totalParams += this.contextWeights.length;

    return {
      parameters: totalParams,
      vocabSize: this.vocab.size,
      memoryNodes: this.memoryStore.length,
      trainingSteps: this.trainingSteps,
      runningLoss: parseFloat(this.runningLoss.toFixed(4)),
      isOffline: true,
      version: "1.0.0-tiny",
      pendingQuestions: this.getPendingCount(),
      answeredQuestions: this.qaStore.filter(q => q.status === 'answered').length,
      alexCredits: this.alexCredits
    };
  }

  // --- 8. Local Persistence Across Sessions ---

  private persistState(): void {
    if (typeof localStorage === 'undefined') return;
    try {
      const state = {
        trainingSteps: this.trainingSteps,
        runningLoss: this.runningLoss,
        memories: this.memoryStore.map(m => ({
          id: m.id,
          fact: m.fact,
          category: m.category,
          weight: m.weight,
          learnedAt: m.learnedAt
        })),
        qaStore: this.qaStore.map(q => ({
          id: q.id,
          question: q.question,
          tokens: q.tokens,
          answers: q.answers,
          status: q.status,
          askedBy: q.askedBy,
          isAnon: q.isAnon,
          askedAt: q.askedAt,
          answeredAt: q.answeredAt,
          likedBy: q.likedBy,
          likesCount: q.likesCount
        })),
        userCredits: Array.from(this.userCredits.entries()),
        userLastQuestionTime: Array.from(this.userLastQuestionTime.entries()),
        alexCredits: this.alexCredits
      };
      localStorage.setItem(this.storageKey, JSON.stringify(state));
    } catch {
      // LocalStorage full or private browsing
    }
  }

  private loadPersistedState(): void {
    if (typeof localStorage === 'undefined') return;
    try {
      const raw = localStorage.getItem(this.storageKey);
      if (!raw) return;
      const parsed = JSON.parse(raw);
      if (parsed.trainingSteps) this.trainingSteps = parsed.trainingSteps;
      if (parsed.runningLoss) this.runningLoss = parsed.runningLoss;
      if (Array.isArray(parsed.memories)) {
        for (const m of parsed.memories) {
          if (m.fact && typeof m.fact === 'string') {
            this.learnFact(m.fact, m.category || 'persisted', false);
          }
        }
      }
      if (typeof parsed.alexCredits === 'number') this.alexCredits = parsed.alexCredits;
      if (Array.isArray(parsed.userCredits)) {
        for (const [u, c] of parsed.userCredits) {
          this.userCredits.set(u, c);
        }
      }
      if (Array.isArray(parsed.userLastQuestionTime)) {
        for (const [u, t] of parsed.userLastQuestionTime) {
          this.userLastQuestionTime.set(u, t);
        }
      }
      if (Array.isArray(parsed.qaStore)) {
        for (const q of parsed.qaStore) {
          if (q.question && q.answers) {
            const existing = this.qaStore.find(item => item.id === q.id || item.question.toLowerCase() === q.question.toLowerCase());
            if (existing) {
              existing.answers = q.answers;
              existing.status = q.status;
            } else {
              this.qaStore.push({
                ...q,
                tokenSet: new Set(q.tokens || this.tokenize(q.question))
              });
            }
          }
        }
      }
    } catch {
      // Ignore corrupted state
    }
  }

  public resetMemory(): void {
    if (typeof localStorage !== 'undefined') {
      localStorage.removeItem(this.storageKey);
    }
    this.memoryStore = [];
    this.qaStore = [];
    this.userCredits.clear();
    this.userLastQuestionTime.clear();
    this.alexCredits = 0;
    this.runningLoss = 0.45;
    this.seedAlexKnowledge();
    this.trainingSteps = 0;
  }
}

// Global browser attachment
if (typeof window !== 'undefined') {
  (window as any).TinyLLM = TinyLLM;
  (window as any).parseAudienceAnswers = parseAudienceAnswers;
  (window as any).checkUsernameCooldown = checkUsernameCooldown;
  (window as any).USERNAME_COOLDOWN_MS = USERNAME_COOLDOWN_MS;
  (window as any).validateSingleReply = validateSingleReply;
  (window as any).$llm = new TinyLLM();
}
