import { describe, test, expect } from 'bun:test';
import { readFileSync, readdirSync, existsSync } from 'fs';
import { gzipSync } from 'zlib';

/**
 * Keeps docs/BENCHMARKS.md honest.
 *
 * The footprint table in §1 drifted ~21 KB out of date once (htmx-bolt.js was
 * documented at 85.84 KB while the shipped artifact was 107,637 B), so these
 * tests compare the published numbers against the real bundles in public/.
 * Raw sizes are asserted exactly — any rebuild that changes bundle size must
 * come with an updated doc. Gzip sizes are asserted within a small tolerance
 * because zlib output can shift a few bytes across Bun/zlib versions.
 */

const DOC_PATH = 'docs/BENCHMARKS.md';
const GZIP_TOLERANCE = 0.05;

const doc = readFileSync(DOC_PATH, 'utf-8');

/** Engine artifacts that actually ship in public/. */
function shippedEngines(): string[] {
  return readdirSync('public')
    .filter((f) => f === 'appz.js' || /^htmx-.*\.js$/.test(f))
    .sort();
}

interface DocRow {
  file: string;
  raw: number;
  gzip: number;
}

function parseEngineRows(): DocRow[] {
  const rows: DocRow[] = [];
  const re = /^\|\s*\*\*([\w.-]+\.js)\*\*\s*\|[^|]*\|\s*([\d,]+) B\s*\|\s*([\d,]+) B\s*\|\s*$/gm;
  for (const m of doc.matchAll(re)) {
    rows.push({
      file: m[1],
      raw: Number(m[2].replace(/,/g, '')),
      gzip: Number(m[3].replace(/,/g, '')),
    });
  }
  return rows;
}

function parseIntCell(label: string): number {
  const re = new RegExp('\\|\\s*\\*\\*' + label + '\\*\\*\\s*\\|[^|]*\\|\\s*\\*\\*([\\d,]+) B\\*\\*\\s*\\|\\s*\\*\\*([\\d,]+) B\\*\\*\\s*\\|');
  const m = doc.match(re);
  if (!m) throw new Error('Could not find the ' + label + ' row in ' + DOC_PATH);
  return Number(m[1].replace(/,/g, ''));
}

describe('BENCHMARKS.md footprint table matches shipped bundles', () => {
  test('1. every shipped engine appears in the published table', () => {
    const documented = parseEngineRows().map((r) => r.file);
    const shipped = shippedEngines();
    expect(shipped.length).toBe(13);
    expect(documented.sort()).toEqual(shipped);
  });

  test('2. documented raw byte counts equal the actual artifacts', () => {
    for (const row of parseEngineRows()) {
      expect(existsSync('public/' + row.file)).toBe(true);
      const actual = readFileSync('public/' + row.file).length;
      if (actual !== row.raw) {
        throw new Error(
          `${row.file} is ${actual} B but docs/BENCHMARKS.md publishes ${row.raw} B. ` +
          'Rebuild and refresh section 1 of docs/BENCHMARKS.md, README.md and llms.txt with the measured sizes.',
        );
      }
    }
  });

  test('3. documented gzip byte counts match within tolerance', () => {
    for (const row of parseEngineRows()) {
      const actual = gzipSync(readFileSync('public/' + row.file)).length;
      const delta = Math.abs(actual - row.gzip) / actual;
      expect(delta).toBeLessThanOrEqual(GZIP_TOLERANCE);
    }
  });

  test('4. the published total equals the sum of the published rows', () => {
    const rows = parseEngineRows();
    const rawSum = rows.reduce((s, r) => s + r.raw, 0);
    const gzipSum = rows.reduce((s, r) => s + r.gzip, 0);
    expect(parseIntCell('TOTAL \\(13 engines\\)')).toBe(rawSum);
    // Published gzip total must be the sum of the per-row gzip figures.
    const totalRow = doc.match(/\|\s*\*\*TOTAL \(13 engines\)\*\*\s*\|[^|]*\|\s*\*\*[\d,]+ B\*\*\s*\|\s*\*\*([\d,]+) B\*\*\s*\|/);
    expect(totalRow).toBeTruthy();
    expect(Number(totalRow![1].replace(/,/g, ''))).toBe(gzipSum);
  });

  test('5. every shipped engine is covered by CHECKSUMS.sha256', () => {
    const sums = readFileSync('CHECKSUMS.sha256', 'utf-8');
    const listed = sums
      .trim()
      .split('\n')
      .map((l) => l.trim().split(/\s+/)[1]?.replace(/^\.\//, ''))
      .filter(Boolean);
    for (const file of shippedEngines()) {
      expect(listed).toContain('public/' + file);
    }
    expect(listed.length).toBe(shippedEngines().length);
  });

  test('6. the README engine table agrees with the published footprint table', () => {
    const readme = readFileSync('README.md', 'utf-8');
    const re = /^\|\s*\*\*[^|]*\*\*\s*\|\s*`[^`]+`\s*\|\s*`([\w.-]+\.js)`\s*\|\s*([\d.]+) KB\s*\|/gm;
    const rows = [...readme.matchAll(re)].map((m) => ({ file: m[1], gzipKb: Number(m[2]) }));
    const published = parseEngineRows();
    expect(rows.length).toBe(published.length);
    for (const row of rows) {
      const doc = published.find((p) => p.file === row.file);
      expect(doc).toBeTruthy();
      const expected = Math.round((doc!.gzip / 1024) * 100) / 100;
      expect(row.gzipKb).toBe(expected);
    }
  });

  test('7. the llms.txt engine table agrees with the published footprint table', () => {
    const llms = readFileSync('llms.txt', 'utf-8');
    const re = /^\|\s*\*\*[^*]+\*\*\s*\|\s*`([\w.-]+\.js)`\s*\|\s*([\d.]+)KB\s*\|/gm;
    const rows = [...llms.matchAll(re)].map((m) => ({ file: m[1], gzipKb: Number(m[2]) }));
    const published = parseEngineRows();
    expect(rows.length).toBe(published.length);
    for (const row of rows) {
      const doc = published.find((p) => p.file === row.file);
      expect(doc).toBeTruthy();
      expect(Math.abs(row.gzipKb - doc!.gzip / 1024)).toBeLessThanOrEqual(0.01);
    }
  });

  test('8. the served and built llms.txt copies are identical to the root source', () => {
    const source = readFileSync('llms.txt');
    for (const copy of ['public/llms.txt', 'dist/llms.txt']) {
      if (!existsSync(copy)) continue;
      expect(Buffer.compare(source, readFileSync(copy))).toBe(0);
    }
  });

  test('9. the static build output is deterministic', () => {
    // The mock invoices used an unseeded PRNG plus Date.now(), which made every
    // build emit a different dist/api/erp-invoices.json. The values below are the
    // frozen seeded output; if the generator changes, update them deliberately.
    if (!existsSync('dist/api/erp-invoices.json')) return;
    const invoices = JSON.parse(readFileSync('dist/api/erp-invoices.json', 'utf-8'));
    expect(invoices.length).toBe(500);
    expect(invoices[0]).toEqual({ id: 'INV-100000', customer: 'Acme Corp', amount: '3871.13', status: 'Paid', date: '2026-01-01' });
    expect(invoices[499]).toEqual({ id: 'INV-100499', customer: 'Umbrella Co', amount: '3653.51', status: 'Processing', date: '2024-08-20' });
  });

  test('10. the documented CI bundle budget matches the workflow threshold', () => {
    const ci = readFileSync('.github/workflows/ci.yml', 'utf-8');
    const m = ci.match(/total_size[^\n]*-gt\s+(\d+)/);
    expect(m).toBeTruthy();
    const budget = Number(m![1]);
    // The doc quotes the CI ceiling (350,000 B) and the current total (295,980 B).
    expect(doc).toContain('**350,000 B**');
    const total = parseIntCell('TOTAL \\(13 engines\\)');
    expect(total).toBeLessThanOrEqual(budget);
  });
});
