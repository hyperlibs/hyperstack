# Hyperstack Official Performance & Benchmark Suite

> Reproducible benchmark suite covering all **13 Hyperstack micro-engines** and the core reactive algorithms.
> Run it locally with `bun run benchmark` (see [How to Reproduce](#how-to-reproduce)).

---

## 1. Micro-Engine Payload Footprint

Measured from the shipped artifacts in `public/` on **2026-10-16** (Bun 1.3.4, Windows x64).
Raw and gzip figures below are exact byte counts, and `tests/benchmarks.test.ts` fails if this table drifts from the actual bundle sizes.

| Micro-Engine | Responsibility | Raw | Gzipped |
| :--- | :--- | ---: | ---: |
| **htmx-bolt.js** | Fine-grained signals, zero-eval parser, auto-fallback | 107,775 B | 23,761 B |
| **htmx-flash.js** | In-memory fuzzy search & TypedArray columnar queries | 17,246 B | 4,353 B |
| **htmx-form.js** | Declarative validation, optimistic UI & undo/redo | 17,999 B | 4,518 B |
| **htmx-vibe.js** | FLIP layout animation & damped-harmonic physics | 11,944 B | 3,335 B |
| **htmx-a11y.js** | Focus trap, roving tabindex & 2D matrix navigation | 11,977 B | 3,124 B |
| **htmx-virtual.js** | 1D/2D bi-directional viewport virtualization | 14,008 B | 3,350 B |
| **htmx-grid.js** | Data grid: column resize, reorder & CSV export | 20,335 B | 5,292 B |
| **htmx-calc.js** | Reactive formula engine & Kahn DAG topological sort | 11,908 B | 3,254 B |
| **htmx-offline.js** | IndexedDB offline cache & 409 conflict resolution | 7,925 B | 2,375 B |
| **htmx-devtools.js** | Signal graph inspector & `@diag` telemetry | 6,303 B | 2,087 B |
| **htmx-canvas.js** | Spatial node diagram engine | 7,525 B | 2,198 B |
| **htmx-sim.js** | Kinetic particle physics, swarm boids & 2D simulation | 25,547 B | 6,145 B |
| **appz.js** | App shell, `app-seg`/`app-layer` stacking, hardware bridge | 35,792 B | 9,155 B |
| **TOTAL (13 engines)** | **Complete full-stack frontend capability** | **296,284 B** | **72,947 B** |

The gzip total is the sum of the per-file gzip figures, which is the honest number because engines are served as separate requests. The CI bundle budget in `.github/workflows/ci.yml` enforces a **350,000 B** ceiling on the uncompressed `public/htmx-*.js` + `public/appz.js` total; the current total is **296,284 B** (~84% of budget).

*For comparison, `react-dom` alone is ~130 KB raw / ~42 KB gzip — before any state management library (Zustand/Redux), router, component primitives, physics, or virtualized grid is added.*

---

## 2. Core Algorithmic Benchmarks

Representative single run of `bun run benchmark` on the machine noted above. **Timings are machine- and load-dependent and swing noticeably between runs** — three consecutive runs of the parser benchmark here measured 1,575 / 2,345 / 3,603 expressions per second. Treat these as order-of-magnitude figures, not guarantees. The bundle sizes in §1, by contrast, are exact and enforced by CI.

### A. 2D Coordinate Sparse Matrix
- **Benchmark**: 10,000 coordinate-keyed cell writes (O(1)) on a live 100×100 tick grid.
- **Latency**: **23.41 ms** total (~0.0023 ms per cell mutation)
- **Throughput**: **~427,000 writes / second** (observed range 427k–724k across runs)

### B. In-Memory Columnar Search (TypedArray Indexing)
- **Benchmark**: 1,000 range-filter queries across a 10,000-row columnar store.
- **Latency**: **0.1461 ms** average query latency
- **Throughput**: **~6,800 multi-column queries / second** (3,220 matches per query)

### C. Kahn's DAG Formula Topological Sort
- **Benchmark**: Evaluating a 500-level deep chained formula dependency graph (A1 = 100, A*n* = A*n-1* × 1.05).
- **Latency**: **0.15 ms** total cascade evaluation
- Circular dependencies are detected and surfaced as `#CYCLE!` rather than recursing.

### D. Zero-Eval Recursive Descent Parser Throughput
- **Benchmark**: 10,000 complex array `.reduce()` arrow closures plus arithmetic expressions, evaluated under strict CSP without `new Function()`.
- **Throughput**: **2,345 expressions / second** on this run (observed range 1,575–3,603)

### E. Heap Footprint (Sparse Signal Cells)
- **Benchmark**: 20,000 coordinate-keyed signal cells over a 50,000 × 50 matrix.
- **Heap delta**: **~5,317 KB** for 20,000 cells (**~266 bytes per cell**)
- Each cell stores an object (`{ value, timestamp }`), which dominates the per-cell cost; the coordinate index itself stays sparse and allocation-free.

---

## How to Reproduce

```bash
bun run build       # rebuild engines, ESM bundle, CSS and static pages
bun run benchmark   # run the full suite above
bun test tests/     # includes the doc-vs-artifact size drift check
```
