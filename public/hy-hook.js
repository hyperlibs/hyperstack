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

  // hyperstack/src/hy-hook.ts
  var exports_hy_hook = {};
  __export(exports_hy_hook, {
    HyperPlane: () => HyperPlane,
    HyperHookEngine: () => HyperHookEngine,
    HyperFrameCodec: () => HyperFrameCodec,
    HyperFlag: () => HyperFlag,
    HyHook: () => HyHook
  });
  var HyperPlane;
  ((HyperPlane2) => {
    HyperPlane2[HyperPlane2["TENSOR"] = 0] = "TENSOR";
    HyperPlane2[HyperPlane2["HYPERMEDIA"] = 1] = "HYPERMEDIA";
    HyperPlane2[HyperPlane2["TELEMETRY"] = 2] = "TELEMETRY";
    HyperPlane2[HyperPlane2["PRECOGNITION"] = 3] = "PRECOGNITION";
    HyperPlane2[HyperPlane2["SWARM"] = 4] = "SWARM";
  })(HyperPlane ||= {});
  var HyperFlag;
  ((HyperFlag2) => {
    HyperFlag2[HyperFlag2["UNRELIABLE"] = 1] = "UNRELIABLE";
    HyperFlag2[HyperFlag2["RELIABLE"] = 2] = "RELIABLE";
    HyperFlag2[HyperFlag2["SPECULATIVE"] = 4] = "SPECULATIVE";
    HyperFlag2[HyperFlag2["EOS"] = 8] = "EOS";
  })(HyperFlag ||= {});

  class HyperFrameCodec {
    static MAGIC_0 = 72;
    static MAGIC_1 = 89;
    static VERSION = 1;
    static HEADER_SIZE = 16;
    static encode(frame) {
      const payloadLen = frame.payload.byteLength;
      const buf = new Uint8Array(this.HEADER_SIZE + payloadLen);
      const view = new DataView(buf.buffer, buf.byteOffset, buf.byteLength);
      buf[0] = this.MAGIC_0;
      buf[1] = this.MAGIC_1;
      buf[2] = this.VERSION;
      buf[3] = frame.plane;
      buf[4] = frame.flags;
      buf[5] = 0;
      view.setUint16(6, frame.channelId, false);
      view.setUint32(8, frame.sequence, false);
      view.setUint32(12, payloadLen, false);
      buf.set(frame.payload, this.HEADER_SIZE);
      return buf;
    }
    static decode(buffer) {
      const bytes = buffer instanceof Uint8Array ? buffer : new Uint8Array(buffer);
      if (bytes.byteLength < this.HEADER_SIZE) {
        throw new Error(`[HyperFrame] Corrupt frame: buffer size ${bytes.byteLength} < header size ${this.HEADER_SIZE}`);
      }
      if (bytes[0] !== this.MAGIC_0 || bytes[1] !== this.MAGIC_1) {
        throw new Error(`[HyperFrame] Invalid magic bytes: expected 'HY', got 0x${bytes[0].toString(16)} 0x${bytes[1].toString(16)}`);
      }
      const version = bytes[2];
      if (version !== this.VERSION) {
        throw new Error(`[HyperFrame] Unsupported version: ${version}`);
      }
      const view = new DataView(bytes.buffer, bytes.byteOffset, bytes.byteLength);
      const plane = bytes[3];
      const flags = bytes[4];
      const channelId = view.getUint16(6, false);
      const sequence = view.getUint32(8, false);
      const payloadLen = view.getUint32(12, false);
      if (bytes.byteLength < this.HEADER_SIZE + payloadLen) {
        throw new Error(`[HyperFrame] Incomplete payload: expected ${payloadLen} bytes, got ${bytes.byteLength - this.HEADER_SIZE}`);
      }
      const payload = new Uint8Array(bytes.buffer, bytes.byteOffset + this.HEADER_SIZE, payloadLen);
      return {
        plane,
        flags,
        channelId,
        sequence,
        timestamp: Date.now(),
        payload
      };
    }
    static encodeTensor(channelId, token, isEos = false, sequence = 0) {
      const enc = new TextEncoder;
      const payload = enc.encode(token);
      return this.encode({
        plane: 0 /* TENSOR */,
        flags: isEos ? 8 /* EOS */ : 2 /* RELIABLE */,
        channelId,
        sequence,
        timestamp: Date.now(),
        payload
      });
    }
    static decodeTensor(payload) {
      const dec = new TextDecoder;
      return { token: dec.decode(payload) };
    }
    static encodeHypermedia(channelId, target, swap, html, sequence = 0) {
      const enc = new TextEncoder;
      const targetBytes = enc.encode(target);
      const htmlBytes = enc.encode(html);
      const swapByte = this.swapStrategyToByte(swap);
      const payload = new Uint8Array(1 + targetBytes.byteLength + 1 + htmlBytes.byteLength);
      payload[0] = targetBytes.byteLength;
      payload.set(targetBytes, 1);
      payload[1 + targetBytes.byteLength] = swapByte;
      payload.set(htmlBytes, 2 + targetBytes.byteLength);
      return this.encode({
        plane: 1 /* HYPERMEDIA */,
        flags: 2 /* RELIABLE */,
        channelId,
        sequence,
        timestamp: Date.now(),
        payload
      });
    }
    static decodeHypermedia(payload) {
      if (payload.byteLength < 2) {
        return { target: "", swap: "outerHTML", html: "" };
      }
      const dec = new TextDecoder;
      const targetLen = payload[0];
      const target = dec.decode(payload.subarray(1, 1 + targetLen));
      const swapByte = payload[1 + targetLen];
      const swap = this.byteToSwapStrategy(swapByte);
      const html = dec.decode(payload.subarray(2 + targetLen));
      return { target, swap, html };
    }
    static encodeTelemetry(channelId, floats, sequence = 0) {
      const floatArr = floats instanceof Float32Array ? floats : new Float32Array(floats);
      const payload = new Uint8Array(floatArr.buffer, floatArr.byteOffset, floatArr.byteLength);
      return this.encode({
        plane: 2 /* TELEMETRY */,
        flags: 1 /* UNRELIABLE */,
        channelId,
        sequence,
        timestamp: Date.now(),
        payload
      });
    }
    static decodeTelemetry(payload) {
      if (payload.byteOffset % 4 === 0) {
        return new Float32Array(payload.buffer, payload.byteOffset, Math.floor(payload.byteLength / 4));
      }
      const clone = new Uint8Array(payload.byteLength);
      clone.set(payload);
      return new Float32Array(clone.buffer, 0, Math.floor(clone.byteLength / 4));
    }
    static encodePrecognition(channelId, key, confidence, html, sequence = 0) {
      const enc = new TextEncoder;
      const keyBytes = enc.encode(key);
      const htmlBytes = enc.encode(html);
      const payload = new Uint8Array(1 + keyBytes.byteLength + 4 + htmlBytes.byteLength);
      payload[0] = keyBytes.byteLength;
      payload.set(keyBytes, 1);
      const view = new DataView(payload.buffer, payload.byteOffset, payload.byteLength);
      view.setFloat32(1 + keyBytes.byteLength, confidence, false);
      payload.set(htmlBytes, 5 + keyBytes.byteLength);
      return this.encode({
        plane: 3 /* PRECOGNITION */,
        flags: 4 /* SPECULATIVE */,
        channelId,
        sequence,
        timestamp: Date.now(),
        payload
      });
    }
    static decodePrecognition(payload) {
      if (payload.byteLength < 5)
        return { key: "", confidence: 0, html: "" };
      const dec = new TextDecoder;
      const keyLen = payload[0];
      const key = dec.decode(payload.subarray(1, 1 + keyLen));
      const view = new DataView(payload.buffer, payload.byteOffset, payload.byteLength);
      const confidence = view.getFloat32(1 + keyLen, false);
      const html = dec.decode(payload.subarray(5 + keyLen));
      return { key, confidence, html };
    }
    static encodeSwarm(channelId, nodeId, vectorClock, data, sequence = 0) {
      const enc = new TextEncoder;
      const nodeBytes = enc.encode(nodeId);
      const dataBytes = typeof data === "string" ? enc.encode(data) : data;
      const payload = new Uint8Array(1 + nodeBytes.byteLength + 4 + dataBytes.byteLength);
      payload[0] = nodeBytes.byteLength;
      payload.set(nodeBytes, 1);
      const view = new DataView(payload.buffer, payload.byteOffset, payload.byteLength);
      view.setUint32(1 + nodeBytes.byteLength, vectorClock, false);
      payload.set(dataBytes, 5 + nodeBytes.byteLength);
      return this.encode({
        plane: 4 /* SWARM */,
        flags: 2 /* RELIABLE */,
        channelId,
        sequence,
        timestamp: Date.now(),
        payload
      });
    }
    static decodeSwarm(payload) {
      if (payload.byteLength < 5)
        return { nodeId: "", vectorClock: 0, data: "" };
      const dec = new TextDecoder;
      const nodeLen = payload[0];
      const nodeId = dec.decode(payload.subarray(1, 1 + nodeLen));
      const view = new DataView(payload.buffer, payload.byteOffset, payload.byteLength);
      const vectorClock = view.getUint32(1 + nodeLen, false);
      const data = dec.decode(payload.subarray(5 + nodeLen));
      return { nodeId, vectorClock, data };
    }
    static swapStrategyToByte(strategy) {
      switch (strategy) {
        case "outerHTML":
          return 1;
        case "innerHTML":
          return 2;
        case "afterbegin":
          return 3;
        case "beforeend":
          return 4;
        case "delete":
          return 5;
        case "none":
          return 6;
        default:
          return 1;
      }
    }
    static byteToSwapStrategy(b) {
      switch (b) {
        case 1:
          return "outerHTML";
        case 2:
          return "innerHTML";
        case 3:
          return "afterbegin";
        case 4:
          return "beforeend";
        case 5:
          return "delete";
        case 6:
          return "none";
        default:
          return "outerHTML";
      }
    }
  }

  class HyperHookEngine {
    subscriptions = new Map;
    eventSource = null;
    broadcastChannel = null;
    gatewayUrl = "/hy-hook/stream";
    reconnectAttempts = 0;
    maxReconnectDelay = 15000;
    reconnectTimer = null;
    isConnecting = false;
    eventsReceived = 0;
    eventsEmitted = 0;
    binaryFramesReceived = 0;
    lastEvent;
    subCounter = 0;
    speculativeCache = new Map;
    swarmNodes = new Map;
    telemetryRenderers = new Map;
    neuralStreams = new Map;
    constructor() {
      if (typeof window !== "undefined") {
        if ("BroadcastChannel" in window) {
          try {
            this.broadcastChannel = new BroadcastChannel("hy-hook-mesh");
            this.broadcastChannel.onmessage = (e) => this.handleBroadcastMessage(e.data);
          } catch {
            this.broadcastChannel = null;
          }
        }
        if (window.HY_HOOK_GATEWAY) {
          this.gatewayUrl = window.HY_HOOK_GATEWAY;
        }
      }
    }
    setGateway(url) {
      if (this.gatewayUrl === url)
        return;
      this.gatewayUrl = url;
      if (this.eventSource) {
        this.disconnect();
        this.connect();
      }
    }
    connect() {
      if (typeof window === "undefined" || typeof EventSource === "undefined")
        return;
      if (this.eventSource || this.isConnecting)
        return;
      this.isConnecting = true;
      const channels = new Set;
      const topics = new Set;
      this.subscriptions.forEach((sub) => {
        channels.add(sub.channel);
        if (sub.topic)
          topics.add(sub.topic);
      });
      const query = new URLSearchParams;
      if (channels.size > 0)
        query.set("channels", Array.from(channels).join(","));
      if (topics.size > 0)
        query.set("topics", Array.from(topics).join(","));
      const fullUrl = this.gatewayUrl + (query.toString() ? `?${query.toString()}` : "");
      try {
        this.eventSource = new EventSource(fullUrl);
        this.eventSource.onopen = () => {
          this.isConnecting = false;
          this.reconnectAttempts = 0;
          this.dispatchDiagEvent("connected");
        };
        this.eventSource.onmessage = (e) => {
          try {
            const payload = JSON.parse(e.data);
            this.routeEvent({
              channel: payload.channel || "default",
              event: payload.event || "message",
              topic: payload.topic,
              html: payload.html,
              data: payload.data,
              timestamp: payload.timestamp || Date.now(),
              origin: "server"
            });
          } catch (err) {
            this.routeRawHtml(e.data);
          }
        };
        this.eventSource.onerror = () => {
          this.isConnecting = false;
          this.disconnect();
          this.scheduleReconnect();
        };
      } catch {
        this.isConnecting = false;
        this.scheduleReconnect();
      }
    }
    disconnect() {
      if (this.eventSource) {
        this.eventSource.close();
        this.eventSource = null;
      }
      if (this.reconnectTimer) {
        clearTimeout(this.reconnectTimer);
        this.reconnectTimer = null;
      }
    }
    scheduleReconnect() {
      this.reconnectAttempts++;
      const delay = Math.min(this.maxReconnectDelay, 1000 * Math.pow(1.5, this.reconnectAttempts));
      this.reconnectTimer = setTimeout(() => {
        this.connect();
      }, delay);
    }
    subscribe(opts) {
      const id = `sub_${++this.subCounter}_${Date.now()}`;
      const sub = {
        id,
        channel: opts.channel,
        event: opts.event,
        topic: opts.topic,
        element: opts.element,
        swap: opts.swap || "outerHTML",
        target: opts.target,
        callback: opts.callback
      };
      this.subscriptions.set(id, sub);
      if (typeof window !== "undefined" && this.subscriptions.size > 0 && !this.eventSource) {
        this.connect();
      }
      return id;
    }
    unsubscribe(id) {
      const removed = this.subscriptions.delete(id);
      if (this.subscriptions.size === 0) {
        this.disconnect();
      }
      return removed;
    }
    emit(channel, event, opts = {}) {
      this.eventsEmitted++;
      const hookEvent = {
        channel,
        event,
        topic: opts.topic,
        html: opts.html,
        data: opts.data,
        timestamp: Date.now(),
        origin: "local"
      };
      this.routeEvent(hookEvent);
      if (opts.broadcast !== false && this.broadcastChannel) {
        try {
          this.broadcastChannel.postMessage({
            ...hookEvent,
            origin: "broadcast"
          });
        } catch {}
      }
      if (typeof window !== "undefined") {
        window.dispatchEvent(new CustomEvent("hy-hook:emitted", { detail: hookEvent }));
      }
      return hookEvent;
    }
    routeEvent(evt) {
      this.eventsReceived++;
      this.lastEvent = evt;
      let matchCount = 0;
      if (evt.html && typeof document !== "undefined") {
        this.processOOB(evt.html);
      }
      this.subscriptions.forEach((sub) => {
        const channelMatches = sub.channel === "*" || sub.channel === evt.channel;
        const eventMatches = sub.event === "*" || sub.event === evt.event;
        const topicMatches = !sub.topic || sub.topic === "*" || sub.topic === evt.topic;
        if (channelMatches && eventMatches && topicMatches) {
          matchCount++;
          if (sub.callback) {
            sub.callback(evt);
          }
          if (evt.html && typeof document !== "undefined") {
            let targetEl = sub.element;
            if (sub.target) {
              const queried = document.querySelector(sub.target);
              if (queried)
                targetEl = queried;
            }
            if (targetEl) {
              this.applySwap(targetEl, evt.html, sub.swap);
            }
          }
        }
      });
      if (typeof window !== "undefined") {
        window.dispatchEvent(new CustomEvent("hy-hook:received", { detail: { event: evt, matchCount } }));
      }
      return matchCount;
    }
    routeBinaryFrame(buffer) {
      const frame = HyperFrameCodec.decode(buffer);
      this.binaryFramesReceived++;
      let handled = false;
      switch (frame.plane) {
        case 0 /* TENSOR */: {
          const { token } = HyperFrameCodec.decodeTensor(frame.payload);
          const isEos = (frame.flags & 8 /* EOS */) !== 0;
          const listener = this.neuralStreams.get(frame.channelId);
          if (listener) {
            listener(token, isEos);
            handled = true;
          }
          if (typeof document !== "undefined") {
            const targets = document.querySelectorAll(`[hy-ai-stream="${frame.channelId}"], [hy-ai-stream="#${frame.channelId}"]`);
            targets.forEach((el) => {
              el.insertAdjacentText("beforeend", token);
              handled = true;
            });
            window.dispatchEvent(new CustomEvent("hy-hook:tensor", {
              detail: { channelId: frame.channelId, token, isEos }
            }));
          }
          break;
        }
        case 1 /* HYPERMEDIA */: {
          const { target, swap, html } = HyperFrameCodec.decodeHypermedia(frame.payload);
          if (typeof document !== "undefined" && target) {
            const el = document.querySelector(target);
            if (el) {
              this.applySwap(el, html, swap);
              handled = true;
            }
            window.dispatchEvent(new CustomEvent("hy-hook:hypermedia", {
              detail: { channelId: frame.channelId, target, swap, html }
            }));
          }
          break;
        }
        case 2 /* TELEMETRY */: {
          const floats = HyperFrameCodec.decodeTelemetry(frame.payload);
          const channelKey = String(frame.channelId);
          const renderer = this.telemetryRenderers.get(channelKey);
          if (renderer) {
            renderer(floats);
            handled = true;
          }
          if (typeof window !== "undefined") {
            window.dispatchEvent(new CustomEvent("hy-hook:telemetry", {
              detail: { channelId: frame.channelId, floats }
            }));
          }
          break;
        }
        case 3 /* PRECOGNITION */: {
          const { key, confidence, html } = HyperFrameCodec.decodePrecognition(frame.payload);
          this.cacheSpeculation(key, confidence, html);
          handled = true;
          if (typeof window !== "undefined") {
            window.dispatchEvent(new CustomEvent("hy-hook:speculative-cached", {
              detail: { channelId: frame.channelId, key, confidence, html }
            }));
          }
          break;
        }
        case 4 /* SWARM */: {
          const { nodeId, vectorClock, data } = HyperFrameCodec.decodeSwarm(frame.payload);
          this.swarmNodes.set(nodeId, {
            nodeId,
            lastSeen: Date.now(),
            vectorClock
          });
          handled = true;
          if (typeof window !== "undefined") {
            window.dispatchEvent(new CustomEvent("hy-hook:swarm-sync", {
              detail: { channelId: frame.channelId, nodeId, vectorClock, data }
            }));
          }
          break;
        }
      }
      return { frame, handled };
    }
    onNeuralStream(channelId, callback) {
      this.neuralStreams.set(channelId, callback);
      return () => this.neuralStreams.delete(channelId);
    }
    registerTelemetryRenderer(channelId, renderer) {
      const key = String(channelId);
      this.telemetryRenderers.set(key, renderer);
      return () => this.telemetryRenderers.delete(key);
    }
    cacheSpeculation(key, confidence, html) {
      this.speculativeCache.set(key, {
        key,
        confidence,
        html,
        timestamp: Date.now()
      });
    }
    commitSpeculation(key, targetSelector, swap = "outerHTML") {
      const cached = this.speculativeCache.get(key);
      if (!cached)
        return false;
      if (typeof document !== "undefined") {
        const selector = targetSelector || `#${key}`;
        const targetEl = document.querySelector(selector);
        if (targetEl) {
          this.applySwap(targetEl, cached.html, swap);
        }
        window.dispatchEvent(new CustomEvent("hy-hook:speculative-committed", {
          detail: { key, cached, targetSelector }
        }));
      }
      this.speculativeCache.delete(key);
      return true;
    }
    getSpeculation(key) {
      return this.speculativeCache.get(key);
    }
    getSwarmNodes() {
      return Array.from(this.swarmNodes.values());
    }
    handleBroadcastMessage(data) {
      if (!data || !data.channel || !data.event)
        return;
      this.routeEvent({
        channel: data.channel,
        event: data.event,
        topic: data.topic,
        html: data.html,
        data: data.data,
        timestamp: data.timestamp || Date.now(),
        origin: "broadcast"
      });
    }
    routeRawHtml(html) {
      if (typeof document === "undefined")
        return;
      const temp = document.createElement("div");
      temp.innerHTML = html.trim();
      const firstEl = temp.firstElementChild;
      if (firstEl) {
        const channel = firstEl.getAttribute("hy-channel") || "*";
        const event = firstEl.getAttribute("hy-event") || "*";
        const topic = firstEl.getAttribute("hy-topic") || undefined;
        this.routeEvent({
          channel,
          event,
          topic,
          html,
          timestamp: Date.now(),
          origin: "server"
        });
      }
    }
    applySwap(target, html, strategy) {
      if (!target || !html && strategy !== "delete")
        return;
      const beforeEvt = new CustomEvent("hy-hook:before-swap", {
        bubbles: true,
        cancelable: true,
        detail: { target, html, strategy }
      });
      target.dispatchEvent(beforeEvt);
      if (beforeEvt.defaultPrevented)
        return;
      switch (strategy) {
        case "innerHTML":
          target.innerHTML = html;
          break;
        case "outerHTML":
          target.outerHTML = html;
          break;
        case "afterbegin":
          target.insertAdjacentHTML("afterbegin", html);
          break;
        case "beforeend":
          target.insertAdjacentHTML("beforeend", html);
          break;
        case "delete":
          target.remove();
          break;
        case "none":
          break;
      }
      if (typeof document !== "undefined") {
        document.dispatchEvent(new CustomEvent("hy-hook:after-swap", {
          detail: { target, strategy }
        }));
      }
    }
    processOOB(html) {
      if (typeof document === "undefined")
        return;
      const parser = document.createElement("div");
      parser.innerHTML = html.trim();
      const oobElements = parser.querySelectorAll("[hy-swap-oob], [hx-swap-oob]");
      oobElements.forEach((oobEl) => {
        const id = oobEl.getAttribute("id");
        if (id) {
          const existing = document.getElementById(id);
          if (existing) {
            const mode = oobEl.getAttribute("hy-swap-oob") || oobEl.getAttribute("hx-swap-oob");
            const strategy = mode === "true" || mode === "outerHTML" ? "outerHTML" : mode || "outerHTML";
            this.applySwap(existing, oobEl.outerHTML, strategy);
          }
        }
      });
    }
    scan(root = document) {
      if (typeof document === "undefined")
        return;
      root.querySelectorAll("[hy-hook]").forEach((el) => {
        if (el.hasAttribute("data-hy-bound"))
          return;
        el.setAttribute("data-hy-bound", "true");
        const rawHook = el.getAttribute("hy-hook") || "";
        const [channel, event] = rawHook.split(":");
        const topic = el.getAttribute("hy-topic") || undefined;
        const swap = el.getAttribute("hy-swap") || "outerHTML";
        const target = el.getAttribute("hy-target") || undefined;
        if (channel && event) {
          this.subscribe({
            channel,
            event,
            topic,
            element: el,
            swap,
            target
          });
        }
      });
      root.querySelectorAll("[hy-hook-listen]").forEach((el) => {
        if (el.hasAttribute("data-hy-listen-bound"))
          return;
        el.setAttribute("data-hy-listen-bound", "true");
        const rawHook = el.getAttribute("hy-hook-listen") || "";
        const [channel, event] = rawHook.split(":");
        const topic = el.getAttribute("hy-topic") || undefined;
        const swap = el.getAttribute("hy-swap") || "afterbegin";
        const target = el.getAttribute("hy-target") || undefined;
        if (channel && event) {
          this.subscribe({
            channel,
            event,
            topic,
            element: el,
            swap,
            target
          });
        }
      });
      root.querySelectorAll("[hy-hook-emit]").forEach((el) => {
        if (el.hasAttribute("data-hy-emit-bound"))
          return;
        el.setAttribute("data-hy-emit-bound", "true");
        const triggerEvent = el.tagName === "FORM" ? "submit" : "click";
        el.addEventListener(triggerEvent, () => {
          const rawHook = el.getAttribute("hy-hook-emit") || "";
          const [channel, event] = rawHook.split(":");
          const topic = el.getAttribute("hy-topic") || undefined;
          const rawData = el.getAttribute("hy-data");
          const rawHtml = el.getAttribute("hy-html");
          let data;
          if (rawData) {
            try {
              data = JSON.parse(rawData);
            } catch {}
          }
          if (channel && event) {
            this.emit(channel, event, {
              topic,
              data,
              html: rawHtml || undefined,
              broadcast: true
            });
          }
        });
      });
      root.querySelectorAll("[hy-speculate-key]").forEach((el) => {
        if (el.hasAttribute("data-hy-speculate-bound"))
          return;
        el.setAttribute("data-hy-speculate-bound", "true");
        const key = el.getAttribute("hy-speculate-key") || "";
        const trigger = el.getAttribute("hy-speculate-trigger") || "mouseenter";
        const target = el.getAttribute("hy-target") || undefined;
        const swap = el.getAttribute("hy-swap") || "outerHTML";
        el.addEventListener(trigger, () => {
          if (key) {
            this.commitSpeculation(key, target, swap);
          }
        });
      });
    }
    diag() {
      return {
        connected: this.eventSource !== null && this.eventSource.readyState === EventSource.OPEN,
        gatewayUrl: this.gatewayUrl,
        activeSubscriptionsCount: this.subscriptions.size,
        eventsReceivedCount: this.eventsReceived,
        eventsEmittedCount: this.eventsEmitted,
        binaryFramesReceivedCount: this.binaryFramesReceived,
        speculativeCacheCount: this.speculativeCache.size,
        swarmNodesCount: this.swarmNodes.size,
        broadcastChannelAvailable: this.broadcastChannel !== null,
        lastEvent: this.lastEvent
      };
    }
    dispatchDiagEvent(status) {
      if (typeof window !== "undefined") {
        window.dispatchEvent(new CustomEvent("hy-hook:diag", {
          detail: { status, diag: this.diag() }
        }));
      }
    }
  }
  var HyHook = new HyperHookEngine;
  if (typeof document !== "undefined") {
    if (document.readyState === "loading") {
      document.addEventListener("DOMContentLoaded", () => HyHook.scan());
    } else {
      HyHook.scan();
    }
    if ("MutationObserver" in window) {
      const observer = new MutationObserver((mutations) => {
        for (const mutation of mutations) {
          mutation.addedNodes.forEach((node) => {
            if (node.nodeType === Node.ELEMENT_NODE) {
              HyHook.scan(node);
            }
          });
        }
      });
      observer.observe(document.body, { childList: true, subtree: true });
    }
  }
  if (typeof window !== "undefined") {
    window.HyHook = HyHook;
    window.HyperPlane = HyperPlane;
    window.HyperFlag = HyperFlag;
    window.HyperFrameCodec = HyperFrameCodec;
  }
})();
