package components

import (
	"strings"
	"testing"
)

func TestUniversalAddressingAndStreamConstructors(t *testing.T) {
	// 1. ObjID and DocID attribute formatting
	objAttr := ObjID("hero-title")
	if objAttr != `objid="hero-title"` {
		t.Errorf("expected objid=\"hero-title\", got %q", objAttr)
	}

	docAttr := DocID("site-main")
	if docAttr != `docid="site-main"` {
		t.Errorf("expected docid=\"site-main\", got %q", docAttr)
	}

	// 2. HxStream attribute helper
	if HxStream() != "hx-stream" {
		t.Errorf("expected hx-stream, got %q", HxStream())
	}
	if HxStream(false) != "" {
		t.Errorf("expected empty string when disabled, got %q", HxStream(false))
	}

	// 3. 2D Layer Matrix constructor
	layer := HxLayer(5, "manipulator", "resize-handles", Raw(`<span class="handle"></span>`))
	rendered := layer.Render()

	if !strings.Contains(rendered, `<hx-layer level="5"`) {
		t.Errorf("expected rendered layer to contain level=\"5\", got %s", rendered)
	}
	if !strings.Contains(rendered, `role="manipulator"`) {
		t.Errorf("expected rendered layer to contain role=\"manipulator\", got %s", rendered)
	}
	if !strings.Contains(rendered, `sub="resize-handles"`) {
		t.Errorf("expected rendered layer to contain sub=\"resize-handles\", got %s", rendered)
	}
	if !strings.Contains(rendered, `<span class="handle"></span>`) {
		t.Errorf("expected rendered layer to contain child, got %s", rendered)
	}

	// 4. Atomic Patch Validation
	validPatch := StreamPatch{
		Target: "btn_save",
		Op:     "set_style",
		Prop:   "background",
		Val:    "#2563eb",
	}
	if err := ValidatePatch(validPatch); err != nil {
		t.Errorf("expected valid patch to pass validation, got %v", err)
	}

	invalidPatch := StreamPatch{
		Target: "",
		Op:     "set_style",
	}
	if err := ValidatePatch(invalidPatch); err == nil {
		t.Errorf("expected patch with missing target to fail validation")
	}

	unrecognizedOpPatch := StreamPatch{
		Target: "box_1",
		Op:     "unknown_mutation",
	}
	if err := ValidatePatch(unrecognizedOpPatch); err == nil {
		t.Errorf("expected patch with unrecognized op to fail validation")
	}
}
