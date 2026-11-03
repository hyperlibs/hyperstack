package components

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// PGXConfig defines connection parameters for communicating with a PGX (pgrust) runtime.
type PGXConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Template string `json:"template"`
	Prefix   string `json:"prefix"` // Defaults to "tdb_"
	SSLMode  string `json:"sslmode"`
}

// DefaultPGXConfig returns standard local PGX development settings.
func DefaultPGXConfig() PGXConfig {
	return PGXConfig{
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "postgres",
		Template: "app",
		Prefix:   "tdb_",
		SSLMode:  "disable",
	}
}

// MintURL returns a Postgres connection string following PGX's Mint-on-Connect specification.
// Connecting to this URL automatically mints an isolated database from the sealed template in ~6.2ms.
func (c PGXConfig) MintURL(token string) string {
	host := c.Host
	if host == "" {
		host = "127.0.0.1"
	}
	port := c.Port
	if port == 0 {
		port = 5432
	}
	user := c.User
	if user == "" {
		user = "postgres"
	}
	prefix := c.Prefix
	if prefix == "" {
		prefix = "tdb_"
	}
	tpl := c.Template
	if tpl == "" {
		tpl = "app"
	}
	ssl := c.SSLMode
	if ssl == "" {
		ssl = "disable"
	}

	// Format: <prefix><template>__<token>
	dbName := fmt.Sprintf("%s%s__%s", prefix, tpl, sanitizePGXToken(token))

	if c.Password != "" {
		return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", user, c.Password, host, port, dbName, ssl)
	}
	return fmt.Sprintf("postgres://%s@%s:%d/%s?sslmode=%s", user, host, port, dbName, ssl)
}

// ParsePGXDatabase inspects a database name and extracts the template and agent token.
// Returns true if the database name adheres to PGX's ephemeral naming convention.
func ParsePGXDatabase(dbName string, prefix string) (template string, token string, isEphemeral bool) {
	if prefix == "" {
		prefix = "tdb_"
	}
	if !strings.HasPrefix(dbName, prefix) {
		return "", "", false
	}
	remainder := strings.TrimPrefix(dbName, prefix)
	parts := strings.Split(remainder, "__")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", false
	}
	return parts[0], parts[1], true
}

func sanitizePGXToken(token string) string {
	var sb strings.Builder
	for _, r := range token {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' || r == '-' {
			sb.WriteRune(r)
		} else {
			sb.WriteRune('_')
		}
	}
	res := sb.String()
	if res == "" {
		return fmt.Sprintf("tok_%d", time.Now().UnixNano()%100000)
	}
	return res
}

// PGXRuntimeStatus represents the telemetry payload emitted by PGX's pgrust_runtime_status().
type PGXRuntimeStatus struct {
	Memory struct {
		RSSBytes         int64 `json:"rss_bytes"`
		SessionLimitMB   int   `json:"session_limit_mb"`
		DatabaseLimitMB  int   `json:"database_limit_mb"`
		RuntimeLimitMB   int   `json:"runtime_limit_mb"`
	} `json:"memory"`
	Connections int `json:"connections"`
	Databases   struct {
		Total     int `json:"total"`
		Ephemeral int `json:"ephemeral"`
		Spares    int `json:"spares"`
		Pinned    int `json:"pinned"`
	} `json:"databases"`
	Janitor struct {
		Prefix       string `json:"prefix"`
		PoolSize     int    `json:"pool_size"`
		PendingMints int    `json:"pending_mints"`
	} `json:"janitor"`
}

// ParsePGXStatus decodes the raw JSON output of SELECT pgrust_runtime_status();
func ParsePGXStatus(rawJSON string) (*PGXRuntimeStatus, error) {
	var status PGXRuntimeStatus
	if err := json.Unmarshal([]byte(rawJSON), &status); err != nil {
		return nil, fmt.Errorf("invalid PGX status JSON: %w", err)
	}
	return &status, nil
}
