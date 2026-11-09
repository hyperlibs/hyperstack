package components

import (
	"fmt"
	"strings"
)

// =============================================================================
// LINK IN BIO: GO COMPONENT DSL
// Pure Mobile-First Link-in-Bio Application using Hyperstack (Light Theme & PWA)
// Conforms to Hyperstack Component Standards & 1:1 Parity Rule
// Tabs: Bio | Feed | Talk (Digital Clone AI Agent) with <app-navbar> PWA Navigation
// =============================================================================

// --- 1. Bio Profile Header ---

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
		p.AvatarURL = "https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&w=280&q=80"
	}
	if p.Name == "" {
		p.Name = "Alex Rivera"
	}
	if p.Handle == "" {
		p.Handle = "@alex_hyperstack"
	}
	if p.Bio == "" {
		p.Bio = "Systems Architect & Open Source Creator. Building lean hyperreactive web systems with zero Virtual DOM."
	}

	verifiedBadge := ""
	if p.Verified {
		verifiedBadge = `<span class="inline-flex items-center justify-center w-4 h-4 rounded-full bg-blue-500 text-white text-[10px] font-bold shadow-sm" title="Verified Creator">✓</span>`
	}

	var socialsSb strings.Builder
	for _, s := range p.Socials {
		socialsSb.WriteString(fmt.Sprintf(`
      <a href="%s" target="_blank" rel="noopener noreferrer" 
         class="w-9 h-9 rounded-full bg-slate-100 hover:bg-slate-200 border border-slate-200 flex items-center justify-center text-sm text-slate-700 transition-colors"
         title="%s">
        %s
      </a>`, s.URL, s.Label, s.Icon))
	}

	html := fmt.Sprintf(`
<div class="flex flex-col items-center text-center p-6 bg-white border border-slate-200 rounded-3xl shadow-sm">
  <div class="relative mb-3">
    <div class="w-24 h-24 rounded-full p-1 bg-gradient-to-tr from-indigo-500 to-purple-600 shadow-md">
      <img src="%s" alt="%s" class="w-full h-full rounded-full object-cover border-2 border-white" />
    </div>
  </div>

  <div class="flex items-center gap-1.5 justify-center mb-0.5">
    <h1 class="text-xl font-bold text-slate-900 tracking-tight">%s</h1>
    %s
  </div>
  <p class="text-xs font-mono font-medium text-indigo-600 mb-2">%s</p>

  <p class="text-xs text-slate-600 max-w-xs mb-4 leading-relaxed font-normal">%s</p>

  <div class="flex items-center justify-center gap-2 mb-4 flex-wrap">
    %s
  </div>

  <button onclick="navigator.clipboard.writeText(window.location.href); alert('Profile link copied to clipboard!');"
          class="w-full py-2.5 px-4 rounded-xl bg-slate-900 hover:bg-slate-800 text-white font-medium text-xs shadow-sm transition-colors flex items-center justify-center gap-1.5 active:scale-95">
    <span>🔗</span> <span>Share Profile</span>
  </button>
</div>`, p.AvatarURL, p.Name, p.Name, verifiedBadge, p.Handle, p.Bio, socialsSb.String())

	return Component{html: html}
}

// --- 2. Bio Link Card ---

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
		badgeHtml = fmt.Sprintf(`<span class="text-[10px] font-bold px-2 py-0.5 rounded-full bg-indigo-50 text-indigo-700 border border-indigo-200">%s</span>`, p.Badge)
	}

	descHtml := ""
	if p.Description != "" {
		descHtml = fmt.Sprintf(`<p class="text-xs text-slate-500 truncate mt-0.5">%s</p>`, p.Description)
	}

	html := fmt.Sprintf(`
<a href="%s" target="_blank" rel="noopener noreferrer"
   class="group flex items-center gap-3.5 p-4 rounded-2xl bg-white border border-slate-200 hover:border-indigo-400 hover:shadow-md transition-all active:scale-[0.99] shadow-sm">
  <div class="w-11 h-11 rounded-xl bg-slate-100 border border-slate-200 flex items-center justify-center text-xl shrink-0 text-slate-700">
    %s
  </div>
  <div class="flex-1 min-w-0 text-left">
    <div class="flex items-center gap-2">
      <h2 class="font-bold text-sm text-slate-900 group-hover:text-indigo-600 transition-colors truncate">%s</h2>
      %s
    </div>
    %s
  </div>
  <div class="w-8 h-8 rounded-full bg-slate-50 flex items-center justify-center text-slate-400 group-hover:text-indigo-600 group-hover:bg-indigo-50 transition-colors shrink-0">
    ↗
  </div>
</a>`, p.URL, p.Icon, p.Title, badgeHtml, descHtml)

	return Component{html: html}
}

// --- 3. Bio Feed Component (Text & Emojis Only) ---

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
		tagColor := "text-slate-500"
		if strings.Contains(strings.ToLower(post.Tag), "now") {
			tagColor = "text-indigo-600 font-semibold"
		}

		postsSb.WriteString(fmt.Sprintf(`
  <article class="p-4 rounded-2xl bg-white border border-slate-200 shadow-sm space-y-2">
    <div class="flex items-center justify-between text-xs text-slate-400">
      <span class="%s">%s</span>
      <span class="font-mono text-[11px]">%s</span>
    </div>
    <p class="text-sm text-slate-800 leading-relaxed font-normal">
      %s
    </p>
    <div class="pt-2 flex items-center justify-between text-xs text-slate-400 border-t border-slate-100">
      <span class="text-[11px]">📍 %s</span>
      <button class="flex items-center gap-1 text-xs text-slate-400 hover:text-rose-500 transition-colors">
        <span>🤍</span> <span>%d</span>
      </button>
    </div>
  </article>`, tagColor, post.Tag, post.TimeAgo, post.Content, post.Location, post.Likes))
	}

	html := fmt.Sprintf(`<section class="space-y-3">%s</section>`, postsSb.String())
	return Component{html: html}
}

// --- 4. Bio Talk Component (Digital Clone AI Agent) ---

type BioTalkProps struct {
	CloneName    string
	Subtitle     string
	Greeting     string
	QuickPrompts []string
}

func BioTalk(p BioTalkProps) Component {
	if p.CloneName == "" {
		p.CloneName = "Alex's Digital Clone"
	}
	if p.Subtitle == "" {
		p.Subtitle = "Responds strictly as per Alex's verified answers"
	}
	if p.Greeting == "" {
		p.Greeting = "Hey! 👋 I'm <strong>Alex's Digital Clone AI</strong>. I answer strictly based on Alex's own personal inputs and real answers."
	}
	if len(p.QuickPrompts) == 0 {
		p.QuickPrompts = []string{
			"Do you have a girlfriend? 👀",
			"Top projects? 🚀",
			"Where do you live? 📍",
			"Tech stack? ⚡",
		}
	}

	var chipsSb strings.Builder
	for _, prompt := range p.QuickPrompts {
		chipsSb.WriteString(fmt.Sprintf(`
    <button class="px-2.5 py-1 rounded-full bg-slate-100 hover:bg-indigo-50 hover:text-indigo-600 border border-slate-200 text-xs font-medium text-slate-700 shrink-0 transition-colors">
      %s
    </button>`, prompt))
	}

	html := fmt.Sprintf(`
<div class="flex flex-col h-[520px] rounded-3xl bg-white border border-slate-200 overflow-hidden shadow-sm">
  <div class="h-14 px-4 bg-slate-50 border-b border-slate-200 flex items-center justify-between shrink-0">
    <div class="flex items-center gap-2.5">
      <div class="relative">
        <div class="w-8 h-8 rounded-full bg-gradient-to-tr from-indigo-500 to-purple-600 flex items-center justify-center text-white text-xs font-bold shadow-sm">
          A
        </div>
        <span class="w-2.5 h-2.5 rounded-full bg-emerald-500 absolute -bottom-0.5 -right-0.5 border-2 border-white"></span>
      </div>
      <div>
        <div class="flex items-center gap-1.5">
          <h3 class="text-xs font-bold text-slate-900 leading-none">%s</h3>
          <span class="text-[9px] font-mono px-1.5 py-0.2 rounded bg-indigo-50 text-indigo-700 border border-indigo-200 font-bold">AI Agent</span>
        </div>
        <p class="text-[10px] text-slate-500 mt-0.5">%s</p>
      </div>
    </div>

    <label class="flex items-center gap-1.5 cursor-pointer select-none" title="Ask without revealing your identity">
      <input type="checkbox" id="anon-checkbox" class="rounded border-slate-300 text-indigo-600 focus:ring-indigo-500 h-3.5 w-3.5 cursor-pointer" />
      <span class="text-xs font-medium text-slate-600 flex items-center gap-1">
        <span>🎭</span> <span>Ask Anon</span>
      </span>
    </label>
  </div>

  <div class="flex-1 overflow-y-auto p-4 space-y-3 bg-slate-50/50">
    <div class="flex items-start gap-2.5 max-w-[88%%]">
      <div class="w-7 h-7 rounded-full bg-indigo-600 flex items-center justify-center text-white text-xs font-bold shrink-0">A</div>
      <div class="p-3 rounded-2xl rounded-tl-sm bg-white border border-slate-200 text-xs text-slate-800 leading-relaxed shadow-sm">
        %s
        <div class="text-[9px] text-slate-400 text-right mt-1 font-mono">09:41</div>
      </div>
    </div>
  </div>

  <div class="px-3 py-2 bg-white border-t border-slate-200 flex items-center gap-1.5 overflow-x-auto shrink-0 scrollbar-none">
    %s
  </div>

  <div class="p-3 bg-white border-t border-slate-200 flex items-center gap-2 shrink-0">
    <input type="text" placeholder="Ask Alex's clone anything..." class="flex-1 bg-slate-50 border border-slate-200 rounded-xl px-3 py-2 text-xs text-slate-900 placeholder:text-slate-400 focus:outline-none focus:border-indigo-500 shadow-inner" />
    <button class="px-4 py-2 rounded-xl bg-slate-900 hover:bg-slate-800 text-white font-semibold text-xs shadow-sm transition-all flex items-center gap-1">
      <span>Send</span> <span>➤</span>
    </button>
  </div>
</div>`, p.CloneName, p.Subtitle, p.Greeting, chipsSb.String())

	return Component{html: html}
}

// --- 5. Bio Navbar Component (<app-navbar>) ---

type BioNavItem struct {
	ID     string
	Label  string
	Icon   string
	Active bool
}

type BioNavbarProps struct {
	Mode     string // "float" | "sticky"
	Behavior string // "always-visible" | "hide-on-scroll"
	Active   string // "bio" | "feed" | "talk"
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
      <button onclick="switchTab('%s')" id="nav-btn-%s" class="nav-tab-item flex flex-col items-center gap-0.5 py-1.5 px-3 rounded-xl transition-all %s">
        <span class="text-base leading-none">%s</span>
        <span class="text-[10px] tracking-tight">%s</span>
      </button>`, it.ID, it.ID, activeClass, it.Icon, it.Label))
	}

	html := fmt.Sprintf(`
<app-navbar mode="%s" behavior="%s" class="fixed bottom-3 inset-x-0 mx-auto max-w-md px-4 z-40 flex items-center justify-between">
  <div class="w-full bg-white/95 backdrop-blur-xl border border-slate-200/90 rounded-2xl shadow-xl shadow-slate-900/5 h-16 px-3 flex items-center justify-around">
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
  <app-shellui view-mode="Mob-P" class="w-full flex-1 flex flex-col items-center justify-start pb-28">
    <main class="w-full max-w-md flex flex-col space-y-4 px-4 pt-4">
      %s
      
      <!-- 3 Tabs: Bio | Feed | Talk -->
      <div class="flex items-center justify-between p-1 bg-white border border-slate-200 rounded-2xl shadow-sm text-xs font-semibold">
        <button class="flex-1 py-2 rounded-xl bg-slate-900 text-white font-bold transition-all shadow-sm flex items-center justify-center gap-1.5">
          <span>🔗</span> <span>Bio</span>
        </button>
        <button class="flex-1 py-2 rounded-xl text-slate-600 hover:text-slate-900 transition-all flex items-center justify-center gap-1.5">
          <span>📝</span> <span>Feed</span>
        </button>
        <button class="flex-1 py-2 rounded-xl text-slate-600 hover:text-slate-900 transition-all flex items-center justify-center gap-1.5">
          <span>💬</span> <span>Talk</span>
        </button>
      </div>

      <!-- Links Stack -->
      <div class="space-y-3">
        %s
      </div>

      <footer class="pt-6 pb-4 text-center">
        <p class="text-xs text-slate-400 font-mono">Built with Hyperstack • Zero Virtual DOM</p>
      </footer>
    </main>

    %s
  </app-shellui>
</body>
</html>`, p.Title, profileComp.Render(), linksSb.String(), navbarComp.Render())

	return Component{html: html}
}
