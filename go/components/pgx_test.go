package components

import (
	"testing"
)

func TestPGXMintURL(t *testing.T) {
	cfg := DefaultPGXConfig()
	url := cfg.MintURL("agent_007")

	expected := "postgres://postgres@127.0.0.1:5432/tdb_app__agent_007?sslmode=disable"
	if url != expected {
		t.Fatalf("expected MintURL %q, got %q", expected, url)
	}

	// Test with password and custom host
	customCfg := PGXConfig{
		Host:     "10.0.0.5",
		Port:     5433,
		User:     "admin",
		Password: "secret",
		Template: "crm_v2",
		Prefix:   "tdb_",
		SSLMode:  "require",
	}
	customURL := customCfg.MintURL("pr-492")
	expectedCustom := "postgres://admin:secret@10.0.0.5:5433/tdb_crm_v2__pr-492?sslmode=require"
	if customURL != expectedCustom {
		t.Fatalf("expected custom MintURL %q, got %q", expectedCustom, customURL)
	}
}

func TestParsePGXDatabase(t *testing.T) {
	tpl, tok, ok := ParsePGXDatabase("tdb_crm_v2__agent-99", "tdb_")
	if !ok {
		t.Fatalf("expected valid ephemeral database")
	}
	if tpl != "crm_v2" || tok != "agent-99" {
		t.Fatalf("expected tpl=crm_v2 tok=agent-99, got tpl=%q tok=%q", tpl, tok)
	}

	// Non-ephemeral
	_, _, ok = ParsePGXDatabase("production_db", "tdb_")
	if ok {
		t.Fatalf("expected production_db to not be recognized as ephemeral")
	}

	// Bad format
	_, _, ok = ParsePGXDatabase("tdb_broken", "tdb_")
	if ok {
		t.Fatalf("expected tdb_broken to fail parsing")
	}
}

func TestParsePGXStatus(t *testing.T) {
	raw := `{
		"memory": {
			"rss_bytes": 142606336,
			"session_limit_mb": 256,
			"database_limit_mb": 512,
			"runtime_limit_mb": 3072
		},
		"connections": 42,
		"databases": {
			"total": 1003,
			"ephemeral": 998,
			"spares": 8,
			"pinned": 2
		},
		"janitor": {
			"prefix": "tdb_",
			"pool_size": 8,
			"pending_mints": 0
		}
	}`

	status, err := ParsePGXStatus(raw)
	if err != nil {
		t.Fatalf("unexpected error parsing PGX status: %v", err)
	}

	if status.Connections != 42 {
		t.Errorf("expected 42 connections, got %d", status.Connections)
	}
	if status.Databases.Total != 1003 {
		t.Errorf("expected 1003 databases, got %d", status.Databases.Total)
	}
	if status.Databases.Ephemeral != 998 {
		t.Errorf("expected 998 ephemeral databases, got %d", status.Databases.Ephemeral)
	}
	if status.Memory.RuntimeLimitMB != 3072 {
		t.Errorf("expected 3072 runtime limit MB, got %d", status.Memory.RuntimeLimitMB)
	}
}
