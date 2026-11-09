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

	t.Run("BioFeed renders creator status updates with text and emojis", func(t *testing.T) {
		feed := BioFeed(BioFeedProps{
			Posts: []BioFeedPost{
				{
					ID:       "test-post-1",
					TimeAgo:  "1h ago",
					Tag:      "Right now",
					Content:  "Shipped Hyperstack v1.1 🚀⚡️",
					Location: "San Francisco",
					Likes:    25,
				},
			},
		})

		html := feed.Render()
		if !strings.Contains(html, "Shipped Hyperstack v1.1 🚀⚡️") {
			t.Errorf("expected feed to contain text and emojis, got %s", html)
		}
		if !strings.Contains(html, "25") {
			t.Errorf("expected likes count, got %s", html)
		}
	})

	t.Run("BioTalk renders Digital Clone AI chat interface with Anonymous toggle", func(t *testing.T) {
		talk := BioTalk(BioTalkProps{
			CloneName: "Alex's Digital Clone",
			Subtitle:  "Responds strictly as per Alex's verified answers",
		})

		html := talk.Render()
		if !strings.Contains(html, "Alex's Digital Clone") {
			t.Errorf("expected clone name, got %s", html)
		}
		if !strings.Contains(html, "Ask Anon") {
			t.Errorf("expected anonymous mode option, got %s", html)
		}
		if !strings.Contains(html, "Do you have a girlfriend? 👀") {
			t.Errorf("expected prompt chip for girlfriend query, got %s", html)
		}
	})

	t.Run("BioPage compiles full mobile-first light HTML shell with 3 tabs", func(t *testing.T) {
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
		if !strings.Contains(html, "Bio") || !strings.Contains(html, "Feed") || !strings.Contains(html, "Talk") {
			t.Errorf("expected 3 tabs (Bio, Feed, Talk), got %s", html)
		}
	})
}
