package components

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"net/http"
	"sync"
	"time"
)

// =============================================================================
// HY-HOOK: GOLANG COMPONENT DSL & HYPERMEDIA WEBHOOK GATEWAY
// =============================================================================

type HyHookProps struct {
	Channel string
	Event   string
	Topic   string
	Swap    string // "outerHTML" | "innerHTML" | "afterbegin" | "beforeend" | "delete"
	Target  string
	Content string
}

// HyHook renders a declarative HyperHook subscriber element matching views/components/hy-hook.html.
func HyHook(p HyHookProps) Component {
	if p.Channel == "" {
		p.Channel = "default"
	}
	if p.Event == "" {
		p.Event = "message"
	}
	if p.Swap == "" {
		p.Swap = "outerHTML"
	}

	topicAttr := ""
	if p.Topic != "" {
		topicAttr = fmt.Sprintf(` hy-topic="%s"`, p.Topic)
	}

	targetAttr := ""
	if p.Target != "" {
		targetAttr = fmt.Sprintf(` hy-target="%s"`, p.Target)
	}

	html := fmt.Sprintf(`<div hy-hook="%s:%s"%s%s hy-swap="%s">
  %s
</div>`, p.Channel, p.Event, topicAttr, targetAttr, p.Swap, p.Content)

	return Component{html: html}
}

// HyHookEmit renders a button or form that broadcasts a hypermedia webhook.
type HyHookEmitProps struct {
	Channel  string
	Event    string
	Topic    string
	DataJSON string
	Label    string
	Class    string
}

func HyHookEmit(p HyHookEmitProps) Component {
	topicAttr := ""
	if p.Topic != "" {
		topicAttr = fmt.Sprintf(` hy-topic="%s"`, p.Topic)
	}
	dataAttr := ""
	if p.DataJSON != "" {
		dataAttr = fmt.Sprintf(` hy-data='%s'`, p.DataJSON)
	}

	html := fmt.Sprintf(`<button hy-hook-emit="%s:%s"%s%s class="%s">
  %s
</button>`, p.Channel, p.Event, topicAttr, dataAttr, p.Class, p.Label)

	return Component{html: html}
}

// =============================================================================
// FORWARD-FOCUSED HYPERHOOK DSL CONSTRUCTORS (EDGE AI & SPATIAL FABRIC)
// =============================================================================

// HyNeuralStreamProps configures an Edge AI token streaming container.
type HyNeuralStreamProps struct {
	ChannelId uint16
	Target    string
	Class     string
	Initial   string
}

// HyNeuralStream renders a real-time neural streaming container (Plane 0).
func HyNeuralStream(p HyNeuralStreamProps) Component {
	targetAttr := ""
	if p.Target != "" {
		targetAttr = fmt.Sprintf(` hy-ai-stream="%s"`, p.Target)
	} else {
		targetAttr = fmt.Sprintf(` hy-ai-stream="%d"`, p.ChannelId)
	}

	html := fmt.Sprintf(`<div id="neural-stream-%d"%s class="%s">
  %s
</div>`, p.ChannelId, targetAttr, p.Class, p.Initial)

	return Component{html: html}
}

// HyTelemetryCanvasProps configures a 60fps binary telemetry canvas (Plane 2).
type HyTelemetryCanvasProps struct {
	ChannelId uint16
	CanvasId  string
	Width     int
	Height    int
	Class     string
}

// HyTelemetryCanvas renders a canvas subscriber for continuous zero-copy float vectors.
func HyTelemetryCanvas(p HyTelemetryCanvasProps) Component {
	html := fmt.Sprintf(`<canvas id="%s" width="%d" height="%d" hy-telemetry="%d" class="%s"></canvas>`,
		p.CanvasId, p.Width, p.Height, p.ChannelId, p.Class)
	return Component{html: html}
}

// HySpeculativeBranchProps configures a speculative pre-cognition commit trigger (Plane 3).
type HySpeculativeBranchProps struct {
	Key     string
	Trigger string // "mouseenter" | "focus" | "click"
	Target  string
	Label   string
	Class   string
}

// HySpeculativeBranch renders a predictive interactive control with 0.0ms state activation.
func HySpeculativeBranch(p HySpeculativeBranchProps) Component {
	if p.Trigger == "" {
		p.Trigger = "mouseenter"
	}
	targetAttr := ""
	if p.Target != "" {
		targetAttr = fmt.Sprintf(` hy-target="%s"`, p.Target)
	}

	html := fmt.Sprintf(`<button hy-speculate-key="%s" hy-speculate-trigger="%s"%s class="%s">
  %s
</button>`, p.Key, p.Trigger, targetAttr, p.Class, p.Label)

	return Component{html: html}
}

// HySwarmClusterProps configures a P2P cluster state node HUD (Plane 4).
type HySwarmClusterProps struct {
	NodeId  string
	Cluster string
	Class   string
}

// HySwarmCluster renders a cluster synchronization HUD.
func HySwarmCluster(p HySwarmClusterProps) Component {
	html := fmt.Sprintf(`<div hy-swarm-node="%s:%s" class="%s">
  <span class="inline-block w-2.5 h-2.5 rounded-full bg-emerald-500 animate-pulse mr-2"></span>
  <span>Cluster: <strong>%s</strong> | Node: <code>%s</code></span>
</div>`, p.Cluster, p.NodeId, p.Class, p.Cluster, p.NodeId)

	return Component{html: html}
}

// =============================================================================
// QUANTUM DUAL-PLANE BINARY FRAMING (GOLANG CODEC)
// =============================================================================

const (
	HyperMagic0    byte = 0x48 // 'H'
	HyperMagic1    byte = 0x59 // 'Y'
	HyperVersion   byte = 0x01
	HyperHeaderLen      = 16

	HyperPlaneTensor       byte = 0x00
	HyperPlaneHypermedia   byte = 0x01
	HyperPlaneTelemetry    byte = 0x02
	HyperPlanePrecognition byte = 0x03
	HyperPlaneSwarm        byte = 0x04

	HyperFlagUnreliable  byte = 0x01
	HyperFlagReliable    byte = 0x02
	HyperFlagSpeculative byte = 0x04
	HyperFlagEOS         byte = 0x08
)

type HyperFrame struct {
	Plane     byte
	Flags     byte
	ChannelId uint16
	Sequence  uint32
	Payload   []byte
}

// EncodeHyperFrame serializes a HyperFrame into a 16-byte fixed header + payload slice.
func EncodeHyperFrame(f HyperFrame) []byte {
	payloadLen := len(f.Payload)
	buf := make([]byte, HyperHeaderLen+payloadLen)

	buf[0] = HyperMagic0
	buf[1] = HyperMagic1
	buf[2] = HyperVersion
	buf[3] = f.Plane
	buf[4] = f.Flags
	buf[5] = 0x00 // reserved
	binary.BigEndian.PutUint16(buf[6:8], f.ChannelId)
	binary.BigEndian.PutUint32(buf[8:12], f.Sequence)
	binary.BigEndian.PutUint32(buf[12:16], uint32(payloadLen))

	copy(buf[HyperHeaderLen:], f.Payload)
	return buf
}

// DecodeHyperFrame unpacks raw bytes into a HyperFrame with zero-copy slicing.
func DecodeHyperFrame(b []byte) (HyperFrame, error) {
	if len(b) < HyperHeaderLen {
		return HyperFrame{}, errors.New("frame length shorter than header")
	}
	if b[0] != HyperMagic0 || b[1] != HyperMagic1 {
		return HyperFrame{}, errors.New("invalid magic bytes")
	}
	if b[2] != HyperVersion {
		return HyperFrame{}, fmt.Errorf("unsupported version: %d", b[2])
	}

	channelId := binary.BigEndian.Uint16(b[6:8])
	sequence := binary.BigEndian.Uint32(b[8:12])
	payloadLen := binary.BigEndian.Uint32(b[12:16])

	if len(b) < HyperHeaderLen+int(payloadLen) {
		return HyperFrame{}, errors.New("incomplete payload")
	}

	return HyperFrame{
		Plane:     b[3],
		Flags:     b[4],
		ChannelId: channelId,
		Sequence:  sequence,
		Payload:   b[HyperHeaderLen : HyperHeaderLen+int(payloadLen)],
	}, nil
}

// EncodeTensorFrame serializes an AI token chunk (Plane 0).
func EncodeTensorFrame(channelId uint16, token string, isEos bool, sequence uint32) []byte {
	flags := HyperFlagReliable
	if isEos {
		flags |= HyperFlagEOS
	}
	return EncodeHyperFrame(HyperFrame{
		Plane:     HyperPlaneTensor,
		Flags:     flags,
		ChannelId: channelId,
		Sequence:  sequence,
		Payload:   []byte(token),
	})
}

// EncodeHypermediaFrame serializes an atomic DOM morph frame (Plane 1).
func EncodeHypermediaFrame(channelId uint16, target, swap, html string, sequence uint32) []byte {
	swapByte := byte(0x01)
	switch swap {
	case "outerHTML":
		swapByte = 0x01
	case "innerHTML":
		swapByte = 0x02
	case "afterbegin":
		swapByte = 0x03
	case "beforeend":
		swapByte = 0x04
	case "delete":
		swapByte = 0x05
	case "none":
		swapByte = 0x06
	}

	targetBytes := []byte(target)
	htmlBytes := []byte(html)
	payload := make([]byte, 1+len(targetBytes)+1+len(htmlBytes))

	payload[0] = byte(len(targetBytes))
	copy(payload[1:], targetBytes)
	payload[1+len(targetBytes)] = swapByte
	copy(payload[2+len(targetBytes):], htmlBytes)

	return EncodeHyperFrame(HyperFrame{
		Plane:     HyperPlaneHypermedia,
		Flags:     HyperFlagReliable,
		ChannelId: channelId,
		Sequence:  sequence,
		Payload:   payload,
	})
}

// EncodeTelemetryFrame serializes a 60fps packed float vector (Plane 2).
func EncodeTelemetryFrame(channelId uint16, floats []float32, sequence uint32) []byte {
	payload := make([]byte, len(floats)*4)
	for i, f := range floats {
		binary.LittleEndian.PutUint32(payload[i*4:], math.Float32bits(f))
	}

	return EncodeHyperFrame(HyperFrame{
		Plane:     HyperPlaneTelemetry,
		Flags:     HyperFlagUnreliable,
		ChannelId: channelId,
		Sequence:  sequence,
		Payload:   payload,
	})
}

// EncodePrecognitionFrame serializes speculative pre-rendered hypermedia (Plane 3).
func EncodePrecognitionFrame(channelId uint16, key string, confidence float32, html string, sequence uint32) []byte {
	keyBytes := []byte(key)
	htmlBytes := []byte(html)
	payload := make([]byte, 1+len(keyBytes)+4+len(htmlBytes))

	payload[0] = byte(len(keyBytes))
	copy(payload[1:], keyBytes)
	binary.BigEndian.PutUint32(payload[1+len(keyBytes):], math.Float32bits(confidence))
	copy(payload[5+len(keyBytes):], htmlBytes)

	return EncodeHyperFrame(HyperFrame{
		Plane:     HyperPlanePrecognition,
		Flags:     HyperFlagSpeculative,
		ChannelId: channelId,
		Sequence:  sequence,
		Payload:   payload,
	})
}

// EncodeSwarmFrame serializes a cluster node heartbeat / gossip frame (Plane 4).
func EncodeSwarmFrame(channelId uint16, nodeId string, clock uint32, data string, sequence uint32) []byte {
	nodeBytes := []byte(nodeId)
	dataBytes := []byte(data)
	payload := make([]byte, 1+len(nodeBytes)+4+len(dataBytes))

	payload[0] = byte(len(nodeBytes))
	copy(payload[1:], nodeBytes)
	binary.BigEndian.PutUint32(payload[1+len(nodeBytes):], clock)
	copy(payload[5+len(nodeBytes):], dataBytes)

	return EncodeHyperFrame(HyperFrame{
		Plane:     HyperPlaneSwarm,
		Flags:     HyperFlagReliable,
		ChannelId: channelId,
		Sequence:  sequence,
		Payload:   payload,
	})
}

// =============================================================================
// GOLANG SERVER HYPERMEDIA WEBHOOK BROKER / GATEWAY
// =============================================================================

type HyperHookBroker struct {
	mu      sync.RWMutex
	clients map[chan string]bool
}

var GlobalHyperHookBroker = &HyperHookBroker{
	clients: make(map[chan string]bool),
}

// Broadcast sends an HTML partial to all connected hypermedia SSE streams.
func (b *HyperHookBroker) Broadcast(channel, event, topic, htmlPartial string) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	msg := fmt.Sprintf(`data: {"channel":"%s","event":"%s","topic":"%s","html":%q,"timestamp":%d}`+"\n\n",
		channel, event, topic, htmlPartial, time.Now().UnixMilli())

	for clientChan := range b.clients {
		select {
		case clientChan <- msg:
		default:
		}
	}
}

// ServeSSE handles SSE connection multiplexing for /hy-hook/stream
func (b *HyperHookBroker) ServeSSE(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	clientChan := make(chan string, 16)
	b.mu.Lock()
	b.clients[clientChan] = true
	b.mu.Unlock()

	defer func() {
		b.mu.Lock()
		delete(b.clients, clientChan)
		close(clientChan)
		b.mu.Unlock()
	}()

	notify := r.Context().Done()

	// Initial ping
	fmt.Fprintf(w, "event: connected\ndata: {\"status\":\"ready\"}\n\n")
	flusher.Flush()

	for {
		select {
		case <-notify:
			return
		case msg := <-clientChan:
			fmt.Fprint(w, msg)
			flusher.Flush()
		}
	}
}
