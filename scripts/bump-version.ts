import fs from "fs";
import path from "path";
import { gzipSync } from "zlib";
import { createHash } from "crypto";
import { execSync } from "child_process";

// Local maintainer version synchronizer for Hyperstack
// Usage: bun run scripts/bump-version.ts [patch|minor|major|<explicit-version>] [--commit]

const rootDir = path.resolve(__dirname, "..");
const pkgPath = path.join(rootDir, "package.json");
const pkg = JSON.parse(fs.readFileSync(pkgPath, "utf8"));
const currentVersion = pkg.version;

const arg = process.argv[2] || "patch";
const autoCommit = process.argv.includes("--commit");

let nextVersion = "";
const semverRegex = /^\d+\.\d+\.\d+$/;

if (semverRegex.test(arg)) {
  nextVersion = arg;
} else {
  const [maj, min, pat] = currentVersion.split(".").map(Number);
  if (arg === "major") {
    nextVersion = `${maj + 1}.0.0`;
  } else if (arg === "minor") {
    nextVersion = `${maj}.${min + 1}.0`;
  } else {
    // default to patch
    nextVersion = `${maj}.${min}.${pat + 1}`;
  }
}

console.log(`\n===============================================================`);
console.log(`  Hyperstack Maintainer Release Sync: v${currentVersion} -> v${nextVersion}`);
console.log(`===============================================================\n`);

// 1. package.json
pkg.version = nextVersion;
fs.writeFileSync(pkgPath, JSON.stringify(pkg, null, 2) + "\n");
console.log(`  ✓ package.json: ${nextVersion}`);

// 2. src/htmx-bolt.ts
const boltPath = path.join(rootDir, "src", "htmx-bolt.ts");
if (fs.existsSync(boltPath)) {
  let bolt = fs.readFileSync(boltPath, "utf8");
  bolt = bolt.replace(/version:\s*['"][^'"]+['"]/, `version: '${nextVersion}'`);
  fs.writeFileSync(boltPath, bolt);
  console.log(`  ✓ src/htmx-bolt.ts: ${nextVersion}`);
}

// 3. schema/hyperstack.json
const schemaPath = path.join(rootDir, "schema", "hyperstack.json");
if (fs.existsSync(schemaPath)) {
  let schema = fs.readFileSync(schemaPath, "utf8");
  schema = schema.replace(/"version":\s*"[^"]+"/, `"version": "${nextVersion}"`);
  fs.writeFileSync(schemaPath, schema);
  console.log(`  ✓ schema/hyperstack.json: ${nextVersion}`);
}

// 4. extension/manifest.json
const manifestPath = path.join(rootDir, "extension", "manifest.json");
if (fs.existsSync(manifestPath)) {
  const manifest = JSON.parse(fs.readFileSync(manifestPath, "utf8"));
  manifest.version = nextVersion;
  fs.writeFileSync(manifestPath, JSON.stringify(manifest, null, 2) + "\n");
  console.log(`  ✓ extension/manifest.json: ${nextVersion}`);
}

// 5. HTML views with version badges
const viewsDir = path.join(rootDir, "views");
if (fs.existsSync(viewsDir)) {
  const viewFiles = fs.readdirSync(viewsDir).filter(f => f.endsWith(".html"));
  for (const f of viewFiles) {
    const p = path.join(viewsDir, f);
    let html = fs.readFileSync(p, "utf8");
    // Replace release badge like v0.10.30 or v1.0.0 (but not CSS/JS URLs or dummy timestamps)
    const updated = html.replace(/(>|\b)v\d+\.\d+\.\d+(<|\b)/g, (match, prefix, suffix) => {
      return `${prefix}v${nextVersion}${suffix}`;
    });
    if (updated !== html) {
      fs.writeFileSync(p, updated);
      console.log(`  ✓ views/${f} badge updated to v${nextVersion}`);
    }
  }
}

// 6. public/docs/index.html
const docsPortalPath = path.join(rootDir, "public", "docs", "index.html");
if (fs.existsSync(docsPortalPath)) {
  let docsHtml = fs.readFileSync(docsPortalPath, "utf8");
  docsHtml = docsHtml.replace(/(>|\b)v\d+\.\d+\.\d+(<|\b)/g, (match, p1, p2) => `${p1}v${nextVersion}${p2}`);
  fs.writeFileSync(docsPortalPath, docsHtml);
  console.log(`  ✓ public/docs/index.html badge updated to v${nextVersion}`);
}

// 7. Rebuild engines & static site
console.log(`\n  Compiling engines and rebuilding static distribution...`);
execSync("bun run build:engines && bun run build:esm && bun run build:css && bun run scripts/build-static.ts", {
  cwd: rootDir,
  stdio: "ignore"
});
console.log(`  ✓ Engines & static site rebuilt.`);

// 8. Calculate exact bundle sizes & sha256 checksums
console.log(`  Recalculating bundle footprints and sha256 checksums...`);
const engines = [
  "public/appz.js",
  "public/htmx-a11y.js",
  "public/htmx-bolt.js",
  "public/htmx-calc.js",
  "public/htmx-canvas.js",
  "public/htmx-devtools.js",
  "public/htmx-flash.js",
  "public/htmx-form.js",
  "public/htmx-grid.js",
  "public/htmx-offline.js",
  "public/htmx-sim.js",
  "public/htmx-vibe.js",
  "public/htmx-virtual.js"
];

const checksumLines = engines.map(rel => {
  const abs = path.join(rootDir, rel);
  const hash = createHash("sha256").update(fs.readFileSync(abs)).digest("hex");
  return `${hash}  ${rel}`;
});
fs.writeFileSync(path.join(rootDir, "CHECKSUMS.sha256"), checksumLines.join("\n") + "\n");
console.log(`  ✓ CHECKSUMS.sha256 synchronized.`);

// 9. Sync docs/BENCHMARKS.md, README.md, and llms.txt footprints
const sizes: Record<string, { raw: number; gzip: number }> = {};
let totalRaw = 0;
let totalGzip = 0;

for (const rel of engines) {
  const fileName = path.basename(rel);
  const buf = fs.readFileSync(path.join(rootDir, rel));
  const gz = gzipSync(buf).length;
  sizes[fileName] = { raw: buf.length, gzip: gz };
  totalRaw += buf.length;
  totalGzip += gz;
}

// Update docs/BENCHMARKS.md table
const benchPath = path.join(rootDir, "docs", "BENCHMARKS.md");
if (fs.existsSync(benchPath)) {
  let doc = fs.readFileSync(benchPath, "utf8");
  for (const [file, sz] of Object.entries(sizes)) {
    const rawFmt = sz.raw.toLocaleString("en-US") + " B";
    const gzipFmt = sz.gzip.toLocaleString("en-US") + " B";
    const re = new RegExp(`(\\|\\s*\\*\\*${file.replace(".", "\\.")}\\*\\*\\s*\\|[^|]*\\|\\s*)([\\d,]+ B)(\\s*\\|\\s*)([\\d,]+ B)(\\s*\\|)`, "g");
    doc = doc.replace(re, `$1${rawFmt}$3${gzipFmt}$5`);
  }
  const totalRawFmt = totalRaw.toLocaleString("en-US") + " B";
  const totalGzipFmt = totalGzip.toLocaleString("en-US") + " B";
  doc = doc.replace(
    /(\|\s*\*\*TOTAL \(13 engines\)\*\*\s*\|[^|]*\|\s*\*\*)([\d,]+ B)(\*\*\s*\|\s*\*\*)([\d,]+ B)(\*\*\s*\|)/,
    `$1${totalRawFmt}$3${totalGzipFmt}$5`
  );
  doc = doc.replace(/current total is \*\*[\d,]+ B\*\*/, `current total is **${totalRawFmt}**`);
  fs.writeFileSync(benchPath, doc);
  console.log(`  ✓ docs/BENCHMARKS.md footprint table synchronized.`);
}

// Update README.md sizes
const readmePath = path.join(rootDir, "README.md");
if (fs.existsSync(readmePath)) {
  let readme = fs.readFileSync(readmePath, "utf8");
  for (const [file, sz] of Object.entries(sizes)) {
    const kb = (Math.round((sz.gzip / 1024) * 100) / 100).toFixed(2);
    const re = new RegExp(`(\\|\\s*\\*\\*[^|]*\\*\\*\\s*\\|\\s*\`[^\`]+\`\\s*\\|\\s*\`${file.replace(".", "\\.")}\`\\s*\\|\\s*)([\\d.]+ KB)(\\s*\\|)`, "g");
    readme = readme.replace(re, `$1${kb} KB$3`);
  }
  const totalRawFmt = totalRaw.toLocaleString("en-US") + " B";
  const totalGzipFmt = totalGzip.toLocaleString("en-US") + " B";
  readme = readme.replace(
    /\*\*Complete 13-Engine Suite\*\*:\s*\*\*[\d,]+ B raw \/ [\d,]+ B gzipped\*\*/,
    `**Complete 13-Engine Suite**: **${totalRawFmt} raw / ${totalGzipFmt} gzipped**`
  );
  fs.writeFileSync(readmePath, readme);
  console.log(`  ✓ README.md sizes synchronized.`);
}

// Update llms.txt sizes
const llmsPath = path.join(rootDir, "llms.txt");
if (fs.existsSync(llmsPath)) {
  let llms = fs.readFileSync(llmsPath, "utf8");
  for (const [file, sz] of Object.entries(sizes)) {
    const kb = (Math.round((sz.gzip / 1024) * 100) / 100).toFixed(2);
    const re = new RegExp(`(\\|\\s*\\*\\*[^*]+\\*\\*\\s*\\|\\s*\`${file.replace(".", "\\.")}\`\\s*\\|\\s*)([\\d.]+KB)(\\s*\\|)`, "g");
    llms = llms.replace(re, `$1${kb}KB$3`);
  }
  fs.writeFileSync(llmsPath, llms);
  console.log(`  ✓ llms.txt sizes synchronized.`);
}

// Re-copy llms.txt to public and dist
fs.copyFileSync(llmsPath, path.join(rootDir, "public", "llms.txt"));
fs.copyFileSync(llmsPath, path.join(rootDir, "dist", "llms.txt"));

// 10. Verify benchmarks test suite
console.log(`\n  Running integrity test suite...`);
try {
  execSync("bun test tests/benchmarks.test.ts", { cwd: rootDir, stdio: "inherit" });
  console.log(`  ✓ All benchmark integrity tests passed with 100% agreement.`);
} catch {
  console.error(`  ❌ Benchmark integrity verification failed.`);
  process.exit(1);
}

if (autoCommit) {
  execSync(`git add -A && git commit -m "chore(release): bump version to v${nextVersion}" && git tag -fa v${nextVersion} -m "Release v${nextVersion}"`, {
    cwd: rootDir,
    stdio: "inherit"
  });
  console.log(`  ✓ Git commit & tag v${nextVersion} created.`);
}

console.log(`\n===============================================================`);
console.log(`  RELEASE v${nextVersion} READY FOR LOCAL VERIFICATION OR PUSH`);
console.log(`===============================================================\n`);
