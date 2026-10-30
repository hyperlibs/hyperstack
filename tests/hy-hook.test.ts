import { describe, it, expect, beforeEach } from "bun:test";
import {
  HyHook,
  HyperHookEngine,
  HyperPlane,
  HyperFlag,
  HyperFrameCodec,
  type HyperFrame,
} from "../src/hy-hook";

describe("HyperHook (hy-hook): Hypermedia Webhooks & Event Mesh Engine", () => {
  let engine: HyperHookEngine;

  beforeEach(() => {
    engine = new HyperHookEngine();
  });

  // 1. Subscription & Channel/Event Routing
  it("subscribes and routes events by channel and event name", () => {
    let receivedPayload: any = null;

    const subId = engine.subscribe({
      channel: "pacs",
      event: "study.completed",
      callback: (evt) => {
        receivedPayload = evt;
      }
    });

    expect(subId).toBeTruthy();

    // Mismatched channel - should not fire
    engine.routeEvent({
      channel: "billing",
      event: "study.completed",
      timestamp: Date.now()
    });
    expect(receivedPayload).toBeNull();

    // Mismatched event - should not fire
    engine.routeEvent({
      channel: "pacs",
      event: "study.started",
      timestamp: Date.now()
    });
    expect(receivedPayload).toBeNull();

    // Exact match - should fire
    engine.routeEvent({
      channel: "pacs",
      event: "study.completed",
      html: "<div id='study-alert'>Study Ready (320 slices)</div>",
      data: { patientId: "P1" },
      timestamp: Date.now()
    });

    expect(receivedPayload).not.toBeNull();
    expect(receivedPayload.channel).toBe("pacs");
    expect(receivedPayload.event).toBe("study.completed");
    expect(receivedPayload.data.patientId).toBe("P1");
  });

  // 2. Topic Scoping (Entity/Tenant Multi-Tenancy)
  it("enforces strict topic scoping when specified", () => {
    const receivedEvents: string[] = [];

    // Subscribe to patient_A only
    engine.subscribe({
      channel: "pacs",
      event: "scan.ready",
      topic: "patient_A",
      callback: (evt) => {
        receivedEvents.push(`patient_A:${evt.data?.sliceId}`);
      }
    });

    // Subscribe to wildcard topic (*)
    engine.subscribe({
      channel: "pacs",
      event: "scan.ready",
      topic: "*",
      callback: (evt) => {
        receivedEvents.push(`wildcard:${evt.topic}`);
      }
    });

    // Fire event for patient_B
    engine.routeEvent({
      channel: "pacs",
      event: "scan.ready",
      topic: "patient_B",
      data: { sliceId: "SL-01" },
      timestamp: Date.now()
    });

    // Only wildcard should match
    expect(receivedEvents).toEqual(["wildcard:patient_B"]);

    // Fire event for patient_A
    engine.routeEvent({
      channel: "pacs",
      event: "scan.ready",
      topic: "patient_A",
      data: { sliceId: "SL-02" },
      timestamp: Date.now()
    });

    expect(receivedEvents).toContain("patient_A:SL-02");
    expect(receivedEvents).toContain("wildcard:patient_A");
  });

  // 3. Unsubscription
  it("unsubscribes cleanly by subscription ID", () => {
    let callCount = 0;
    const subId = engine.subscribe({
      channel: "iot",
      event: "telemetry",
      callback: () => { callCount++; }
    });

    engine.routeEvent({ channel: "iot", event: "telemetry", timestamp: Date.now() });
    expect(callCount).toBe(1);

    const unsubSuccess = engine.unsubscribe(subId);
    expect(unsubSuccess).toBe(true);

    engine.routeEvent({ channel: "iot", event: "telemetry", timestamp: Date.now() });
    expect(callCount).toBe(1); // No additional invocation
  });

  // 4. Safe Zero-Eval DOM Swapping Strategies
  it("performs safe zero-eval DOM swaps with various strategies", () => {
    const container: any = {
      innerHTML: "<span>initial content</span>",
      outerHTML: "<div id='target'><span>initial content</span></div>",
      removeCalled: false,
      dispatchEvent: () => true,
      insertAdjacentHTML(position: string, text: string) {
        if (position === "afterbegin") this.innerHTML = text + this.innerHTML;
        if (position === "beforeend") this.innerHTML = this.innerHTML + text;
      },
      remove() {
        this.removeCalled = true;
      }
    };

    // innerHTML strategy
    engine.applySwap(container, "<strong>updated inner</strong>", "innerHTML");
    expect(container.innerHTML).toBe("<strong>updated inner</strong>");

    // afterbegin strategy
    engine.applySwap(container, "<em>prefix</em>", "afterbegin");
    expect(container.innerHTML).toBe("<em>prefix</em><strong>updated inner</strong>");

    // beforeend strategy
    engine.applySwap(container, "<small>suffix</small>", "beforeend");
    expect(container.innerHTML).toBe("<em>prefix</em><strong>updated inner</strong><small>suffix</small>");

    // delete strategy
    engine.applySwap(container, "", "delete");
    expect((container as any).removeCalled).toBe(true);
  });

  // 5. Local Event Emission & Diagnostics
  it("emits events and updates telemetry counters in diag()", () => {
    const diagInitial = engine.diag();
    expect(diagInitial.eventsEmittedCount).toBe(0);
    expect(diagInitial.eventsReceivedCount).toBe(0);

    let emittedEvent: any = null;
    engine.subscribe({
      channel: "pos",
      event: "checkout",
      callback: (evt) => { emittedEvent = evt; }
    });

    engine.emit("pos", "checkout", {
      data: { total: 149.99 },
      html: "<div id='receipt'>Total: $149.99</div>",
      broadcast: false
    });

    expect(emittedEvent).not.toBeNull();
    expect(emittedEvent.data.total).toBe(149.99);

    const diagAfter = engine.diag();
    expect(diagAfter.eventsEmittedCount).toBe(1);
    expect(diagAfter.eventsReceivedCount).toBe(1);
    expect(diagAfter.activeSubscriptionsCount).toBe(1);
    expect(diagAfter.lastEvent?.channel).toBe("pos");
  });

  // ===========================================================================
  // QUANTUM BINARY FRAMING (FORWARD-FOCUSED SUITE)
  // ===========================================================================

  // 6. Binary Framing: Header Integrity & Corrupt Frame Handling
  it("encodes and decodes quantum binary frames with strict header verification", () => {
    const rawPayload = new Uint8Array([1, 2, 3, 4, 5, 6, 7, 8]);
    const frame: HyperFrame = {
      plane: HyperPlane.TENSOR,
      flags: HyperFlag.RELIABLE,
      channelId: 42,
      sequence: 1001,
      timestamp: Date.now(),
      payload: rawPayload
    };

    const encoded = HyperFrameCodec.encode(frame);
    expect(encoded.byteLength).toBe(16 + rawPayload.byteLength);
    expect(encoded[0]).toBe(0x48); // 'H'
    expect(encoded[1]).toBe(0x59); // 'Y'
    expect(encoded[2]).toBe(0x01); // Version
    expect(encoded[3]).toBe(HyperPlane.TENSOR);

    const decoded = HyperFrameCodec.decode(encoded);
    expect(decoded.plane).toBe(HyperPlane.TENSOR);
    expect(decoded.flags).toBe(HyperFlag.RELIABLE);
    expect(decoded.channelId).toBe(42);
    expect(decoded.sequence).toBe(1001);
    expect(Array.from(decoded.payload)).toEqual([1, 2, 3, 4, 5, 6, 7, 8]);

    // Corrupt magic check
    const corruptMagic = new Uint8Array(encoded);
    corruptMagic[0] = 0x00;
    expect(() => HyperFrameCodec.decode(corruptMagic)).toThrow(/Invalid magic bytes/);

    // Truncated buffer check
    expect(() => HyperFrameCodec.decode(new Uint8Array([0x48, 0x59]))).toThrow(/Corrupt frame/);
  });

  // 7. Plane 0 (Tensor / Neural Stream): AI Token Chunking & EOS
  it("encodes and routes Plane 0 neural stream tokens into listeners", () => {
    const tokensReceived: string[] = [];
    let eosReached = false;

    engine.onNeuralStream(101, (token, isEos) => {
      tokensReceived.push(token);
      if (isEos) eosReached = true;
    });

    const frame1 = HyperFrameCodec.encodeTensor(101, "Deep ", false, 1);
    const frame2 = HyperFrameCodec.encodeTensor(101, "Thinking...", false, 2);
    const frame3 = HyperFrameCodec.encodeTensor(101, " Done.", true, 3);

    engine.routeBinaryFrame(frame1);
    engine.routeBinaryFrame(frame2);
    engine.routeBinaryFrame(frame3);

    expect(tokensReceived).toEqual(["Deep ", "Thinking...", " Done."]);
    expect(eosReached).toBe(true);
    expect(engine.diag().binaryFramesReceivedCount).toBe(3);
  });

  // 8. Plane 1 (Hypermedia): Atomic Target & Swap Codec
  it("encodes and decodes Plane 1 atomic DOM morph frames", () => {
    const encoded = HyperFrameCodec.encodeHypermedia(
      200,
      "#patient-status",
      "innerHTML",
      "<span class='text-emerald-500'>Vitals Normal</span>",
      50
    );

    const frame = HyperFrameCodec.decode(encoded);
    expect(frame.plane).toBe(HyperPlane.HYPERMEDIA);

    const mutation = HyperFrameCodec.decodeHypermedia(frame.payload);
    expect(mutation.target).toBe("#patient-status");
    expect(mutation.swap).toBe("innerHTML");
    expect(mutation.html).toContain("Vitals Normal");
  });

  // 9. Plane 2 (Telemetry): 60fps Float32Array Zero-Copy Streaming
  it("streams packed Float32Array vectors with zero serialization allocations", () => {
    let capturedFloats: Float32Array | null = null;

    engine.registerTelemetryRenderer(300, (floats) => {
      capturedFloats = floats;
    });

    const inputFloats = new Float32Array([12.5, -4.25, 0.75, 99.125]);
    const encoded = HyperFrameCodec.encodeTelemetry(300, inputFloats, 1);

    const res = engine.routeBinaryFrame(encoded);
    expect(res.handled).toBe(true);
    expect(capturedFloats).not.toBeNull();
    expect(capturedFloats![0]).toBeCloseTo(12.5);
    expect(capturedFloats![1]).toBeCloseTo(-4.25);
    expect(capturedFloats![2]).toBeCloseTo(0.75);
    expect(capturedFloats![3]).toBeCloseTo(99.125);
  });

  // 10. Plane 3 (Speculative Pre-Cognition): 0.0ms Latency State Activation
  it("caches speculative hypermedia deltas and commits instantly", () => {
    const preRenderedHTML = "<div id='spec-card'>Instant Pre-Materialized State</div>";
    const encoded = HyperFrameCodec.encodePrecognition(400, "user_intent_tab_billing", 0.94, preRenderedHTML);

    engine.routeBinaryFrame(encoded);

    const cached = engine.getSpeculation("user_intent_tab_billing");
    expect(cached).toBeDefined();
    expect(cached?.confidence).toBeCloseTo(0.94);
    expect(cached?.html).toBe(preRenderedHTML);
    expect(engine.diag().speculativeCacheCount).toBe(1);

    // Commit speculation
    const committed = engine.commitSpeculation("user_intent_tab_billing");
    expect(committed).toBe(true);
    expect(engine.getSpeculation("user_intent_tab_billing")).toBeUndefined();
    expect(engine.diag().speculativeCacheCount).toBe(0);
  });

  // 11. Plane 4 (Swarm Cluster Mesh): Node Heartbeat & Vector Clock Sync
  it("synchronizes cluster swarm node state and vector clocks", () => {
    const encoded = HyperFrameCodec.encodeSwarm(500, "edge-node-eu-01", 1420, '{"cluster":"medical-pacs"}');

    engine.routeBinaryFrame(encoded);

    const swarmNodes = engine.getSwarmNodes();
    expect(swarmNodes.length).toBe(1);
    expect(swarmNodes[0].nodeId).toBe("edge-node-eu-01");
    expect(swarmNodes[0].vectorClock).toBe(1420);
    expect(engine.diag().swarmNodesCount).toBe(1);
  });
});
