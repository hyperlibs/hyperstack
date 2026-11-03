package components

import (
	"strings"
	"testing"
)

func TestPGXHUDRender(t *testing.T) {
	comp := PGXHUD(PGXHUDProps{
		ID:              "custom-pgx-hud",
		DefaultTemplate: "crm",
		PollInterval:    "3s",
	})

	html := comp.Render()
	if !strings.Contains(html, `id="custom-pgx-hud"`) {
		t.Errorf("expected custom id in rendered HTML")
	}
	if !strings.Contains(html, `hx-trigger="every 3s"`) {
		t.Errorf("expected poll interval trigger in rendered HTML")
	}
	if !strings.Contains(html, `PGX Ephemeral Runtime HUD`) {
		t.Errorf("expected HUD title in rendered HTML")
	}
	if !strings.Contains(html, `tdb_crm__agent-alpha`) {
		t.Errorf("expected default template to be crm")
	}
}
