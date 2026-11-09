package components

import (
	"fmt"
	"strings"
)

// =============================================================================
// LINK IN BIO: GO COMPONENT DSL
// Pure Mobile-First Link-in-Bio Application using Hyperstack (Ultra-Compact Light Theme & PWA)
// Conforms to Hyperstack Component Standards & 1:1 Parity Rule
// Tabs: Bio | Feed | Talk (Digital Clone AI Agent) with <app-navbar> PWA Navigation
// =============================================================================

// --- 1. Bio Profile Header (Ultra-Compact) ---

type SocialLink struct {
	Platform string
	URL      string
	Icon     string
	Label    string
}

type BioProfileProps struct {
	AvatarURL string
	Name      string
	Handle    string
	Verified  bool
	Bio       string
	Socials   []SocialLink
}

func BioProfile(p BioProfileProps) Component {
	if p.AvatarURL == "" {
		p.AvatarURL = "https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&w=140&q=80"
	}
	if p.Name == "" {
		p.Name = "Alex Rivera"
	}
	if p.Handle == "" {
		p.Handle = "@alex_hyperstack"
	}
	if p.Bio == "" {
		p.Bio = "Systems Architect • Building Hyperstack (Zero VDOM)"
	}

	verifiedBadge := ""
	if p.Verified {
		verifiedBadge = `<span class="w-3.5 h-3.5 rounded-full bg-blue-500 text-white text-[8px] font-bold absolute -bottom-0.5 -right-0.5 flex items-center justify-center border border-white" title="Verified Creator">✓</span>`
	}

	var socialsSb strings.Builder
	for _, s := range p.Socials {
		socialsSb.WriteString(fmt.Sprintf(`<a href="%s" target="_blank" rel="noopener noreferrer" class="hover:text-slate-900 transition-colors" title="%s">%s %s</a>`, s.URL, s.Label, s.Icon, s.Label))
	}

	html := fmt.Sprintf(`
<div class="p-3 bg-white border border-slate-200/80 rounded-2xl shadow-xs">
  <div class="flex items-center gap-3">
    <div class="relative shrink-0">
      <img src="%s" alt="%s" class="w-12 h-12 rounded-full object-cover border border-slate-200" />
      %s
    </div>

    <div class="flex-1 min-w-0">
      <div class="flex items-center justify-between">
        <h1 class="text-sm font-bold text-slate-900 truncate">%s</h1>
        <button onclick="navigator.clipboard.writeText(window.location.href); alert('Profile link copied!');" 
                title="Share Profile"
                class="text-[11px] text-slate-500 hover:text-slate-900 px-2 py-0.5 rounded-md hover:bg-slate-100 transition-colors flex items-center gap-1">
          <span>🔗</span> <span>Share Profile</span>
        </button>
      </div>
      <p class="text-[11px] font-mono text-indigo-600 font-medium">%s</p>
      <p class="text-[11px] text-slate-600 truncate mt-0.5">%s</p>
    </div>
  </div>

  <div class="mt-2.5 pt-2 border-t border-slate-100 flex items-center justify-between text-xs text-slate-500">
    <div class="flex items-center gap-3">
      %s
    </div>
    <a href="mailto:alex@hyperstack.dev" class="text-indigo-600 font-medium hover:underline text-[11px]">alex@hyperstack.dev</a>
  </div>
</div>`, p.AvatarURL, p.Name, verifiedBadge, p.Name, p.Handle, p.Bio, socialsSb.String())

	return Component{html: html}
}

// --- 2. Bio Link Card (Ultra-Compact) ---

type BioLinkProps struct {
	ID          string
	Title       string
	Description string
	URL         string
	Icon        string
	Badge       string
}

func BioLinkCard(p BioLinkProps) Component {
	badgeHtml := ""
	if p.Badge != "" {
		badgeHtml = fmt.Sprintf(`<span class="text-[9px] font-bold px-1.5 py-0.2 rounded bg-indigo-50 text-indigo-700 border border-indigo-200">%s</span>`, p.Badge)
	}

	descHtml := ""
	if p.Description != "" {
		descHtml = fmt.Sprintf(`<p class="text-[10px] text-slate-500 truncate">%s</p>`, p.Description)
	}

	html := fmt.Sprintf(`
<a href="%s" target="_blank" rel="noopener noreferrer"
   class="bio-link-item group flex items-center gap-2.5 px-3 py-2 rounded-xl bg-white border border-slate-200/80 hover:border-indigo-400 hover:bg-slate-50/50 transition-all active:scale-[0.99] shadow-2xs">
  <span class="text-base shrink-0">%s</span>
  <div class="flex-1 min-w-0 text-left">
    <div class="flex items-center gap-1.5">
      <h2 class="font-bold text-xs text-slate-900 group-hover:text-indigo-600 transition-colors truncate">%s</h2>
      %s
    </div>
    %s
  </div>
  <span class="text-slate-300 group-hover:text-indigo-600 text-xs transition-colors shrink-0">↗</span>
</a>`, p.URL, p.Icon, p.Title, badgeHtml, descHtml)

	return Component{html: html}
}

// --- 3. Bio Feed Component (Ultra-Compact Text & Emojis Only) ---

type BioFeedPost struct {
	ID       string
	TimeAgo  string
	Tag      string
	Content  string
	Location string
	Likes    int
}

type BioFeedProps struct {
	Posts []BioFeedPost
}

func BioFeed(p BioFeedProps) Component {
	if len(p.Posts) == 0 {
		p.Posts = []BioFeedPost{
			{
				ID:       "post-1",
				TimeAgo:  "2h ago",
				Tag:      "Right now",
				Content:  "Shipped Hyperstack v1.1 release! Zero-eval signals are 2x faster now 🚀⚡️ Benchmarks looking super clean 📈✨",
				Location: "San Francisco, CA",
				Likes:    24,
			},
			{
				ID:       "post-2",
				TimeAgo:  "6h ago",
				Tag:      "Earlier today",
				Content:  "Brewing fresh Ethiopian natural pour-over ☕️✨ Reading distributed consensus papers on Sunday morning 📖🎧",
				Location: "Mission District",
				Likes:    19,
			},
			{
				ID:       "post-3",
				TimeAgo:  "Yesterday",
				Tag:      "Activity",
				Content:  "Morning 5km run done 🏃‍♂️💨 Beautiful sunny morning over the bay ☀️🌉 Feeling energized for deep work 🔥💻",
				Location: "Embarcadero",
				Likes:    42,
			},
		}
	}

	var postsSb strings.Builder
	for _, post := range p.Posts {
		tagColor := "text-slate-400"
		if strings.Contains(strings.ToLower(post.Tag), "now") {
			tagColor = "text-indigo-600 font-semibold"
		}

		postsSb.WriteString(fmt.Sprintf(`
  <article class="p-2.5 rounded-xl bg-white border border-slate-200/80 shadow-2xs space-y-1">
    <div class="flex items-center justify-between text-[10px] text-slate-400">
      <span class="%s">%s</span>
      <span class="font-mono">%s</span>
    </div>
    <p class="text-xs text-slate-800 leading-snug font-normal">
      %s
    </p>
    <div class="pt-1 flex items-center justify-between text-[10px] text-slate-400 border-t border-slate-100">
      <span>📍 %s</span>
      <button class="flex items-center gap-1 text-[10px] text-slate-400 hover:text-rose-500 transition-colors">
        <span>🤍</span> <span>%d</span>
      </button>
    </div>
  </article>`, tagColor, post.Tag, post.TimeAgo, post.Content, post.Location, post.Likes))
	}

	html := fmt.Sprintf(`<section class="space-y-1.5">%s</section>`, postsSb.String())
	return Component{html: html}
}

// --- 4. Bio Talk Component (Ultra-Compact Digital Clone AI Agent) ---

type BioTalkProps struct {
	CloneName    string
	Subtitle     string
	Greeting     string
	QuickPrompts []string
}

func BioTalk(p BioTalkProps) Component {
	if p.CloneName == "" {
		p.CloneName = "Alex's Clone"
	}
	if p.Subtitle == "" {
		p.Subtitle = "Verified answers"
	}
	if p.Greeting == "" {
		p.Greeting = "Hey! 👋 I'm <strong>Alex's Digital Clone AI</strong>. I answer strictly based on Alex's own personal inputs."
	}
	if len(p.QuickPrompts) == 0 {
		p.QuickPrompts = []string{
			"Do you have a girlfriend? 👀",
			"Top projects? 🚀",
			"Where do you live? 📍",
		}
	}

	var chipsSb strings.Builder
	for _, prompt := range p.QuickPrompts {
		chipsSb.WriteString(fmt.Sprintf(`
    <button class="px-2 py-0.5 rounded-full bg-slate-100 hover:bg-indigo-50 hover:text-indigo-600 border border-slate-200 text-[10px] font-medium text-slate-700 shrink-0 transition-colors">
      %s
    </button>`, prompt))
	}

	html := fmt.Sprintf(`
<div class="flex flex-col h-[380px] rounded-2xl bg-white border border-slate-200/80 overflow-hidden shadow-2xs">
  <div class="h-10 px-3 bg-slate-50 border-b border-slate-200/70 flex items-center justify-between shrink-0">
    <div class="flex items-center gap-1.5">
      <div class="w-5 h-5 rounded-full bg-gradient-to-tr from-indigo-500 to-purple-600 flex items-center justify-center text-white text-[9px] font-bold">
        A
      </div>
      <span class="text-xs font-bold text-slate-900">%s</span>
      <span class="text-[9px] text-indigo-600 font-mono font-bold bg-indigo-50 px-1 rounded">AI</span>
    </div>

    <label class="flex items-center gap-1 cursor-pointer select-none">
      <input type="checkbox" id="anon-checkbox" class="rounded border-slate-300 text-indigo-600 focus:ring-indigo-500 h-3 w-3" />
      <span class="text-[10px] font-medium text-slate-600 flex items-center gap-0.5">
        <span>🎭</span> <span>Ask Anon</span>
      </span>
    </label>
  </div>

  <div class="flex-1 overflow-y-auto p-2.5 space-y-2 bg-slate-50/40 text-xs">
    <div class="flex items-start gap-1.5 max-w-[92%%]">
      <div class="w-5 h-5 rounded-full bg-indigo-600 flex items-center justify-center text-white text-[9px] font-bold shrink-0">A</div>
      <div class="p-2 rounded-xl rounded-tl-xs bg-white border border-slate-200/80 text-xs text-slate-800 leading-snug shadow-2xs">
        %s
        <div class="text-[8px] text-slate-400 text-right mt-0.5 font-mono">09:41</div>
      </div>
    </div>
  </div>

  <div class="px-2 py-1 bg-white border-t border-slate-200/70 flex items-center gap-1 overflow-x-auto shrink-0 scrollbar-none">
    %s
  </div>

  <div class="p-1.5 bg-white border-t border-slate-200/70 flex items-center gap-1.5 shrink-0">
    <input type="text" placeholder="Ask Alex's clone anything..." class="flex-1 bg-slate-50 border border-slate-200/80 rounded-lg px-2.5 py-1 text-xs text-slate-900 placeholder:text-slate-400 focus:outline-none focus:border-indigo-500 shadow-2xs" />
    <button class="px-3 py-1 rounded-lg bg-slate-900 hover:bg-slate-800 text-white font-semibold text-xs transition-all">
      ➤
    </button>
  </div>
</div>`, p.CloneName, p.Greeting, chipsSb.String())

	return Component{html: html}
}

// --- 5. Bio Navbar Component (Ultra-Compact <app-navbar>) ---

type BioNavItem struct {
	ID     string
	Label  string
	Icon   string
	Active bool
}

type BioNavbarProps struct {
	Mode     string
	Behavior string
	Active   string
	Items    []BioNavItem
}

func BioNavbar(p BioNavbarProps) Component {
	if p.Mode == "" {
		p.Mode = "float"
	}
	if p.Behavior == "" {
		p.Behavior = "always-visible"
	}
	if p.Active == "" {
		p.Active = "bio"
	}
	if len(p.Items) == 0 {
		p.Items = []BioNavItem{
			{ID: "bio", Label: "Bio", Icon: "🔗", Active: p.Active == "bio"},
			{ID: "feed", Label: "Feed", Icon: "📝", Active: p.Active == "feed"},
			{ID: "talk", Label: "Talk", Icon: "💬", Active: p.Active == "talk"},
			{ID: "install", Label: "Install", Icon: "📲", Active: false},
		}
	}

	var itemsSb strings.Builder
	for _, it := range p.Items {
		activeClass := "text-slate-400 hover:text-slate-700"
		if it.Active {
			activeClass = "text-indigo-600 font-bold bg-indigo-50/80"
		}
		itemsSb.WriteString(fmt.Sprintf(`
      <button onclick="switchTab('%s')" id="nav-btn-%s" class="nav-tab-item flex flex-col items-center gap-0.5 py-1 px-3 rounded-xl transition-all %s">
        <span class="text-sm leading-none">%s</span>
        <span class="text-[9px] tracking-tight">%s</span>
      </button>`, it.ID, it.ID, activeClass, it.Icon, it.Label))
	}

	html := fmt.Sprintf(`
<app-navbar mode="%s" behavior="%s" class="fixed bottom-2.5 inset-x-0 mx-auto max-w-sm px-3 z-40 flex items-center justify-between">
  <div class="w-full bg-white/95 backdrop-blur-md border border-slate-200/90 rounded-2xl shadow-lg shadow-slate-900/5 h-13 px-2 flex items-center justify-around">
    %s
  </div>
</app-navbar>`, p.Mode, p.Behavior, itemsSb.String())

	return Component{html: html}
}

// --- 6. Full Mobile-First Link-in-Bio Page Component ---

type BioPageProps struct {
	Title   string
	Profile BioProfileProps
	Links   []BioLinkProps
	Feed    BioFeedProps
	Talk    BioTalkProps
	Navbar  BioNavbarProps
}

func BioPage(p BioPageProps) Component {
	if p.Title == "" {
		p.Title = "Alex Rivera • Bio, Feed & Talk"
	}

	profileComp := BioProfile(p.Profile)
	navbarComp := BioNavbar(p.Navbar)

	var linksSb strings.Builder
	for _, l := range p.Links {
		linksSb.WriteString(BioLinkCard(l).Render())
	}

	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en" class="h-full bg-slate-50 text-slate-900 antialiased select-none">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no, viewport-fit=cover">
  <link rel="manifest" href="/manifest.json">
  <meta name="mobile-web-app-capable" content="yes">
  <meta name="apple-mobile-web-app-capable" content="yes">
  <meta name="apple-mobile-web-app-status-bar-style" content="default">
  <meta name="theme-color" content="#f8fafc">
  <title>%s</title>
  <script src="/vendor/htmx.min.js"></script>
  <script src="/htmx-bolt.js"></script>
  <script src="/htmx-vibe.js"></script>
  <script src="/appz.js"></script>
  <script src="https://cdn.tailwindcss.com"></script>
  <link rel="stylesheet" href="/styles.css">
  <style>
    body { padding-top: env(safe-area-inset-top); padding-bottom: env(safe-area-inset-bottom); }
  </style>
</head>
<body class="min-h-full flex flex-col items-center justify-start bg-slate-50 text-slate-900">
  <app-shellui view-mode="Mob-P" class="w-full max-w-sm flex-1 flex flex-col items-center justify-start pb-20">
    <main class="w-full flex flex-col space-y-2.5 px-3 pt-2.5">
      %s
      
      <!-- 3 Tabs: Bio | Feed | Talk -->
      <div class="flex items-center p-0.5 bg-slate-200/60 rounded-xl text-xs font-semibold">
        <button class="flex-1 py-1 rounded-lg bg-white text-slate-900 font-bold transition-all shadow-xs flex items-center justify-center gap-1">
          <span>🔗</span> <span>Bio</span>
        </button>
        <button class="flex-1 py-1 rounded-lg text-slate-600 hover:text-slate-900 transition-all flex items-center justify-center gap-1">
          <span>📝</span> <span>Feed</span>
        </button>
        <button class="flex-1 py-1 rounded-lg text-slate-600 hover:text-slate-900 transition-all flex items-center justify-center gap-1">
          <span>💬</span> <span>Talk</span>
        </button>
      </div>

      <!-- Links Stack -->
      <div class="space-y-1.5">
        %s
      </div>

      <footer class="pt-2 pb-1 text-center">
        <p class="text-[10px] text-slate-400 font-mono">Built with Hyperstack • Zero VDOM</p>
      </footer>
    </main>

    %s
  </app-shellui>
</body>
</html>`, p.Title, profileComp.Render(), linksSb.String(), navbarComp.Render())

	return Component{html: html}
}
