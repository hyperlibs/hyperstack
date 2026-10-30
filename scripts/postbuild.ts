import { readdirSync, readFileSync, writeFileSync } from 'node:fs';

for (const dir of ['public', 'dist']) {
  try {
    for (const f of readdirSync(dir)) {
      if (f.endsWith('.js') || f.endsWith('.mjs')) {
        const p = `${dir}/${f}`;
        const content = readFileSync(p, 'utf8');
        if (content.includes('../htmxpro/')) {
          writeFileSync(p, content.replaceAll('../htmxpro/', 'hyperstack/'));
        }
      }
    }
  } catch (e) {}
}
