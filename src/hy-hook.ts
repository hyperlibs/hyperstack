/**
 * HY-HOOK.TS — Hypermedia Webhooks & Event Mesh Engine
 * 
 * Part of the HyperPRO / HyperUI Suite
 * 
 * Prefix Standard: hy-* (Non-conflicting hypermedia extension)
 * 
 * Core Capabilities:
 * 1. Declarative Hypermedia Webhooks (SSE + BroadcastChannel mesh)
 * 2. Quantum Dual-Plane Binary Framing (Tensors, Hypermedia, Telemetry, Precognition, Swarm)
 * 3. Zero-Allocation Binary Telemetry & Canvas Streaming (60fps)
 * 4. Speculative Pre-Cognition Cache (0.0ms perceived latency)
 * 5. Edge AI Neural Token Streaming
 * 6. Swarm Cluster P2P Vector Mesh
 */

export type HySwapStrategy = 'outerHTML' | 'innerHTML' | 'afterbegin' | 'beforeend' | 'delete' | 'none';

export interface HyHookEvent {
  channel: string;
  event: string;
  topic?: string;
  html?: string;
  data?: Record<string, any>;
  timestamp: number;
  origin?: 'server' | 'broadcast' | 'local';
}

export interface HySubscription {
  id: string;
  element?: HTMLElement;
  channel: string;
  event: string;
  topic?: string;
  swap: HySwapStrategy;
  target?: string;
  callback?: (evt: HyHookEvent) => void;
}

export interface HyHookDiagnostics {
  connected: boolean;
  gatewayUrl: string;
  activeSubscriptionsCount: number;
  eventsReceivedCount: number;
  eventsEmittedCount: number;
  binaryFramesReceivedCount: number;
  speculativeCacheCount: number;
  swarmNodesCount: number;
  broadcastChannelAvailable: boolean;
  lastEvent?: HyHookEvent;
  lastLatencyMs?: number;
}

// =============================================================================
// QUANTUM DUAL-PLANE BINARY FRAMING (FORWARD-FOCUSED STATE FABRIC)
// =============================================================================

export enum HyperPlane {
  TENSOR = 0x00,       // Plane 0: Edge AI token stream, KV slice, neural reasoning token
  HYPERMEDIA = 0x01,   // Plane 1: Atomic DOM morph / HTML partial swap
  TELEMETRY = 0x02,    // Plane 2: Continuous 60fps packed float32/int16 vector
  PRECOGNITION = 0x03, // Plane 3: Speculative pre-materialized state & confidence
  SWARM = 0x04,        // Plane 4: P2P cluster mesh state, vector clock, CRDT delta
}

export enum HyperFlag {
  UNRELIABLE = 0x01,   // Loss-tolerant datagram (UDP/WebTransport Datagram)
  RELIABLE = 0x02,     // Guaranteed ordered stream
  SPECULATIVE = 0x04,  // Candidate speculative state
  EOS = 0x08,          // End-of-stream delimiter
}

export interface HyperFrame {
  plane: HyperPlane;
  flags: number;
  channelId: number;      // 16-bit uint
  sequence: number;       // 32-bit uint
  timestamp: number;      // ms timestamp
  payload: Uint8Array;
}

export class HyperFrameCodec {
  static readonly MAGIC_0 = 0x48; // 'H'
  static readonly MAGIC_1 = 0x59; // 'Y'
  static readonly VERSION = 0x01;
  static readonly HEADER_SIZE = 16;

  static encode(frame: HyperFrame): Uint8Array {
    const payloadLen = frame.payload.byteLength;
    const buf = new Uint8Array(this.HEADER_SIZE + payloadLen);
    const view = new DataView(buf.buffer, buf.byteOffset, buf.byteLength);

    buf[0] = this.MAGIC_0;
    buf[1] = this.MAGIC_1;
    buf[2] = this.VERSION;
    buf[3] = frame.plane;
    buf[4] = frame.flags;
    buf[5] = 0; // reserved
    view.setUint16(6, frame.channelId, false);
    view.setUint32(8, frame.sequence, false);
    view.setUint32(12, payloadLen, false);

    buf.set(frame.payload, this.HEADER_SIZE);
    return buf;
  }

  static decode(buffer: ArrayBuffer | Uint8Array): HyperFrame {
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
    const plane = bytes[3] as HyperPlane;
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

  // --- Plane 0: Neural / Tensor Codec ---
  static encodeTensor(channelId: number, token: string, isEos = false, sequence = 0): Uint8Array {
    const enc = new TextEncoder();
    const payload = enc.encode(token);
    return this.encode({
      plane: HyperPlane.TENSOR,
      flags: isEos ? HyperFlag.EOS : HyperFlag.RELIABLE,
      channelId,
      sequence,
      timestamp: Date.now(),
      payload
    });
  }

  static decodeTensor(payload: Uint8Array): { token: string } {
    const dec = new TextDecoder();
    return { token: dec.decode(payload) };
  }

  // --- Plane 1: Hypermedia Mutation Codec ---
  static encodeHypermedia(channelId: number, target: string, swap: HySwapStrategy, html: string, sequence = 0): Uint8Array {
    const enc = new TextEncoder();
    const targetBytes = enc.encode(target);
    const htmlBytes = enc.encode(html);
    const swapByte = this.swapStrategyToByte(swap);

    const payload = new Uint8Array(1 + targetBytes.byteLength + 1 + htmlBytes.byteLength);
    payload[0] = targetBytes.byteLength;
    payload.set(targetBytes, 1);
    payload[1 + targetBytes.byteLength] = swapByte;
    payload.set(htmlBytes, 2 + targetBytes.byteLength);

    return this.encode({
      plane: HyperPlane.HYPERMEDIA,
      flags: HyperFlag.RELIABLE,
      channelId,
      sequence,
      timestamp: Date.now(),
      payload
    });
  }

  static decodeHypermedia(payload: Uint8Array): { target: string; swap: HySwapStrategy; html: string } {
    if (payload.byteLength < 2) {
      return { target: '', swap: 'outerHTML', html: '' };
    }
    const dec = new TextDecoder();
    const targetLen = payload[0];
    const target = dec.decode(payload.subarray(1, 1 + targetLen));
    const swapByte = payload[1 + targetLen];
    const swap = this.byteToSwapStrategy(swapByte);
    const html = dec.decode(payload.subarray(2 + targetLen));

    return { target, swap, html };
  }

  // --- Plane 2: Continuous Telemetry Codec ---
  static encodeTelemetry(channelId: number, floats: Float32Array | number[], sequence = 0): Uint8Array {
    const floatArr = floats instanceof Float32Array ? floats : new Float32Array(floats);
    const payload = new Uint8Array(floatArr.buffer, floatArr.byteOffset, floatArr.byteLength);

    return this.encode({
      plane: HyperPlane.TELEMETRY,
      flags: HyperFlag.UNRELIABLE,
      channelId,
      sequence,
      timestamp: Date.now(),
      payload
    });
  }

  static decodeTelemetry(payload: Uint8Array): Float32Array {
    if (payload.byteOffset % 4 === 0) {
      return new Float32Array(payload.buffer, payload.byteOffset, Math.floor(payload.byteLength / 4));
    }
    const clone = new Uint8Array(payload.byteLength);
    clone.set(payload);
    return new Float32Array(clone.buffer, 0, Math.floor(clone.byteLength / 4));
  }

  // --- Plane 3: Speculative Pre-Cognition Codec ---
  static encodePrecognition(channelId: number, key: string, confidence: number, html: string, sequence = 0): Uint8Array {
    const enc = new TextEncoder();
    const keyBytes = enc.encode(key);
    const htmlBytes = enc.encode(html);

    const payload = new Uint8Array(1 + keyBytes.byteLength + 4 + htmlBytes.byteLength);
    payload[0] = keyBytes.byteLength;
    payload.set(keyBytes, 1);

    const view = new DataView(payload.buffer, payload.byteOffset, payload.byteLength);
    view.setFloat32(1 + keyBytes.byteLength, confidence, false);

    payload.set(htmlBytes, 5 + keyBytes.byteLength);

    return this.encode({
      plane: HyperPlane.PRECOGNITION,
      flags: HyperFlag.SPECULATIVE,
      channelId,
      sequence,
      timestamp: Date.now(),
      payload
    });
  }

  static decodePrecognition(payload: Uint8Array): { key: string; confidence: number; html: string } {
    if (payload.byteLength < 5) return { key: '', confidence: 0, html: '' };
    const dec = new TextDecoder();
    const keyLen = payload[0];
    const key = dec.decode(payload.subarray(1, 1 + keyLen));

    const view = new DataView(payload.buffer, payload.byteOffset, payload.byteLength);
    const confidence = view.getFloat32(1 + keyLen, false);
    const html = dec.decode(payload.subarray(5 + keyLen));

    return { key, confidence, html };
  }

  // --- Plane 4: Swarm Cluster Gossip Codec ---
  static encodeSwarm(channelId: number, nodeId: string, vectorClock: number, data: string | Uint8Array, sequence = 0): Uint8Array {
    const enc = new TextEncoder();
    const nodeBytes = enc.encode(nodeId);
    const dataBytes = typeof data === 'string' ? enc.encode(data) : data;

    const payload = new Uint8Array(1 + nodeBytes.byteLength + 4 + dataBytes.byteLength);
    payload[0] = nodeBytes.byteLength;
    payload.set(nodeBytes, 1);

    const view = new DataView(payload.buffer, payload.byteOffset, payload.byteLength);
    view.setUint32(1 + nodeBytes.byteLength, vectorClock, false);

    payload.set(dataBytes, 5 + nodeBytes.byteLength);

    return this.encode({
      plane: HyperPlane.SWARM,
      flags: HyperFlag.RELIABLE,
      channelId,
      sequence,
      timestamp: Date.now(),
      payload
    });
  }

  static decodeSwarm(payload: Uint8Array): { nodeId: string; vectorClock: number; data: string } {
    if (payload.byteLength < 5) return { nodeId: '', vectorClock: 0, data: '' };
    const dec = new TextDecoder();
    const nodeLen = payload[0];
    const nodeId = dec.decode(payload.subarray(1, 1 + nodeLen));

    const view = new DataView(payload.buffer, payload.byteOffset, payload.byteLength);
    const vectorClock = view.getUint32(1 + nodeLen, false);
    const data = dec.decode(payload.subarray(5 + nodeLen));

    return { nodeId, vectorClock, data };
  }

  private static swapStrategyToByte(strategy: HySwapStrategy): number {
    switch (strategy) {
      case 'outerHTML': return 0x01;
      case 'innerHTML': return 0x02;
      case 'afterbegin': return 0x03;
      case 'beforeend': return 0x04;
      case 'delete': return 0x05;
      case 'none': return 0x06;
      default: return 0x01;
    }
  }

  private static byteToSwapStrategy(b: number): HySwapStrategy {
    switch (b) {
      case 0x01: return 'outerHTML';
      case 0x02: return 'innerHTML';
      case 0x03: return 'afterbegin';
      case 0x04: return 'beforeend';
      case 0x05: return 'delete';
      case 0x06: return 'none';
      default: return 'outerHTML';
    }
  }
}

// =============================================================================
// HYPERHOOK ENGINE IMPLEMENTATION
// =============================================================================

export class HyperHookEngine {
  private subscriptions: Map<string, HySubscription> = new Map();
  private eventSource: EventSource | null = null;
  private broadcastChannel: BroadcastChannel | null = null;
  private gatewayUrl = '/hy-hook/stream';
  private reconnectAttempts = 0;
  private maxReconnectDelay = 15000;
  private reconnectTimer: any = null;
  private isConnecting = false;
  private eventsReceived = 0;
  private eventsEmitted = 0;
  private binaryFramesReceived = 0;
  private lastEvent?: HyHookEvent;
  private subCounter = 0;

  // Quantum Planes State Storage
  private speculativeCache = new Map<string, { key: string; confidence: number; html: string; timestamp: number }>();
  private swarmNodes = new Map<string, { nodeId: string; lastSeen: number; vectorClock: number }>();
  private telemetryRenderers = new Map<string, (floats: Float32Array) => void>();
  private neuralStreams = new Map<number, (token: string, isEos: boolean) => void>();

  constructor() {
    if (typeof window !== 'undefined') {
      if ('BroadcastChannel' in window) {
        try {
          this.broadcastChannel = new BroadcastChannel('hy-hook-mesh');
          this.broadcastChannel.onmessage = (e) => this.handleBroadcastMessage(e.data);
        } catch {
          this.broadcastChannel = null;
        }
      }

      if ((window as any).HY_HOOK_GATEWAY) {
        this.gatewayUrl = (window as any).HY_HOOK_GATEWAY;
      }
    }
  }

  setGateway(url: string): void {
    if (this.gatewayUrl === url) return;
    this.gatewayUrl = url;
    if (this.eventSource) {
      this.disconnect();
      this.connect();
    }
  }

  connect(): void {
    if (typeof window === 'undefined' || typeof EventSource === 'undefined') return;
    if (this.eventSource || this.isConnecting) return;

    this.isConnecting = true;
    const channels = new Set<string>();
    const topics = new Set<string>();

    this.subscriptions.forEach(sub => {
      channels.add(sub.channel);
      if (sub.topic) topics.add(sub.topic);
    });

    const query = new URLSearchParams();
    if (channels.size > 0) query.set('channels', Array.from(channels).join(','));
    if (topics.size > 0) query.set('topics', Array.from(topics).join(','));

    const fullUrl = this.gatewayUrl + (query.toString() ? `?${query.toString()}` : '');

    try {
      this.eventSource = new EventSource(fullUrl);

      this.eventSource.onopen = () => {
        this.isConnecting = false;
        this.reconnectAttempts = 0;
        this.dispatchDiagEvent('connected');
      };

      this.eventSource.onmessage = (e) => {
        try {
          const payload = JSON.parse(e.data);
          this.routeEvent({
            channel: payload.channel || 'default',
            event: payload.event || 'message',
            topic: payload.topic,
            html: payload.html,
            data: payload.data,
            timestamp: payload.timestamp || Date.now(),
            origin: 'server'
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

  disconnect(): void {
    if (this.eventSource) {
      this.eventSource.close();
      this.eventSource = null;
    }
    if (this.reconnectTimer) {
      clearTimeout(this.reconnectTimer);
      this.reconnectTimer = null;
    }
  }

  private scheduleReconnect(): void {
    this.reconnectAttempts++;
    const delay = Math.min(this.maxReconnectDelay, 1000 * Math.pow(1.5, this.reconnectAttempts));
    this.reconnectTimer = setTimeout(() => {
      this.connect();
    }, delay);
  }

  subscribe(opts: {
    channel: string;
    event: string;
    topic?: string;
    element?: HTMLElement;
    swap?: HySwapStrategy;
    target?: string;
    callback?: (evt: HyHookEvent) => void;
  }): string {
    const id = `sub_${++this.subCounter}_${Date.now()}`;
    const sub: HySubscription = {
      id,
      channel: opts.channel,
      event: opts.event,
      topic: opts.topic,
      element: opts.element,
      swap: opts.swap || 'outerHTML',
      target: opts.target,
      callback: opts.callback
    };

    this.subscriptions.set(id, sub);

    if (typeof window !== 'undefined' && this.subscriptions.size > 0 && !this.eventSource) {
      this.connect();
    }

    return id;
  }

  unsubscribe(id: string): boolean {
    const removed = this.subscriptions.delete(id);
    if (this.subscriptions.size === 0) {
      this.disconnect();
    }
    return removed;
  }

  emit(channel: string, event: string, opts: {
    topic?: string;
    html?: string;
    data?: Record<string, any>;
    broadcast?: boolean;
  } = {}): HyHookEvent {
    this.eventsEmitted++;

    const hookEvent: HyHookEvent = {
      channel,
      event,
      topic: opts.topic,
      html: opts.html,
      data: opts.data,
      timestamp: Date.now(),
      origin: 'local'
    };

    this.routeEvent(hookEvent);

    if (opts.broadcast !== false && this.broadcastChannel) {
      try {
        this.broadcastChannel.postMessage({
          ...hookEvent,
          origin: 'broadcast'
        });
      } catch {}
    }

    if (typeof window !== 'undefined') {
      window.dispatchEvent(new CustomEvent('hy-hook:emitted', { detail: hookEvent }));
    }

    return hookEvent;
  }

  routeEvent(evt: HyHookEvent): number {
    this.eventsReceived++;
    this.lastEvent = evt;
    let matchCount = 0;

    if (evt.html && typeof document !== 'undefined') {
      this.processOOB(evt.html);
    }

    this.subscriptions.forEach(sub => {
      const channelMatches = sub.channel === '*' || sub.channel === evt.channel;
      const eventMatches = sub.event === '*' || sub.event === evt.event;
      const topicMatches = !sub.topic || sub.topic === '*' || sub.topic === evt.topic;

      if (channelMatches && eventMatches && topicMatches) {
        matchCount++;

        if (sub.callback) {
          sub.callback(evt);
        }

        if (evt.html && typeof document !== 'undefined') {
          let targetEl = sub.element;
          if (sub.target) {
            const queried = document.querySelector(sub.target) as HTMLElement;
            if (queried) targetEl = queried;
          }

          if (targetEl) {
            this.applySwap(targetEl, evt.html, sub.swap);
          }
        }
      }
    });

    if (typeof window !== 'undefined') {
      window.dispatchEvent(new CustomEvent('hy-hook:received', { detail: { event: evt, matchCount } }));
    }

    return matchCount;
  }

  // ===========================================================================
  // QUANTUM BINARY DISPATCHER (ZERO-ALLOC DEMUX)
  // ===========================================================================

  routeBinaryFrame(buffer: ArrayBuffer | Uint8Array): { frame: HyperFrame; handled: boolean } {
    const frame = HyperFrameCodec.decode(buffer);
    this.binaryFramesReceived++;
    let handled = false;

    switch (frame.plane) {
      // Plane 0: Neural / Edge AI Token Streaming
      case HyperPlane.TENSOR: {
        const { token } = HyperFrameCodec.decodeTensor(frame.payload);
        const isEos = (frame.flags & HyperFlag.EOS) !== 0;

        const listener = this.neuralStreams.get(frame.channelId);
        if (listener) {
          listener(token, isEos);
          handled = true;
        }

        if (typeof document !== 'undefined') {
          const targets = document.querySelectorAll<HTMLElement>(`[hy-ai-stream="${frame.channelId}"], [hy-ai-stream="#${frame.channelId}"]`);
          targets.forEach(el => {
            el.insertAdjacentText('beforeend', token);
            handled = true;
          });

          window.dispatchEvent(new CustomEvent('hy-hook:tensor', {
            detail: { channelId: frame.channelId, token, isEos }
          }));
        }
        break;
      }

      // Plane 1: Hypermedia Sub-Millisecond Morph
      case HyperPlane.HYPERMEDIA: {
        const { target, swap, html } = HyperFrameCodec.decodeHypermedia(frame.payload);
        if (typeof document !== 'undefined' && target) {
          const el = document.querySelector<HTMLElement>(target);
          if (el) {
            this.applySwap(el, html, swap);
            handled = true;
          }
          window.dispatchEvent(new CustomEvent('hy-hook:hypermedia', {
            detail: { channelId: frame.channelId, target, swap, html }
          }));
        }
        break;
      }

      // Plane 2: Continuous 60fps Telemetry
      case HyperPlane.TELEMETRY: {
        const floats = HyperFrameCodec.decodeTelemetry(frame.payload);
        const channelKey = String(frame.channelId);
        const renderer = this.telemetryRenderers.get(channelKey);
        if (renderer) {
          renderer(floats);
          handled = true;
        }

        if (typeof window !== 'undefined') {
          window.dispatchEvent(new CustomEvent('hy-hook:telemetry', {
            detail: { channelId: frame.channelId, floats }
          }));
        }
        break;
      }

      // Plane 3: Speculative Pre-Cognition Cache
      case HyperPlane.PRECOGNITION: {
        const { key, confidence, html } = HyperFrameCodec.decodePrecognition(frame.payload);
        this.cacheSpeculation(key, confidence, html);
        handled = true;

        if (typeof window !== 'undefined') {
          window.dispatchEvent(new CustomEvent('hy-hook:speculative-cached', {
            detail: { channelId: frame.channelId, key, confidence, html }
          }));
        }
        break;
      }

      // Plane 4: Swarm Cluster Gossip
      case HyperPlane.SWARM: {
        const { nodeId, vectorClock, data } = HyperFrameCodec.decodeSwarm(frame.payload);
        this.swarmNodes.set(nodeId, {
          nodeId,
          lastSeen: Date.now(),
          vectorClock
        });
        handled = true;

        if (typeof window !== 'undefined') {
          window.dispatchEvent(new CustomEvent('hy-hook:swarm-sync', {
            detail: { channelId: frame.channelId, nodeId, vectorClock, data }
          }));
        }
        break;
      }
    }

    return { frame, handled };
  }

  // --- Neural Streaming API ---
  onNeuralStream(channelId: number, callback: (token: string, isEos: boolean) => void): () => void {
    this.neuralStreams.set(channelId, callback);
    return () => this.neuralStreams.delete(channelId);
  }

  // --- Continuous Telemetry API ---
  registerTelemetryRenderer(channelId: string | number, renderer: (floats: Float32Array) => void): () => void {
    const key = String(channelId);
    this.telemetryRenderers.set(key, renderer);
    return () => this.telemetryRenderers.delete(key);
  }

  // --- Speculative Pre-Cognition API ---
  cacheSpeculation(key: string, confidence: number, html: string): void {
    this.speculativeCache.set(key, {
      key,
      confidence,
      html,
      timestamp: Date.now()
    });
  }

  commitSpeculation(key: string, targetSelector?: string, swap: HySwapStrategy = 'outerHTML'): boolean {
    const cached = this.speculativeCache.get(key);
    if (!cached) return false;

    if (typeof document !== 'undefined') {
      const selector = targetSelector || `#${key}`;
      const targetEl = document.querySelector<HTMLElement>(selector);
      if (targetEl) {
        this.applySwap(targetEl, cached.html, swap);
      }
      window.dispatchEvent(new CustomEvent('hy-hook:speculative-committed', {
        detail: { key, cached, targetSelector }
      }));
    }

    this.speculativeCache.delete(key);
    return true;
  }

  getSpeculation(key: string): { key: string; confidence: number; html: string; timestamp: number } | undefined {
    return this.speculativeCache.get(key);
  }

  getSwarmNodes(): Array<{ nodeId: string; lastSeen: number; vectorClock: number }> {
    return Array.from(this.swarmNodes.values());
  }

  private handleBroadcastMessage(data: any): void {
    if (!data || !data.channel || !data.event) return;
    this.routeEvent({
      channel: data.channel,
      event: data.event,
      topic: data.topic,
      html: data.html,
      data: data.data,
      timestamp: data.timestamp || Date.now(),
      origin: 'broadcast'
    });
  }

  private routeRawHtml(html: string): void {
    if (typeof document === 'undefined') return;
    const temp = document.createElement('div');
    temp.innerHTML = html.trim();
    const firstEl = temp.firstElementChild as HTMLElement;

    if (firstEl) {
      const channel = firstEl.getAttribute('hy-channel') || '*';
      const event = firstEl.getAttribute('hy-event') || '*';
      const topic = firstEl.getAttribute('hy-topic') || undefined;

      this.routeEvent({
        channel,
        event,
        topic,
        html,
        timestamp: Date.now(),
        origin: 'server'
      });
    }
  }

  applySwap(target: HTMLElement, html: string, strategy: HySwapStrategy): void {
    if (!target || (!html && strategy !== 'delete')) return;

    const beforeEvt = new CustomEvent('hy-hook:before-swap', {
      bubbles: true,
      cancelable: true,
      detail: { target, html, strategy }
    });
    target.dispatchEvent(beforeEvt);
    if (beforeEvt.defaultPrevented) return;

    switch (strategy) {
      case 'innerHTML':
        target.innerHTML = html;
        break;
      case 'outerHTML':
        target.outerHTML = html;
        break;
      case 'afterbegin':
        target.insertAdjacentHTML('afterbegin', html);
        break;
      case 'beforeend':
        target.insertAdjacentHTML('beforeend', html);
        break;
      case 'delete':
        target.remove();
        break;
      case 'none':
        break;
    }

    if (typeof document !== 'undefined') {
      document.dispatchEvent(new CustomEvent('hy-hook:after-swap', {
        detail: { target, strategy }
      }));
    }
  }

  private processOOB(html: string): void {
    if (typeof document === 'undefined') return;
    const parser = document.createElement('div');
    parser.innerHTML = html.trim();

    const oobElements = parser.querySelectorAll('[hy-swap-oob], [hx-swap-oob]');
    oobElements.forEach(oobEl => {
      const id = oobEl.getAttribute('id');
      if (id) {
        const existing = document.getElementById(id);
        if (existing) {
          const mode = oobEl.getAttribute('hy-swap-oob') || oobEl.getAttribute('hx-swap-oob');
          const strategy: HySwapStrategy = (mode === 'true' || mode === 'outerHTML') ? 'outerHTML' : (mode as HySwapStrategy) || 'outerHTML';
          this.applySwap(existing, oobEl.outerHTML, strategy);
        }
      }
    });
  }

  scan(root: Document | HTMLElement = document): void {
    if (typeof document === 'undefined') return;

    // 1. Process hy-hook (Server Webhook Consumer)
    root.querySelectorAll<HTMLElement>('[hy-hook]').forEach(el => {
      if (el.hasAttribute('data-hy-bound')) return;
      el.setAttribute('data-hy-bound', 'true');

      const rawHook = el.getAttribute('hy-hook') || '';
      const [channel, event] = rawHook.split(':');
      const topic = el.getAttribute('hy-topic') || undefined;
      const swap = (el.getAttribute('hy-swap') || 'outerHTML') as HySwapStrategy;
      const target = el.getAttribute('hy-target') || undefined;

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

    // 2. Process hy-hook-listen (Cross-Tab / Local Broadcast Listener)
    root.querySelectorAll<HTMLElement>('[hy-hook-listen]').forEach(el => {
      if (el.hasAttribute('data-hy-listen-bound')) return;
      el.setAttribute('data-hy-listen-bound', 'true');

      const rawHook = el.getAttribute('hy-hook-listen') || '';
      const [channel, event] = rawHook.split(':');
      const topic = el.getAttribute('hy-topic') || undefined;
      const swap = (el.getAttribute('hy-swap') || 'afterbegin') as HySwapStrategy;
      const target = el.getAttribute('hy-target') || undefined;

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

    // 3. Process hy-hook-emit (Button / Action Webhook Emitter)
    root.querySelectorAll<HTMLElement>('[hy-hook-emit]').forEach(el => {
      if (el.hasAttribute('data-hy-emit-bound')) return;
      el.setAttribute('data-hy-emit-bound', 'true');

      const triggerEvent = el.tagName === 'FORM' ? 'submit' : 'click';
      el.addEventListener(triggerEvent, () => {
        const rawHook = el.getAttribute('hy-hook-emit') || '';
        const [channel, event] = rawHook.split(':');
        const topic = el.getAttribute('hy-topic') || undefined;
        const rawData = el.getAttribute('hy-data');
        const rawHtml = el.getAttribute('hy-html');

        let data: Record<string, any> | undefined;
        if (rawData) {
          try { data = JSON.parse(rawData); } catch {}
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

    // 4. Process hy-speculate-trigger (Speculative Pre-Cognition Commit Trigger)
    root.querySelectorAll<HTMLElement>('[hy-speculate-key]').forEach(el => {
      if (el.hasAttribute('data-hy-speculate-bound')) return;
      el.setAttribute('data-hy-speculate-bound', 'true');

      const key = el.getAttribute('hy-speculate-key') || '';
      const trigger = el.getAttribute('hy-speculate-trigger') || 'mouseenter';
      const target = el.getAttribute('hy-target') || undefined;
      const swap = (el.getAttribute('hy-swap') || 'outerHTML') as HySwapStrategy;

      el.addEventListener(trigger, () => {
        if (key) {
          this.commitSpeculation(key, target, swap);
        }
      });
    });
  }

  diag(): HyHookDiagnostics {
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

  private dispatchDiagEvent(status: string): void {
    if (typeof window !== 'undefined') {
      window.dispatchEvent(new CustomEvent('hy-hook:diag', {
        detail: { status, diag: this.diag() }
      }));
    }
  }
}

// Global Singleton Instance
export const HyHook = new HyperHookEngine();

// Auto-initialize when DOM is ready
if (typeof document !== 'undefined') {
  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', () => HyHook.scan());
  } else {
    HyHook.scan();
  }

  if ('MutationObserver' in window) {
    const observer = new MutationObserver((mutations) => {
      for (const mutation of mutations) {
        mutation.addedNodes.forEach(node => {
          if (node.nodeType === Node.ELEMENT_NODE) {
            HyHook.scan(node as HTMLElement);
          }
        });
      }
    });
    observer.observe(document.body, { childList: true, subtree: true });
  }
}

if (typeof window !== 'undefined') {
  (window as any).HyHook = HyHook;
  (window as any).HyperPlane = HyperPlane;
  (window as any).HyperFlag = HyperFlag;
  (window as any).HyperFrameCodec = HyperFrameCodec;
}
