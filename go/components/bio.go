package components

import (
	"fmt"
	"strings"
)

// =============================================================================
// LINK IN BIO: GO COMPONENT DSL
// Pure Mobile-First Link-in-Bio Application using Hyperstack (Ultra-Compact Light Theme & PWA)
// Conforms to Hyperstack Component Standards & 1:1 Parity Rule
// Tabs: Bio | Links | Feed | Talk (Digital Clone AI Agent with Embedded Tiny Self-Learning LLM)
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

// --- 2. Bio Details Component (Name, Relation with You, Hobbies, Interests, Goals, Achievements) ---

type AchievementItem struct {
	Icon        string
	Title       string
	Description string
}

type BioDetailsProps struct {
	RelationTitle string
	RelationBadge string
	RelationDesc  string
	Hobbies       []string
	Interests     []string
	Goals         []string
	Achievements  []AchievementItem
}

func BioDetails(p BioDetailsProps) Component {
	if p.RelationTitle == "" {
		p.RelationTitle = "Relation with You"
	}
	if p.RelationBadge == "" {
		p.RelationBadge = "Visitor (Public)"
	}
	if p.RelationDesc == "" {
		p.RelationDesc = "You are viewing Alex's bio as a <strong>public visitor</strong>. If you received an invite code, enter it to reveal your personalized connection."
	}
	if len(p.Hobbies) == 0 {
		p.Hobbies = []string{
			"☕️ Pour-over coffee (V60 Ethiopian)",
			"🏃‍♂️ Morning 5km runs around the bay",
			"♟️ Blitz Chess (1800 Elo)",
			"⌨️ Custom mechanical keyboards",
		}
	}
	if len(p.Interests) == 0 {
		p.Interests = []string{
			"⚡ Hypermedia & Zero-VDOM",
			"🌐 Distributed Consensus & Raft",
			"📱 Local-First & Spatial Operating Systems",
			"📚 Hard Sci-Fi Literature",
		}
	}
	if len(p.Goals) == 0 {
		p.Goals = []string{
			"Empower 10,000+ developers to build hyperreactive web apps without Node.js or heavy SPAs.",
			"Run a sub-3:30 marathon in San Francisco.",
			"Ship Hyperstack v2.0 single-binary full-stack runtime with instant zero-allocation clustering.",
		}
	}
	if len(p.Achievements) == 0 {
		p.Achievements = []AchievementItem{
			{
				Icon:        "⚡",
				Title:       "Created Hyperstack Framework",
				Description: "Brought reactive signals, zero-eval runtime, and spatial layouts to HTMX with 0 Node.js footprint.",
			},
			{
				Icon:        "📊",
				Title:       "Authored Hypersheet DAG Core",
				Description: "High-performance topological spreadsheet engine in Go with sub-millisecond recalculations.",
			},
			{
				Icon:        "🚀",
				Title:       "Scaled Systems to 1,000,000 RPS",
				Description: "Engineered zero-allocation SSE streaming architecture in Go with kernel-level zero-copy sockets.",
			},
		}
	}

	var hobbiesSb strings.Builder
	for _, h := range p.Hobbies {
		hobbiesSb.WriteString(fmt.Sprintf(`<span class="px-2 py-1 rounded-lg bg-slate-100 text-slate-700 text-[11px] font-medium">%s</span> `, h))
	}

	var interestsSb strings.Builder
	for _, in := range p.Interests {
		interestsSb.WriteString(fmt.Sprintf(`<span class="px-2 py-1 rounded-lg bg-indigo-50 text-indigo-700 border border-indigo-100 text-[11px] font-medium">%s</span> `, in))
	}

	var goalsSb strings.Builder
	for _, g := range p.Goals {
		goalsSb.WriteString(fmt.Sprintf(`
      <li class="flex items-start gap-2">
        <span class="text-indigo-600 font-bold shrink-0 mt-0.5">▹</span>
        <span>%s</span>
      </li>`, g))
	}

	var achSb strings.Builder
	for _, a := range p.Achievements {
		achSb.WriteString(fmt.Sprintf(`
      <div class="p-2 rounded-xl bg-slate-50 border border-slate-200/70 flex items-start gap-2">
        <span class="text-base shrink-0">%s</span>
        <div>
          <div class="text-xs font-bold text-slate-900">%s</div>
          <p class="text-[10px] text-slate-500 mt-0.5">%s</p>
        </div>
      </div>`, a.Icon, a.Title, a.Description))
	}

	html := fmt.Sprintf(`
<div id="content-bio" class="space-y-2.5">
  <!-- Relation With You Card (Dynamic based on active collection/persona) -->
  <div class="p-3 rounded-2xl bg-gradient-to-r from-indigo-50/80 to-purple-50/80 border border-indigo-100 shadow-2xs">
    <div class="flex items-center justify-between">
      <span class="text-[10px] font-bold text-indigo-600 uppercase tracking-wider flex items-center gap-1">
        <span>🤝</span> <span>%s</span>
      </span>
      <span id="bio-relation-badge" class="px-2 py-0.5 rounded-full text-[10px] font-bold bg-white text-indigo-700 border border-indigo-200 shadow-2xs">
        %s
      </span>
    </div>
    <p id="bio-relation-desc" class="text-xs text-slate-700 mt-1.5 leading-relaxed font-medium">
      %s
    </p>
  </div>

  <!-- Hobbies & Daily Rituals -->
  <div class="p-3 rounded-2xl bg-white border border-slate-200/80 shadow-2xs space-y-1.5">
    <h3 class="text-xs font-bold text-slate-900 flex items-center gap-1.5">
      <span>☕️</span> <span>Hobbies & Rituals</span>
    </h3>
    <div class="flex flex-wrap gap-1.5 pt-0.5">
      %s
    </div>
  </div>

  <!-- Core Interests -->
  <div class="p-3 rounded-2xl bg-white border border-slate-200/80 shadow-2xs space-y-1.5">
    <h3 class="text-xs font-bold text-slate-900 flex items-center gap-1.5">
      <span>🔬</span> <span>Interests</span>
    </h3>
    <div class="flex flex-wrap gap-1.5 pt-0.5">
      %s
    </div>
  </div>

  <!-- Future Goals -->
  <div class="p-3 rounded-2xl bg-white border border-slate-200/80 shadow-2xs space-y-2">
    <h3 class="text-xs font-bold text-slate-900 flex items-center gap-1.5">
      <span>🎯</span> <span>2026 Goals</span>
    </h3>
    <ul class="space-y-1.5 text-xs text-slate-700">
      %s
    </ul>
  </div>

  <!-- Key Achievements -->
  <div class="p-3 rounded-2xl bg-white border border-slate-200/80 shadow-2xs space-y-2">
    <h3 class="text-xs font-bold text-slate-900 flex items-center gap-1.5">
      <span>🏆</span> <span>Achievements</span>
    </h3>
    <div class="space-y-1.5">
      %s
    </div>
  </div>
</div>`, p.RelationTitle, p.RelationBadge, p.RelationDesc, hobbiesSb.String(), interestsSb.String(), goalsSb.String(), achSb.String())

	return Component{html: html}
}

// --- 3. Bio Link Card (Ultra-Compact) ---

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

// --- 4. Bio Feed Component (Ultra-Compact Text & Emojis Only) ---

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

// --- 5. Bio Talk Component (Digital Clone AI Agent with Embedded Tiny Self-Learning LLM) ---

type BioTalkProps struct {
	CloneName    string
	Subtitle     string
	Greeting     string
	QuickPrompts []string
	IsOfflineLLM bool
}

func BioTalk(p BioTalkProps) Component {
	if p.CloneName == "" {
		p.CloneName = "Alex's Clone"
	}
	if p.Subtitle == "" {
		p.Subtitle = "Verified answers"
	}
	if p.Greeting == "" {
		p.Greeting = "Hey! 👋 I'm <strong>AlexAI</strong>, Alex Rivera's digital clone middleman. I only respond as per Alex's provided answers for your group.<p class=\"text-slate-500 text-[10px] mt-1\">If Alex hasn't answered a question yet, it is forwarded directly to Alex's inbox 📨. Alex can provide multiple answers per audience tag (@all, @family, @friends, @bffs).</p>"
	}
	if len(p.QuickPrompts) == 0 {
		p.QuickPrompts = []string{
			"Where are you now? 📍",
			"Do you have a girlfriend? 👀",
			"Top projects? 🚀",
			"Coffee ritual ☕️",
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
<div id="talk-container" class="flex flex-col h-[520px] rounded-2xl bg-white border border-slate-200/80 overflow-hidden shadow-2xs">
  <div class="h-10 px-3 bg-slate-50 border-b border-slate-200/70 flex items-center justify-between shrink-0">
    <div class="flex items-center gap-1.5">
      <div class="w-5 h-5 rounded-full bg-gradient-to-tr from-indigo-500 to-purple-600 flex items-center justify-center text-white text-[9px] font-bold">
        A
      </div>
      <span class="text-xs font-bold text-slate-900">%s</span>
      <span id="clone-role-badge" class="text-[9px] text-indigo-600 font-mono font-bold bg-indigo-50 px-1 rounded">AI</span>
      <span id="clone-llm-badge" class="text-[9px] text-emerald-700 font-mono font-bold bg-emerald-50 border border-emerald-200/80 px-1 rounded flex items-center gap-0.5" title="Tiny Self-Learning LLM Middleman Running Locally On-Device">
        <span class="w-1.5 h-1.5 rounded-full bg-emerald-500 animate-pulse"></span>
        <span>Offline LLM</span>
      </span>
    </div>

    <div class="flex items-center gap-1.5">
      <button onclick="openUsernameModal()" id="talk-user-badge-btn" type="button"
              class="px-2 py-0.5 rounded-md bg-white hover:bg-slate-100 border border-slate-200 text-[10px] font-mono font-bold text-slate-700 flex items-center gap-1 transition-colors shadow-2xs"
              title="Your username for identification (can be changed once a year)">
        <span class="w-1.5 h-1.5 rounded-full bg-emerald-500"></span>
        <span id="active-user-handle">@visitor</span>
        <span class="text-[9px] text-slate-400">✏️</span>
      </button>

      <button onclick="toggleInboxDrawer()" id="inbox-drawer-btn" type="button"
              class="text-[10px] font-bold text-amber-700 hover:text-amber-900 bg-amber-50 hover:bg-amber-100 border border-amber-200/80 px-1.5 py-0.5 rounded-md transition-colors flex items-center gap-1" title="Alex's Question Inbox (Middleman Control)">
        <span>📥</span> <span>Inbox</span>
        <span id="inbox-pending-count" class="bg-amber-600 text-white rounded-full px-1 text-[8px] font-bold">2</span>
      </button>

      <button onclick="toggleTeachDrawer()" id="teach-drawer-btn" type="button"
              class="text-[10px] font-bold text-indigo-600 hover:text-indigo-800 bg-indigo-50 hover:bg-indigo-100 px-1.5 py-0.5 rounded-md transition-colors flex items-center gap-0.5" title="Teach new facts to the self-learning LLM">
        <span>💡</span> <span>Teach</span>
      </button>

      <label class="flex items-center gap-1 cursor-pointer select-none" title="Ask without identity">
        <input type="checkbox" id="anon-toggle" class="rounded border-slate-300 text-indigo-600 focus:ring-indigo-500 h-3 w-3 cursor-pointer" />
        <span class="text-[10px] font-medium text-slate-600 flex items-center gap-0.5">
          <span>🎭</span> <span>Ask Anon</span>
        </span>
      </label>
    </div>
  </div>

  <!-- 2 Sub-Tabs Controller: Talk | Talk back -->
  <div class="px-3 py-1.5 bg-slate-100/70 border-b border-slate-200 flex items-center justify-between shrink-0">
    <div class="flex items-center gap-1 bg-white p-0.5 rounded-xl border border-slate-200/80 shadow-2xs text-xs font-semibold">
      <button type="button" onclick="switchTalkSubTab('talk')" id="talk-subtab-talk" 
              class="px-3 py-1 rounded-lg bg-slate-900 text-white font-bold transition-all shadow-xs flex items-center gap-1.5 text-xs">
        <span>💬</span> <span>Talk</span>
        <span id="talk-unanswered-badge" class="bg-amber-500 text-white rounded-full px-1.5 text-[9px] font-bold">2</span>
      </button>
      <button type="button" onclick="switchTalkSubTab('talkback')" id="talk-subtab-talkback" 
              class="px-3 py-1 rounded-lg text-slate-600 hover:text-slate-900 font-bold transition-all flex items-center gap-1.5 text-xs">
        <span>↩️</span> <span>Talk back</span>
        <span class="text-[9px] text-slate-400 font-normal">Alex Answers</span>
      </button>
    </div>
  </div>

  <!-- Collapsible Alex's Question Inbox Drawer -->
  <div id="inbox-drawer" class="hidden px-3 py-2 bg-amber-50/70 border-b border-amber-200/80 text-xs space-y-2 animate-slide-down">
    <div class="flex items-center justify-between text-[10px]">
      <span class="font-bold text-amber-900 flex items-center gap-1">
        <span>👑</span> <span>Alex's Inbox (Middleman Control)</span>
      </span>
      <button type="button" onclick="toggleInboxDrawer()" class="text-amber-600 hover:text-amber-900 ml-1 font-bold">✕</button>
    </div>
    <div id="inbox-answer-box" class="p-2 rounded-xl bg-white border border-amber-200/90 shadow-2xs space-y-1.5">
      <div class="text-[10px] font-bold text-slate-800">Middleman Question Queue Active</div>
      <button type="button" onclick="submitInboxAnswer()" class="hidden">Save</button>
    </div>
  </div>

  <!-- Collapsible Self-Learning Teach Drawer -->
  <div id="teach-drawer" class="hidden px-3 py-2 bg-indigo-50/70 border-b border-indigo-100 text-xs space-y-1.5 animate-slide-down">
    <div class="flex items-center justify-between text-[10px]">
      <span class="font-bold text-indigo-900 flex items-center gap-1">
        <span>🧠</span> <span>Teach Tiny LLM (On-Device Self-Learning)</span>
      </span>
      <span id="llm-telemetry-pill" class="font-mono text-[9px] text-indigo-700 bg-white/80 px-1.5 py-0.2 rounded border border-indigo-200">
        Steps: <span id="telemetry-step">0</span> | Loss: <span id="telemetry-loss">0.038</span>
      </span>
    </div>
    <form onsubmit="handleTeachSubmit(event)" class="flex items-center gap-1.5">
      <input type="text" id="teach-input" placeholder="e.g. Alex loves building mechanical keyboards with Boba U4T switches" 
             class="flex-1 bg-white border border-indigo-200 rounded-lg px-2.5 py-1 text-xs text-slate-900 placeholder:text-slate-400 focus:outline-none focus:border-indigo-600 shadow-2xs" />
      <button type="submit" class="px-2.5 py-1 rounded-lg bg-indigo-600 hover:bg-indigo-700 text-white font-bold text-[11px] shrink-0 transition-all shadow-xs">
        Teach & Train
      </button>
    </form>
  </div>

  <!-- SUB-PANE 1: TALK (Recent Unanswered Questions FIRST + Chat Stream) -->
  <div id="subpane-talk" class="flex-1 flex flex-col min-h-0">
    <div id="unanswered-questions-section" class="border-b border-amber-200/80 bg-amber-50/40 p-2.5 space-y-1.5 shrink-0 max-h-48 overflow-y-auto custom-scroll">
      <div class="flex items-center justify-between text-[11px]">
        <div class="flex items-center gap-1.5 font-bold text-amber-950">
          <span class="w-2 h-2 rounded-full bg-amber-500 animate-pulse"></span>
          <span>Recent Unanswered Questions</span>
          <span id="unanswered-count-badge" class="px-1.5 py-0.2 rounded-full bg-amber-200 text-amber-900 text-[9px] font-mono font-bold">2</span>
        </div>
        <span class="text-[9px] text-amber-800 font-mono">AlexAI silent until answered</span>
      </div>
      <div id="unanswered-questions-list" class="space-y-1.5"></div>
    </div>

    <div id="talk-stream" class="flex-1 overflow-y-auto p-2.5 space-y-2 bg-slate-50/40 custom-scroll text-xs">
      <div class="flex items-start gap-1.5 max-w-[92%%]">
        <div class="w-5 h-5 rounded-full bg-indigo-600 flex items-center justify-center text-white text-[9px] font-bold shrink-0">A</div>
        <div class="p-2 rounded-xl rounded-tl-xs bg-white border border-slate-200/80 text-xs text-slate-800 leading-snug shadow-2xs">
          %s
          <div class="text-[8px] text-slate-400 text-right mt-1 font-mono flex items-center justify-end gap-1">
            <span class="text-emerald-600">● Middleman Proxy</span> • <span>Multi-Audience</span>
          </div>
        </div>
      </div>
    </div>

    <div class="px-2 py-1 bg-white border-t border-slate-200/70 flex items-center gap-1 overflow-x-auto shrink-0 scrollbar-none">
      %s
    </div>

    <form id="talk-form" onsubmit="handleCloneSubmit(event)" class="p-1.5 bg-white border-t border-slate-200/70 flex items-center gap-1.5 shrink-0">
      <input type="text" id="clone-input" placeholder="Ask Alex's offline clone anything..." class="flex-1 bg-slate-50 border border-slate-200/80 rounded-lg px-2.5 py-1 text-xs text-slate-900 placeholder:text-slate-400 focus:outline-none focus:border-indigo-500 shadow-2xs" />
      <button type="submit" class="px-3 py-1 rounded-lg bg-slate-900 hover:bg-slate-800 text-white font-semibold text-xs transition-all">
        ➤
      </button>
    </form>
  </div>

  <!-- SUB-PANE 2: TALK BACK (Answer Station: Multi-People & Usernames) -->
  <div id="subpane-talkback" class="hidden flex-1 flex flex-col min-h-0 bg-slate-50/50 overflow-y-auto custom-scroll p-3 space-y-3">
    <div class="p-2.5 rounded-xl bg-white border border-amber-200 shadow-2xs space-y-2">
      <div class="flex items-center justify-between text-xs">
        <span class="font-bold text-amber-950 flex items-center gap-1.5">
          <span>👑</span> <span>Talk back — Alex's Answer Station</span>
        </span>
        <span class="text-[9px] px-1.5 py-0.2 rounded bg-amber-100 text-amber-900 font-mono font-bold">Middleman Mode</span>
      </div>
      <div class="space-y-1">
        <select id="talkback-question-select" onchange="onTalkBackQuestionSelected(this.value)" class="w-full bg-amber-50/70 border border-amber-300 rounded-lg px-2 py-1 text-xs font-bold text-slate-900"></select>
      </div>
    </div>

    <div class="p-3 rounded-xl bg-white border border-slate-200 shadow-2xs space-y-2.5">
      <div class="flex items-center gap-1 flex-wrap text-[10px]">
        <span class="font-bold text-slate-600">People:</span>
        <span class="font-mono text-emerald-800 bg-emerald-50 px-1 rounded">@marcus</span>
        <span class="font-mono text-purple-800 bg-purple-50 px-1 rounded">@sarah</span>
        <span class="font-mono text-blue-800 bg-blue-50 px-1 rounded">@priya</span>
        <span class="font-mono text-rose-800 bg-rose-50 px-1 rounded">@mom</span>
      </div>
      <textarea id="talkback-answers-editor" rows="4" placeholder="@all im at work, chennai&#10;@friends, @marcus im at home with my family" class="w-full bg-slate-50 border border-slate-200 rounded-xl p-2.5 text-xs text-slate-900 font-mono"></textarea>
      <div class="flex items-center justify-end">
        <button type="button" onclick="submitTalkBackAnswer()" class="px-3 py-1.5 rounded-xl bg-amber-600 text-white font-bold text-xs shadow-xs">
          💾 Save Answers & Publish to AlexAI
        </button>
      </div>
    </div>

    <div id="talkback-answered-directory" class="space-y-1.5"></div>
  </div>
</div>`, p.CloneName, p.Greeting, chipsSb.String())

	return Component{html: html}
}

// --- 6. Bio Invite & Collections Access Rules Component ---

type CollectionRuleItem struct {
	ID    string
	Name  string
	Bio   bool
	Links bool
	Feed  bool
	Talk  bool
}

type BioInviteProps struct {
	Collections []CollectionRuleItem
}

func BioInvite(p BioInviteProps) Component {
	if len(p.Collections) == 0 {
		p.Collections = []CollectionRuleItem{
			{ID: "friends", Name: "Friends", Bio: true, Links: true, Feed: true, Talk: true},
			{ID: "bffs", Name: "BFFs", Bio: true, Links: true, Feed: true, Talk: true},
			{ID: "relationship", Name: "Relationship", Bio: true, Links: true, Feed: true, Talk: true},
			{ID: "family", Name: "Family", Bio: true, Links: true, Feed: true, Talk: false},
			{ID: "work", Name: "Work", Bio: true, Links: true, Feed: false, Talk: false},
			{ID: "church", Name: "Church", Bio: true, Links: false, Feed: true, Talk: false},
		}
	}

	var chipsSb strings.Builder
	var rulesSb strings.Builder

	for _, c := range p.Collections {
		chipsSb.WriteString(fmt.Sprintf(`
      <button type="button" class="px-2.5 py-1 rounded-lg text-[11px] transition-all flex items-center gap-1 bg-slate-100 hover:bg-slate-200 text-slate-700">
        <span>%s</span>
      </button>`, c.Name))

		rulesSb.WriteString(fmt.Sprintf(`
      <div class="p-2.5 rounded-xl bg-white border border-slate-200/80 shadow-2xs space-y-2">
        <div class="flex items-center justify-between">
          <input type="text" value="%s" class="text-xs font-bold text-slate-900 bg-transparent border-b border-dashed border-slate-300 focus:outline-none focus:border-indigo-600 px-0.5" />
          <span class="text-[9px] font-mono text-slate-400">#%s</span>
        </div>
        <div class="grid grid-cols-4 gap-1 text-[10px] text-slate-600">
          <label class="flex items-center gap-1">
            <input type="checkbox" %s class="rounded text-indigo-600" /> Bio
          </label>
          <label class="flex items-center gap-1">
            <input type="checkbox" %s class="rounded text-indigo-600" /> Links
          </label>
          <label class="flex items-center gap-1">
            <input type="checkbox" %s class="rounded text-indigo-600" /> Feed
          </label>
          <label class="flex items-center gap-1">
            <input type="checkbox" %s class="rounded text-indigo-600" /> Talk
          </label>
        </div>
      </div>`, c.Name, c.ID, checkedAttr(c.Bio), checkedAttr(c.Links), checkedAttr(c.Feed), checkedAttr(c.Talk)))
	}

	html := fmt.Sprintf(`
<div id="invite-modal" class="fixed inset-0 z-50 bg-slate-900/40 backdrop-blur-xs flex items-end sm:items-center justify-center p-0 sm:p-4 opacity-0 pointer-events-none transition-opacity duration-200">
  <div class="w-full max-w-md bg-white border border-slate-200 rounded-t-3xl sm:rounded-3xl shadow-2xl max-h-[90vh] flex flex-col overflow-hidden animate-slide-up">
    
    <!-- Modal Header -->
    <div class="px-4 py-3 bg-slate-50 border-b border-slate-200 flex items-center justify-between shrink-0">
      <div class="flex items-center gap-2">
        <span class="w-7 h-7 rounded-xl bg-indigo-50 border border-indigo-100 text-indigo-600 flex items-center justify-center text-xs font-bold">👥</span>
        <div>
          <h3 class="text-xs font-bold text-slate-900 leading-none">Invite & Collections</h3>
          <p class="text-[10px] text-slate-500 mt-0.5">Tag people & configure tab access rules</p>
        </div>
      </div>
      <button onclick="closeInviteModal()" class="w-7 h-7 rounded-full bg-slate-200/70 hover:bg-slate-300 text-slate-600 flex items-center justify-center text-xs font-bold transition-colors">✕</button>
    </div>

    <!-- Sub Tabs: 1. Invite People | 2. Collections & Access Rules -->
    <div class="p-2 border-b border-slate-100 bg-white flex items-center gap-1.5 shrink-0 text-xs font-semibold">
      <button onclick="switchInviteSubTab('invite')" id="subtab-btn-invite" 
              class="flex-1 py-1.5 rounded-xl bg-slate-900 text-white font-bold transition-all shadow-xs flex items-center justify-center gap-1">
        <span>✉️</span> <span>Invite People</span>
      </button>
      <button onclick="switchInviteSubTab('rules')" id="subtab-btn-rules" 
              class="flex-1 py-1.5 rounded-xl text-slate-600 hover:text-slate-900 transition-all flex items-center justify-center gap-1">
        <span>🔒</span> <span>Access Rules</span>
      </button>
    </div>

    <!-- Modal Content Stream -->
    <div class="flex-1 overflow-y-auto p-4 space-y-4 custom-scroll text-xs">
      
      <!-- SUBTAB 1: INVITE PEOPLE -->
      <div id="subpane-invite" class="space-y-3.5">
        <div class="space-y-1.5">
          <label class="text-[11px] font-bold text-slate-700 block">Person's Name or Email</label>
          <input type="text" id="invitee-name" placeholder="e.g. Sarah Jenkins or sarah@example.com" 
                 class="w-full bg-slate-50 border border-slate-200 rounded-xl px-3 py-2 text-xs text-slate-900 placeholder:text-slate-400 focus:outline-none focus:border-indigo-500 focus:bg-white transition-all shadow-2xs" />
        </div>

        <!-- Tag in Collections -->
        <div class="space-y-1.5">
          <label class="text-[11px] font-bold text-slate-700 flex items-center justify-between">
            <span>Tag in Collections</span>
            <span class="text-[10px] font-normal text-slate-400">Select one or more</span>
          </label>
          <div id="invite-collection-chips" class="flex flex-wrap gap-1.5">
            %s
          </div>
        </div>

        <button onclick="generateInvite()" 
                class="w-full py-2 px-3 rounded-xl bg-indigo-600 hover:bg-indigo-700 text-white font-bold text-xs shadow-xs transition-all active:scale-95 flex items-center justify-center gap-1.5">
          <span>🔗</span> <span>Generate & Copy Invite Link</span>
        </button>

        <!-- Invited People List -->
        <div class="pt-2 border-t border-slate-100 space-y-2">
          <div class="flex items-center justify-between text-[11px] font-bold text-slate-700">
            <span>Invited Members</span>
            <span id="invited-count" class="text-[10px] font-mono text-indigo-600">3 active</span>
          </div>
          <div id="invited-list" class="space-y-1.5">
            <!-- Dynamic Invited Members -->
          </div>
        </div>
      </div>

      <!-- SUBTAB 2: COLLECTIONS & ACCESS RULES -->
      <div id="subpane-rules" class="hidden space-y-3.5">
        <div class="flex items-center justify-between">
          <span class="text-[11px] text-slate-500">Configure which tabs each group can access:</span>
          <button onclick="addNewCollectionPrompt()" 
                  class="px-2 py-0.5 rounded-lg bg-indigo-50 hover:bg-indigo-100 text-indigo-700 text-[10px] font-bold border border-indigo-200 flex items-center gap-0.5 transition-colors">
            <span>+</span> <span>New Collection</span>
          </button>
        </div>

        <!-- Collections Table / Cards -->
        <div id="collections-rules-list" class="space-y-2">
          %s
        </div>

        <div class="p-2.5 rounded-xl bg-slate-50 border border-slate-200 text-[10px] text-slate-500 leading-relaxed">
          💡 <strong>Tip:</strong> You can click any collection name above to <strong>rename it however you want</strong> (e.g. rename <em>relationship</em> to <em>dating</em> or <em>work</em> to <em>clients</em>).
        </div>
      </div>

    </div>

    <!-- Modal Footer -->
    <div class="p-3 bg-slate-50 border-t border-slate-200 flex items-center justify-between shrink-0">
      <span class="text-[10px] text-slate-400">Rules apply instantly to active sessions</span>
      <button onclick="closeInviteModal()" 
              class="px-4 py-1.5 rounded-xl bg-slate-900 hover:bg-slate-800 text-white font-semibold text-xs transition-all">
        Done
      </button>
    </div>

  </div>
</div>`, chipsSb.String(), rulesSb.String())

	return Component{html: html}
}

func checkedAttr(b bool) string {
	if b {
		return "checked"
	}
	return ""
}

// --- 7. Bio Navbar Component (2x Increased Size <app-navbar>) ---

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
			{ID: "bio", Label: "Bio", Icon: "👤", Active: p.Active == "bio"},
			{ID: "links", Label: "Links", Icon: "🔗", Active: p.Active == "links"},
			{ID: "feed", Label: "Feed", Icon: "📝", Active: p.Active == "feed"},
			{ID: "talk", Label: "Talk", Icon: "💬", Active: p.Active == "talk"},
			{ID: "invite", Label: "Invite", Icon: "👥", Active: false},
		}
	}

	var itemsSb strings.Builder
	for _, it := range p.Items {
		activeClass := "text-slate-400 hover:text-slate-700"
		if it.Active {
			activeClass = "text-indigo-600 font-bold bg-indigo-50/80"
		}
		onClick := fmt.Sprintf("switchTab('%s')", it.ID)
		if it.ID == "invite" {
			onClick = "openInviteModal()"
		}
		itemsSb.WriteString(fmt.Sprintf(`
    <button onclick="%s" id="nav-btn-%s" 
            class="nav-tab-item flex flex-col items-center gap-1 py-1 px-2.5 rounded-xl transition-all %s active:scale-95">
      <span class="text-2xl leading-none">%s</span>
      <span class="text-xs font-bold tracking-tight">%s</span>
    </button>`, onClick, it.ID, activeClass, it.Icon, it.Label))
	}

	html := fmt.Sprintf(`
<app-navbar mode="%s" behavior="%s" 
            class="fixed bottom-3 inset-x-0 mx-auto max-w-sm px-2.5 z-40 flex items-center justify-between">
  <div class="w-full bg-white/95 backdrop-blur-md border border-slate-200/90 rounded-2xl shadow-xl shadow-slate-900/5 h-17 px-1.5 flex items-center justify-around">
    %s
  </div>
</app-navbar>`, p.Mode, p.Behavior, itemsSb.String())

	return Component{html: html}
}

// --- 8. Full Mobile-First Link-in-Bio Page Component ---

type BioPageProps struct {
	Title   string
	Profile BioProfileProps
	Details BioDetailsProps
	Links   []BioLinkProps
	Feed    BioFeedProps
	Talk    BioTalkProps
	Invite  BioInviteProps
	Navbar  BioNavbarProps
}

func BioPage(p BioPageProps) Component {
	if p.Title == "" {
		p.Title = "Alex Rivera • Bio, Links, Feed & Talk"
	}

	profileComp := BioProfile(p.Profile)
	detailsComp := BioDetails(p.Details)
	inviteComp := BioInvite(p.Invite)
	feedComp := BioFeed(p.Feed)
	talkComp := BioTalk(p.Talk)
	navbarComp := BioNavbar(p.Navbar)

	var linksSb strings.Builder
	if len(p.Links) == 0 {
		p.Links = []BioLinkProps{
			{
				ID:          "hyperstack",
				Title:       "Hyperstack Framework",
				Description: "HTMX on Steroids • Zero Virtual DOM",
				URL:         "https://github.com/hyperlibs/hyperstack",
				Icon:        "⚡",
				Badge:       "v1.0",
			},
			{
				ID:          "hypersheet",
				Title:       "Hypersheet Live Cloud",
				Description: "Zero-NPM Reactive Spreadsheet Core",
				URL:         "/demo/hypersheet",
				Icon:        "📊",
				Badge:       "Demo",
			},
			{
				ID:          "appzos",
				Title:       "AppzOS Mobile Platform",
				Description: "8 Viewport Modes • 0–9 Z-Index Layers",
				URL:         "/demo/mobile",
				Icon:        "📱",
				Badge:       "Mobile",
			},
		}
	}
	for _, l := range p.Links {
		linksSb.WriteString(BioLinkCard(l).Render())
	}

	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en" class="h-full bg-slate-50 text-slate-900 antialiased font-sans select-none">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no, viewport-fit=cover">
  <link rel="manifest" href="/manifest.json">
  <meta name="mobile-web-app-capable" content="yes">
  <meta name="apple-mobile-web-app-capable" content="yes">
  <meta name="apple-mobile-web-app-status-bar-style" content="default">
  <meta name="apple-mobile-web-app-title" content="Alex Rivera">
  <link rel="apple-touch-icon" href="/icons/icon-192.svg">
  <meta name="theme-color" content="#f8fafc">
  <title>%s</title>
  <script src="/vendor/htmx.min.js"></script>
  <script src="/htmx-bolt.js"></script>
  <script src="/htmx-vibe.js"></script>
  <script src="/appz.js"></script>
  <script src="/tiny-llm.js"></script>
  <script src="https://cdn.tailwindcss.com"></script>
  <link rel="stylesheet" href="/styles.css">
  <style>
    body {
      padding-top: env(safe-area-inset-top);
      padding-bottom: env(safe-area-inset-bottom);
      padding-left: env(safe-area-inset-left);
      padding-right: env(safe-area-inset-right);
    }
    .custom-scroll::-webkit-scrollbar { width: 3px; }
    .custom-scroll::-webkit-scrollbar-track { background: transparent; }
    .custom-scroll::-webkit-scrollbar-thumb { background: rgba(148, 163, 184, 0.25); border-radius: 3px; }
  </style>
</head>
<body class="min-h-full flex flex-col items-center justify-start bg-slate-50 text-slate-900">
  <app-shellui view-mode="Mob-P" class="w-full max-w-sm flex-1 flex flex-col items-center justify-start pb-28">

    <!-- Top App Bar with Role Switcher & Invite Trigger -->
    <app-bar class="w-full sticky top-0 z-30 bg-white/90 backdrop-blur-md border-b border-slate-200/70 px-3.5 h-11 flex items-center justify-between">
      <div class="flex items-center gap-2">
        <span class="text-xs font-bold text-slate-900 flex items-center gap-1">
          Alex Rivera <span class="text-blue-500 text-[10px]">✓</span>
        </span>
      </div>
      <div class="flex items-center gap-2">
        <button onclick="openInviteModal()" class="px-2 py-0.5 rounded-lg bg-indigo-50 text-indigo-700 text-[11px] font-bold border border-indigo-200">
          👥 Invite
        </button>
      </div>
    </app-bar>

    <main class="w-full flex flex-col space-y-2.5 px-3 pt-2.5">
      <!-- Profile Header -->
      %s

      <!-- 4 Segmented Tabs: Bio | Links | Feed | Talk -->
      <div class="flex items-center p-0.5 bg-slate-200/60 rounded-xl text-xs font-semibold">
        <button onclick="switchTab('bio')" id="tab-btn-bio" 
                class="flex-1 py-1 rounded-lg bg-white text-slate-900 font-bold transition-all shadow-xs flex items-center justify-center gap-1">
          <span>👤</span> <span>Bio</span>
        </button>
        <button onclick="switchTab('links')" id="tab-btn-links" 
                class="flex-1 py-1 rounded-lg text-slate-600 hover:text-slate-900 transition-all flex items-center justify-center gap-1">
          <span>🔗</span> <span>Links</span>
        </button>
        <button onclick="switchTab('feed')" id="tab-btn-feed" 
                class="flex-1 py-1 rounded-lg text-slate-600 hover:text-slate-900 transition-all flex items-center justify-center gap-1">
          <span>📝</span> <span>Feed</span>
        </button>
        <button onclick="switchTab('talk')" id="tab-btn-talk" 
                class="flex-1 py-1 rounded-lg text-slate-600 hover:text-slate-900 transition-all flex items-center justify-center gap-1">
          <span>💬</span> <span>Talk</span>
        </button>
      </div>

      <!-- Tab 1: Bio Details -->
      <section id="pane-bio" class="space-y-2">
        %s
      </section>

      <!-- Tab 2: Links Stack -->
      <section id="pane-links" class="hidden space-y-1.5">
        <div class="relative">
          <input type="text" placeholder="Filter links..." class="w-full bg-white border border-slate-200/80 rounded-xl pl-7 pr-3 py-1.5 text-xs text-slate-900 placeholder:text-slate-400 focus:outline-none focus:border-indigo-500 shadow-2xs" />
          <span class="absolute left-2.5 top-1.5 text-[11px] text-slate-400">🔍</span>
        </div>
        <div class="space-y-1.5">
          %s
        </div>
      </section>

      <!-- Tab 3: Feed -->
      <section id="pane-feed" class="hidden space-y-2">
        %s
      </section>

      <!-- Tab 4: Talk -->
      <section id="pane-talk" class="hidden space-y-2">
        %s
      </section>

      <footer class="pt-2 pb-1 text-center">
        <p class="text-[10px] text-slate-400 font-mono">Built with Hyperstack • Zero VDOM</p>
      </footer>
    </main>

    <!-- Floating 2x App Navbar -->
    %s

    <!-- Invite & Collections Modal -->
    %s

  </app-shellui>
</body>
</html>`, p.Title, profileComp.Render(), detailsComp.Render(), linksSb.String(), feedComp.Render(), talkComp.Render(), navbarComp.Render(), inviteComp.Render())

	return Component{html: html}
}
