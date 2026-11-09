import { describe, expect, test } from "bun:test";
import { existsSync, readFileSync } from "node:fs";

describe("Link in Bio Mobile-First Application Suite (Light Theme)", () => {
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
    const components = ["bio-profile", "bio-links"];
    for (const c of components) {
      const p = `views/components/${c}.html`;
      expect(existsSync(p)).toBe(true);
      const content = readFileSync(p, "utf-8");
      expect(content).toContain('hx-ext="reactive"');
    }
  });

  test("9. Go DSL parity: go/components/bio.go implements 1:1 constructors", () => {
    const goBioPath = "go/components/bio.go";
    expect(existsSync(goBioPath)).toBe(true);
    const goCode = readFileSync(goBioPath, "utf-8");
    expect(goCode).toContain("func BioProfile(");
    expect(goCode).toContain("func BioLinkCard(");
    expect(goCode).toContain("func BioPage(");
    expect(goCode).not.toContain("func BioChat(");
    expect(goCode).not.toContain("func BioTipJar(");
  });
});
