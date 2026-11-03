import { readdirSync, readFileSync, writeFileSync } from 'node:fs';

for (const dir of ['public', 'dist']) {
  try {
    for (const f of readdirSync(dir)) {
      if (f.endsWith('.js') || f.endsWith('.mjs')) {
        const p = `${dir}/${f}`;
        const content = readFileSync(p, 'utf8');
        // Normalize any relative or workspace prefix before src/ to canonical "// src/"
        // Ensures identical byte-for-byte bundles across Windows junctions, Linux & macOS CI runners.
        const normalized = content.replace(/\/\/ [^\r\n]*?src\//g, '// src/');
        if (normalized !== content) {
          writeFileSync(p, normalized);
        }
      }
    }
  } catch (e) {}
}
