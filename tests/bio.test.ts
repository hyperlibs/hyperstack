import { describe, expect, test } from "bun:test";
import { existsSync, readFileSync } from "node:fs";

describe("ABC — A Bio Chat Mobile First Application Suite", () => {
  const bioHtmlPath = "views/app-bio.html";

  test("1. views/app-bio.html exists and is non-empty", () => {
    expect(existsSync(bioHtmlPath)).toBe(true);
    const html = readFileSync(bioHtmlPath, "utf-8");
    expect(html.length).toBeGreaterThan(1000);
  });

  test("2. AppzOS mobile custom elements are present", () => {
    const html = readFileSync(bioHtmlPath, "utf-8");
    expect(html).toContain("<app-shellui");
    expect(html).toContain('view-mode="Mob-P"');
    expect(html).toContain("<app-statusbar");
    expect(html).toContain("<app-bar");
    expect(html).toContain("<app-navbar");
    expect(html).toContain("<app-bottom-sheet");
  });

  test("3. Core Hyperstack runtime micro-engines are loaded", () => {
    const html = readFileSync(bioHtmlPath, "utf-8");
    expect(html).toContain("/vendor/htmx.min.js");
    expect(html).toContain("/htmx-bolt.js");
    expect(html).toContain("/htmx-flash.js");
    expect(html).toContain("/appz.js");
    expect(html).toContain("/htmx-vibe.js");
  });

  test("4. Bio Profile Header includes creator details and verified badge", () => {
    const html = readFileSync(bioHtmlPath, "utf-8");
    expect(html).toContain("Alex Rivera");
    expect(html).toContain("@alex_hyperstack");
    expect(html).toContain("Verified Creator");
    expect(html).toContain("Open for Advisory");
  });

  test("5. Real-time 0ms search & category filtering are configured", () => {
    const html = readFileSync(bioHtmlPath, "utf-8");
    expect(html).toContain('placeholder="Search links, articles, projects..."');
    expect(html).toContain("data-cat-btn=\"projects\"");
    expect(html).toContain("data-cat-btn=\"articles\"");
    expect(html).toContain("data-cat-btn=\"media\"");
    expect(html).toContain("data-cat-btn=\"store\"");
  });

  test("6. Bio Chat interactive AMA messenger has input form and socket indicators", () => {
    const html = readFileSync(bioHtmlPath, "utf-8");
    expect(html).toContain('id="chat-messages-stream"');
    expect(html).toContain('id="chat-input-form"');
    expect(html).toContain("Alex Rivera (AMA)");
    expect(html).toContain("typing-indicator");
    expect(html).toContain("Top Projects?");
  });

  test("7. Tip Jar has micro-patronage chips and physics shockwave blast", () => {
    const html = readFileSync(bioHtmlPath, "utf-8");
    expect(html).toContain("Fuel the Open Source Journey");
    expect(html).toContain("tip-btn");
    expect(html).toContain("selectTip(15)");
    expect(html).toContain("$blast");
  });

  test("8. Spring bottom sheets exist for Booking, Tip Jar and QR Share", () => {
    const html = readFileSync(bioHtmlPath, "utf-8");
    expect(html).toContain('id="sheet-booking"');
    expect(html).toContain('id="sheet-tip"');
    expect(html).toContain('id="sheet-share"');
  });

  test("9. 1:1 Component HTML partials exist in views/components/", () => {
    const components = ["bio-profile", "bio-links", "bio-chat", "bio-tipjar"];
    for (const c of components) {
      const p = `views/components/${c}.html`;
      expect(existsSync(p)).toBe(true);
      const content = readFileSync(p, "utf-8");
      expect(content).toContain('hx-ext="reactive"');
    }
  });

  test("10. Go DSL parity: go/components/bio.go implements 1:1 constructors", () => {
    const goBioPath = "go/components/bio.go";
    expect(existsSync(goBioPath)).toBe(true);
    const goCode = readFileSync(goBioPath, "utf-8");
    expect(goCode).toContain("func BioProfile(");
    expect(goCode).toContain("func BioLinkCard(");
    expect(goCode).toContain("func BioChat(");
    expect(goCode).toContain("func BioTipJar(");
    expect(goCode).toContain("func BioPage(");
  });
});
