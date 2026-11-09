import { describe, expect, test } from "bun:test";
import { existsSync, readFileSync } from "node:fs";

describe("Link in Bio Mobile-First Application Suite (Light Theme & PWA)", () => {
  const bioHtmlPath = "views/app-bio.html";

  test("1. views/app-bio.html exists and is non-empty", () => {
    expect(existsSync(bioHtmlPath)).toBe(true);
    const html = readFileSync(bioHtmlPath, "utf-8");
    expect(html.length).toBeGreaterThan(500);
  });

  test("2. Light theme styles and backgrounds are applied", () => {
    const html = readFileSync(bioHtmlPath, "utf-8");
    expect(html).toContain("bg-slate-50");
    expect(html).toContain("text-slate-900");
    expect(html).toContain("bg-white");
  });

  test("3. Core Hyperstack runtime micro-engines are loaded", () => {
    const html = readFileSync(bioHtmlPath, "utf-8");
    expect(html).toContain("/vendor/htmx.min.js");
    expect(html).toContain("/htmx-bolt.js");
    expect(html).toContain("/appz.js");
  });

  test("4. Bio Profile Header includes creator details and verified badge", () => {
    const html = readFileSync(bioHtmlPath, "utf-8");
    expect(html).toContain("Alex Rivera");
    expect(html).toContain("@alex_hyperstack");
    expect(html).toContain("Verified Creator");
  });

  test("5. Curated link items stack exists with destinations", () => {
    const html = readFileSync(bioHtmlPath, "utf-8");
    expect(html).toContain("Hyperstack Framework");
    expect(html).toContain("Hypersheet Live Cloud");
    expect(html).toContain("AppzOS Mobile Platform");
    expect(html).toContain("Hypermedia Discord");
  });

  test("6. One-click share profile action is available", () => {
    const html = readFileSync(bioHtmlPath, "utf-8");
    expect(html).toContain("Share Profile");
    expect(html).toContain("copyBioUrl");
  });

  test("7. Uninstructed clutter features are completely removed", () => {
    const html = readFileSync(bioHtmlPath, "utf-8");
    expect(html).not.toContain("Fuel the Open Source Journey");
    expect(html).not.toContain("Alex Rivera (AMA)");
    expect(html).not.toContain("selectBookingDay");
    expect(html).not.toContain("toggleAudioPreview");
    expect(html).not.toContain("<app-bottom-sheet");
  });

  test("8. 1:1 Component HTML partials exist in views/components/", () => {
    const components = ["bio-profile", "bio-links", "bio-feed", "bio-talk", "bio-navbar"];
    for (const c of components) {
      const p = `views/components/${c}.html`;
      expect(existsSync(p)).toBe(true);
      const content = readFileSync(p, "utf-8");
      expect(content.length).toBeGreaterThan(50);
    }
  });

  test("9. Go DSL parity: go/components/bio.go implements 1:1 constructors", () => {
    const goBioPath = "go/components/bio.go";
    expect(existsSync(goBioPath)).toBe(true);
    const goCode = readFileSync(goBioPath, "utf-8");
    expect(goCode).toContain("func BioProfile(");
    expect(goCode).toContain("func BioLinkCard(");
    expect(goCode).toContain("func BioFeed(");
    expect(goCode).toContain("func BioTalk(");
    expect(goCode).toContain("func BioNavbar(");
    expect(goCode).toContain("func BioPage(");
    expect(goCode).not.toContain("func BioChat(");
    expect(goCode).not.toContain("func BioTipJar(");
  });

  test("10. 3 Tabs under profile: Bio, Feed, and Talk segmented controller", () => {
    const html = readFileSync(bioHtmlPath, "utf-8");
    expect(html).toContain('switchTab(\'bio\')');
    expect(html).toContain('switchTab(\'feed\')');
    expect(html).toContain('switchTab(\'talk\')');
    expect(html).toContain('id="pane-bio"');
    expect(html).toContain('id="pane-feed"');
    expect(html).toContain('id="pane-talk"');
  });

  test("11. Feed tab: 'what i do now' with only text and emojis", () => {
    const html = readFileSync(bioHtmlPath, "utf-8");
    expect(html).toContain("Shipped Hyperstack v1.1 release! Zero-eval signals are 2x faster now 🚀⚡️");
    expect(html).toContain("Brewing fresh Ethiopian natural pour-over ☕️✨");
    expect(html).toContain("Morning 5km run done 🏃‍♂️💨");
    expect(html).toContain("toggleLike");
  });

  test("12. Talk tab: Digital Clone AI agent answering verified inputs & anonymous toggle", () => {
    const html = readFileSync(bioHtmlPath, "utf-8");
    expect(html).toContain("Alex's Digital Clone");
    expect(html).toContain("Ask Anon");
    expect(html).toContain("anon-toggle");
    expect(html).toContain("Do you have a girlfriend?");
    expect(html).toContain("Haha nope, currently single!");
    expect(html).toContain("alexInputs");
  });

  test("13. PWA Manifest & Service Worker exist and are properly configured", () => {
    const manifestPath = "public/manifest.json";
    const swPath = "public/sw.js";
    expect(existsSync(manifestPath)).toBe(true);
    expect(existsSync(swPath)).toBe(true);

    const manifest = JSON.parse(readFileSync(manifestPath, "utf-8"));
    expect(manifest.name).toBe("Alex Rivera • Bio, Feed & Talk");
    expect(manifest.display).toBe("standalone");
    expect(manifest.start_url).toBe("/bio");
    expect(manifest.theme_color).toBe("#f8fafc");

    const sw = readFileSync(swPath, "utf-8");
    expect(sw).toContain("caches.open");
    expect(sw).toContain("/bio");
  });

  test("14. Mobile App Shell UI includes <app-shellui>, <app-bar>, and <app-navbar>", () => {
    const html = readFileSync(bioHtmlPath, "utf-8");
    expect(html).toContain("<app-shellui");
    expect(html).toContain("<app-bar");
    expect(html).toContain("<app-navbar");
    expect(html).toContain('mode="float"');
    expect(html).toContain('behavior="always-visible"');
    expect(html).toContain('id="nav-btn-bio"');
    expect(html).toContain('id="nav-btn-feed"');
    expect(html).toContain('id="nav-btn-talk"');
    expect(html).toContain('id="nav-btn-install"');
  });

  test("15. PWA Install & Offline indicators are implemented in page script", () => {
    const html = readFileSync(bioHtmlPath, "utf-8");
    expect(html).toContain("serviceWorker.register('/sw.js')");
    expect(html).toContain("beforeinstallprompt");
    expect(html).toContain("triggerPwaInstall");
    expect(html).toContain("updateNetworkStatus");
  });
});
