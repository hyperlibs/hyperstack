package components

import (
	"strings"
	"testing"
)

func TestBioComponents(t *testing.T) {
	t.Run("BioProfile renders avatar, name, handle, and social links in light theme", func(t *testing.T) {
		profile := BioProfile(BioProfileProps{
			Name:     "Alex Rivera",
			Handle:   "@alex_hyperstack",
			Verified: true,
			Bio:      "Systems Architect & Open Source Creator",
			Socials: []SocialLink{
				{Platform: "github", URL: "https://github.com/hyperlibs/hyperstack", Icon: "🐙", Label: "GitHub"},
				{Platform: "x", URL: "https://x.com/", Icon: "𝕏", Label: "X / Twitter"},
			},
		})

		html := profile.Render()
		if !strings.Contains(html, "Alex Rivera") {
			t.Errorf("expected profile to contain name, got %s", html)
		}
		if !strings.Contains(html, "@alex_hyperstack") {
			t.Errorf("expected profile to contain handle, got %s", html)
		}
		if !strings.Contains(html, "Verified Creator") {
			t.Errorf("expected verified badge, got %s", html)
		}
		if !strings.Contains(html, "bg-white") {
			t.Errorf("expected light theme white card background, got %s", html)
		}
	})

	t.Run("BioLinkCard renders link with destination and badge", func(t *testing.T) {
		link := BioLinkCard(BioLinkProps{
			ID:          "link-1",
			Title:       "Hyperstack Framework",
			Description: "HTMX on Steroids",
			URL:         "https://github.com/hyperlibs/hyperstack",
			Icon:        "⚡",
			Badge:       "v1.0",
		})

		html := link.Render()
		if !strings.Contains(html, "Hyperstack Framework") {
			t.Errorf("expected link card to contain title, got %s", html)
		}
		if !strings.Contains(html, "https://github.com/hyperlibs/hyperstack") {
			t.Errorf("expected link card to contain href, got %s", html)
		}
		if !strings.Contains(html, "bg-white") {
			t.Errorf("expected light theme white card, got %s", html)
		}
	})

	t.Run("BioPage compiles full mobile-first light HTML shell", func(t *testing.T) {
		page := BioPage(BioPageProps{
			Title: "Test Bio",
		})

		html := page.Render()
		if !strings.Contains(html, "<!DOCTYPE html>") {
			t.Errorf("expected doctype in full page, got %s", html)
		}
		if !strings.Contains(html, "bg-slate-50") {
			t.Errorf("expected light background class, got %s", html)
		}
		if !strings.Contains(html, "/vendor/htmx.min.js") {
			t.Errorf("expected htmx script tag, got %s", html)
		}
	})
}
