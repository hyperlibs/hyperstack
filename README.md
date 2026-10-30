# Hyperstack — The Hypermedia Platform for Go, Rust, Python & .NET

<div align="center">
  <p><strong>A hyperreactive, lean, hypermedia-first frontend platform built natively on HTMX 4 for backend engineers and AI agent swarms.</strong></p>
  <p><em>Zero Virtual DOM. Zero Hydration Mismatches. Zero Node/NPM build step required. Server remains the single source of truth.</em></p>
  <p>
    <a href="https://github.com/hyperlibs/hyperstack/actions/workflows/ci.yml"><img src="https://github.com/hyperlibs/hyperstack/actions/workflows/ci.yml/badge.svg" alt="CI Status"></a>
    <a href="CHANGELOG.md"><img src="https://img.shields.io/badge/release-v0.10.30-blue?style=flat-square" alt="Release"></a>
    <a href="https://unpkg.com/htmx.org@4.0.0/"><img src="https://img.shields.io/badge/htmx-4.0.0_core-black?style=flat-square&logo=htmx" alt="HTMX 4 Core"></a>
    <a href="https://hyperstack.cgb.workers.dev/demo/"><img src="https://img.shields.io/badge/Cloudflare-Live_Deployment-F38020?logo=cloudflare&style=flat-square" alt="Cloudflare Deployment"></a>
    <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-purple?style=flat-square" alt="License"></a>
    <a href="CHECKSUMS.sha256"><img src="https://img.shields.io/badge/SHA--256-13%2F13_Verified-emerald?style=flat-square" alt="Checksums"></a>
  </p>
  <p>
    <a href="https://hyperstack.cgb.workers.dev/demo/"><strong>⚡ Live Demos on Cloudflare</strong></a> &nbsp;•&nbsp; 
    <a href="https://hyperstack.cgb.workers.dev/demo/mobile/"><strong>📱 AppzOS Mobile Simulator</strong></a> &nbsp;•&nbsp; 
    <a href="https://hyperstack.cgb.workers.dev/docs/components/introduction/"><strong>📖 Documentation</strong></a> &nbsp;•&nbsp; 
    <a href="#-about-us--the-hyperlibs-mission"><strong>👥 About Us</strong></a> &nbsp;•&nbsp; 
    <a href="https://github.com/hyperlibs/hyperstack"><strong>⭐ GitHub</strong></a>
  </p>
</div>

---

> [!IMPORTANT]
> ### 📢 The Modern Web Without JavaScript Fatigue
> **Skip Complex JS. Avoid Node.js Bloat. Build with Hyperstack.**  
> The web was meant to be simple. While single-page apps (SPAs) introduced 500MB `node_modules`, fragile build steps, and hydration mismatches, **Hyperstack** delivers a hyperreactive, zero-build platform that puts power back into your backend. 
> 
> Built on top of **HTMX 4**, Hyperstack combines the simplicity of declarative hypermedia with fine-grained signal reactivity, sub-millisecond in-memory columnar search, real-time webhooks, and an enterprise component catalog. Whether you write **Python, Go, C# (.NET), Rust, PHP, or Java**, Hyperstack lets you build high-performance, interactive user interfaces directly from your server with zero client-side glue code.

---

## ⚡ The Foundation: Native HTMX 4 Core & Extension Suite

Hyperstack uses the official **HTMX 4** (`4.0.0`) runtime as its underlying hypermedia engine. HTMX 4 introduces a modern fetch-based core, native DOM tree reconciliation, and enhanced security primitives:

| Extension / Module | Distribution Path | Purpose & Capability |
|---|---|---|
| **HTMX 4 Core** | `public/vendor/htmx.min.js` | Modern `fetch()` request pipeline, OOB fragment swaps, history, and DOM events. |
| **`hx-csp`** | `public/vendor/hx-csp.min.js` | Strict Content Security Policy (Level 3) enforcement with zero `eval()` or `new Function()`. |
| **`hx-live`** | `public/vendor/hx-live.min.js` | DOM-based live reactivity and fine-grained DOM state synchronization. |
| **`hx-alpine-compat`** | `public/vendor/hx-alpine-compat.min.js` | Direct drop-in syntax compatibility for Alpine.js reactive attribute patterns. |
| **`htmx-2-compat`** | `public/vendor/htmx-2-compat.min.js` | Backward compatibility layer for legacy HTMX 2 attributes and behaviors. |

---

## 🪝 Hyperhooks (`hy-hook`) — Real-Time Hypermedia Webhooks & Streaming

Hyperstack introduces **Hyperhooks (`hy-hook`)**, a declarative hypermedia protocol for real-time webhooks, SSE, WebSockets, and WebTransport streams:

- **Declarative Wire Subscriptions**: Connect any HTML container to live server webhooks with zero custom JavaScript:
  ```html
  <div hy-hook="/api/live-orders" 
       hy-action="merge" 
       hy-filter="status == 'active'" 
       hy-target="#order-list"
       hy-codec="json">
    <div id="order-list" class="divide-y"></div>
  </div>
  ```
- **Asymmetric Transport with Fallback**: Transparently connects over UDP/WebTransport or WebSockets with automated, zero-drop fallback to Server-Sent Events (SSE) and HTTP streaming.
- **Bit-Packed Zero-Copy Framing**: Compact serialization format designed for edge AI, mobile POS, and enterprise ERP event streams.
- **Go Dispatcher DSL**: Native Go server integration via `hy_hook.NewDispatcher()` in `go/components/hy_hook.go`.

---

## 📱 `appz` — 4-Layer Mobile OS & Spatial Gaming Stack

For mobile web apps, embedded screens, POS terminals, and spatial gaming HUDs, Hyperstack provides the standardized **`appz`** custom elements:

- **`<app-shellui>`**: The top-level mobile/game application shell container.
- **`<app-bar>`**: Top application header with leading navigation, titles, and trailing action slots.
- **`<app-seg id="segN">`**: Independent scrollable or spatial viewport segments.
- **`<app-layer level="0..9">`**: **Photoshop-style 0–9 Z-index layer stacking**:
  - `level="0"`: Base background canvas / content feed.
  - `level="1"`: Overlapping cards / floating HUD panels.
  - `level="2..8"`: Secondary floating badges / contextual controls.
  - `level="9"`: Topmost overlay / alert / dynamic action button.
- **`<app-navbar>`**: Bottom navigation with `mode="float|sticky|hidden"` and `behavior="hide-on-scroll"`.
- **`<app-bottom-sheet>`**: Spring-animated modal bottom sheet (`$appz.sheet('#id')`).
- **`<app-joystick>`**: Virtual twin-stick HUD controller for mobile games and robotics.

---

## ⚡ The Equation

```text
Hyperstack = HTMX 4 (Modern DOM & Fetch Hypermedia Core)
       + Bolt (Signal-Based Reactive State & 2D Sparse Matrix)
       + Flash (In-Memory Typo-Tolerant Search & Columnar Store)
       + Form (Declarative Validation & Cell Transaction Rollbacks)
       + Virtual (Bi-Directional 2D Scroller with Pinned Columns)
       + Calc (Topological Formula DAG & Cycle Detection)
       + Vibe (60fps Spring Physics & FLIP Layout Animations)
       + A11y (WAI-ARIA Focus Trapping & 2D Matrix Nav)
       + Sim (Particle Kinetics, Boids Flocking & Conway Life)
       + Offline (IndexedDB Mutation Queues & Auto-Sync)
       + Canvas (Layer-Indexed Spatial Node Graphs & Live Béziers)
       + DevTools (Telemetry Bridge, Real-Time Signal HUD & @diag)
       + Appz (4-Layer Mobile OS, Photoshop 0–9 Layers & Hardware Bridge)
       + Hy-Hook (Hypermedia Webhooks & Streaming Binary Protocol)
       + 200+ Production-Grade Components (100 Basic + 100 Smart + 10 Design Blocks)
```

---

## 📦 The 13 Modular Core Micro-Engines (~71 KB Total Gzipped)

> Sizes below are exact gzipped bytes of the shipped bundles and are enforced by `tests/benchmarks.test.ts`; see [docs/BENCHMARKS.md](docs/BENCHMARKS.md) for the full raw/gzip table and methodology.

| Engine | Source (`src/`) | Distribution (`public/`) | Size (Gzipped) | Mathematical & Algorithmic Foundation | Capabilities |
|---|---|---|---|---|---|
| **⚡ Bolt** | `htmx-bolt.ts` | `htmx-bolt.js` | 23.2 KB | $O(1)$ coordinate-keyed sparse matrix & microtask batching | Signals, dependency tracking via deep Proxies, computed properties, `$store`, `HxBolt.matrix`, 120 FPS ticker, zero-eval safe evaluator. |
| **🔍 Flash** | `htmx-flash.ts` | `htmx-flash.js` | 4.25 KB | TypedArray columnar indexing & bitmask query filters | In-memory fuzzy search (`hx-flash-search`), multi-column filters, TypedArray columnar store (`createColumnStore`). |
| **🪟 Virtual** | `htmx-virtual.ts` | `htmx-virtual.js` | 3.27 KB | $O(1)$ spatial slicing & binary row-height bisecting | Bi-directional 2D windowing virtualization with pinned rows/columns for million-row tables. |
| **📊 Calc** | `htmx-calc.ts` | `htmx-calc.js` | 3.18 KB | Kahn's algorithm topological DAG sort & cycle detection | Reactive formula DAG engine with coordinate references (`A1`, `SUM(A1:B10)`) and `#CYCLE!` detection. |
| **📝 Form** | `htmx-form.ts` | `htmx-form.js` | 4.41 KB | Deterministic finite-state machine (FSM) validation | Declarative validation (`hx-validate="required\|email"`), optimistic cell transactions with undo/redo (`Ctrl+Z`). |
| **🌊 Vibe** | `htmx-vibe.ts` | `htmx-vibe.js` | 3.26 KB | Damped harmonic oscillator kinematics: $m \ddot{x} + c \dot{x} + kx = 0$ | FLIP layout animation engine, spring physics solver (`SpringSolver`), stagger sequences, scroll triggers. |
| **📋 Grid** | `htmx-grid.ts` | `htmx-grid.js` | 5.17 KB | Frozen-pane offset math over virtualized column indices | Enterprise data grid: column resize, drag reorder, pinned columns, CSV export. |
| **🧬 Sim** | `htmx-sim.ts` | `htmx-sim.js` | 6.00 KB | Particle physics, Boids Reynolds rules, 2D Barnes-Hut gravity | Swarm intelligence flocking, N-Body orbital dynamics, elastic momentum, 23/3 Conway Life, slime mould sensors. |
| **♿ A11y** | `htmx-a11y.ts` | `htmx-a11y.js` | 3.05 KB | WAI-ARIA roving coordinate space navigation | Modal focus trapping (`hx-trap-focus`), roving tabindex, 2D matrix navigation, Excel range selection. |
| **📶 Offline** | `htmx-offline.ts` | `htmx-offline.js` | 2.32 KB | FIFO transaction journal with exponential backoff | IndexedDB mutation queues with exponential retry backoff and online sync. |
| **🛠️ DevTools** | `htmx-devtools.ts` | `htmx-devtools.js` | 2.04 KB | Telemetry bridge & SARIF 2.1.0 diagnostic stream | Real-time signal inspector, swap latency timing metrics, `@diag` HUD. |
| **🖼️ Canvas** | `htmx-canvas.ts` | `htmx-canvas.js` | 2.15 KB | Layer-indexed spatial node graph with live connector routing | Draggable node diagrams, automatic connector/bézier routing between nodes, per-node layer resolution. |
| **📱 Appz** | `appz.ts` | `appz.js` | 8.94 KB | 0–9 Photoshop-style layer stack, viewport OS modes & coroutine channels | `app-shellui`, `app-bar`, `app-seg`, `app-layer`, `app-navbar`, bottom sheets, joystick HUD, haptics, hardware/network/battery bridge. |

---

## ⚡ Verified Performance & Automated Benchmarks

Hyperstack includes an automated, reproducible benchmark suite (`bun run benchmark`) measuring real runtime throughput and payload efficiency. Full table, methodology and variance notes: [docs/BENCHMARKS.md](docs/BENCHMARKS.md).

- **Complete 13-Engine Suite**: **296,284 B raw / 72,947 B gzipped** — smaller than `react-dom` alone before components or state libraries.
- **Atomic 2D Sparse Matrix**: **427,000+ cell writes / second** in $O(1)$ coordinate space.
- **In-Memory Columnar Search**: **0.146 ms query latency** across 10,000 rows.
- **DAG Topological Solver**: **0.15 ms cascade resolution** across 500-level formula dependencies.
- **Zero-`eval` Strict CSP Parser**: **1,500–3,600 complex expressions & arrow closures / second** with zero dynamic `new Function()`.
- **Automated Verification**: CI runs `bun run typecheck`, all engine builds, `sha256sum -c CHECKSUMS.sha256`, the full `bun test` suite (109 tests, 100% pass) and a 350 KB bundle budget on Linux, macOS and Windows.

---

## 🎮 Live Flagship Demos

Experience complete, production-grade applications built with zero React/Node runtime:

> [!TIP]
> **⚡ Live Cloudflare Deployment**: [hyperstack.cgb.workers.dev/demo/](https://hyperstack.cgb.workers.dev/demo/)  
> **💻 Instant Local Run**: `bun run dev` (or `go run server.go` $\to$ `http://localhost:3000/demo/`).  
> **🚀 Deploy to Cloudflare**: `bun run deploy:pages` (or `bun run deploy:worker`).

| Flagship Demo | Live Cloudflare URL | Source View | Capabilities Demonstrated | Stack / Engines |
| :--- | :--- | :--- | :--- | :--- |
| **🎯 Verifiable Demos Hub** | [**/demo/**](https://hyperstack.cgb.workers.dev/demo/) | [`views/demos-hub.html`](views/demos-hub.html) | Master index of live interactive previews, stress tests, and showcases. | All Micro-Engines |
| **📱 AppzOS Mobile Simulator** | [**/demo/mobile/**](https://hyperstack.cgb.workers.dev/demo/mobile/) | [`views/app-mobile.html`](views/app-mobile.html) | 8 ergonomic mobile/tablet/TV viewports, Photoshop-style 0-9 layer stacking, coroutines, and haptics. | `appz.js`, `HxBolt`, `HxVibe` |
| **📊 HyperAdmin Dashboard** | [**/demo/tailadmin/**](https://hyperstack.cgb.workers.dev/demo/tailadmin/) | [`views/app-tailadmin.html`](views/app-tailadmin.html) | 100% pure Hyperstack enterprise admin panel with dynamic sales/profit analytics, dark/light theme, and reactive sidebar. | `HxBolt`, `HxFlash`, `HxVibe` |
| **🏢 Enterprise HyperERP** | [**/demo/erp/**](https://hyperstack.cgb.workers.dev/demo/erp/) | [`views/app-erp.html`](views/app-erp.html) | 10,000+ virtualized rows, zero-latency in-memory fuzzy search, and multi-tier analytics. | `HxVirtual`, `HxFlash`, `HxGrid` |
| **📈 HyperSheet Spreadsheet** | [**/demo/hypersheet/**](https://hyperstack.cgb.workers.dev/demo/hypersheet/) | [`views/app-hypersheet.html`](views/app-hypersheet.html) | 2D reactive spreadsheet with topological formula DAGs, coordinate resolution (`SUM(A1:B5)`), cycle detection, and roving 2D tabindex. | `HxCalc`, `HxA11y`, `HxBolt.matrix` |
| **🌌 Cosmic 3D Universe** | [**/demo/universe/**](https://hyperstack.cgb.workers.dev/demo/universe/) | [`views/app-universe.html`](views/app-universe.html) | Autonomous procedural universe viewport with domain-warped solar convection shaders and relativistic accretion disk simulations. | `HxBolt`, `HxSim`, `HxVibe` |

---

## 🔬 10 Research & Industrial Design Blocks

Hyperstack includes 10 domain-specific, zero-dependency engineering workstations built for specialized industries:

| Block Component | Source | Target Domain | Key Capabilities |
|---|---|---|---|
| **Pathology Lab** | [`views/components/block-pathology-lab.html`](views/components/block-pathology-lab.html) | Histopathology & Oncology | Whole-slide gigapixel pyramid viewer, stain normalization, tumor cellularity scoring. |
| **Radiology DICOM** | [`views/components/block-radiology-dicom.html`](views/components/block-radiology-dicom.html) | Medical Imaging & RECIST 1.1 | Multi-planar CT/MRI reconstruction, Hounsfield windowing, lesion measurement. |
| **CRISPR Studio** | [`views/components/block-crispr-studio.html`](views/components/block-crispr-studio.html) | Gene Editing & Genomics | sgRNA on/off-target scoring (CFD/Hsu), PAM recognition, cleavage visualization. |
| **Genomic Map** | [`views/components/block-genomic-map.html`](views/components/block-genomic-map.html) | Bioinformatics | Multi-track chromosome browser, cytoband rendering, VCF variant density heatmaps. |
| **Biology Atlas** | [`views/components/block-biology-atlas.html`](views/components/block-biology-atlas.html) | Single-Cell Transcriptomics | UMAP/t-SNE cell cluster projections, gene expression overlays, cell-type clustering. |
| **Civil Architect** | [`views/components/block-civil-architect.html`](views/components/block-civil-architect.html) | Structural & BIM CAD | IFC/BIM element hierarchy, finite element stress heatmaps, deflection diagrams. |
| **Marine Navigation** | [`views/components/block-marine-navigation.html`](views/components/block-marine-navigation.html) | Maritime Telemetry | ECDIS nautical chart plotter, AIS vessel tracking, radar sweep overlays. |
| **Precision Agriculture** | [`views/components/block-agriculture-precision.html`](views/components/block-agriculture-precision.html) | Agronomy & Remote Sensing | Multi-spectral NDVI satellite heatmaps, soil moisture grids, variable-rate fertilizer zones. |
| **Energy SCADA** | [`views/components/block-energy-scada.html`](views/components/block-energy-scada.html) | Industrial IoT & Smart Grids | Substation single-line diagrams (SLD), busbar voltage telemetry, breaker trip commands. |
| **Semiconductor PCB** | [`views/components/block-pcb-semiconductor.html`](views/components/block-pcb-semiconductor.html) | Electronics & EDA CAD | Multi-layer Gerber viewer (copper, silkscreen, solder mask), DRC netlist inspector. |

---

## 🐹 Go DSL Component Architecture (1:1 Parity)

Every HTML component in Hyperstack has an exact 1:1 Go DSL counterpart in `go/components/`:

```go
package main

import (
    "fmt"
    "net/http"
    "github.com/hyperlibs/hyperstack/go/components"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        page := components.Page(components.PageProps{
            Title: "Hyperstack Enterprise Dashboard",
            Content: components.VStack(
                components.StatCard(components.StatCardProps{
                    Title: "Quarterly Revenue",
                    Value: "$1,429,200",
                    Trend: "+18.4%",
                }),
                components.DataTable(components.DataTableProps{
                    ID: "orders-grid",
                    Endpoint: "/api/orders",
                    Columns: []string{"ID", "Customer", "Amount", "Status"},
                }),
            ),
        })
        fmt.Fprint(w, page.Render())
    })

    http.ListenAndServe(":3000", nil)
}
```

---

## 📖 Specifications & Guides

- [🔌 **Backend Wire Protocol Specification**](docs/BACKEND_SPEC.md): Copy-paste server implementations for Go, Rust, Python, and .NET.
- [🛡️ **Security & Hostile-Input Defense**](docs/SECURITY.md): CSP Level 3 headers, anti-CSRF token docking, and OOB injection mitigation.
- [🤖 **Machine-Actionable `@diag` Protocol**](docs/DIAGNOSTICS.md): Self-healing diagnostics, SARIF 2.1.0 schema, and error codes.
- [📊 **Official Benchmark Suite**](docs/BENCHMARKS.md): Measured bundle footprints and core algorithmic throughput.
- [📜 **Release Notes & Changelog**](CHANGELOG.md): Version history and upgrades.
- [🔒 **Release Checksums**](CHECKSUMS.sha256): Cryptographic verification hashes for all shipped micro-engines (13/13 verified).

---

## 🚀 Quick Start

### 1. Include the Scripts

```html
<!-- 1. HTMX 4 Core & Official Extensions -->
<script src="/vendor/htmx.min.js"></script>
<script src="/vendor/hx-csp.min.js"></script>
<script src="/vendor/hx-live.min.js"></script>

<!-- 2. Hyperstack Reactive Signal Engine -->
<script src="/htmx-bolt.js"></script>

<!-- 3. Optional Micro-Engines -->
<script src="/htmx-flash.js"></script>     <!-- In-Memory Fuzzy Search & Columnar Index -->
<script src="/htmx-form.js"></script>      <!-- Declarative Validation & Optimistic UI -->
<script src="/htmx-vibe.js"></script>      <!-- FLIP Layout Animations & Spring Physics -->
<script src="/htmx-a11y.js"></script>      <!-- Modal Focus Trap & 2D Matrix Tabindex -->
<script src="/htmx-virtual.js"></script>   <!-- 1D/2D Viewport Virtualization -->
<script src="/htmx-grid.js"></script>      <!-- Enterprise Data Grid with Column Drag -->
<script src="/htmx-offline.js"></script>   <!-- IndexedDB Offline Mutation Queue -->
<script src="/htmx-calc.js"></script>      <!-- Topological DAG Formula Solver -->
<script src="/htmx-sim.js"></script>       <!-- Kinetic Particle Physics & Swarms -->
<script src="/htmx-canvas.js"></script>    <!-- Spatial Node Editor & Bézier Routing -->
<script src="/htmx-devtools.js"></script>  <!-- Real-Time Signal HUD & Telemetry -->
<script src="/appz.js"></script>           <!-- 4-Layer Mobile OS & Hardware Bridge -->
<script src="/hy-hook.js"></script>        <!-- Hyperhooks Webhooks & Streaming Protocol -->
```

### 2. Reactive Component Example

```html
<div class="p-6 border border-border rounded-xl space-y-4 bg-card"
     hx-computed='{ totalPrice: items.reduce((sum, i) => sum + (i.price * i.qty), 0) }'>

  <!-- Initial State Declaration -->
  <script type="application/json" hx-state>
  {
    "newItem": "",
    "items": [
      { "id": 1, "name": "Mechanical Keyboard", "price": 120, "qty": 1 },
      { "id": 2, "name": "Wireless Mouse", "price": 60, "qty": 2 }
    ]
  }
  </script>

  <!-- Two-Way Binding -->
  <div class="flex gap-2">
    <input type="text" hx-model="newItem" placeholder="Item name..." class="px-3 py-2 border rounded-md flex-1 bg-background text-foreground">
    <button hx-on:click='if(newItem.trim()) { items.push({ id: Date.now(), name: newItem, price: 40, qty: 1 }); newItem = ""; }' class="bg-primary text-primary-foreground px-4 py-2 rounded-md font-medium">
      Add
    </button>
  </div>

  <!-- Structural List Rendering -->
  <div class="divide-y border rounded-md overflow-hidden">
    <template hx-for="(item, idx) in items">
      <div class="p-3 flex items-center justify-between bg-card">
        <span hx-text="item.name" class="font-medium text-foreground"></span>
        <div class="flex items-center gap-3">
          <span class="text-muted-foreground">$<span hx-text="item.price"></span></span>
          <input type="number" hx-model.number="item.qty" min="1" class="w-16 px-2 py-1 border rounded text-xs bg-background">
          <button hx-on:click="items.splice(idx, 1)" class="text-destructive text-xs hover:underline">Remove</button>
        </div>
      </div>
    </template>
  </div>

  <!-- Computed Reactive Total -->
  <div class="flex justify-between items-center font-bold text-base pt-2 text-foreground">
    <span>Total:</span>
    <span class="text-primary">$<span hx-text="totalPrice">0</span></span>
  </div>
</div>
```

---

## 🛠️ Local Development & Contributing

> **Strict Maintainer Principle**: Zero Node.js and Zero NPM. Hyperstack uses `bun` for TypeScript/bundling and `go` for native servers.

```bash
# Install dependencies
bun install

# Start development server
bun run dev

# Or run the native single-binary Go server
go run server.go

# Compile TypeScript engines to public distribution bundles
bun run build:engines

# Run the automated unit test suite (109 tests, 100% pass)
bun test

# Verify SHA-256 release checksums
bun run scripts/verify-checksums.ts
```

---

## 👥 About Us & The Hyperlibs Mission

[**Hyperlibs**](https://github.com/hyperlibs) is an open-source systems research and engineering collective dedicated to liberating backend engineers, systems developers, and AI agents from the complexity of modern client-side JavaScript stacks.

### 🧭 Core Principles
1. **Zero Virtual DOM**: HTML elements are the primary components; the server remains the single source of truth.
2. **Native HTMX 4 Core**: Modern fetch pipeline, strict CSP Level 3 enforcement, and zero-eval reactivity.
3. **Lean & Hyperreactive**: All 13 core micro-engines together weigh ~71 KB gzipped — smaller than a typical React hydration runtime.
4. **Deterministic Structure**: Predictable declarative attributes (`hx-state`, `hx-action`, `hx-text`, `hx-show`, `hy-hook`, `@diag`) make UI easily inspectable and maintainable.
5. **Backend Freedom**: First-class support for Go, Rust, Python, .NET, PHP, and Bun with zero client build toolchains.

- 🚀 **Live Previews**: [Verifiable Live Demos — Hyperstack](https://hyperstack.cgb.workers.dev/demo/)
- 📖 **Read More**: [Introduction & Architecture Documentation](https://hyperstack.cgb.workers.dev/docs/components/introduction/)
- 💬 **Community & Discussions**: [GitHub Discussions](https://github.com/hyperlibs/hyperstack/discussions)
- 🤝 **Organization**: [github.com/hyperlibs](https://github.com/hyperlibs)

---

## 📄 License

MIT © [Hyperlibs](https://github.com/hyperlibs)
