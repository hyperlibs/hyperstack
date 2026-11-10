import { describe, expect, test } from "bun:test";
import { TinyLLM } from "../src/tiny-llm";

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
});
