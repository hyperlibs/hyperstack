import { readFileSync, appendFileSync } from 'node:fs';
import { createHash } from 'node:crypto';

const lines = readFileSync('CHECKSUMS.sha256', 'utf-8').trim().split('\n');
let failed = false;
let summary = '### Checksum Verification Results\n\n| File | Status | Expected | Actual |\n| --- | --- | --- | --- |\n';

for (const line of lines) {
  const [expectedHash, file] = line.trim().split(/\s+/);
  if (!expectedHash || !file) continue;
  let actualHash = '';
  try {
    const content = readFileSync(file);
    actualHash = createHash('sha256').update(content).digest('hex');
  } catch (err: any) {
    actualHash = 'FILE_READ_ERROR: ' + err.message;
  }
  if (expectedHash !== actualHash) {
    failed = true;
    console.error(`::error file=${file}::SHA256 mismatch for ${file}: expected ${expectedHash}, got ${actualHash}`);
    summary += `| \`${file}\` | ❌ Failed | \`${expectedHash.slice(0, 10)}...\` | \`${actualHash.slice(0, 10)}...\` |\n`;
  } else {
    summary += `| \`${file}\` | ✅ Match | \`${expectedHash.slice(0, 10)}...\` | \`${actualHash.slice(0, 10)}...\` |\n`;
  }
}

if (process.env.GITHUB_STEP_SUMMARY) {
  try {
    appendFileSync(process.env.GITHUB_STEP_SUMMARY, summary);
  } catch {}
}

if (failed) {
  process.exit(1);
}
console.log(`Verified ${lines.length} release checksums successfully!`);
