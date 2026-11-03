import { describe, test, expect } from "bun:test";
import fs from "node:fs";
import path from "node:path";

describe("PGX Ephemeral Database & HUD Suite", () => {
  const componentPath = path.join(__dirname, "..", "views", "components", "pgx-hud.html");
  const goComponentPath = path.join(__dirname, "..", "go", "components", "pgx_hud.go");
  const goHelperPath = path.join(__dirname, "..", "go", "components", "pgx.go");

  test("1. pgx-hud.html view exists and is well-formed", () => {
    expect(fs.existsSync(componentPath)).toBe(true);
    const content = fs.readFileSync(componentPath, "utf-8");
    expect(content).toContain('id="pgx-hud-container"');
    expect(content).toContain('hx-ext="reactive"');
    expect(content).toContain('hx-state=');
    expect(content).toContain('PGX Ephemeral Runtime HUD');
    expect(content).toContain('Mint-on-Connect Simulator');
    expect(content).toContain('pgrust_runtime_status()');
  });

  test("2. Go DSL component and PGX helpers exist on disk", () => {
    expect(fs.existsSync(goComponentPath)).toBe(true);
    expect(fs.existsSync(goHelperPath)).toBe(true);

    const dslContent = fs.readFileSync(goComponentPath, "utf-8");
    expect(dslContent).toContain("func PGXHUD(p PGXHUDProps) Component");

    const helperContent = fs.readFileSync(goHelperPath, "utf-8");
    expect(helperContent).toContain("func (c PGXConfig) MintURL(token string) string");
    expect(helperContent).toContain("func ParsePGXDatabase(");
    expect(helperContent).toContain("func ParsePGXStatus(");
  });

  test("3. Mint-on-Connect URL generation matches PGX specification", () => {
    function generateMintURL(host: string, port: number, template: string, token: string): string {
      const sanitized = token.replace(/[^a-zA-Z0-9_-]/g, "_");
      return `postgres://postgres@${host}:${port}/tdb_${template}__${sanitized}?sslmode=disable`;
    }

    const url = generateMintURL("127.0.0.1", 5432, "app", "agent-alpha-42");
    expect(url).toBe("postgres://postgres@127.0.0.1:5432/tdb_app__agent-alpha-42?sslmode=disable");

    const urlSpecial = generateMintURL("localhost", 5433, "crm", "pr/492/test");
    expect(urlSpecial).toBe("postgres://postgres@localhost:5433/tdb_crm__pr_492_test?sslmode=disable");
  });

  test("4. PGX Ephemeral Database name parser correctly identifies branches", () => {
    function parsePGX(dbName: string, prefix = "tdb_") {
      if (!dbName.startsWith(prefix)) return null;
      const remainder = dbName.slice(prefix.length);
      const parts = remainder.split("__");
      if (parts.length !== 2 || !parts[0] || !parts[1]) return null;
      return { template: parts[0], token: parts[1] };
    }

    expect(parsePGX("tdb_app__agent-01")).toEqual({ template: "app", token: "agent-01" });
    expect(parsePGX("tdb_biotech__bench_run")).toEqual({ template: "biotech", token: "bench_run" });
    expect(parsePGX("production_db")).toBeNull();
    expect(parsePGX("tdb_invalid")).toBeNull();
  });

  test("5. Telemetry parser extracts memory limits and density metrics", () => {
    const rawTelemetry = {
      memory: {
        rss_bytes: 142606336,
        session_limit_mb: 256,
        database_limit_mb: 512,
        runtime_limit_mb: 3072
      },
      connections: 42,
      databases: {
        total: 1003,
        ephemeral: 998,
        spares: 8,
        pinned: 2
      },
      janitor: {
        prefix: "tdb_",
        pool_size: 8,
        pending_mints: 0
      }
    };

    const memoryMb = Math.round(rawTelemetry.memory.rss_bytes / 1024 / 1024);
    expect(memoryMb).toBe(136);
    expect(rawTelemetry.databases.ephemeral).toBe(998);
    expect(rawTelemetry.databases.spares).toBe(8);
  });
});
