(() => {
  var __defProp = Object.defineProperty;
  var __getOwnPropNames = Object.getOwnPropertyNames;
  var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
  var __hasOwnProp = Object.prototype.hasOwnProperty;
  var __moduleCache = /* @__PURE__ */ new WeakMap;
  var __toCommonJS = (from) => {
    var entry = __moduleCache.get(from), desc;
    if (entry)
      return entry;
    entry = __defProp({}, "__esModule", { value: true });
    if (from && typeof from === "object" || typeof from === "function")
      __getOwnPropNames(from).map((key) => !__hasOwnProp.call(entry, key) && __defProp(entry, key, {
        get: () => from[key],
        enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable
      }));
    __moduleCache.set(from, entry);
    return entry;
  };
  var __export = (target, all) => {
    for (var name in all)
      __defProp(target, name, {
        get: all[name],
        enumerable: true,
        configurable: true,
        set: (newValue) => all[name] = () => newValue
      });
  };

  // src/tiny-llm.ts
  var exports_tiny_llm = {};
  __export(exports_tiny_llm, {
    validateSingleReply: () => validateSingleReply,
    parseAudienceAnswers: () => parseAudienceAnswers,
    checkUsernameCooldown: () => checkUsernameCooldown,
    USERNAME_COOLDOWN_MS: () => USERNAME_COOLDOWN_MS,
    USERNAME_COOLDOWN_DAYS: () => USERNAME_COOLDOWN_DAYS,
    TinyLLM: () => TinyLLM,
    QUESTION_RATE_LIMIT_MS: () => QUESTION_RATE_LIMIT_MS,
    DEFAULT_INITIAL_CREDITS: () => DEFAULT_INITIAL_CREDITS
  });
  var USERNAME_COOLDOWN_DAYS = 365;
  var USERNAME_COOLDOWN_MS = 365 * 24 * 60 * 60 * 1000;
  var QUESTION_RATE_LIMIT_MS = 60 * 1000;
  var DEFAULT_INITIAL_CREDITS = 3;
  function checkUsernameCooldown(lastChangeTimestamp, now = Date.now()) {
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
  function validateSingleReply(rawText, explicitTag) {
    const trimmed = rawText.trim();
    if (!trimmed) {
      return { valid: false, error: "Reply text cannot be empty" };
    }
    const lines = trimmed.split(`
`).map((l) => l.trim()).filter(Boolean);
    let tagLinesCount = 0;
    for (const line of lines) {
      if (line.startsWith("@")) {
        tagLinesCount++;
      }
    }
    if (tagLinesCount > 1 || lines.length > 1) {
      return {
        valid: false,
        error: "Multiple answers in a single reply are not allowed. Please submit one answer per reply."
      };
    }
    const allTags = (trimmed.match(/@[a-zA-Z0-9_-]+/g) || []).map((t) => t.replace("@", "").toLowerCase());
    const uniqueTags = Array.from(new Set(allTags));
    const hasExplicitTag = Boolean(explicitTag && explicitTag.trim());
    const cleanExplicitTag = hasExplicitTag ? explicitTag.replace(/^@/, "").toLowerCase().trim() : "";
    if (uniqueTags.length > 1) {
      return {
        valid: false,
        error: "Each reply can have only one @list tagging"
      };
    }
    if (hasExplicitTag && uniqueTags.length === 1 && uniqueTags[0] !== cleanExplicitTag) {
      return {
        valid: false,
        error: "Conflicting tags: inline tag does not match selected tag"
      };
    }
    const selectedTag = cleanExplicitTag || (uniqueTags[0] || "all");
    let cleanAnswer = trimmed;
    if (uniqueTags.length === 1) {
      cleanAnswer = cleanAnswer.replace(new RegExp(`@${uniqueTags[0]}\\s*`, "i"), "").trim();
    }
    if (!cleanAnswer) {
      return { valid: false, error: "Please provide an answer for this reply" };
    }
    return {
      valid: true,
      tag: selectedTag,
      answer: cleanAnswer
    };
  }
  function parseAudienceAnswers(rawText) {
    const answers = {};
    const lines = rawText.split(`
`).map((l) => l.trim()).filter(Boolean);
    let currentTags = ["all"];
    for (const line of lines) {
      const match = line.match(/^((?:@[a-zA-Z0-9_-]+[,\s]*)+)(.*)$/);
      if (match && match[1].trim()) {
        const tagPart = match[1];
        const content = match[2].trim();
        const tags = (tagPart.match(/@[a-zA-Z0-9_-]+/g) || []).map((t) => t.replace("@", "").toLowerCase());
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
      for (const tag of currentTags) {
        if (!answers[tag]) {
          answers[tag] = line;
        } else {
          answers[tag] += " " + line;
        }
      }
    }
    if (Object.keys(answers).length === 0 && rawText.trim()) {
      answers["all"] = rawText.trim();
    }
    return answers;
  }

  class TinyLLM {
    vocab = new Map;
    invVocab = new Map;
    nextTokenId = 0;
    embeddingDim = 32;
    maxContextLength = 32;
    learningRate = 0.05;
    embeddings = new Map;
    transitions = new Map;
    contextWeights;
    memoryStore = [];
    qaStore = [];
    userCredits = new Map;
    userLastQuestionTime = new Map;
    alexCredits = 0;
    trainingSteps = 0;
    runningLoss = 0.45;
    storageKey = "hyperstack:tiny_llm_state";
    constructor() {
      this.contextWeights = new Float32Array(this.embeddingDim);
      for (let i = 0;i < this.embeddingDim; i++) {
        this.contextWeights[i] = (Math.random() - 0.5) * 0.1;
      }
      this.initDefaultVocab();
      this.seedAlexKnowledge();
      this.trainingSteps = 0;
      this.loadPersistedState();
    }
    tokenize(text) {
      const matches = text.toLowerCase().match(/[\p{L}\p{N}]+|[^\s\p{L}\p{N}]/gu);
      return matches || [];
    }
    getOrAddToken(token) {
      if (this.vocab.has(token)) {
        return this.vocab.get(token);
      }
      const id = this.nextTokenId++;
      this.vocab.set(token, id);
      this.invVocab.set(id, token);
      const emb = new Float32Array(this.embeddingDim);
      for (let i = 0;i < this.embeddingDim; i++) {
        emb[i] = (Math.random() - 0.5) * 0.2;
      }
      this.embeddings.set(id, emb);
      this.transitions.set(id, new Map);
      return id;
    }
    initDefaultVocab() {
      const specials = ["<pad>", "<unk>", "<s>", "</s>", "<user>", "<ai>", "<sys>"];
      for (const s of specials) {
        this.getOrAddToken(s);
      }
    }
    seedAlexKnowledge() {
      this.answerQuestion("Where are you now?", {
        all: "im at work, chennai",
        family: "im going for shopping today",
        friends: "im at home with my family",
        bffs: "im dating with name, quiet a time so far",
        marcus: "im at home with my family (hey Marcus! grab a slice if you swing by \uD83C\uDF55)",
        sarah: "im dating with name, quiet a time so far (tell you all the tea later Sarah! \uD83E\uDD2B)"
      });
      this.addPendingQuestion("What are your plans for this weekend?", "@sarah", false);
      this.addPendingQuestion("Are you going to the tech meetup tomorrow?", "@marcus", false);
      this.answerQuestion("Do you have a girlfriend?", {
        all: "Haha nope, currently single! 100% focused on engineering open-source systems and building Hyperstack right now \uD83D\uDE80",
        bffs: "Haha you know my dating life inside out! Focused on engineering right now \uD83D\uDE09",
        relationship: "Haha, what a question — you already know you have my whole heart! ❤️ \uD83E\uDD70 So glad you're here chatting with my offline clone."
      });
      this.answerQuestion("What is your morning coffee routine?", {
        all: "Huge pour-over enthusiast! V60 with light roast Ethiopian Yirgacheffe natural beans is his daily ritual ☕️✨"
      });
      this.answerQuestion("Where do you live?", {
        all: "Alex lives and works in San Francisco, California \uD83C\uDF09"
      });
      this.addPendingQuestion("What are your plans for this weekend?", "Anonymous", true);
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
    getQAStore() {
      return this.qaStore;
    }
    getPendingCount() {
      return this.qaStore.filter((q) => q.status === "pending").length;
    }
    addPendingQuestion(question, askedBy = "Anonymous", isAnon = true) {
      const cleanQ = question.trim();
      const existing = this.findQAByQuestion(cleanQ);
      if (existing) {
        return existing;
      }
      const tokens = this.tokenize(cleanQ);
      const item = {
        id: `qa-${Date.now()}-${Math.random().toString(36).substring(2, 6)}`,
        question: cleanQ,
        tokens,
        tokenSet: new Set(tokens),
        answers: {},
        status: "pending",
        askedBy: isAnon ? "Anonymous" : askedBy || "Visitor",
        isAnon,
        askedAt: Date.now()
      };
      this.qaStore.unshift(item);
      this.persistState();
      return item;
    }
    answerQuestion(questionIdOrText, answersInput) {
      let item = this.qaStore.find((q) => q.id === questionIdOrText);
      if (!item) {
        item = this.findQAByQuestion(questionIdOrText);
      }
      const answersObj = typeof answersInput === "string" ? parseAudienceAnswers(answersInput) : answersInput;
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
          status: "answered",
          askedBy: "Visitor",
          isAnon: false,
          askedAt: Date.now(),
          answeredAt: Date.now()
        };
        this.qaStore.unshift(item);
      } else {
        item.answers = { ...item.answers, ...answersObj };
        item.status = "answered";
        item.answeredAt = Date.now();
      }
      if (!item.replies)
        item.replies = [];
      for (const [tag, ans] of Object.entries(answersObj)) {
        const existingRep = item.replies.find((r) => r.tag === tag);
        if (existingRep) {
          existingRep.text = ans;
          existingRep.timestamp = Date.now();
        } else {
          item.replies.push({
            id: `rep-${Date.now()}-${Math.random().toString(36).substring(2, 6)}`,
            tag,
            text: ans,
            author: "Alex",
            timestamp: Date.now()
          });
        }
      }
      for (const [tag, ans] of Object.entries(item.answers)) {
        this.learnFact(`[${tag}] ${item.question} -> ${ans}`, "qa", false);
      }
      this.persistState();
      return item;
    }
    normalizeQuestion(q) {
      return q.toLowerCase().replace(/^alex\s*ai\s*[:,]?\s*/i, "").replace(/^alex\s*[:,]?\s*/i, "").replace(/^hey\s*alex\s*[:,]?\s*/i, "").replace(/^hi\s*alex\s*[:,]?\s*/i, "").replace(/[^a-z0-9\s]/g, "").replace(/\s+/g, " ").trim();
    }
    findQAByQuestion(query) {
      const normQuery = this.normalizeQuestion(query);
      if (!normQuery)
        return null;
      for (const item of this.qaStore) {
        if (this.normalizeQuestion(item.question) === normQuery) {
          return item;
        }
      }
      const queryTokens = this.tokenize(normQuery);
      if (queryTokens.length === 0)
        return null;
      let bestItem = null;
      let bestScore = 0;
      for (const item of this.qaStore) {
        const itemTokens = this.tokenize(this.normalizeQuestion(item.question));
        if (itemTokens.length === 0)
          continue;
        let overlap = 0;
        for (const t of queryTokens) {
          if (item.tokenSet.has(t))
            overlap++;
        }
        const score = overlap / Math.max(queryTokens.length, itemTokens.length);
        if (score > bestScore) {
          bestScore = score;
          bestItem = item;
        }
      }
      return bestScore >= 0.7 ? bestItem : null;
    }
    getContextualAnswer(query, audienceRole = "all", username) {
      const item = this.findQAByQuestion(query);
      if (!item) {
        return { matched: false, isPending: true };
      }
      if (item.status === "pending") {
        return {
          matched: true,
          isPending: true,
          question: item.question,
          qaItem: item
        };
      }
      if (username) {
        const userKey = username.replace(/^@/, "").toLowerCase().trim();
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
      const roleKey = (audienceRole || "all").toLowerCase() === "public" ? "all" : (audienceRole || "all").toLowerCase();
      let answer = item.answers[roleKey];
      let matchedAudience = roleKey;
      if (!answer && item.answers["all"]) {
        answer = item.answers["all"];
        matchedAudience = "all";
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
        matchedAudience: "none",
        qaItem: item
      };
    }
    getUnansweredQuestions() {
      return this.qaStore.filter((q) => q.status === "pending").sort((a, b) => b.askedAt - a.askedAt);
    }
    getAnsweredQuestions() {
      return this.qaStore.filter((q) => q.status === "answered").sort((a, b) => (b.answeredAt || b.askedAt) - (a.answeredAt || a.askedAt));
    }
    addSingleReply(questionIdOrText, replyText, explicitTag, author = "Alex") {
      const validation = validateSingleReply(replyText, explicitTag);
      if (!validation.valid || !validation.tag || !validation.answer) {
        throw new Error(validation.error || "Invalid reply");
      }
      let item = this.qaStore.find((q) => q.id === questionIdOrText);
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
          status: "answered",
          askedBy: "Visitor",
          isAnon: false,
          askedAt: Date.now(),
          answeredAt: Date.now()
        };
        this.qaStore.unshift(item);
      }
      if (!item.replies) {
        item.replies = [];
      }
      const threadReply = {
        id: `rep-${Date.now()}-${Math.random().toString(36).substring(2, 6)}`,
        tag: validation.tag,
        text: validation.answer,
        author,
        timestamp: Date.now()
      };
      item.replies.push(threadReply);
      item.answers[validation.tag] = validation.answer;
      item.status = "answered";
      item.answeredAt = Date.now();
      this.learnFact(`[${validation.tag}] ${item.question} -> ${validation.answer}`, "qa", false);
      this.persistState();
      return threadReply;
    }
    getQuestionReplies(questionIdOrText) {
      const item = this.qaStore.find((q) => q.id === questionIdOrText) || this.findQAByQuestion(questionIdOrText);
      if (!item)
        return [];
      if (item.replies && item.replies.length > 0)
        return item.replies;
      return Object.entries(item.answers).map(([tag, text], idx) => ({
        id: `rep-legacy-${idx}`,
        tag,
        text,
        author: "Alex",
        timestamp: item.answeredAt || item.askedAt
      }));
    }
    getUserCredits(username) {
      const cleanUser = (username || "visitor").replace(/^@/, "").toLowerCase().trim();
      if (!this.userCredits.has(cleanUser)) {
        this.userCredits.set(cleanUser, DEFAULT_INITIAL_CREDITS);
      }
      return this.userCredits.get(cleanUser);
    }
    setUserCredits(username, credits) {
      const cleanUser = (username || "visitor").replace(/^@/, "").toLowerCase().trim();
      this.userCredits.set(cleanUser, Math.max(0, credits));
      this.persistState();
    }
    getAlexCredits() {
      return this.alexCredits;
    }
    setAlexCredits(credits) {
      this.alexCredits = Math.max(0, credits);
      this.persistState();
    }
    canUserAskQuestion(username, now = Date.now()) {
      const cleanUser = (username || "visitor").replace(/^@/, "").toLowerCase().trim();
      const credits = this.getUserCredits(cleanUser);
      if (credits <= 0) {
        return {
          allowed: false,
          reason: `Insufficient credits: You have 0 credits. You need at least 1 credit to ask a question.`,
          cooldownSeconds: 0,
          credits: 0,
          alexCredits: this.alexCredits
        };
      }
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
    recordQuestionAsked(username, now = Date.now()) {
      const cleanUser = (username || "visitor").replace(/^@/, "").toLowerCase().trim();
      this.userLastQuestionTime.set(cleanUser, now);
      this.persistState();
    }
    likeAnswer(questionIdOrText, tag = "all", username) {
      const cleanUser = (username || "visitor").replace(/^@/, "").toLowerCase().trim();
      const cleanTag = tag.replace(/^@/, "").toLowerCase().trim() || "all";
      let item = this.qaStore.find((q) => q.id === questionIdOrText) || this.findQAByQuestion(questionIdOrText);
      if (!item) {
        return {
          success: false,
          userCredits: this.getUserCredits(cleanUser),
          alexCredits: this.alexCredits,
          error: "Question not found"
        };
      }
      item.likedBy = item.likedBy || [];
      const likeKey = `${cleanUser}:${cleanTag}`;
      if (item.likedBy.includes(likeKey) || item.likedBy.includes(cleanUser)) {
        return {
          success: false,
          userCredits: this.getUserCredits(cleanUser),
          alexCredits: this.alexCredits,
          error: "You have already liked this answer",
          qaItem: item
        };
      }
      const currentCredits = this.getUserCredits(cleanUser);
      if (currentCredits <= 0) {
        return {
          success: false,
          userCredits: 0,
          alexCredits: this.alexCredits,
          error: "Insufficient credits: You need at least 1 credit to like an answer and reward Alex",
          qaItem: item
        };
      }
      this.userCredits.set(cleanUser, currentCredits - 1);
      this.alexCredits += 1;
      item.likedBy.push(likeKey);
      item.likesCount = (item.likesCount || 0) + 1;
      const answerText = item.answers[cleanTag] || item.answers["all"] || "";
      if (answerText) {
        this.feedback(item.question, answerText, 1);
      }
      this.persistState();
      return {
        success: true,
        userCredits: this.getUserCredits(cleanUser),
        alexCredits: this.alexCredits,
        qaItem: item
      };
    }
    isAnswerLiked(questionIdOrText, tag = "all", username) {
      const cleanUser = (username || "visitor").replace(/^@/, "").toLowerCase().trim();
      const cleanTag = tag.replace(/^@/, "").toLowerCase().trim() || "all";
      const item = this.qaStore.find((q) => q.id === questionIdOrText) || this.findQAByQuestion(questionIdOrText);
      if (!item || !item.likedBy)
        return false;
      const likeKey = `${cleanUser}:${cleanTag}`;
      return item.likedBy.includes(likeKey) || item.likedBy.includes(cleanUser);
    }
    textEmbedding(tokens) {
      const pooled = new Float32Array(this.embeddingDim);
      if (tokens.length === 0)
        return pooled;
      for (const id of tokens) {
        const emb = this.embeddings.get(id);
        if (emb) {
          for (let i = 0;i < this.embeddingDim; i++) {
            pooled[i] += emb[i];
          }
        }
      }
      let norm = 0;
      for (let i = 0;i < this.embeddingDim; i++) {
        norm += pooled[i] * pooled[i];
      }
      norm = Math.sqrt(norm);
      if (norm > 0) {
        for (let i = 0;i < this.embeddingDim; i++) {
          pooled[i] /= norm;
        }
      }
      return pooled;
    }
    cosineSimilarity(a, b) {
      let dot = 0;
      let normA = 0;
      let normB = 0;
      for (let i = 0;i < this.embeddingDim; i++) {
        dot += a[i] * b[i];
        normA += a[i] * a[i];
        normB += b[i] * b[i];
      }
      const denom = Math.sqrt(normA) * Math.sqrt(normB);
      return denom > 0 ? dot / denom : 0;
    }
    computeRelevance(queryTokenSet, queryEmb, mem) {
      let tokenOverlap = 0;
      for (const t of mem.tokenSet) {
        if (queryTokenSet.has(t)) {
          tokenOverlap++;
        }
      }
      const jaccard = mem.tokenSet.size > 0 && queryTokenSet.size > 0 ? tokenOverlap / Math.sqrt(mem.tokenSet.size * queryTokenSet.size) : 0;
      const cosine = this.cosineSimilarity(queryEmb, mem.embedding);
      return jaccard * 0.75 + cosine * 0.25;
    }
    learnFact(fact, category = "custom", persist = true) {
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
      const tokenIds = rawTokens.map((t) => this.getOrAddToken(t));
      const tokenSet = new Set(tokenIds);
      for (let i = 0;i < tokenIds.length - 1; i++) {
        const current = tokenIds[i];
        const next = tokenIds[i + 1];
        const transMap = this.transitions.get(current) || new Map;
        const currentScore = transMap.get(next) || 0;
        transMap.set(next, currentScore + 1);
        this.transitions.set(current, transMap);
        if (i + 2 < tokenIds.length) {
          const next2 = tokenIds[i + 2];
          const skipScore = transMap.get(next2) || 0;
          transMap.set(next2, skipScore + 0.5);
        }
      }
      const embedding = this.textEmbedding(tokenIds);
      let matched = false;
      for (const mem of this.memoryStore) {
        if (this.cosineSimilarity(mem.embedding, embedding) > 0.88) {
          mem.fact = fact;
          mem.weight += 1;
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
          weight: 1,
          learnedAt: Date.now()
        });
      }
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
    onlineGradientStep(tokens) {
      const lr = this.learningRate;
      for (let i = 0;i < tokens.length - 1; i++) {
        const a = this.embeddings.get(tokens[i]);
        const b = this.embeddings.get(tokens[i + 1]);
        if (a && b) {
          for (let d = 0;d < this.embeddingDim; d++) {
            const diff = b[d] - a[d];
            a[d] += lr * diff * 0.1;
            b[d] -= lr * diff * 0.1;
          }
        }
      }
    }
    feedback(turnPrompt, turnResponse, reward = 1) {
      const text = `${turnPrompt} ${turnResponse}`;
      return this.learnFact(text, "reinforcement", true);
    }
    async generate(prompt, options = {}) {
      const startTime = performance.now();
      const role = options.contextRole ?? "public";
      const cleanPrompt = prompt.replace(/^alex\s*ai\s*[:,]?\s*/i, "").replace(/^alex\s*[:,]?\s*/i, "").replace(/^hey\s*alex\s*[:,]?\s*/i, "").replace(/^hi\s*alex\s*[:,]?\s*/i, "").trim();
      const contextual = this.getContextualAnswer(cleanPrompt || prompt, role, options.username);
      if (contextual.matched && !contextual.isPending && contextual.answer) {
        return {
          text: contextual.answer,
          tokens: this.tokenize(contextual.answer),
          confidence: 1,
          trainingStep: this.trainingSteps,
          retrievedMemory: `Contextual Answer (@${contextual.matchedAudience})`,
          loss: this.runningLoss,
          latencyMs: Math.round(performance.now() - startTime),
          isPending: false,
          matchedAudience: contextual.matchedAudience
        };
      }
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
      const queryTokens = this.tokenize(cleanPrompt || prompt);
      const queryIds = queryTokens.map((t) => this.getOrAddToken(t));
      const queryTokenSet = new Set(queryIds);
      const queryEmb = this.textEmbedding(queryIds);
      let bestMem = null;
      let maxRel = -1;
      for (const mem of this.memoryStore) {
        const rel = this.computeRelevance(queryTokenSet, queryEmb, mem);
        if (rel > maxRel) {
          maxRel = rel;
          bestMem = mem;
        }
      }
      if (bestMem && maxRel >= 0.28 && (bestMem.category === "custom" || bestMem.category === "hobby")) {
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
      this.addPendingQuestion(cleanPrompt || prompt, options.contextRole || "public", options.isAnon ?? false);
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
    synthesizeFromMemory(prompt, fact) {
      const lower = prompt.toLowerCase();
      if (lower.includes("project") || lower.includes("hyperstack")) {
        return "Alex is currently building Hyperstack (zero-eval reactive hypermedia framework) and Hypersheet (topological spreadsheet core in Go) ⚡";
      }
      if (lower.includes("run") || lower.includes("marathon")) {
        return "Alex runs 5km every morning around the Embarcadero and is training for a sub-3:30 marathon in SF \uD83C\uDFC3‍♂️\uD83D\uDCA8";
      }
      if (lower.includes("hobby") || lower.includes("chess") || lower.includes("keyboard")) {
        return "Pour-over coffee ☕️, morning 5km runs \uD83C\uDFC3‍♂️, blitz chess (1800 Elo) ♟️, and custom mechanical keyboards ⌨️!";
      }
      return `${fact} \uD83D\uDCA1`;
    }
    getStats() {
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
        answeredQuestions: this.qaStore.filter((q) => q.status === "answered").length,
        alexCredits: this.alexCredits
      };
    }
    persistState() {
      if (typeof localStorage === "undefined")
        return;
      try {
        const state = {
          trainingSteps: this.trainingSteps,
          runningLoss: this.runningLoss,
          memories: this.memoryStore.map((m) => ({
            id: m.id,
            fact: m.fact,
            category: m.category,
            weight: m.weight,
            learnedAt: m.learnedAt
          })),
          qaStore: this.qaStore.map((q) => ({
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
      } catch {}
    }
    loadPersistedState() {
      if (typeof localStorage === "undefined")
        return;
      try {
        const raw = localStorage.getItem(this.storageKey);
        if (!raw)
          return;
        const parsed = JSON.parse(raw);
        if (parsed.trainingSteps)
          this.trainingSteps = parsed.trainingSteps;
        if (parsed.runningLoss)
          this.runningLoss = parsed.runningLoss;
        if (Array.isArray(parsed.memories)) {
          for (const m of parsed.memories) {
            if (m.fact && typeof m.fact === "string") {
              this.learnFact(m.fact, m.category || "persisted", false);
            }
          }
        }
        if (typeof parsed.alexCredits === "number")
          this.alexCredits = parsed.alexCredits;
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
              const existing = this.qaStore.find((item) => item.id === q.id || item.question.toLowerCase() === q.question.toLowerCase());
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
      } catch {}
    }
    resetMemory() {
      if (typeof localStorage !== "undefined") {
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
  if (typeof window !== "undefined") {
    window.TinyLLM = TinyLLM;
    window.parseAudienceAnswers = parseAudienceAnswers;
    window.checkUsernameCooldown = checkUsernameCooldown;
    window.USERNAME_COOLDOWN_MS = USERNAME_COOLDOWN_MS;
    window.validateSingleReply = validateSingleReply;
    window.$llm = new TinyLLM;
  }
})();
