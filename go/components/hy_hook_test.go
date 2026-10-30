package components

import (
	"strings"
	"testing"
)

func TestHyperFrameEncodingDecoding(t *testing.T) {
	payload := []byte("hello-quantum-wire")
	frame := HyperFrame{
		Plane:     HyperPlaneTensor,
		Flags:     HyperFlagReliable,
		ChannelId: 77,
		Sequence:  404,
		Payload:   payload,
	}

	encoded := EncodeHyperFrame(frame)
	if len(encoded) != HyperHeaderLen+len(payload) {
		t.Fatalf("expected length %d, got %d", HyperHeaderLen+len(payload), len(encoded))
	}

	decoded, err := DecodeHyperFrame(encoded)
	if err != nil {
		t.Fatalf("unexpected error decoding frame: %v", err)
	}

	if decoded.Plane != HyperPlaneTensor {
		t.Errorf("expected plane %d, got %d", HyperPlaneTensor, decoded.Plane)
	}
	if decoded.Flags != HyperFlagReliable {
		t.Errorf("expected flags %d, got %d", HyperFlagReliable, decoded.Flags)
	}
	if decoded.ChannelId != 77 {
		t.Errorf("expected channelId 77, got %d", decoded.ChannelId)
	}
	if decoded.Sequence != 404 {
		t.Errorf("expected sequence 404, got %d", decoded.Sequence)
	}
	if string(decoded.Payload) != "hello-quantum-wire" {
		t.Errorf("expected payload 'hello-quantum-wire', got %q", string(decoded.Payload))
	}
}

func TestSpecializedPlaneEncoders(t *testing.T) {
	// Plane 0: Tensor
	tensorBytes := EncodeTensorFrame(10, "inference token", true, 1)
	tensorFrame, err := DecodeHyperFrame(tensorBytes)
	if err != nil {
		t.Fatalf("decode tensor failed: %v", err)
	}
	if tensorFrame.Plane != HyperPlaneTensor {
		t.Errorf("expected PlaneTensor, got %d", tensorFrame.Plane)
	}
	if (tensorFrame.Flags & HyperFlagEOS) == 0 {
		t.Errorf("expected EOS flag set")
	}

	// Plane 1: Hypermedia
	hmBytes := EncodeHypermediaFrame(20, "#target", "innerHTML", "<b>Bold</b>", 2)
	hmFrame, err := DecodeHyperFrame(hmBytes)
	if err != nil {
		t.Fatalf("decode hypermedia failed: %v", err)
	}
	if hmFrame.Plane != HyperPlaneHypermedia {
		t.Errorf("expected PlaneHypermedia, got %d", hmFrame.Plane)
	}

	// Plane 2: Telemetry
	telBytes := EncodeTelemetryFrame(30, []float32{1.5, -2.5, 3.14}, 3)
	telFrame, err := DecodeHyperFrame(telBytes)
	if err != nil {
		t.Fatalf("decode telemetry failed: %v", err)
	}
	if telFrame.Plane != HyperPlaneTelemetry {
		t.Errorf("expected PlaneTelemetry, got %d", telFrame.Plane)
	}

	// Plane 3: Precognition
	precogBytes := EncodePrecognitionFrame(40, "intent_nav", 0.95, "<div>Prefetched</div>", 4)
	precogFrame, err := DecodeHyperFrame(precogBytes)
	if err != nil {
		t.Fatalf("decode precognition failed: %v", err)
	}
	if precogFrame.Plane != HyperPlanePrecognition {
		t.Errorf("expected PlanePrecognition, got %d", precogFrame.Plane)
	}

	// Plane 4: Swarm
	swarmBytes := EncodeSwarmFrame(50, "node-alpha", 999, "{}", 5)
	swarmFrame, err := DecodeHyperFrame(swarmBytes)
	if err != nil {
		t.Fatalf("decode swarm failed: %v", err)
	}
	if swarmFrame.Plane != HyperPlaneSwarm {
		t.Errorf("expected PlaneSwarm, got %d", swarmFrame.Plane)
	}
}

func TestGoDSLForwardConstructors(t *testing.T) {
	// Neural Stream DSL
	neuralComp := HyNeuralStream(HyNeuralStreamProps{
		ChannelId: 101,
		Class:     "bg-neutral-900",
		Initial:   "Ready...",
	})
	if !strings.Contains(neuralComp.Render(), `hy-ai-stream="101"`) {
		t.Errorf("expected hy-ai-stream attribute, got %s", neuralComp.Render())
	}

	// Telemetry Canvas DSL
	telComp := HyTelemetryCanvas(HyTelemetryCanvasProps{
		ChannelId: 202,
		CanvasId:  "gyro-canvas",
		Width:     400,
		Height:    300,
	})
	if !strings.Contains(telComp.Render(), `hy-telemetry="202"`) {
		t.Errorf("expected hy-telemetry attribute, got %s", telComp.Render())
	}

	// Speculative Branch DSL
	specComp := HySpeculativeBranch(HySpeculativeBranchProps{
		Key:     "tab_analytics",
		Trigger: "mouseenter",
		Label:   "Analytics",
	})
	if !strings.Contains(specComp.Render(), `hy-speculate-key="tab_analytics"`) {
		t.Errorf("expected hy-speculate-key attribute, got %s", specComp.Render())
	}

	// Swarm Cluster DSL
	swarmComp := HySwarmCluster(HySwarmClusterProps{
		NodeId:  "node-01",
		Cluster: "edge-grid",
	})
	if !strings.Contains(swarmComp.Render(), `hy-swarm-node="edge-grid:node-01"`) {
		t.Errorf("expected hy-swarm-node attribute, got %s", swarmComp.Render())
	}
}
