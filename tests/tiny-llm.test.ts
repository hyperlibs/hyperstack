import { describe, expect, test } from "bun:test";
import { TinyLLM, checkUsernameCooldown, USERNAME_COOLDOWN_DAYS, USERNAME_COOLDOWN_MS, validateSingleReply, ThreadReply, QUESTION_RATE_LIMIT_MS, DEFAULT_INITIAL_CREDITS } from "../src/tiny-llm";

describe("TinyLLM Embedded On-Device Self-Learning Language Model", () => {
  test("1. Model initializes with baseline vocab, embeddings, and parameters", () => {
    const llm = new TinyLLM();
    const stats = llm.getStats();

    expect(stats.parameters).toBeGreaterThan(1000);
    expect(stats.vocabSize).toBeGreaterThan(50);
    expect(stats.memoryNodes).toBeGreaterThanOrEqual(10);
    expect(stats.isOffline).toBe(true);
    expect(stats.version).toBe("1.0.0-tiny");
  });

  test("2. Generates verified answers to baseline questions offline", async () => {
    const llm = new TinyLLM();

    const gfRes = await llm.generate("Do you have a girlfriend?");
    expect(gfRes.text.toLowerCase()).toContain("single");
    expect(gfRes.confidence).toBeGreaterThan(0.4);
    expect(gfRes.latencyMs).toBeLessThan(100);

    const coffeeRes = await llm.generate("What is your morning coffee routine?");
    expect(coffeeRes.text.toLowerCase()).toContain("pour-over");
    expect(coffeeRes.text.toLowerCase()).toContain("ethiopian");

    const locRes = await llm.generate("Where do you live?");
    expect(locRes.text.toLowerCase()).toContain("san francisco");
  });

  test("3. Adapts responses dynamically for Relationship collection role", async () => {
    const llm = new TinyLLM();

    const publicRes = await llm.generate("Do you have a girlfriend?", { contextRole: "public" });
    expect(publicRes.text.toLowerCase()).toContain("single");

    const partnerRes = await llm.generate("Do you have a girlfriend?", { contextRole: "relationship" });
    expect(partnerRes.text).toContain("you already know you have my whole heart! ❤️");
    expect(partnerRes.confidence).toBe(1.0);
  });

  test("4. Self-Learning: learnFact absorbs new knowledge and updates weights in real-time", async () => {
    const llm = new TinyLLM();
    const initialStats = llm.getStats();

    // Teach a brand new unique fact
    const newFact = "Alex builds modular analog synthesizers and eurorack audio gear in his spare time.";
    const learnRes = llm.learnFact(newFact, "hobby", false);

    expect(learnRes.step).toBeGreaterThan(initialStats.trainingSteps);
    expect(learnRes.tokensLearned).toBeGreaterThan(5);
    expect(learnRes.memoryNodesCount).toBeGreaterThan(initialStats.memoryNodes);

    // Query the LLM about the newly taught topic
    const queryRes = await llm.generate("What audio gear or modular synthesizers do you build?");
    expect(queryRes.text).toContain("modular analog synthesizers");
    expect(queryRes.retrievedMemory).toBe(newFact);
  });

  test("5. Online gradient step and loss decay over training steps", () => {
    const llm = new TinyLLM();
    const lossBefore = llm.getStats().runningLoss;

    llm.learnFact("Zero-eval hypermedia systems provide optimal memory efficiency.", "tech", false);
    llm.learnFact("Hypersheet calculates topological spreadsheets under one millisecond.", "tech", false);

    const lossAfter = llm.getStats().runningLoss;
    expect(lossAfter).toBeLessThan(lossBefore);
  });

  test("6. Turn feedback reinforces learning", () => {
    const llm = new TinyLLM();
    const initialSteps = llm.getStats().trainingSteps;

    const res = llm.feedback("What is your favorite editor?", "Alex exclusively codes in Neovim with Lua configs.");
    expect(res.step).toBe(initialSteps + 1);
  });

  test("7. Reset restores clean baseline knowledge", () => {
    const llm = new TinyLLM();
    llm.learnFact("Temporary testing fact 12345", "test", false);
    llm.resetMemory();

    const stats = llm.getStats();
    expect(stats.trainingSteps).toBe(0);
    expect(stats.memoryNodes).toBeGreaterThanOrEqual(10);
  });

  test("8. Contextual Multi-Audience Responses: 'Where are you now?' answers per collection", async () => {
    const llm = new TinyLLM();

    // @all / public
    const allRes = await llm.generate("Where are you now?", { contextRole: "all" });
    expect(allRes.text).toBe("im at work, chennai");
    expect(allRes.matchedAudience).toBe("all");
    expect(allRes.isPending).toBe(false);

    // @family
    const familyRes = await llm.generate("Where are you now?", { contextRole: "family" });
    expect(familyRes.text).toBe("im going for shopping today");
    expect(familyRes.matchedAudience).toBe("family");

    // @friends
    const friendsRes = await llm.generate("Where are you now?", { contextRole: "friends" });
    expect(friendsRes.text).toBe("im at home with my family");
    expect(friendsRes.matchedAudience).toBe("friends");

    // @bffs
    const bffsRes = await llm.generate("Where are you now?", { contextRole: "bffs" });
    expect(bffsRes.text).toBe("im dating with name, quiet a time so far");
    expect(bffsRes.matchedAudience).toBe("bffs");

    // Natural variation with Alexai prefix
    const prefixedRes = await llm.generate("Alexai Where are you now?", { contextRole: "family" });
    expect(prefixedRes.text).toBe("im going for shopping today");
  });

  test("9. Unanswered Question Protocol: AlexAI does not reply at all and queues to Alex's Inbox", async () => {
    const llm = new TinyLLM();
    const initialPending = llm.getPendingCount();

    // Ask an unknown question
    const unansweredRes = await llm.generate("What car do you drive?");
    expect(unansweredRes.text).toBe("");
    expect(unansweredRes.confidence).toBe(0);
    expect(unansweredRes.isPending).toBe(true);

    // Verify it was added to pending questions
    expect(llm.getPendingCount()).toBe(initialPending + 1);
  });

  test("10. Middleman Answer Flow: Alex provides multi-tag answers and AlexAI responds contextually", async () => {
    const llm = new TinyLLM();

    const question = "What did you do over the weekend?";
    // First, user asks -> AlexAI does not reply
    const beforeRes = await llm.generate(question);
    expect(beforeRes.isPending).toBe(true);
    expect(beforeRes.text).toBe("");

    // Alex answers with multi-tag format
    llm.answerQuestion(question, `@all Went hiking at Marin Headlands 🥾\n@family Attended cousin's wedding 💒\n@friends Played poker night until 3am 🃏`);

    // Verify AlexAI now responds according to the asker's context
    const publicRes = await llm.generate(question, { contextRole: "public" });
    expect(publicRes.text).toBe("Went hiking at Marin Headlands 🥾");

    const familyRes = await llm.generate(question, { contextRole: "family" });
    expect(familyRes.text).toBe("Attended cousin's wedding 💒");

    const friendsRes = await llm.generate(question, { contextRole: "friends" });
    expect(friendsRes.text).toBe("Played poker night until 3am 🃏");
  });

  test("11. People usernames used for identification & multi-audience routing", async () => {
    const llm = new TinyLLM();
    const q = "Where are you now?";

    // Explicitly configure contextual answers with specific username overrides
    llm.answerQuestion(q, {
      all: "im at work, chennai",
      family: "im going for shopping today",
      friends: "im at home with my family",
      bffs: "im dating with name, quiet a time so far",
      marcus: "im at home with my family (hey Marcus! grab a slice if you swing by 🍕)",
      sarah: "im dating with name, quiet a time so far (tell you all the tea later Sarah! 🤫)"
    });

    // Query as @marcus -> matches specific @marcus answer
    const marcusRes = await llm.generate(q, { username: "marcus", contextRole: "friends" });
    expect(marcusRes.text).toContain("grab a slice if you swing by 🍕");
    expect(marcusRes.matchedAudience).toBe("@marcus");

    // Query as @sarah -> matches specific @sarah answer
    const sarahRes = await llm.generate(q, { username: "sarah", contextRole: "bffs" });
    expect(sarahRes.text).toContain("tell you all the tea later Sarah! 🤫");
    expect(sarahRes.matchedAudience).toBe("@sarah");

    // Multi-tag answer with multiple comma-separated usernames & groups
    const newQ = "What project are you shipping today?";
    llm.answerQuestion(newQ, `@all Shipping Hyperstack v1.1 🚀\n@friends, @marcus, @priya Shipping with custom audio synthesis 🎹\n@bffs, @sarah Shipping secret surprise feature 🤫`);

    const priyaRes = await llm.generate(newQ, { username: "priya" });
    expect(priyaRes.text).toBe("Shipping with custom audio synthesis 🎹");
    expect(priyaRes.matchedAudience).toBe("@priya");

    const generalFriendRes = await llm.generate(newQ, { username: "dave", contextRole: "friends" });
    expect(generalFriendRes.text).toBe("Shipping with custom audio synthesis 🎹");
    expect(generalFriendRes.matchedAudience).toBe("friends");

    const strangerRes = await llm.generate(newQ, { username: "stranger", contextRole: "public" });
    expect(strangerRes.text).toBe("Shipping Hyperstack v1.1 🚀");
    expect(strangerRes.matchedAudience).toBe("all");
  });

  test("12. Username change policy: Usernames can be changed only once a year (365 days cooldown)", () => {
    expect(USERNAME_COOLDOWN_DAYS).toBe(365);
    expect(USERNAME_COOLDOWN_MS).toBe(31536000000);

    // Initial check (no prior change) -> Allowed
    const initialCheck = checkUsernameCooldown(0);
    expect(initialCheck.allowed).toBe(true);
    expect(initialCheck.daysRemaining).toBe(0);

    // Changed 100 days ago -> Blocked with remaining days
    const hundredDaysAgo = Date.now() - (100 * 24 * 60 * 60 * 1000);
    const blockedCheck = checkUsernameCooldown(hundredDaysAgo);
    expect(blockedCheck.allowed).toBe(false);
    expect(blockedCheck.daysRemaining).toBe(265);
    expect(blockedCheck.reason).toContain("Usernames can only be changed once a year");

    // Changed 366 days ago -> Cooldown expired, allowed again
    const overOneYearAgo = Date.now() - (366 * 24 * 60 * 60 * 1000);
    const allowedCheck = checkUsernameCooldown(overOneYearAgo);
    expect(allowedCheck.allowed).toBe(true);
    expect(allowedCheck.daysRemaining).toBe(0);
  });
  test("13. validateSingleReply enforces strictly one answer and only one @list tag per reply", () => {
    // 1. Multiple @list tags in single string -> REJECT
    const multiTagRes = validateSingleReply("@family, @friends im going out today");
    expect(multiTagRes.valid).toBe(false);
    expect(multiTagRes.error).toContain("Each reply can have only one @list tagging");

    // 2. Multiple answers in separate lines -> REJECT
    const multiAnswerRes = validateSingleReply("@all im at work\n@family im shopping");
    expect(multiAnswerRes.valid).toBe(false);
    expect(multiAnswerRes.error).toContain("Multiple answers in a single reply are not allowed");

    // 3. Valid single reply with inline @list tag -> ACCEPT
    const validInlineRes = validateSingleReply("@family im going for shopping today");
    expect(validInlineRes.valid).toBe(true);
    expect(validInlineRes.tag).toBe("family");
    expect(validInlineRes.answer).toBe("im going for shopping today");

    // 4. Valid single reply with explicit tag button -> ACCEPT
    const validExplicitRes = validateSingleReply("im dating with name, quiet a time so far", "bffs");
    expect(validExplicitRes.valid).toBe(true);
    expect(validExplicitRes.tag).toBe("bffs");
    expect(validExplicitRes.answer).toBe("im dating with name, quiet a time so far");

    // 5. Conflicting explicit tag vs inline tag -> REJECT
    const conflictRes = validateSingleReply("@family at home", "friends");
    expect(conflictRes.valid).toBe(false);
    expect(conflictRes.error).toContain("Conflicting tags");
  });

  test("14. addSingleReply appends reply comment threads to questions and routes contextually", async () => {
    const llm = new TinyLLM();
    const q = "Where are you spending the holidays?";

    // Post single reply 1 for @all
    const r1 = llm.addSingleReply(q, "@all im at work, chennai");
    expect(r1.tag).toBe("all");
    expect(r1.text).toBe("im at work, chennai");
    expect(r1.author).toBe("Alex");

    // Post single reply 2 for @family
    const r2 = llm.addSingleReply(q, "im going for shopping today", "family");
    expect(r2.tag).toBe("family");
    expect(r2.text).toBe("im going for shopping today");

    // Post single reply 3 for @friends
    const r3 = llm.addSingleReply(q, "@friends im at home with my family");
    expect(r3.tag).toBe("friends");

    // Post single reply 4 for @bffs
    const r4 = llm.addSingleReply(q, "im dating with name, quiet a time so far", "bffs");
    expect(r4.tag).toBe("bffs");

    // Attempting to post multiple answers in one reply must throw
    expect(() => {
      llm.addSingleReply(q, "@all line 1\n@family line 2");
    }).toThrow("Multiple answers in a single reply are not allowed");

    // Retrieve thread replies
    const thread = llm.getQuestionReplies(q);
    expect(thread.length).toBe(4);
    expect(thread.map(t => t.tag)).toEqual(["all", "family", "friends", "bffs"]);

    // Verify contextual generation routes properly
    const familyQuery = await llm.generate(q, { contextRole: "family" });
    expect(familyQuery.text).toBe("im going for shopping today");

    const friendsQuery = await llm.generate(q, { contextRole: "friends" });
    expect(friendsQuery.text).toBe("im at home with my family");

    const bffsQuery = await llm.generate(q, { contextRole: "bffs" });
    expect(bffsQuery.text).toBe("im dating with name, quiet a time so far");

    const publicQuery = await llm.generate(q, { contextRole: "public" });
    expect(publicQuery.text).toBe("im at work, chennai");
  });
  test("15. Rate limiting: restrict users to ask 1 question per minute", () => {
    const llm = new TinyLLM();
    const username = "persong";
    const startTime = 1000000;

    expect(QUESTION_RATE_LIMIT_MS).toBe(60000);

    // Initial check -> allowed
    const check1 = llm.canUserAskQuestion(username, startTime);
    expect(check1.allowed).toBe(true);

    // User asks question at startTime
    llm.recordQuestionAsked(username, startTime);

    // User attempts to ask another question 20 seconds later -> BLOCKED
    const check2 = llm.canUserAskQuestion(username, startTime + 20000);
    expect(check2.allowed).toBe(false);
    expect(check2.cooldownSeconds).toBe(40);
    expect(check2.reason).toContain("Rate limit: You can only ask 1 question per minute");

    // User attempts to ask another question 59.5 seconds later -> BLOCKED (1s remaining)
    const check3 = llm.canUserAskQuestion(username, startTime + 59500);
    expect(check3.allowed).toBe(false);
    expect(check3.cooldownSeconds).toBe(1);

    // User attempts to ask another question 61 seconds later -> ALLOWED
    const check4 = llm.canUserAskQuestion(username, startTime + 61000);
    expect(check4.allowed).toBe(true);
    expect(check4.cooldownSeconds).toBeUndefined();
  });

  test("16. Credit economics: Liking answers transfers 1 credit from questioning user to Alex & users with 0 credits cannot ask", () => {
    const llm = new TinyLLM();
    const userG = "persong"; // Person G

    expect(DEFAULT_INITIAL_CREDITS).toBe(3);

    // 1. Initial credit balances
    expect(llm.getUserCredits(userG)).toBe(3);
    expect(llm.getAlexCredits()).toBe(0);

    // 2. Person G has credits, can ask question
    const q1Check = llm.canUserAskQuestion(userG);
    expect(q1Check.allowed).toBe(true);
    expect(q1Check.credits).toBe(3);

    // Person G asks a question and Alex answers
    const question = "Are you available for consulting next week?";
    llm.addPendingQuestion(question, "@persong", false);
    llm.addSingleReply(question, "Yes, open for consulting on distributed systems!", "all", "Alex");

    // 3. Person G likes Alex's answer -> Alex earns 1 credit from Person G
    const likeRes1 = llm.likeAnswer(question, "all", userG);
    expect(likeRes1.success).toBe(true);
    expect(likeRes1.userCredits).toBe(2); // Person G balance: 3 - 1 = 2
    expect(likeRes1.alexCredits).toBe(1); // Alex balance: 0 + 1 = 1
    expect(llm.getUserCredits(userG)).toBe(2);
    expect(llm.getAlexCredits()).toBe(1);
    expect(llm.isAnswerLiked(question, "all", userG)).toBe(true);

    // Duplicate like attempt on same answer is rejected
    const dupLike = llm.likeAnswer(question, "all", userG);
    expect(dupLike.success).toBe(false);
    expect(dupLike.error).toContain("already liked");
    expect(llm.getUserCredits(userG)).toBe(2);
    expect(llm.getAlexCredits()).toBe(1);

    // 4. Person G likes 2 more answers until balance reaches 0
    const q2 = "What framework is this?";
    llm.addSingleReply(q2, "Hyperstack framework!", "all", "Alex");
    const likeRes2 = llm.likeAnswer(q2, "all", userG);
    expect(likeRes2.success).toBe(true);
    expect(likeRes2.userCredits).toBe(1);
    expect(likeRes2.alexCredits).toBe(2);

    const q3 = "Is Node.js used?";
    llm.addSingleReply(q3, "Zero Node.js, 100% Bun & Go!", "all", "Alex");
    const likeRes3 = llm.likeAnswer(q3, "all", userG);
    expect(likeRes3.success).toBe(true);
    expect(likeRes3.userCredits).toBe(0); // 0 credits left
    expect(likeRes3.alexCredits).toBe(3);

    // 5. "if person G doesnt have credits they cant ask questions"
    const zeroCreditsCheck = llm.canUserAskQuestion(userG, Date.now() + 100000);
    expect(zeroCreditsCheck.allowed).toBe(false);
    expect(zeroCreditsCheck.credits).toBe(0);
    expect(zeroCreditsCheck.reason).toContain("Insufficient credits: You have 0 credits. You need at least 1 credit to ask a question.");

    // Cannot like another answer when balance is 0
    const q4 = "Another question?";
    llm.addSingleReply(q4, "Another answer", "all", "Alex");
    const zeroLike = llm.likeAnswer(q4, "all", userG);
    expect(zeroLike.success).toBe(false);
    expect(zeroLike.error).toContain("Insufficient credits");
  });
});