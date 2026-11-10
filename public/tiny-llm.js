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
    TinyLLM: () => TinyLLM
  });

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
      const temp = options.temperature ?? 0.65;
      const role = options.contextRole ?? "public";
      const lowerPrompt = prompt.toLowerCase();
      if (role === "relationship" && (lowerPrompt.includes("girlfriend") || lowerPrompt.includes("dating") || lowerPrompt.includes("love") || lowerPrompt.includes("relationship"))) {
        const answer = "Haha, what a question — you already know you have my whole heart! ❤️ \uD83E\uDD70 So glad you're here chatting with my offline clone.";
        return {
          text: answer,
          tokens: this.tokenize(answer),
          confidence: 0.99,
          trainingStep: this.trainingSteps,
          retrievedMemory: "Partner / Significant Other Relationship",
          loss: this.runningLoss,
          latencyMs: Math.round(performance.now() - startTime)
        };
      }
      if (lowerPrompt.includes("girlfriend") || lowerPrompt.includes("dating") || lowerPrompt.includes("single")) {
        const answer = "Haha nope, currently single! 100% focused on engineering open-source systems and building Hyperstack right now \uD83D\uDE80";
        return {
          text: answer,
          tokens: this.tokenize(answer),
          confidence: 0.98,
          trainingStep: this.trainingSteps,
          retrievedMemory: "Relationship Status: Single",
          loss: this.runningLoss,
          latencyMs: Math.round(performance.now() - startTime)
        };
      }
      if (lowerPrompt.includes("coffee") || lowerPrompt.includes("v60") || lowerPrompt.includes("morning") && lowerPrompt.includes("routine")) {
        const answer = "Huge pour-over enthusiast! V60 with light roast Ethiopian Yirgacheffe natural beans is his daily ritual ☕️✨";
        return {
          text: answer,
          tokens: this.tokenize(answer),
          confidence: 0.98,
          trainingStep: this.trainingSteps,
          retrievedMemory: "Daily Coffee Ritual: Ethiopian V60",
          loss: this.runningLoss,
          latencyMs: Math.round(performance.now() - startTime)
        };
      }
      if (lowerPrompt.includes("where") || lowerPrompt.includes("live") || lowerPrompt.includes("location")) {
        const answer = "Alex lives and works in San Francisco, California \uD83C\uDF09";
        return {
          text: answer,
          tokens: this.tokenize(answer),
          confidence: 0.98,
          trainingStep: this.trainingSteps,
          retrievedMemory: "Location: San Francisco, CA",
          loss: this.runningLoss,
          latencyMs: Math.round(performance.now() - startTime)
        };
      }
      const queryTokens = this.tokenize(prompt);
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
      let responseText = "";
      if (bestMem && maxRel > 0.12) {
        responseText = this.synthesizeFromMemory(prompt, bestMem.fact);
      } else {
        responseText = this.predictSequence(queryIds, temp);
      }
      if (!responseText || responseText.length < 5) {
        responseText = "Alex's offline clone is continuously learning! Ask about his projects (Hyperstack, Hypersheet), morning routine, or hobbies ☕️\uD83D\uDE80";
      }
      this.learnFact(`${prompt} -> ${responseText}`, "conversation", true);
      const latencyMs = Math.round(performance.now() - startTime);
      return {
        text: responseText,
        tokens: this.tokenize(responseText),
        confidence: Math.max(0.65, Math.min(0.98, maxRel > 0 ? maxRel : 0.7)),
        trainingStep: this.trainingSteps,
        retrievedMemory: bestMem ? bestMem.fact : null,
        loss: this.runningLoss,
        latencyMs
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
    predictSequence(seedTokens, temperature) {
      if (seedTokens.length === 0)
        return "";
      const generated = [];
      let currentId = seedTokens[seedTokens.length - 1];
      const maxGen = 14;
      for (let step = 0;step < maxGen; step++) {
        const trans = this.transitions.get(currentId);
        if (!trans || trans.size === 0)
          break;
        const entries = Array.from(trans.entries());
        const logits = [];
        let sumExp = 0;
        for (const [nextId, count] of entries) {
          const score = count / temperature;
          const expScore = Math.exp(Math.min(10, score));
          logits.push({ id: nextId, prob: expScore });
          sumExp += expScore;
        }
        if (sumExp === 0)
          break;
        let r = Math.random() * sumExp;
        let selectedId = entries[0][0];
        for (const l of logits) {
          r -= l.prob;
          if (r <= 0) {
            selectedId = l.id;
            break;
          }
        }
        const word = this.invVocab.get(selectedId);
        if (!word || word === "</s>" || word === "<pad>")
          break;
        generated.push(word);
        currentId = selectedId;
      }
      if (generated.length === 0) {
        return "Alex lives in San Francisco and builds Hyperstack ⚡";
      }
      return generated.join(" ");
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
        version: "1.0.0-tiny"
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
          }))
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
      } catch {}
    }
    resetMemory() {
      if (typeof localStorage !== "undefined") {
        localStorage.removeItem(this.storageKey);
      }
      this.memoryStore = [];
      this.runningLoss = 0.45;
      this.seedAlexKnowledge();
      this.trainingSteps = 0;
    }
  }
  if (typeof window !== "undefined") {
    window.TinyLLM = TinyLLM;
    window.$llm = new TinyLLM;
  }
})();
