import { spawn } from "child_process";
import { writeFileSync, mkdirSync } from "fs";
import { join } from "path";

const ARTIFACT_DIR = "C:\\Users\\david\\.gemini\\antigravity\\brain\\1b368cdc-48e1-442c-aacc-eac9aeb2b87b\\screenshots";
mkdirSync(ARTIFACT_DIR, { recursive: true });

const edgePath = "C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe";
const port = 9224;
const profileDir = `C:\\Users\\david\\AppData\\Local\\Temp\\edge_cdp_profile_${Date.now()}`;

const proc = spawn(edgePath, [
  "--headless=new",
  `--remote-debugging-port=${port}`,
  "--disable-gpu",
  "--no-first-run",
  "--no-default-browser-check",
  `--user-data-dir=${profileDir}`,
  "about:blank"
], { detached: true });

// Poll until CDP server is ready
let ready = false;
for (let i = 0; i < 20; i++) {
  try {
    const res = await fetch(`http://localhost:${port}/json/version`);
    if (res.ok) {
      ready = true;
      break;
    }
  } catch {}
  await new Promise(r => setTimeout(r, 300));
}

if (!ready) {
  proc.kill("SIGKILL");
  throw new Error(`Edge CDP did not start on port ${port}`);
}

class CDPClient {
  ws: WebSocket;
  id = 0;
  callbacks = new Map<number, (res: any) => void>();

  constructor(wsUrl: string) {
    this.ws = new WebSocket(wsUrl);
  }

  async connect() {
    return new Promise((resolve) => {
      this.ws.onopen = resolve;
      this.ws.onmessage = (event) => {
        const msg = JSON.parse(event.data.toString());
        if (msg.id && this.callbacks.has(msg.id)) {
          const cb = this.callbacks.get(msg.id)!;
          this.callbacks.delete(msg.id);
          cb(msg.result);
        }
      };
    });
  }

  async send(method: string, params: any = {}): Promise<any> {
    const id = ++this.id;
    return new Promise((resolve) => {
      this.callbacks.set(id, resolve);
      this.ws.send(JSON.stringify({ id, method, params }));
    });
  }

  async evaluate(expression: string) {
    return this.send("Runtime.evaluate", { expression, awaitPromise: true, returnByValue: true });
  }

  async screenshot(filename: string) {
    await new Promise(r => setTimeout(r, 600));
    const res = await this.send("Page.captureScreenshot", { format: "png", fromSurface: true });
    const buffer = Buffer.from(res.data, "base64");
    const outPath = join(ARTIFACT_DIR, filename);
    writeFileSync(outPath, buffer);
    console.log(`Saved screenshot: ${filename}`);
  }
}

try {
  const resp = await fetch(`http://localhost:${port}/json`);
  const tabs = await resp.json();
  const pageTab = tabs.find((t: any) => t.type === "page") || tabs[0];
  const client = new CDPClient(pageTab.webSocketDebuggerUrl);
  await client.connect();

  await client.send("Page.enable");
  await client.send("DOM.enable");
  await client.send("Emulation.setDeviceMetricsOverride", {
    width: 412,
    height: 915,
    deviceScaleFactor: 2,
    mobile: true
  });

  // Navigate to /bio
  await client.send("Page.navigate", { url: "http://localhost:3000/bio" });
  await new Promise(r => setTimeout(r, 2000));

  // 1. Bio Public
  await client.evaluate(`switchTab('bio')`);
  await client.screenshot("01_bio_public.png");

  // 2. Links Public
  await client.evaluate(`switchTab('links')`);
  await client.screenshot("02_links_public.png");

  // 3. Feed Locked (Public)
  await client.evaluate(`switchTab('feed')`);
  await client.screenshot("03_feed_locked.png");

  // 4. Talk (Public Visitor)
  await client.evaluate(`switchTab('talk')`);
  await client.screenshot("04_talk_public.png");

  // Switch persona to BFFs
  await client.evaluate(`changePersona('bffs')`);
  // Wait for role toast to settle
  await new Promise(r => setTimeout(r, 2500));

  // 5. Bio with BFFs relation
  await client.evaluate(`switchTab('bio')`);
  await client.screenshot("05_bio_bffs.png");

  // 6. Feed Unlocked
  await client.evaluate(`switchTab('feed')`);
  await client.screenshot("06_feed_unlocked.png");

  // 7. Talk Main (WhatsApp Chat Stream)
  await client.evaluate(`switchTab('talk'); switchTalkSubTab('talk')`);
  await client.screenshot("07_talk_main.png");

  // 8. Talk Back Sub-Tab (Thread replies station)
  await client.evaluate(`switchTalkSubTab('talkback')`);
  await client.screenshot("08_talk_talkback.png");

  // 9. Teach Drawer Open
  await client.evaluate(`switchTalkSubTab('talk'); toggleTeachDrawer()`);
  await client.screenshot("09_teach_drawer.png");
  await client.evaluate(`toggleTeachDrawer()`);

  // 10. Inbox Drawer Open
  await client.evaluate(`toggleInboxDrawer()`);
  await client.screenshot("10_inbox_drawer.png");
  await client.evaluate(`toggleInboxDrawer()`);

  // 11. Username Modal Open
  await client.evaluate(`openUsernameModal()`);
  await client.screenshot("11_username_modal.png");
  await client.evaluate(`closeUsernameModal()`);

  // 12. Invite Modal Open
  await client.evaluate(`openInviteModal()`);
  await client.screenshot("12_invite_modal.png");
  await client.evaluate(`closeInviteModal()`);

  // 13. Trigger No Credits Error Banner in subpane-talk
  await client.evaluate(`
    switchTab('talk');
    switchTalkSubTab('talk');
    if (window.$llm) {
      window.$llm.setUserCredits('visitor', 0);
      updateCreditsDisplay();
      showNoCreditsErrorBanner();
    }
  `);
  await client.screenshot("13_no_credits_banner.png");

  console.log("All screenshots successfully captured!");
} catch (e) {
  console.error("Screenshot capture failed:", e);
} finally {
  proc.kill("SIGKILL");
}
