package components

import (
	"strings"
	"testing"
)

func TestBioComponents(t *testing.T) {
	t.Run("BioProfile renders avatar, name, and social handles", func(t *testing.T) {
		profile := BioProfile(BioProfileProps{
			Name:       "Alex Rivera",
			Handle:     "@alex_abc",
			Verified:   true,
			Tagline:    "Systems Architect & Hypermedia Engineer",
			Location:   "San Francisco, CA",
			StatusText: "Building Lean Apps",
			Socials: []SocialLink{
				{Platform: "github", URL: "https://github.com/hyperlibs/hyperstack", Icon: "🐙", Label: "GitHub"},
				{Platform: "x", URL: "https://x.com/", Icon: "𝕏", Label: "X / Twitter"},
			},
		})

		html := profile.Render()
		if !strings.Contains(html, "Alex Rivera") {
			t.Errorf("expected profile to contain name, got %s", html)
		}
		if !strings.Contains(html, "@alex_abc") {
			t.Errorf("expected profile to contain handle, got %s", html)
		}
		if !strings.Contains(html, "Verified Creator") {
			t.Errorf("expected verified badge in profile, got %s", html)
		}
		if !strings.Contains(html, "https://github.com/hyperlibs/hyperstack") {
			t.Errorf("expected github link in profile, got %s", html)
		}
	})

	t.Run("BioLinkCard renders link properties with hx-vibe-card", func(t *testing.T) {
		link := BioLinkCard(BioLinkProps{
			ID:          "link-1",
			Title:       "Hyperstack Core",
			Description: "Lean hypermedia UI framework",
			URL:         "https://github.com/hyperlibs/hyperstack",
			Icon:        "⚡",
			Category:    "projects",
			Badge:       "Open Source",
			Highlighted: true,
			Metrics:     "8.4k stars",
		})

		html := link.Render()
		if !strings.Contains(html, "Hyperstack Core") {
			t.Errorf("expected link card to contain title, got %s", html)
		}
		if !strings.Contains(html, "hx-vibe-card") {
			t.Errorf("expected link card to have hx-vibe-card attribute, got %s", html)
		}
		if !strings.Contains(html, "8.4k stars") {
			t.Errorf("expected link card to contain metrics, got %s", html)
		}
	})

	t.Run("BioChat renders chat stream and form endpoint", func(t *testing.T) {
		chat := BioChat(BioChatProps{
			InitialMessages: []ChatMessage{
				{ID: 1, Sender: "alex", Text: "Hello world!", Timestamp: "12:00"},
				{ID: 2, Sender: "visitor", Text: "Can I book you?", Timestamp: "12:01"},
			},
		})

		html := chat.Render()
		if !strings.Contains(html, "bio-chat-container") {
			t.Errorf("expected container class, got %s", html)
		}
		if !strings.Contains(html, `hx-post="/api/bio/chat"`) {
			t.Errorf("expected hx-post chat endpoint, got %s", html)
		}
		if !strings.Contains(html, "Hello world!") {
			t.Errorf("expected initial message to be rendered, got %s", html)
		}
	})

	t.Run("BioTipJar renders preset amounts and shockwave action", func(t *testing.T) {
		tip := BioTipJar(BioTipJarProps{
			Amounts: []int{3, 5, 15, 50},
		})

		html := tip.Render()
		if !strings.Contains(html, "Fuel the Open Source Journey") {
			t.Errorf("expected tip jar title, got %s", html)
		}
		if !strings.Contains(html, "$blast") {
			t.Errorf("expected $blast physics shockwave action, got %s", html)
		}
	})

	t.Run("BioPage compiles full HTML shell with Hyperstack engines", func(t *testing.T) {
		page := BioPage(BioPageProps{
			Title: "Test Bio",
		})

		html := page.Render()
		if !strings.Contains(html, "<!DOCTYPE html>") {
			t.Errorf("expected doctype in full page, got %s", html)
		}
		if !strings.Contains(html, "/vendor/htmx.min.js") {
			t.Errorf("expected htmx script tag, got %s", html)
		}
		if !strings.Contains(html, "/htmx-bolt.js") {
			t.Errorf("expected htmx-bolt script tag, got %s", html)
		}
		if !strings.Contains(html, "/appz.js") {
			t.Errorf("expected appz script tag, got %s", html)
		}
	})
}
