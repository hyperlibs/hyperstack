# Hyperstack × PGX: High-Density Ephemeral Sandboxes

> **Architecture Blueprint**: Running thousands of isolated, disposable PostgreSQL databases alongside native Hypermedia backends for AI coding agents, preview environments, and high-density multi-tenant workstations.

---

## 1. Executive Overview

Modern web frameworks (e.g. Next.js, Remix) struggle with full-stack isolation. Giving every AI agent, CI job, or pull-request preview its own database typically requires either:
- **Heavier container overhead**: 1,000 Docker containers running PostgreSQL ($\sim$16 MB idle RAM and 30s cold boot per container).
- **Expensive cloud subscriptions**: Multi-tenant cloud database branching (Neon, Supabase) with high operational costs and latency.

**Hyperstack + PGX** eliminates this constraint entirely.

```
┌────────────────────────────────────────────────────────────────────────┐
│                        AI Coding Agent Swarm                           │
│     "Create an invoice dashboard with a new schema and live feed"      │
└───────────────────┬────────────────────────────────┬───────────────────┘
                    │                                │
                    ▼                                ▼
       ┌────────────────────────┐       ┌────────────────────────┐
       │     PGX (Rust)         │       │  Hyperstack (Go / Bun) │
       │  • Ephemeral DB (6ms)  │       │  • Zero-Build SSR      │
       │  • CoW Branch (+13 MB) │◄─────►│  • Live Hyperhooks     │
       │  • ~70 KB Idle RAM     │       │  • <1ms UI Render      │
       │  • Real PG 18 Dialect  │       │  • @diag Reflection    │
       └────────────────────────┘       └────────────────────────┘
                    │                                │
                    └────────────────┬───────────────┘
                                     ▼
        ┌────────────────────────────────────────────────────────┐
        │  1,000 Live Isolated Full-Stack Tenant Environments    │
        │      Running on a Single $8 / Month VPS (<250 MB RAM)  │
        └────────────────────────────────────────────────────────┘
```

| Metric | Standard PostgreSQL 18 | PGX (Rust Runtime) | Advantage |
| :--- | :--- | :--- | :--- |
| **Idle Memory per Database** | $\sim$16.6 MB (server process) | **$\sim$70 KB** ($\sim$128 KB active) | **230× memory density** |
| **Database Creation Latency** | 2,000–5,000 ms (`initdb` / container) | **6.2 ms warm** / 74 ms cold | **300× faster** |
| **Branch Storage Cost** | Full storage copy | **$\sim$13.6 MB delta** (CoW block clone) | **99% storage savings** |
| **1,000 Idle Databases** | $\sim$16.6 GB RAM | **128–155 MB RAM** | Runs on a low-end VPS |
| **SQL Compatibility** | Reference standard | **219 / 231 byte-exact** (0 result diffs) | 100% ORM / client parity |

---

## 2. Mint-on-Connect Architecture

Instead of orchestrating database creation via external scripts or API calls, PGX introduces **Mint-on-Connect**:

Connecting to an ephemeral connection string automatically mints the database from a sealed template:

$$\text{postgres://postgres:5432/}\mathbf{tdb\_}\langle\text{template}\rangle\mathbf{\_\_}\langle\text{token}\rangle\text{?sslmode=disable}$$

### Example Wire Flow:
1. Agent session connects: `psql -d tdb_app__agent_007`.
2. PGX detects the `tdb_` prefix and immediately mints an isolated database from the sealed `app` template in **6.2 milliseconds** from its pre-created warm pool.
3. The agent executes destructive migrations, modifies schemas, and tests UI state.
4. When the session disconnects, the built-in PGX janitor automatically destroys the branch after its TTL expires.

---

## 3. Go Backend Integration Recipe

Use the built-in Hyperstack Go DSL PGX adapter (`go/components/pgx.go`):

```go
package main

import (
    "database/sql"
    "fmt"
    "net/http"
    
    _ "github.com/lib/pq"
    "github.com/hyperlibs/hyperstack/go/components"
)

func main() {
    pgxConfig := components.DefaultPGXConfig()
    // Default template: "app"

    http.HandleFunc("/sandbox", func(w http.ResponseWriter, r *http.Request) {
        agentToken := r.Header.Get("X-Agent-ID")
        if agentToken == "" {
            agentToken = "anon_session"
        }

        // Generate the 6.2ms Mint-on-Connect URL
        dbURL := pgxConfig.MintURL(agentToken)
        db, err := sql.Open("postgres", dbURL)
        if err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        defer db.Close()

        // Render Hyperstack UI directly from the ephemeral tenant DB
        hud := components.PGXHUD(components.PGXHUDProps{
            ID:              "agent-hud",
            DefaultTemplate: "app",
            PollInterval:    "2s",
        })

        w.Header().Set("Content-Type", "text/html")
        w.Write([]byte(hud.Render()))
    })

    http.ListenAndServe(":8080", nil)
}
```

---

## 4. Bun Backend Integration Recipe

In `server.ts` or standalone Bun backends using `Bun.sql`:

```typescript
import { SQL } from "bun";

const PGX_HOST = process.env.PGX_HOST || "127.0.0.1";
const PGX_PORT = process.env.PGX_PORT || "5432";
const TEMPLATE = process.env.PGX_TEMPLATE || "app";

export function getAgentSQL(agentToken: string) {
  const sanitized = agentToken.replace(/[^a-zA-Z0-9_-]/g, "_");
  const dbName = `tdb_${TEMPLATE}__${sanitized}`;
  
  return new SQL({
    url: `postgres://postgres@${PGX_HOST}:${PGX_PORT}/${dbName}?sslmode=disable`
  });
}

// Example Request Handler
export async function handleAgentRequest(req: Request) {
  const agentId = req.headers.get("X-Agent-ID") || "dev";
  const sql = getAgentSQL(agentId);

  // Queries execute against the agent's private copy-on-write branch
  const invoices = await sql`SELECT id, customer, amount, status FROM invoices ORDER BY id DESC LIMIT 50`;
  
  // Return Hyperstack HTML table partial directly
  return new Response(
    `<table class="w-full text-xs font-mono">
       ${invoices.map(inv => `<tr><td>${inv.id}</td><td>${inv.customer}</td><td>$${inv.amount}</td></tr>`).join("")}
     </table>`,
    { headers: { "Content-Type": "text/html" } }
  );
}
```

---

## 5. Native Coding Agent Workflow (Claude Code / Cursor / Codex / Antigravity)

When using Hyperstack with AI coding agents:

1. **Step 1: Mint Agent Branch**:
   ```bash
   export DATABASE_URL="$(pgx branch app agent-$AGENT_ID | tail -1)"
   ```
2. **Step 2: Run Destructive Migrations & Seed Data**:
   The agent executes raw SQL or migration scripts directly against its isolated branch.
3. **Step 3: Live Hypermedia Verification**:
   The agent boots `bun run dev` or `go run server.go`. Because Hyperstack renders server-side HTML directly from the active connection, changes appear immediately in the browser without SPA re-compilation.
4. **Step 4: Automatic Janitor Cleanup**:
   When the agent finishes or the process exits, PGX reclaims the copy-on-write blocks.

---

## 6. Real-Time Telemetry & `@diag` Architecture

PGX provides internal telemetry via `SELECT pgrust_runtime_status();`. Hyperstack's `<app-pgx-hud>` component polls or streams this telemetry into reactive signals (`hx-state`):

```json
{
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
}
```

If an agent runs an unbounded query, PGX traps the allocation at the session boundary and returns **SQLSTATE 53200 (`out of memory`)**. Hyperstack's `@diag` taxonomy captures this error and dispatches a healing prompt directly to the coding agent without crashing the parent host daemon.
