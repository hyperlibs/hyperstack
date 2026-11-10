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

	t.Run("BioDetails renders relation with you, hobbies, interests, goals, achievements", func(t *testing.T) {
		details := BioDetails(BioDetailsProps{
			RelationTitle: "Relation with You",
			RelationBadge: "Close Friend 🤝",
			RelationDesc:  "You and Alex are close friends.",
			Hobbies:       []string{"☕️ Pour-over coffee", "🏃‍♂️ Morning 5km runs"},
			Interests:     []string{"⚡ Hypermedia", "🌐 Distributed Systems"},
			Goals:         []string{"Ship Hyperstack v2.0"},
			Achievements: []AchievementItem{
				{Icon: "⚡", Title: "Created Hyperstack Framework", Description: "Zero Node.js footprint."},
			},
		})

		html := details.Render()
		if !strings.Contains(html, "Relation with You") {
			t.Errorf("expected Relation with You card, got %s", html)
		}
		if !strings.Contains(html, "Close Friend 🤝") {
			t.Errorf("expected relation badge, got %s", html)
		}
		if !strings.Contains(html, "Hobbies & Rituals") {
			t.Errorf("expected hobbies section, got %s", html)
		}
		if !strings.Contains(html, "☕️ Pour-over coffee") {
			t.Errorf("expected hobby item, got %s", html)
		}
		if !strings.Contains(html, "Interests") {
			t.Errorf("expected interests section, got %s", html)
		}
		if !strings.Contains(html, "2026 Goals") {
			t.Errorf("expected goals section, got %s", html)
		}
		if !strings.Contains(html, "Achievements") {
			t.Errorf("expected achievements section, got %s", html)
		}
		if !strings.Contains(html, "Created Hyperstack Framework") {
			t.Errorf("expected achievement item, got %s", html)
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

	t.Run("BioTalk renders Digital Clone AI chat interface with Offline LLM badge, Teach drawer & Anonymous toggle", func(t *testing.T) {
		talk := BioTalk(BioTalkProps{
			CloneName: "Alex's Digital Clone",
			Subtitle:  "Responds strictly as per Alex's verified answers",
		})

		html := talk.Render()
		if !strings.Contains(html, "Alex's Digital Clone") {
			t.Errorf("expected clone name, got %s", html)
		}
		if !strings.Contains(html, "Offline LLM") {
			t.Errorf("expected Offline LLM badge, got %s", html)
		}
		if !strings.Contains(html, "Teach") {
			t.Errorf("expected teach button, got %s", html)
		}
		if !strings.Contains(html, "teach-drawer") {
			t.Errorf("expected teach drawer, got %s", html)
		}
		if !strings.Contains(html, "Ask Anon") {
			t.Errorf("expected anonymous mode option, got %s", html)
		}
		if !strings.Contains(html, "Do you have a girlfriend? 👀") {
			t.Errorf("expected prompt chip for girlfriend query, got %s", html)
		}
	})

	t.Run("BioInvite renders collections and access rule controls", func(t *testing.T) {
		invite := BioInvite(BioInviteProps{
			Collections: []CollectionRuleItem{
				{ID: "friends", Name: "Friends", Bio: true, Links: true, Feed: true, Talk: true},
				{ID: "work", Name: "Work", Bio: true, Links: true, Feed: false, Talk: false},
			},
		})

		html := invite.Render()
		if !strings.Contains(html, "Invite & Collections") {
			t.Errorf("expected Invite & Collections header, got %s", html)
		}
		if !strings.Contains(html, "Friends") {
			t.Errorf("expected Friends collection, got %s", html)
		}
		if !strings.Contains(html, "Work") {
			t.Errorf("expected Work collection, got %s", html)
		}
		if !strings.Contains(html, "Generate & Copy Invite Link") {
			t.Errorf("expected generate invite button, got %s", html)
		}
	})

	t.Run("BioNavbar renders PWA <app-navbar> with 2x increased icons & text and 5 tabs", func(t *testing.T) {
		navbar := BioNavbar(BioNavbarProps{
			Mode:   "float",
			Active: "bio",
		})

		html := navbar.Render()
		if !strings.Contains(html, "<app-navbar") {
			t.Errorf("expected <app-navbar> custom element, got %s", html)
		}
		if !strings.Contains(html, "mode=\"float\"") {
			t.Errorf("expected mode=float, got %s", html)
		}
		if !strings.Contains(html, "text-2xl") {
			t.Errorf("expected 2x increased icon size text-2xl, got %s", html)
		}
		if !strings.Contains(html, "text-xs font-bold") {
			t.Errorf("expected 2x increased font size text-xs font-bold, got %s", html)
		}
		if !strings.Contains(html, "switchTab('bio')") {
			t.Errorf("expected switchTab('bio'), got %s", html)
		}
		if !strings.Contains(html, "switchTab('links')") {
			t.Errorf("expected switchTab('links'), got %s", html)
		}
		if !strings.Contains(html, "switchTab('feed')") {
			t.Errorf("expected switchTab('feed'), got %s", html)
		}
		if !strings.Contains(html, "switchTab('talk')") {
			t.Errorf("expected switchTab('talk'), got %s", html)
		}
		if !strings.Contains(html, "openInviteModal()") || !strings.Contains(html, "Invite") {
			t.Errorf("expected Invite tab in navbar, got %s", html)
		}
	})

	t.Run("BioPage compiles full mobile-first light HTML shell with PWA manifest & <app-shellui>", func(t *testing.T) {
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
		if !strings.Contains(html, "/manifest.json") {
			t.Errorf("expected PWA manifest link, got %s", html)
		}
		if !strings.Contains(html, "/tiny-llm.js") {
			t.Errorf("expected tiny-llm.js in full page, got %s", html)
		}
		if !strings.Contains(html, "<app-shellui") {
			t.Errorf("expected <app-shellui>, got %s", html)
		}
		if !strings.Contains(html, "<app-navbar") {
			t.Errorf("expected <app-navbar>, got %s", html)
		}
		if !strings.Contains(html, "pane-bio") {
			t.Errorf("expected pane-bio, got %s", html)
		}
		if !strings.Contains(html, "pane-links") {
			t.Errorf("expected pane-links, got %s", html)
		}
		if !strings.Contains(html, "pane-feed") {
			t.Errorf("expected pane-feed, got %s", html)
		}
		if !strings.Contains(html, "pane-talk") {
			t.Errorf("expected pane-talk, got %s", html)
		}
		if !strings.Contains(html, "invite-modal") {
			t.Errorf("expected invite-modal, got %s", html)
		}
	})
}
