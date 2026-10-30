import { readFileSync } from 'node:fs';
import { createHash } from 'node:crypto';

const lines = readFileSync('CHECKSUMS.sha256', 'utf-8').trim().split('\n');
for (const line of lines) {
  const [expectedHash, file] = line.trim().split(/\s+/);
  if (!expectedHash || !file) continue;
  const content = readFileSync(file);
  const actualHash = createHash('sha256').update(content).digest('hex');
  if (expectedHash !== actualHash) {
    console.error(`SHA256 mismatch for ${file}:\n  expected: ${expectedHash}\n  actual:   ${actualHash}`);
    process.exit(1);
  }
}
console.log(`Verified ${lines.length} release checksums successfully!`);
