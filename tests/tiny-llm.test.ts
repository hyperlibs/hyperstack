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
    expect(partnerRes.confidence).toBe(0.99);
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
    expect(stats.memoryNodes).toBe(10);
  });
});
