package components

import (
	"fmt"
	"strings"
)

// PGXHUDProps defines configuration properties for rendering the PGX Ephemeral Database HUD.
type PGXHUDProps struct {
	ID             string
	DefaultTemplate string
	RuntimeLimitMB int
	PollInterval   string // e.g. "2s", "5s"
	Class          string
}

// PGXHUD renders the interactive PGX ephemeral runtime HUD matching views/components/pgx-hud.html.
func PGXHUD(p PGXHUDProps) Component {
	id := p.ID
	if id == "" {
		id = "pgx-hud-container"
	}
	tpl := p.DefaultTemplate
	if tpl == "" {
		tpl = "app"
	}

	pollAttr := ""
	if p.PollInterval != "" {
		pollAttr = fmt.Sprintf(` hx-get="/api/pgx/status" hx-trigger="every %s" hx-swap="outerHTML"`, p.PollInterval)
	}

	classAttr := "rounded-3xl bg-slate-900 border border-slate-800 shadow-2xl overflow-hidden flex flex-col select-none text-slate-100"
	if p.Class != "" {
		classAttr += " " + p.Class
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(`<div class="%s" id="%s" hx-ext="reactive"%s hx-state='{"rssMb":136,"totalDbs":1003,"ephemeralDbs":998,"spares":8,"connections":42,"selectedTpl":"%s","mintCount":38,"recentBranch":"tdb_%s__agent-alpha","branches":[{"name":"tdb_app__agent-01","tpl":"app","ttl":"1h 54m","size":"13.6 MB","status":"Active"},{"name":"tdb_app__agent-02","tpl":"app","ttl":"1h 58m","size":"14.1 MB","status":"Active"},{"name":"tdb_crm__pr-402","tpl":"crm","ttl":"45m","size":"15.2 MB","status":"Active"},{"name":"tdb_biotech__bench","tpl":"biotech","ttl":"12m","size":"18.0 MB","status":"Expiring"}]}'>`,
		classAttr, id, pollAttr, tpl, tpl))

	sb.WriteString(`
  <!-- ZONE 1: HEADER -->
  <div class="px-5 py-4 bg-slate-950 border-b border-slate-800 flex flex-wrap items-center justify-between gap-3">
    <div class="flex items-center gap-3">
      <div class="w-10 h-10 rounded-xl bg-blue-500/20 border border-blue-500/40 flex items-center justify-center text-blue-400 font-black text-xl shadow-md">
        🐘
      </div>
      <div>
        <div class="flex items-center gap-2">
          <h2 class="text-base font-black text-white tracking-tight flex items-center gap-2">
            <span>PGX Ephemeral Runtime HUD</span>
            <span class="text-[10px] font-mono text-blue-400 bg-blue-500/10 px-2 py-0.5 rounded-full border border-blue-500/20">Rust Engine v0.1</span>
          </h2>
          <span class="px-2 py-0.5 rounded text-[10px] font-bold bg-emerald-500/10 text-emerald-400 border border-emerald-500/20">PG 18.6 PROTOCOL</span>
        </div>
        <p class="text-xs text-slate-400 mt-0.5 font-sans">
          One runtime holding thousands of isolated Postgres databases. 6.2ms warm mint, copy-on-write ZFS storage, zero Docker.
        </p>
      </div>
    </div>
    <div class="flex items-center gap-2 bg-slate-950 px-3 py-1.5 rounded-xl border border-slate-800 font-mono text-xs">
      <span class="w-2 h-2 rounded-full bg-emerald-400 animate-pulse"></span>
      <span class="text-slate-400">Runtime RSS:</span>
      <span class="text-blue-400 font-bold" hx-text="$state.rssMb + ' MB'">136 MB</span>
      <span class="text-slate-600">|</span>
      <span class="text-slate-400">Total DBs:</span>
      <span class="text-emerald-400 font-bold" hx-text="$state.totalDbs">1,003</span>
      <span class="text-slate-600">|</span>
      <span class="text-slate-400">Mint P50:</span>
      <span class="text-purple-400 font-bold">6.2 ms</span>
    </div>
  </div>

  <!-- ZONE 2: METRIC CARDS -->
  <div class="grid grid-cols-2 md:grid-cols-4 gap-4 p-5 bg-slate-900/40 border-b border-slate-800">
    <div class="p-3.5 rounded-2xl bg-slate-950 border border-slate-800">
      <div class="text-[10px] uppercase font-bold text-slate-400 tracking-wider">Memory Density</div>
      <div class="text-xl font-black text-blue-400 mt-1 flex items-baseline gap-1">
        <span>~70 KB</span>
        <span class="text-[10px] font-normal text-slate-500">/ idle database</span>
      </div>
      <div class="text-[10px] text-slate-400 mt-1">vs 16.6 MB standard Postgres</div>
    </div>
    <div class="p-3.5 rounded-2xl bg-slate-950 border border-slate-800">
      <div class="text-[10px] uppercase font-bold text-slate-400 tracking-wider">Active Ephemeral DBs</div>
      <div class="text-xl font-black text-emerald-400 mt-1 flex items-baseline gap-1">
        <span hx-text="$state.ephemeralDbs">998</span>
        <span class="text-[10px] font-normal text-slate-500">sandboxes</span>
      </div>
      <div class="text-[10px] text-slate-400 mt-1">Automatic janitor TTL expiry</div>
    </div>
    <div class="p-3.5 rounded-2xl bg-slate-950 border border-slate-800">
      <div class="text-[10px] uppercase font-bold text-slate-400 tracking-wider">Spare Pool</div>
      <div class="text-xl font-black text-purple-400 mt-1 flex items-baseline gap-1">
        <span hx-text="$state.spares">8</span>
        <span class="text-[10px] font-normal text-slate-500">pre-created</span>
      </div>
      <div class="text-[10px] text-slate-400 mt-1">Instant 6ms mint-on-connect</div>
    </div>
    <div class="p-3.5 rounded-2xl bg-slate-950 border border-slate-800">
      <div class="text-[10px] uppercase font-bold text-slate-400 tracking-wider">Copy-on-Write</div>
      <div class="text-xl font-black text-amber-400 mt-1 flex items-baseline gap-1">
        <span>+13.6 MB</span>
        <span class="text-[10px] font-normal text-slate-500">new blocks</span>
      </div>
      <div class="text-[10px] text-slate-400 mt-1">10 GB base template in 674ms</div>
    </div>
  </div>

  <!-- ZONE 3: AGENT MINT CONTROLLER & LIVE BRANCH TABLE -->
  <div class="p-6 grid grid-cols-1 lg:grid-cols-12 gap-6 bg-slate-900/60">
    <div class="lg:col-span-5 flex flex-col space-y-4">
      <div class="p-4 rounded-2xl bg-slate-950 border border-slate-800 space-y-3">
        <div class="flex items-center justify-between">
          <span class="text-xs font-bold text-slate-300 uppercase tracking-wider flex items-center gap-1.5">
            <span>⚡</span> Mint-on-Connect Simulator
          </span>
          <span class="text-[10px] font-mono text-blue-400 bg-blue-500/10 px-2 py-0.5 rounded border border-blue-500/20">Zero Docker</span>
        </div>
        <p class="text-xs text-slate-400 leading-relaxed font-sans">
          Connect with database pattern <code class="text-blue-300 font-mono">tdb_&lt;template&gt;__&lt;agent_id&gt;</code>. PGX clones the template in 6.2ms without provisioning scripts.
        </p>
        <div class="space-y-2">
          <label class="text-[11px] font-bold text-slate-400 uppercase tracking-wider">Golden Template</label>
          <div class="grid grid-cols-3 gap-2">
            <button type="button"
                    class="px-2.5 py-1.5 rounded-xl border text-xs font-mono font-medium transition-all"
                    hx-class="$state.selectedTpl === 'app' ? 'bg-blue-600/30 text-blue-300 border-blue-500' : 'bg-slate-900 text-slate-400 border-slate-800 hover:border-slate-700'"
                    hx-action="$state.selectedTpl = 'app'">
              app (31MB)
            </button>
            <button type="button"
                    class="px-2.5 py-1.5 rounded-xl border text-xs font-mono font-medium transition-all"
                    hx-class="$state.selectedTpl === 'crm' ? 'bg-blue-600/30 text-blue-300 border-blue-500' : 'bg-slate-900 text-slate-400 border-slate-800 hover:border-slate-700'"
                    hx-action="$state.selectedTpl = 'crm'">
              crm (250MB)
            </button>
            <button type="button"
                    class="px-2.5 py-1.5 rounded-xl border text-xs font-mono font-medium transition-all"
                    hx-class="$state.selectedTpl === 'biotech' ? 'bg-blue-600/30 text-blue-300 border-blue-500' : 'bg-slate-900 text-slate-400 border-slate-800 hover:border-slate-700'"
                    hx-action="$state.selectedTpl = 'biotech'">
              biotech (10GB)
            </button>
          </div>
        </div>
        <button type="button"
                id="btn-mint-branch"
                class="w-full py-2.5 px-4 rounded-xl font-bold text-xs bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-500 hover:to-indigo-500 text-white shadow-lg shadow-blue-500/20 active:scale-95 transition-all flex items-center justify-center gap-2 cursor-pointer"
                hx-action="$state.mintCount++; $state.ephemeralDbs++; $state.totalDbs++; const id = 'agent-' + Math.floor(Math.random() * 900 + 100); const branch = 'tdb_' + $state.selectedTpl + '__' + id; $state.recentBranch = branch; $state.branches.unshift({ name: branch, tpl: $state.selectedTpl, ttl: '2h 00m', size: '13.6 MB', status: 'Active' }); HyperFX.toast('Minted ' + branch + ' in 6.2ms!', 'success');">
          <span>🚀</span> Mint Ephemeral Sandbox (6.2ms)
        </button>
        <div class="p-2.5 rounded-xl bg-slate-900 border border-slate-800 space-y-1">
          <div class="text-[10px] text-slate-500 font-mono">DATABASE_URL (Ready for Agent)</div>
          <div class="text-xs font-mono text-teal-300 truncate"
               hx-text="'postgres://postgres:5432/' + $state.recentBranch + '?sslmode=disable'">
            postgres://postgres:5432/tdb_app__agent-alpha?sslmode=disable
          </div>
        </div>
      </div>
    </div>
    <div class="lg:col-span-7 flex flex-col space-y-4">
      <div class="p-4 rounded-2xl bg-slate-950 border border-slate-800 flex flex-col h-full space-y-3">
        <div class="flex items-center justify-between">
          <span class="text-xs font-bold text-slate-300 uppercase tracking-wider flex items-center gap-1.5">
            <span>📦</span> Active Ephemeral Branches
          </span>
          <span class="text-[10px] font-mono text-emerald-400 bg-emerald-500/10 px-2 py-0.5 rounded border border-emerald-500/20">Janitor: 8 Spares</span>
        </div>
        <div class="overflow-x-auto flex-1">
          <table class="w-full text-left font-mono text-xs">
            <thead>
              <tr class="border-b border-slate-800 text-slate-500 text-[10px] uppercase">
                <th class="pb-2">Database Name</th>
                <th class="pb-2">Template</th>
                <th class="pb-2">Delta Size</th>
                <th class="pb-2">TTL Left</th>
                <th class="pb-2 text-right">Action</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-800/60" id="pgx-branch-table-body">
              <tr class="hover:bg-slate-900/50 transition-colors">
                <td class="py-2.5 text-blue-300 font-bold">tdb_app__agent-01</td>
                <td class="py-2.5 text-slate-400">app</td>
                <td class="py-2.5 text-slate-300">13.6 MB</td>
                <td class="py-2.5 text-emerald-400">1h 54m</td>
                <td class="py-2.5 text-right">
                  <span class="px-2 py-0.5 rounded text-[10px] bg-slate-800 text-slate-300 hover:text-white cursor-pointer"
                        hx-action="HyperFX.toast('Dropped database tdb_app__agent-01', 'info')">Drop</span>
                </td>
              </tr>
              <tr class="hover:bg-slate-900/50 transition-colors">
                <td class="py-2.5 text-blue-300 font-bold">tdb_app__agent-02</td>
                <td class="py-2.5 text-slate-400">app</td>
                <td class="py-2.5 text-slate-300">14.1 MB</td>
                <td class="py-2.5 text-emerald-400">1h 58m</td>
                <td class="py-2.5 text-right">
                  <span class="px-2 py-0.5 rounded text-[10px] bg-slate-800 text-slate-300 hover:text-white cursor-pointer">Drop</span>
                </td>
              </tr>
              <tr class="hover:bg-slate-900/50 transition-colors">
                <td class="py-2.5 text-blue-300 font-bold">tdb_crm__pr-402</td>
                <td class="py-2.5 text-slate-400">crm</td>
                <td class="py-2.5 text-slate-300">15.2 MB</td>
                <td class="py-2.5 text-amber-400">45m</td>
                <td class="py-2.5 text-right">
                  <span class="px-2 py-0.5 rounded text-[10px] bg-slate-800 text-slate-300 hover:text-white cursor-pointer">Drop</span>
                </td>
              </tr>
              <tr class="hover:bg-slate-900/50 transition-colors">
                <td class="py-2.5 text-blue-300 font-bold">tdb_biotech__bench</td>
                <td class="py-2.5 text-slate-400">biotech</td>
                <td class="py-2.5 text-slate-300">18.0 MB</td>
                <td class="py-2.5 text-rose-400">12m</td>
                <td class="py-2.5 text-right">
                  <span class="px-2 py-0.5 rounded text-[10px] bg-slate-800 text-slate-300 hover:text-white cursor-pointer">Drop</span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <div class="p-2.5 rounded-xl bg-slate-900/80 border border-slate-800/80 text-[11px] text-slate-400 flex items-center justify-between">
          <span class="flex items-center gap-1.5">
            <span class="text-blue-400 font-bold">Tip:</span> AI agent processes automatically inherit <code class="text-slate-300">DATABASE_URL</code> via PGX Skill.
          </span>
          <span class="text-[10px] text-slate-500 font-mono">pgrust_runtime_status()</span>
        </div>
      </div>
    </div>
  </div>
</div>`)

	return Component{html: sb.String()}
}
