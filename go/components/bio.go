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

// --- 5. Bio Talk Component (WhatsApp Chat with Reply Comment Threads & Single @list Reply) ---

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
		p.Subtitle = "online • WhatsApp middleman"
	}
	if p.Greeting == "" {
		p.Greeting = "Hey! 👋 I'm <strong>AlexAI</strong>, Alex's clone middleman. I reply strictly according to answers provided by Alex for your group or username."
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
	for i, prompt := range p.QuickPrompts {
		bgClass := "bg-slate-100 hover:bg-emerald-50 hover:text-emerald-800 text-slate-700"
		if i == 0 {
			bgClass = "bg-emerald-50 hover:bg-emerald-100 text-emerald-800 font-bold"
		}
		chipsSb.WriteString(fmt.Sprintf(`
        <button onclick="askClone('%s')" class="px-2.5 py-1 rounded-full %s border border-slate-200 text-xs shrink-0 transition-colors">
          %s
        </button>`, prompt, bgClass, prompt))
	}

	html := fmt.Sprintf(`
<!-- 
  Component: bio-talk.html
  Theme: Modern Threads Question & Reply Tree with Branch Connector Lines & Flutter TabBar Text Tabs
  Strict Constraints:
  1. Modern Threads UI layout with curved vertical branch connector lines linking questions to replies.
  2. Dont allow multiple answers in single reply (strict 1 answer, 1 @list tag).
  3. Each reply can have only one @list tagging (@all, @family, @friends, @bffs, @work, @marcus).
  4. Zero scrollbars everywhere, native Flutter mobile ergonomics.
  5. 2x font size scaling for supreme readability while maintaining compact layout.
  6. Zero overlapping elements, authentic Threads/WhatsApp mobile styling, pristine breathing room.
-->
<div id="talk-container" 
     class="flex flex-col h-[calc(100vh-210px)] min-h-[520px] max-h-[720px] flex-1 rounded-2xl bg-white border border-slate-200/90 overflow-hidden shadow-sm">
  
  <!-- Modern App Bar Header (Clean, Professional, Zero Collision & 2x Readability) -->
  <div class="px-3 py-2 bg-[#075e54] text-white flex items-center justify-between shrink-0 shadow-xs select-none">
    <div class="flex items-center gap-2 min-w-0">
      <!-- Avatar with Active Status Dot -->
      <div class="relative shrink-0">
        <div class="w-8 h-8 rounded-full bg-gradient-to-tr from-emerald-400 to-teal-200 flex items-center justify-center text-slate-900 text-xs font-black shadow-xs">
          A
        </div>
        <span class="w-2 h-2 rounded-full bg-emerald-400 absolute bottom-0 right-0 border-2 border-[#075e54]"></span>
      </div>

      <!-- Contact Info (2x Readable Font, Zero Truncation) -->
      <div class="min-w-0">
        <div class="flex items-center gap-1.5">
          <span class="text-sm font-bold whitespace-nowrap text-white">%s</span>
          <span id="clone-llm-badge" class="hidden sm:inline-flex text-[9px] bg-emerald-400/20 text-emerald-200 font-mono px-1 py-0.2 rounded items-center gap-1 font-bold">
            <span class="w-1.5 h-1.5 rounded-full bg-emerald-400 animate-pulse"></span>
            <span>Offline LLM</span>
          </span>
        </div>
        <p class="text-[11px] text-emerald-100/90 truncate flex items-center gap-1">
          <span>online</span> • <span class="font-mono text-[10px]">%s</span>
          <span id="clone-role-badge" class="text-[9px] bg-white/20 text-white font-mono px-1 py-0.2 rounded font-bold">Public</span>
        </p>
      </div>
    </div>

    <!-- Right Controls: Username identity badge & actions (Zero Overlap) -->
    <div class="flex items-center gap-1.5 shrink-0">
      <!-- User Identification Badge (Can change once a year) -->
      <button onclick="openUsernameModal()" id="talk-user-badge-btn" type="button"
              class="px-2 py-0.5 rounded-lg bg-white/10 hover:bg-white/20 border border-white/20 text-xs font-mono font-bold text-white flex items-center gap-1 transition-colors"
              title="Your username for identification (can be changed once a year)">
        <span class="w-1.5 h-1.5 rounded-full bg-emerald-400"></span>
        <span id="active-user-handle">@visitor</span>
        <span class="text-[9px] text-white/70">✏️</span>
      </button>

      <!-- User Question Credits Badge -->
      <div id="talk-user-credits-badge" class="px-1.5 py-0.5 rounded-lg bg-amber-400/20 border border-amber-300/40 text-xs font-mono font-bold text-amber-200 flex items-center gap-0.5" title="Your Question Credits: 1 credit required to ask. Liking Alex's answer awards 1 credit to Alex.">
        <span>🪙</span>
        <span id="talk-user-credits">3</span>
        <span class="text-[9px] text-amber-200/80">cr</span>
      </div>

      <!-- Alex Creator Credits Badge -->
      <div id="talk-alex-credits-badge" class="hidden sm:flex px-1.5 py-0.5 rounded-lg bg-emerald-400/20 border border-emerald-300/40 text-xs font-mono font-bold text-emerald-200 items-center gap-0.5" title="Alex's total credits earned from answer likes">
        <span>👑</span>
        <span id="talk-alex-credits">0</span>
        <span class="text-[9px] text-emerald-200/80">earned</span>
      </div>

      <!-- Teach Trigger -->
      <button onclick="toggleTeachDrawer()" id="teach-drawer-btn" type="button"
              class="text-xs font-bold text-white bg-white/15 hover:bg-white/25 px-2 py-0.5 rounded-lg transition-colors flex items-center gap-1" title="Teach new facts to offline LLM">
        <span>💡</span> <span class="hidden sm:inline">Teach</span>
      </button>
    </div>
  </div>

  <!-- Clean Professional Text Tabs Bar (Flutter TabBar Style, 2x Font Size, Zero Line Wrapping) -->
  <div class="bg-white border-b border-slate-200/80 flex items-stretch px-2 shrink-0 select-none">
    <button type="button" onclick="switchTalkSubTab('talk')" id="talk-subtab-talk" 
            class="flex-1 py-2.5 text-center text-sm font-bold border-b-2 border-[#008069] text-[#008069] transition-all flex items-center justify-center gap-1.5 whitespace-nowrap">
      <span>Talk</span>
      <span id="talk-unanswered-badge" class="bg-amber-400 text-slate-900 rounded-full px-2 py-0.5 text-xs font-bold">2</span>
    </button>
    <button type="button" onclick="switchTalkSubTab('talkback')" id="talk-subtab-talkback" 
            class="flex-1 py-2.5 text-center text-sm font-semibold border-b-2 border-transparent text-slate-500 hover:text-slate-800 transition-all flex items-center justify-center gap-1.5 whitespace-nowrap">
      <span>Talk back</span>
      <span class="hidden sm:inline text-xs text-slate-400 font-normal">Thread Replies</span>
    </button>
    <button type="button" onclick="toggleInboxDrawer()" id="inbox-drawer-btn" 
            class="py-2.5 px-3.5 text-center text-sm font-semibold border-b-2 border-transparent text-slate-500 hover:text-slate-800 transition-all flex items-center justify-center gap-1.5 whitespace-nowrap" title="Question Queue">
      <span>Inbox</span>
      <span id="inbox-pending-count" class="bg-emerald-600 text-white rounded-full px-2 py-0.5 text-xs font-bold">2</span>
    </button>
  </div>

  <!-- Collapsible Alex's Question Inbox Drawer (Compatibility) -->
  <div id="inbox-drawer" class="hidden px-3.5 py-2.5 bg-amber-50/80 border-b border-amber-200 text-sm space-y-2 animate-slide-down">
    <div class="flex items-center justify-between text-xs">
      <span class="font-bold text-amber-900 flex items-center gap-1.5">
        <span>👑</span> <span>Alex's Question Queue (Middleman Control)</span>
      </span>
      <button type="button" onclick="toggleInboxDrawer()" class="text-amber-600 hover:text-amber-900 font-bold text-sm">✕</button>
    </div>
    <div id="inbox-answer-box" class="p-2.5 rounded-xl bg-white border border-amber-200/90 shadow-2xs space-y-1">
      <div class="text-xs font-bold text-slate-800">Middleman Question Queue Active</div>
      <button type="button" onclick="submitInboxAnswer()" class="hidden">Save</button>
    </div>
  </div>

  <!-- Collapsible Self-Learning Teach Drawer -->
  <div id="teach-drawer" class="hidden px-3.5 py-2.5 bg-emerald-50/90 border-b border-emerald-200 text-sm space-y-2 animate-slide-down">
    <div class="flex items-center justify-between text-xs">
      <span class="font-bold text-emerald-900 flex items-center gap-1.5">
        <span>🧠</span> <span>Teach Tiny LLM (On-Device Self-Learning)</span>
      </span>
      <span id="llm-telemetry-pill" class="font-mono text-xs text-emerald-800 bg-white/90 px-2 py-0.5 rounded border border-emerald-300">
        Steps: <span id="telemetry-step">0</span> | Loss: <span id="telemetry-loss">0.038</span>
      </span>
    </div>
    <form onsubmit="handleTeachSubmit(event)" class="flex items-center gap-2">
      <input type="text" id="teach-input" placeholder="e.g. Alex loves building mechanical keyboards with Boba U4T switches" 
             class="flex-1 bg-white border border-emerald-300 rounded-xl px-3 py-1.5 text-sm text-slate-900 placeholder:text-slate-400 focus:outline-none focus:border-emerald-600 shadow-2xs" />
      <button type="submit" class="px-3 py-1.5 rounded-xl bg-[#008069] hover:bg-[#075e54] text-white font-bold text-xs shrink-0 transition-all shadow-xs">
        Teach & Train
      </button>
    </form>
  </div>

  <!-- ===================================================================== -->
  <!-- SUB-PANE 1: TALK (Threads Question & Reply Stream + Recent Unanswered Questions) -->
  <!-- ===================================================================== -->
  <div id="subpane-talk" class="flex-1 flex flex-col min-h-0 bg-white overflow-hidden">
    
    <!-- RECENT UNANSWERED QUESTIONS SECTION (LISTED FIRST, THREAD STYLE) -->
    <div id="unanswered-questions-section" class="border-b border-slate-200 bg-amber-50/60 p-2 space-y-1 shrink-0 max-h-32 overflow-y-auto custom-scroll scrollbar-none shadow-2xs select-none">
      <div class="flex items-center justify-between text-xs">
        <div class="flex items-center gap-1.5 font-bold text-slate-800">
          <span class="w-2 h-2 rounded-full bg-amber-500 animate-pulse"></span>
          <span>Recent Unanswered Questions</span>
          <span id="unanswered-count-badge" class="px-1.5 py-0.2 rounded-full bg-amber-200 text-amber-900 text-[10px] font-mono font-bold">2</span>
        </div>
        <span class="text-[10px] text-slate-500 font-mono">AlexAI silent until answered</span>
      </div>
      <p class="text-[11px] text-slate-500 leading-tight">
        Questions waiting for Alex's answer. Click <strong>Talk back</strong> on any question to reply with a single @list tag!
      </p>

      <!-- Dynamic Unanswered Questions Cards List -->
      <div id="unanswered-questions-list" class="space-y-1.5 pt-0.5">
        <!-- Rendered dynamically by renderUnansweredQuestions() -->
      </div>
    </div>

    <!-- Threads Question & Reply Messages Stream (Clean White Threads UI, Tree Connector Lines, Zero Scrollbars) -->
    <div id="talk-stream" class="flex-1 min-h-0 overflow-y-auto p-3 space-y-3 custom-scroll scrollbar-none text-sm relative"
         style="background-color: #efeae2; background-image: radial-gradient(#d5d1c8 1px, transparent 1px); background-size: 24px 24px;">
      
      <!-- Date Divider -->
      <div class="flex justify-center my-1 select-none">
        <span class="bg-white/95 border border-slate-200/80 text-[10px] uppercase font-bold text-slate-500 px-2.5 py-0.5 rounded-full shadow-2xs">
          TODAY
        </span>
      </div>

      <!-- Welcome Thread Post from AlexAI -->
      <article class="bg-white rounded-2xl p-3 border border-slate-200/80 shadow-2xs space-y-2">
        <div class="flex items-start gap-2.5">
          <div class="flex flex-col items-center shrink-0">
            <div class="w-8 h-8 rounded-full bg-gradient-to-tr from-emerald-500 to-teal-300 flex items-center justify-center text-slate-900 text-xs font-black shadow-xs">
              A
            </div>
            <div class="w-0.5 flex-1 bg-slate-200 my-1 rounded-full min-h-[20px]"></div>
          </div>
          <div class="flex-1 min-w-0">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-1.5 flex-wrap">
                <span class="text-xs font-bold text-slate-900">AlexAI (Digital Clone)</span>
                <span class="text-blue-500 text-[10px]">✓</span>
                <span class="text-[10px] text-slate-400 font-mono">10:40 AM •</span>
                <span class="text-[10px] text-slate-500 font-bold">Middleman Proxy</span>
              </div>
              <span class="text-emerald-600 font-bold text-xs font-mono">✓✓</span>
            </div>
            <p class="text-sm text-slate-800 leading-relaxed font-sans mt-1">
              %s
            </p>
            <div class="mt-2 text-[11px] text-slate-600 bg-slate-50 p-2 rounded-xl border border-slate-100 space-y-0.5">
              <span class="font-bold text-slate-700">💬 Reply Comment Threads:</span>
              <p class="leading-relaxed">Every reply has strictly <strong>ONE @list tag</strong>. Alex can reply multiple times per question, but cannot mix multiple answers into one reply.</p>
            </div>
            <div class="flex items-center justify-between text-xs text-slate-500 pt-2 select-none border-t border-slate-100 mt-2">
              <div class="flex items-center gap-3.5 text-[11px]">
                <button type="button" onclick="likeAnswerInWhatsApp('qa-init', 'all', this)" class="like-answer-btn flex items-center gap-1 text-slate-600 hover:text-rose-600 transition-colors" title="Like Alex's answer (awards 1 credit to Alex)">
                  <span>♡</span> <span class="font-bold text-rose-600">Like (+1 cr)</span>
                </button>
                <button type="button" class="flex items-center gap-1 hover:text-[#008069] transition-colors">
                  <span>💬</span> <span>1</span>
                </button>
                <button type="button" class="hover:text-slate-700"><span>🔁</span></button>
                <button type="button" class="hover:text-slate-700"><span>➤</span></button>
              </div>
              <span class="text-[10px] text-emerald-700 font-bold font-mono">● Welcome</span>
            </div>
          </div>
        </div>
      </article>

      <!-- Preloaded Thread 1: Where are you now? -->
      <article class="bg-white rounded-2xl p-3 border border-slate-200/80 shadow-2xs space-y-2">
        <!-- Question Row -->
        <div class="flex items-start gap-2.5">
          <div class="flex flex-col items-center shrink-0">
            <div class="w-8 h-8 rounded-full bg-slate-200 flex items-center justify-center font-bold text-xs text-slate-700">
              S
            </div>
            <!-- Tree Connector Line -->
            <div class="w-0.5 flex-1 bg-slate-200 my-1 rounded-full min-h-[30px]"></div>
          </div>
          <div class="flex-1 min-w-0">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-1.5 flex-wrap">
                <span class="text-xs font-bold text-slate-900">@sarah</span>
                <span class="text-[10px] text-slate-400 font-mono">10:15 AM</span>
                <span class="text-[9px] font-mono px-1.5 py-0.2 rounded bg-slate-100 text-slate-600 font-medium">Question</span>
              </div>
              <span class="text-slate-400 text-xs">⋯</span>
            </div>
            <p class="text-xs text-slate-900 leading-snug mt-1 font-medium font-sans">
              Where are you now? 📍
            </p>
            <div class="flex items-center justify-between text-xs text-slate-500 pt-1.5 select-none">
              <div class="flex items-center gap-3 text-[11px]">
                <button type="button" class="flex items-center gap-1 hover:text-rose-600 transition-colors"><span>♡</span> <span>2</span></button>
                <button type="button" class="flex items-center gap-1 hover:text-[#008069] transition-colors"><span>💬</span> <span>1</span></button>
                <button type="button" class="hover:text-slate-700"><span>🔁</span></button>
                <button type="button" class="hover:text-slate-700"><span>➤</span></button>
              </div>
              <button type="button" onclick="replyToQuestionInWhatsApp('qa-preload-1', '@sarah', 'Where are you now? 📍')" class="text-xs font-semibold text-slate-600 hover:text-[#008069] transition-colors">
                Reply
              </button>
            </div>
          </div>
        </div>

        <!-- Nested Reply Row (Connected by Tree Line) -->
        <div class="flex items-start gap-2.5 pt-0.5">
          <div class="flex items-start shrink-0 relative">
            <div class="absolute -top-3 left-[15px] w-3 h-4 border-b-2 border-l-2 border-slate-200 rounded-bl-lg pointer-events-none"></div>
            <div class="relative ml-4">
              <div class="w-7 h-7 rounded-full bg-gradient-to-tr from-emerald-500 to-teal-300 flex items-center justify-center text-slate-900 text-xs font-black shadow-xs">A</div>
              <span class="w-1.5 h-1.5 rounded-full bg-emerald-400 absolute bottom-0 right-0 border border-white"></span>
            </div>
          </div>
          <div class="flex-1 min-w-0 bg-slate-50/80 rounded-xl p-2.5 border border-slate-100/90 shadow-2xs">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-1.5 flex-wrap">
                <span class="text-xs font-bold text-slate-900">AlexAI</span>
                <span class="text-blue-500 text-[10px]">✓</span>
                <span class="text-[10px] text-slate-400 font-mono">10:16 AM •</span>
                <span class="text-[10px] text-slate-600 font-bold">Author</span>
                <span class="text-[9px] bg-[#008069] text-white font-mono px-1.5 py-0.2 rounded font-bold">@all</span>
              </div>
              <div class="flex items-center gap-1 text-xs"><span class="text-emerald-600 font-bold">✓✓</span></div>
            </div>
            <p class="text-sm text-slate-800 leading-relaxed font-sans mt-1">
              im at work, chennai
            </p>
            <div class="flex items-center justify-between text-xs text-slate-500 pt-1.5 select-none border-t border-slate-100 mt-2">
              <div class="flex items-center gap-3 text-[11px]">
                <button type="button" onclick="likeAnswerInWhatsApp('qa-preload-1', 'all', this)" class="like-answer-btn flex items-center gap-1 text-slate-600 hover:text-rose-600 transition-colors" title="Like Alex's answer (awards 1 credit to Alex)">
                  <span>♡</span> <span class="font-bold text-rose-600">Like (+1 cr)</span>
                </button>
                <button type="button" class="flex items-center gap-1 hover:text-[#008069] transition-colors"><span>💬</span> <span>1</span></button>
              </div>
              <button type="button" onclick="replyToQuestionInWhatsApp('qa-preload-1', '@sarah', 'Where are you now? 📍')" class="text-xs font-semibold text-[#008069] hover:underline">
                Reply in thread
              </button>
            </div>
          </div>
        </div>

        <!-- Stacked Avatars & Show Replies Link -->
        <div class="flex items-center gap-2 pt-1 pl-6 select-none">
          <div class="flex -space-x-1.5 items-center">
            <div class="w-4 h-4 rounded-full bg-slate-300 border border-white"></div>
            <div class="w-4 h-4 rounded-full bg-emerald-400 border border-white"></div>
          </div>
          <button type="button" onclick="selectQuestionForTalkBack('Where are you now? 📍')" class="text-[11px] text-slate-400 hover:text-slate-700 font-medium">
            Show replies
          </button>
        </div>
      </article>

      <!-- Preloaded Thread 2: Coffee routine -->
      <article class="bg-white rounded-2xl p-3 border border-slate-200/80 shadow-2xs space-y-2">
        <div class="flex items-start gap-2.5">
          <div class="flex flex-col items-center shrink-0">
            <div class="w-8 h-8 rounded-full bg-amber-200 flex items-center justify-center font-bold text-xs text-amber-800">M</div>
            <div class="w-0.5 flex-1 bg-slate-200 my-1 rounded-full min-h-[30px]"></div>
          </div>
          <div class="flex-1 min-w-0">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-1.5 flex-wrap">
                <span class="text-xs font-bold text-slate-900">@marcus</span>
                <span class="text-[10px] text-slate-400 font-mono">10:20 AM</span>
                <span class="text-[9px] font-mono px-1.5 py-0.2 rounded bg-slate-100 text-slate-600 font-medium">Question</span>
              </div>
              <span class="text-slate-400 text-xs">⋯</span>
            </div>
            <p class="text-xs text-slate-900 leading-snug mt-1 font-medium font-sans">
              What is your morning coffee routine? ☕️
            </p>
            <div class="flex items-center justify-between text-xs text-slate-500 pt-1.5 select-none">
              <div class="flex items-center gap-3 text-[11px]">
                <button type="button" class="flex items-center gap-1 hover:text-rose-600 transition-colors"><span>♡</span> <span>1</span></button>
                <button type="button" class="flex items-center gap-1 hover:text-[#008069] transition-colors"><span>💬</span> <span>1</span></button>
                <button type="button" class="hover:text-slate-700"><span>🔁</span></button>
                <button type="button" class="hover:text-slate-700"><span>➤</span></button>
              </div>
              <button type="button" onclick="replyToQuestionInWhatsApp('qa-preload-2', '@marcus', 'What is your morning coffee routine? ☕️')" class="text-xs font-semibold text-slate-600 hover:text-[#008069] transition-colors">
                Reply
              </button>
            </div>
          </div>
        </div>

        <div class="flex items-start gap-2.5 pt-0.5">
          <div class="flex items-start shrink-0 relative">
            <div class="absolute -top-3 left-[15px] w-3 h-4 border-b-2 border-l-2 border-slate-200 rounded-bl-lg pointer-events-none"></div>
            <div class="relative ml-4">
              <div class="w-7 h-7 rounded-full bg-gradient-to-tr from-emerald-500 to-teal-300 flex items-center justify-center text-slate-900 text-xs font-black shadow-xs">A</div>
              <span class="w-1.5 h-1.5 rounded-full bg-emerald-400 absolute bottom-0 right-0 border border-white"></span>
            </div>
          </div>
          <div class="flex-1 min-w-0 bg-slate-50/80 rounded-xl p-2.5 border border-slate-100/90 shadow-2xs">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-1.5 flex-wrap">
                <span class="text-xs font-bold text-slate-900">AlexAI</span>
                <span class="text-blue-500 text-[10px]">✓</span>
                <span class="text-[10px] text-slate-400 font-mono">10:21 AM •</span>
                <span class="text-[10px] text-slate-600 font-bold">Author</span>
                <span class="text-[9px] bg-[#008069] text-white font-mono px-1.5 py-0.2 rounded font-bold">@all</span>
              </div>
              <div class="flex items-center gap-1 text-xs"><span class="text-emerald-600 font-bold">✓✓</span></div>
            </div>
            <p class="text-sm text-slate-800 leading-relaxed font-sans mt-1">
              Fresh pour-over using light-roast Ethiopian Yirgacheffe beans, 1:16 ratio on a V60 with Fellow Ode grinder ☕️
            </p>
            <div class="flex items-center justify-between text-xs text-slate-500 pt-1.5 select-none border-t border-slate-100 mt-2">
              <div class="flex items-center gap-3 text-[11px]">
                <button type="button" onclick="likeAnswerInWhatsApp('qa-preload-2', 'all', this)" class="like-answer-btn flex items-center gap-1 text-slate-600 hover:text-rose-600 transition-colors" title="Like Alex's answer (awards 1 credit to Alex)">
                  <span>♡</span> <span class="font-bold text-rose-600">Like (+1 cr)</span>
                </button>
                <button type="button" class="flex items-center gap-1 hover:text-[#008069] transition-colors"><span>💬</span> <span>1</span></button>
              </div>
              <button type="button" onclick="replyToQuestionInWhatsApp('qa-preload-2', '@marcus', 'What is your morning coffee routine? ☕️')" class="text-xs font-semibold text-[#008069] hover:underline">
                Reply in thread
              </button>
            </div>
          </div>
        </div>
      </article>

      <!-- Preloaded Thread 3: Girlfriend query -->
      <article class="bg-white rounded-2xl p-3 border border-slate-200/80 shadow-2xs space-y-2">
        <div class="flex items-start gap-2.5">
          <div class="flex flex-col items-center shrink-0">
            <div class="w-8 h-8 rounded-full bg-rose-200 flex items-center justify-center font-bold text-xs text-rose-800">P</div>
            <div class="w-0.5 flex-1 bg-slate-200 my-1 rounded-full min-h-[30px]"></div>
          </div>
          <div class="flex-1 min-w-0">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-1.5 flex-wrap">
                <span class="text-xs font-bold text-slate-900">@priya</span>
                <span class="text-[10px] text-slate-400 font-mono">10:25 AM</span>
                <span class="text-[9px] font-mono px-1.5 py-0.2 rounded bg-slate-100 text-slate-600 font-medium">Question</span>
              </div>
              <span class="text-slate-400 text-xs">⋯</span>
            </div>
            <p class="text-xs text-slate-900 leading-snug mt-1 font-medium font-sans">
              Do you have a girlfriend? 👀
            </p>
            <div class="flex items-center justify-between text-xs text-slate-500 pt-1.5 select-none">
              <div class="flex items-center gap-3 text-[11px]">
                <button type="button" class="flex items-center gap-1 hover:text-rose-600 transition-colors"><span>♡</span> <span>4</span></button>
                <button type="button" class="flex items-center gap-1 hover:text-[#008069] transition-colors"><span>💬</span> <span>1</span></button>
              </div>
              <button type="button" onclick="replyToQuestionInWhatsApp('qa-preload-3', '@priya', 'Do you have a girlfriend? 👀')" class="text-xs font-semibold text-slate-600 hover:text-[#008069] transition-colors">
                Reply
              </button>
            </div>
          </div>
        </div>

        <div class="flex items-start gap-2.5 pt-0.5">
          <div class="flex items-start shrink-0 relative">
            <div class="absolute -top-3 left-[15px] w-3 h-4 border-b-2 border-l-2 border-slate-200 rounded-bl-lg pointer-events-none"></div>
            <div class="relative ml-4">
              <div class="w-7 h-7 rounded-full bg-gradient-to-tr from-emerald-500 to-teal-300 flex items-center justify-center text-slate-900 text-xs font-black shadow-xs">A</div>
              <span class="w-1.5 h-1.5 rounded-full bg-emerald-400 absolute bottom-0 right-0 border border-white"></span>
            </div>
          </div>
          <div class="flex-1 min-w-0 bg-slate-50/80 rounded-xl p-2.5 border border-slate-100/90 shadow-2xs">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-1.5 flex-wrap">
                <span class="text-xs font-bold text-slate-900">AlexAI</span>
                <span class="text-blue-500 text-[10px]">✓</span>
                <span class="text-[10px] text-slate-400 font-mono">10:26 AM •</span>
                <span class="text-[10px] text-slate-600 font-bold">Author</span>
                <span class="text-[9px] bg-[#008069] text-white font-mono px-1.5 py-0.2 rounded font-bold">@all</span>
              </div>
              <div class="flex items-center gap-1 text-xs"><span class="text-emerald-600 font-bold">✓✓</span></div>
            </div>
            <p class="text-sm text-slate-800 leading-relaxed font-sans mt-1">
              Haha nope, currently single! 100%% focused on engineering open-source systems and building Hyperstack right now 🚀
            </p>
            <div class="flex items-center justify-between text-xs text-slate-500 pt-1.5 select-none border-t border-slate-100 mt-2">
              <div class="flex items-center gap-3 text-[11px]">
                <button type="button" onclick="likeAnswerInWhatsApp('qa-preload-3', 'all', this)" class="like-answer-btn flex items-center gap-1 text-slate-600 hover:text-rose-600 transition-colors" title="Like Alex's answer (awards 1 credit to Alex)">
                  <span>♡</span> <span class="font-bold text-rose-600">Like (+1 cr)</span>
                </button>
                <button type="button" class="flex items-center gap-1 hover:text-[#008069] transition-colors"><span>💬</span> <span>1</span></button>
              </div>
              <button type="button" onclick="replyToQuestionInWhatsApp('qa-preload-3', '@priya', 'Do you have a girlfriend? 👀')" class="text-xs font-semibold text-[#008069] hover:underline">
                Reply in thread
              </button>
            </div>
          </div>
        </div>
      </article>

      <!-- Preloaded Thread 4: Top projects -->
      <article class="bg-white rounded-2xl p-3 border border-slate-200/80 shadow-2xs space-y-2">
        <div class="flex items-start gap-2.5">
          <div class="flex flex-col items-center shrink-0">
            <div class="w-8 h-8 rounded-full bg-blue-200 flex items-center justify-center font-bold text-xs text-blue-800">D</div>
            <div class="w-0.5 flex-1 bg-slate-200 my-1 rounded-full min-h-[30px]"></div>
          </div>
          <div class="flex-1 min-w-0">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-1.5 flex-wrap">
                <span class="text-xs font-bold text-slate-900">@dev_alex</span>
                <span class="text-[10px] text-slate-400 font-mono">10:30 AM</span>
                <span class="text-[9px] font-mono px-1.5 py-0.2 rounded bg-slate-100 text-slate-600 font-medium">Question</span>
              </div>
              <span class="text-slate-400 text-xs">⋯</span>
            </div>
            <p class="text-xs text-slate-900 leading-snug mt-1 font-medium font-sans">
              Top projects you are shipping? 🚀
            </p>
            <div class="flex items-center justify-between text-xs text-slate-500 pt-1.5 select-none">
              <div class="flex items-center gap-3 text-[11px]">
                <button type="button" class="flex items-center gap-1 hover:text-rose-600 transition-colors"><span>♡</span> <span>3</span></button>
                <button type="button" class="flex items-center gap-1 hover:text-[#008069] transition-colors"><span>💬</span> <span>1</span></button>
              </div>
              <button type="button" onclick="replyToQuestionInWhatsApp('qa-preload-4', '@dev_alex', 'Top projects you are shipping? 🚀')" class="text-xs font-semibold text-slate-600 hover:text-[#008069] transition-colors">
                Reply
              </button>
            </div>
          </div>
        </div>

        <div class="flex items-start gap-2.5 pt-0.5">
          <div class="flex items-start shrink-0 relative">
            <div class="absolute -top-3 left-[15px] w-3 h-4 border-b-2 border-l-2 border-slate-200 rounded-bl-lg pointer-events-none"></div>
            <div class="relative ml-4">
              <div class="w-7 h-7 rounded-full bg-gradient-to-tr from-emerald-500 to-teal-300 flex items-center justify-center text-slate-900 text-xs font-black shadow-xs">A</div>
              <span class="w-1.5 h-1.5 rounded-full bg-emerald-400 absolute bottom-0 right-0 border border-white"></span>
            </div>
          </div>
          <div class="flex-1 min-w-0 bg-slate-50/80 rounded-xl p-2.5 border border-slate-100/90 shadow-2xs">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-1.5 flex-wrap">
                <span class="text-xs font-bold text-slate-900">AlexAI</span>
                <span class="text-blue-500 text-[10px]">✓</span>
                <span class="text-[10px] text-slate-400 font-mono">10:31 AM •</span>
                <span class="text-[10px] text-slate-600 font-bold">Author</span>
                <span class="text-[9px] bg-[#008069] text-white font-mono px-1.5 py-0.2 rounded font-bold">@all</span>
              </div>
              <div class="flex items-center gap-1 text-xs"><span class="text-emerald-600 font-bold">✓✓</span></div>
            </div>
            <p class="text-sm text-slate-800 leading-relaxed font-sans mt-1">
              Shipping Hyperstack v1.1 zero-dependency hypermedia framework and AppzOS mobile engine ⚡️
            </p>
            <div class="flex items-center justify-between text-xs text-slate-500 pt-1.5 select-none border-t border-slate-100 mt-2">
              <div class="flex items-center gap-3 text-[11px]">
                <button type="button" onclick="likeAnswerInWhatsApp('qa-preload-4', 'all', this)" class="like-answer-btn flex items-center gap-1 text-slate-600 hover:text-rose-600 transition-colors" title="Like Alex's answer (awards 1 credit to Alex)">
                  <span>♡</span> <span class="font-bold text-rose-600">Like (+1 cr)</span>
                </button>
                <button type="button" class="flex items-center gap-1 hover:text-[#008069] transition-colors"><span>💬</span> <span>1</span></button>
              </div>
              <button type="button" onclick="replyToQuestionInWhatsApp('qa-preload-4', '@dev_alex', 'Top projects you are shipping? 🚀')" class="text-xs font-semibold text-[#008069] hover:underline">
                Reply in thread
              </button>
            </div>
          </div>
        </div>
      </article>

      <!-- Preloaded Thread 5: Keyboard switches -->
      <article class="bg-white rounded-2xl p-3 border border-slate-200/80 shadow-2xs space-y-2">
        <div class="flex items-start gap-2.5">
          <div class="flex flex-col items-center shrink-0">
            <div class="w-8 h-8 rounded-full bg-purple-200 flex items-center justify-center font-bold text-xs text-purple-800">K</div>
            <div class="w-0.5 flex-1 bg-slate-200 my-1 rounded-full min-h-[30px]"></div>
          </div>
          <div class="flex-1 min-w-0">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-1.5 flex-wrap">
                <span class="text-xs font-bold text-slate-900">@keeb_fan</span>
                <span class="text-[10px] text-slate-400 font-mono">10:34 AM</span>
                <span class="text-[9px] font-mono px-1.5 py-0.2 rounded bg-slate-100 text-slate-600 font-medium">Question</span>
              </div>
              <span class="text-slate-400 text-xs">⋯</span>
            </div>
            <p class="text-xs text-slate-900 leading-snug mt-1 font-medium font-sans">
              What mechanical keyboard switches do you use? ⌨️
            </p>
            <div class="flex items-center justify-between text-xs text-slate-500 pt-1.5 select-none">
              <div class="flex items-center gap-3 text-[11px]">
                <button type="button" class="flex items-center gap-1 hover:text-rose-600 transition-colors"><span>♡</span> <span>5</span></button>
                <button type="button" class="flex items-center gap-1 hover:text-[#008069] transition-colors"><span>💬</span> <span>1</span></button>
              </div>
              <button type="button" onclick="replyToQuestionInWhatsApp('qa-preload-5', '@keeb_fan', 'What mechanical keyboard switches do you use? ⌨️')" class="text-xs font-semibold text-slate-600 hover:text-[#008069] transition-colors">
                Reply
              </button>
            </div>
          </div>
        </div>

        <div class="flex items-start gap-2.5 pt-0.5">
          <div class="flex items-start shrink-0 relative">
            <div class="absolute -top-3 left-[15px] w-3 h-4 border-b-2 border-l-2 border-slate-200 rounded-bl-lg pointer-events-none"></div>
            <div class="relative ml-4">
              <div class="w-7 h-7 rounded-full bg-gradient-to-tr from-emerald-500 to-teal-300 flex items-center justify-center text-slate-900 text-xs font-black shadow-xs">A</div>
              <span class="w-1.5 h-1.5 rounded-full bg-emerald-400 absolute bottom-0 right-0 border border-white"></span>
            </div>
          </div>
          <div class="flex-1 min-w-0 bg-slate-50/80 rounded-xl p-2.5 border border-slate-100/90 shadow-2xs">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-1.5 flex-wrap">
                <span class="text-xs font-bold text-slate-900">AlexAI</span>
                <span class="text-blue-500 text-[10px]">✓</span>
                <span class="text-[10px] text-slate-400 font-mono">10:35 AM •</span>
                <span class="text-[10px] text-slate-600 font-bold">Author</span>
                <span class="text-[9px] bg-[#008069] text-white font-mono px-1.5 py-0.2 rounded font-bold">@all</span>
              </div>
              <div class="flex items-center gap-1 text-xs"><span class="text-emerald-600 font-bold">✓✓</span></div>
            </div>
            <p class="text-sm text-slate-800 leading-relaxed font-sans mt-1">
              Custom Boba U4T tactile switches lubed with TriboSys 3203 on a gasket-mounted aluminum board 🎧
            </p>
            <div class="flex items-center justify-between text-xs text-slate-500 pt-1.5 select-none border-t border-slate-100 mt-2">
              <div class="flex items-center gap-3 text-[11px]">
                <button type="button" onclick="likeAnswerInWhatsApp('qa-preload-5', 'all', this)" class="like-answer-btn flex items-center gap-1 text-slate-600 hover:text-rose-600 transition-colors" title="Like Alex's answer (awards 1 credit to Alex)">
                  <span>♡</span> <span class="font-bold text-rose-600">Like (+1 cr)</span>
                </button>
                <button type="button" class="flex items-center gap-1 hover:text-[#008069] transition-colors"><span>💬</span> <span>1</span></button>
              </div>
              <button type="button" onclick="replyToQuestionInWhatsApp('qa-preload-5', '@keeb_fan', 'What mechanical keyboard switches do you use? ⌨️')" class="text-xs font-semibold text-[#008069] hover:underline">
                Reply in thread
              </button>
            </div>
          </div>
        </div>
      </article>

      <!-- Preloaded Thread 6: Tech stack -->
      <article class="bg-white rounded-2xl p-3 border border-slate-200/80 shadow-2xs space-y-2">
        <div class="flex items-start gap-2.5">
          <div class="flex flex-col items-center shrink-0">
            <div class="w-8 h-8 rounded-full bg-amber-200 flex items-center justify-center font-bold text-xs text-amber-900">G</div>
            <div class="w-0.5 flex-1 bg-slate-200 my-1 rounded-full min-h-[30px]"></div>
          </div>
          <div class="flex-1 min-w-0">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-1.5 flex-wrap">
                <span class="text-xs font-bold text-slate-900">@persong</span>
                <span class="text-[10px] text-slate-400 font-mono">10:38 AM</span>
                <span class="text-[9px] font-mono px-1.5 py-0.2 rounded bg-amber-100 text-amber-800 font-medium">Person G</span>
              </div>
              <span class="text-slate-400 text-xs">⋯</span>
            </div>
            <p class="text-xs text-slate-900 leading-snug mt-1 font-medium font-sans">
              Favorite tech stack for 2026? 💻
            </p>
            <div class="flex items-center justify-between text-xs text-slate-500 pt-1.5 select-none">
              <div class="flex items-center gap-3 text-[11px]">
                <button type="button" class="flex items-center gap-1 hover:text-rose-600 transition-colors"><span>♡</span> <span>3</span></button>
                <button type="button" class="flex items-center gap-1 hover:text-[#008069] transition-colors"><span>💬</span> <span>1</span></button>
              </div>
              <button type="button" onclick="replyToQuestionInWhatsApp('qa-preload-6', '@persong', 'Favorite tech stack for 2026? 💻')" class="text-xs font-semibold text-slate-600 hover:text-[#008069] transition-colors">
                Reply
              </button>
            </div>
          </div>
        </div>

        <div class="flex items-start gap-2.5 pt-0.5">
          <div class="flex items-start shrink-0 relative">
            <div class="absolute -top-3 left-[15px] w-3 h-4 border-b-2 border-l-2 border-slate-200 rounded-bl-lg pointer-events-none"></div>
            <div class="relative ml-4">
              <div class="w-7 h-7 rounded-full bg-gradient-to-tr from-emerald-500 to-teal-300 flex items-center justify-center text-slate-900 text-xs font-black shadow-xs">A</div>
              <span class="w-1.5 h-1.5 rounded-full bg-emerald-400 absolute bottom-0 right-0 border border-white"></span>
            </div>
          </div>
          <div class="flex-1 min-w-0 bg-slate-50/80 rounded-xl p-2.5 border border-slate-100/90 shadow-2xs">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-1.5 flex-wrap">
                <span class="text-xs font-bold text-slate-900">AlexAI</span>
                <span class="text-blue-500 text-[10px]">✓</span>
                <span class="text-[10px] text-slate-400 font-mono">10:39 AM •</span>
                <span class="text-[10px] text-slate-600 font-bold">Author</span>
                <span class="text-[9px] bg-[#008069] text-white font-mono px-1.5 py-0.2 rounded font-bold">@all</span>
              </div>
              <div class="flex items-center gap-1 text-xs"><span class="text-emerald-600 font-bold">✓✓</span></div>
            </div>
            <p class="text-sm text-slate-800 leading-relaxed font-sans mt-1">
              Zero Node.js, 100%% Bun & Go with Hypermedia partials! Blazing fast single binaries 🔥
            </p>
            <div class="flex items-center justify-between text-xs text-slate-500 pt-1.5 select-none border-t border-slate-100 mt-2">
              <div class="flex items-center gap-3 text-[11px]">
                <button type="button" onclick="likeAnswerInWhatsApp('qa-preload-6', 'all', this)" class="like-answer-btn flex items-center gap-1 text-slate-600 hover:text-rose-600 transition-colors" title="Like Alex's answer (awards 1 credit to Alex)">
                  <span>♡</span> <span class="font-bold text-rose-600">Like (+1 cr)</span>
                </button>
                <button type="button" class="flex items-center gap-1 hover:text-[#008069] transition-colors"><span>💬</span> <span>1</span></button>
              </div>
              <button type="button" onclick="replyToQuestionInWhatsApp('qa-preload-6', '@persong', 'Favorite tech stack for 2026? 💻')" class="text-xs font-semibold text-[#008069] hover:underline">
                Reply in thread
              </button>
            </div>
          </div>
        </div>
      </article>

      <!-- Preloaded Thread 7: Systems conference -->
      <article class="bg-white rounded-2xl p-3 border border-slate-200/80 shadow-2xs space-y-2">
        <div class="flex items-start gap-2.5">
          <div class="flex flex-col items-center shrink-0">
            <div class="w-8 h-8 rounded-full bg-teal-200 flex items-center justify-center font-bold text-xs text-teal-800">O</div>
            <div class="w-0.5 flex-1 bg-slate-200 my-1 rounded-full min-h-[30px]"></div>
          </div>
          <div class="flex-1 min-w-0">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-1.5 flex-wrap">
                <span class="text-xs font-bold text-slate-900">@organizer</span>
                <span class="text-[10px] text-slate-400 font-mono">10:41 AM</span>
                <span class="text-[9px] font-mono px-1.5 py-0.2 rounded bg-slate-100 text-slate-600 font-medium">Question</span>
              </div>
              <span class="text-slate-400 text-xs">⋯</span>
            </div>
            <p class="text-xs text-slate-900 leading-snug mt-1 font-medium font-sans">
              Are you free for speaking at the systems conference? 🎤
            </p>
            <div class="flex items-center justify-between text-xs text-slate-500 pt-1.5 select-none">
              <div class="flex items-center gap-3 text-[11px]">
                <button type="button" class="flex items-center gap-1 hover:text-rose-600 transition-colors"><span>♡</span> <span>2</span></button>
                <button type="button" class="flex items-center gap-1 hover:text-[#008069] transition-colors"><span>💬</span> <span>1</span></button>
              </div>
              <button type="button" onclick="replyToQuestionInWhatsApp('qa-preload-7', '@organizer', 'Are you free for speaking at the systems conference? 🎤')" class="text-xs font-semibold text-slate-600 hover:text-[#008069] transition-colors">
                Reply
              </button>
            </div>
          </div>
        </div>

        <div class="flex items-start gap-2.5 pt-0.5">
          <div class="flex items-start shrink-0 relative">
            <div class="absolute -top-3 left-[15px] w-3 h-4 border-b-2 border-l-2 border-slate-200 rounded-bl-lg pointer-events-none"></div>
            <div class="relative ml-4">
              <div class="w-7 h-7 rounded-full bg-gradient-to-tr from-emerald-500 to-teal-300 flex items-center justify-center text-slate-900 text-xs font-black shadow-xs">A</div>
              <span class="w-1.5 h-1.5 rounded-full bg-emerald-400 absolute bottom-0 right-0 border border-white"></span>
            </div>
          </div>
          <div class="flex-1 min-w-0 bg-slate-50/80 rounded-xl p-2.5 border border-slate-100/90 shadow-2xs">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-1.5 flex-wrap">
                <span class="text-xs font-bold text-slate-900">AlexAI</span>
                <span class="text-blue-500 text-[10px]">✓</span>
                <span class="text-[10px] text-slate-400 font-mono">10:42 AM •</span>
                <span class="text-[10px] text-slate-600 font-bold">Author</span>
                <span class="text-[9px] bg-[#008069] text-white font-mono px-1.5 py-0.2 rounded font-bold">@all</span>
              </div>
              <div class="flex items-center gap-1 text-xs"><span class="text-emerald-600 font-bold">✓✓</span></div>
            </div>
            <p class="text-sm text-slate-800 leading-relaxed font-sans mt-1">
              Yes! DM my inbox with conference dates, schedule, and topic details 📅
            </p>
            <div class="flex items-center justify-between text-xs text-slate-500 pt-1.5 select-none border-t border-slate-100 mt-2">
              <div class="flex items-center gap-3 text-[11px]">
                <button type="button" onclick="likeAnswerInWhatsApp('qa-preload-7', 'all', this)" class="like-answer-btn flex items-center gap-1 text-slate-600 hover:text-rose-600 transition-colors" title="Like Alex's answer (awards 1 credit to Alex)">
                  <span>♡</span> <span class="font-bold text-rose-600">Like (+1 cr)</span>
                </button>
                <button type="button" class="flex items-center gap-1 hover:text-[#008069] transition-colors"><span>💬</span> <span>1</span></button>
              </div>
              <button type="button" onclick="replyToQuestionInWhatsApp('qa-preload-7', '@organizer', 'Are you free for speaking at the systems conference? 🎤')" class="text-xs font-semibold text-[#008069] hover:underline">
                Reply in thread
              </button>
            </div>
          </div>
        </div>
      </article>

    </div>

    <!-- Active Quoted Reply Bar -->
    <div id="whatsapp-reply-bar" class="hidden px-3 py-2 bg-emerald-50 border-t border-emerald-200 flex items-center justify-between animate-slide-down select-none">
      <div class="flex items-center gap-2 min-w-0 flex-1">
        <div class="w-1 h-7 rounded-full bg-[#008069] shrink-0"></div>
        <div class="min-w-0 flex-1">
          <div class="text-xs font-bold text-[#008069] flex items-center gap-1">
            <span>Replying to</span>
            <span id="reply-quote-author" class="font-mono text-slate-700">@sarah</span>
          </div>
          <p id="reply-quote-text" class="text-[11px] text-slate-600 truncate font-mono">"Where are you now? 📍"</p>
        </div>
      </div>
      <button type="button" onclick="cancelWhatsAppReply()" class="text-slate-400 hover:text-slate-600 font-bold px-1.5 text-base">✕</button>
    </div>

    <!-- Micro Prompt Quick Chips (2x Readable font) + Anonymous Mode Toggle -->
    <div class="px-2.5 py-1.5 bg-white/95 border-t border-slate-200/80 flex items-center justify-between gap-1.5 shrink-0 select-none">
      <div class="flex items-center gap-1.5 overflow-x-auto shrink-0 scrollbar-none py-0.5">
        %s
      </div>

      <!-- Anonymous Toggle -->
      <label class="flex items-center gap-1 cursor-pointer select-none shrink-0 text-slate-600 hover:text-slate-900 px-2 py-0.5 rounded-lg hover:bg-slate-100 transition-colors" title="Ask without identity">
        <input type="checkbox" id="anon-toggle" onchange="toggleAnonMode(this.checked)" class="rounded border-slate-300 text-emerald-600 focus:ring-emerald-500 h-3.5 w-3.5 cursor-pointer" />
        <span class="text-xs font-semibold flex items-center gap-1">
          <span>🎭</span> <span>Ask Anon</span>
        </span>
      </label>
    </div>

    <!-- Real-Time Rate Limit Banner (1 question per minute) -->
    <div id="rate-limit-banner" class="hidden px-3.5 py-1.5 bg-amber-100/90 border-t border-amber-200 text-xs text-amber-900 font-medium flex items-center justify-between animate-slide-down">
      <div class="flex items-center gap-1.5">
        <span class="animate-pulse">⏳</span>
        <span id="rate-limit-text">1 question/min limit: please wait before asking again</span>
      </div>
      <span id="rate-limit-countdown" class="font-mono font-bold text-amber-950 bg-amber-200/80 px-2 py-0.5 rounded text-xs">60s</span>
    </div>

    <!-- No Credits Error Banner (Visible when user has 0 credits) -->
    <div id="no-credits-error-banner" class="hidden px-3.5 py-2 bg-rose-50 border-t border-rose-200 text-rose-900 text-xs sm:text-sm flex items-center justify-between animate-slide-down">
      <div class="flex items-center gap-1.5 font-medium">
        <span>⚠️</span>
        <span id="no-credits-error-text"><strong>Unable to send message:</strong> No credits available! You need at least 1 credit to ask questions.</span>
      </div>
      <button type="button" onclick="claimStarterCredits()" 
              class="px-2.5 py-1 rounded-lg bg-rose-600 hover:bg-rose-700 text-white font-bold text-xs shrink-0 active:scale-95 transition-all shadow-2xs">
        +3 Credits (Top up)
      </button>
    </div>

    <!-- Bottom Input Bar (2x Readable Text) -->
    <form id="talk-form" onsubmit="handleCloneSubmit(event)" 
          class="p-2.5 bg-slate-50 border-t border-slate-200 flex items-center gap-2 shrink-0">
      <span class="text-xl select-none">😊</span>
      <input type="text" id="clone-input" placeholder="Type a message or question..." required autocomplete="off"
             class="flex-1 bg-white border border-slate-200 rounded-2xl px-3.5 py-2 text-sm text-slate-900 placeholder:text-slate-400 focus:outline-none focus:border-[#008069] shadow-2xs" />
      <button type="submit" 
              class="w-9 h-9 rounded-full bg-[#008069] hover:bg-[#075e54] text-white flex items-center justify-center transition-all active:scale-95 shadow-xs shrink-0" title="Send Question">
        <span class="text-sm">➤</span>
      </button>
    </form>
  </div>

  <!-- ===================================================================== -->
  <!-- SUB-PANE 2: TALK BACK (Threads Reply Comment Tree Station) -->
  <!-- ===================================================================== -->
  <div id="subpane-talkback" class="hidden flex-1 flex flex-col min-h-0 bg-white overflow-hidden">
    
    <!-- Talk Back Header -->
    <div class="p-3 bg-white border-b border-slate-200 shadow-2xs space-y-2 shrink-0 select-none">
      <div class="flex items-center justify-between text-sm">
        <span class="font-bold text-slate-900 flex items-center gap-1.5">
          <span class="text-[#008069]">↩️</span> <span>Threads Reply Station</span>
        </span>
        <span class="text-xs px-2 py-0.5 rounded bg-emerald-100 text-emerald-900 font-mono font-bold">1 @tag per reply</span>
      </div>

      <!-- Question Selector Dropdown -->
      <div class="flex items-center gap-2">
        <select id="talkback-question-select" onchange="onTalkBackQuestionSelected(this.value)"
                class="flex-1 bg-slate-50 border border-slate-200 rounded-lg px-2.5 py-1.5 text-xs font-bold text-slate-900 focus:outline-none focus:border-[#008069] shadow-2xs"></select>
        <button type="button" onclick="promptCustomQuestionForTalkBack()" 
                class="px-2.5 py-1.5 rounded-lg bg-slate-100 hover:bg-slate-200 text-slate-700 font-bold text-xs shrink-0" title="Add new question">
          + New Q
        </button>
      </div>

      <!-- Question Meta -->
      <div id="talkback-q-meta" class="p-2 rounded-lg bg-slate-50 border border-slate-200 text-xs flex items-center justify-between">
        <span class="text-slate-600" id="talkback-target-asker">Thread Author: <strong>Sarah (@sarah)</strong></span>
        <span class="text-xs font-mono text-emerald-700 font-bold" id="talkback-target-status">⏳ Unanswered</span>
      </div>
    </div>

    <!-- Active Question's Reply Comment Threads Stream (Tree Connector Lines) -->
    <div id="talkback-thread-stream" class="flex-1 min-h-0 overflow-y-auto p-3 space-y-2.5 custom-scroll scrollbar-none bg-[#f7f9fa]"></div>

    <!-- STRICT SINGLE REPLY COMPOSER (Strictly 1 answer, 1 @list tag) -->
    <div class="p-3 bg-white border-t border-slate-200 shadow-sm space-y-2 shrink-0 select-none">
      <div class="flex items-center justify-between text-xs">
        <label class="font-bold text-slate-900 flex items-center gap-1">
          <span>✏️</span> <span>Post Single Reply</span>
        </label>
        <span class="font-mono text-emerald-700 font-bold text-[11px]">1 @list tag only • 1 answer</span>
      </div>

      <!-- Single Tag Picker -->
      <div class="space-y-1">
        <span class="text-[10px] font-bold text-slate-400 block uppercase tracking-wider">Select exactly one @list tag:</span>
        <div id="talkback-single-tags-container" class="flex items-center gap-1.5 overflow-x-auto pb-1 scrollbar-none">
          <button type="button" onclick="selectSingleReplyTag('all')" data-tag="all" class="talkback-tag-chip px-2.5 py-1 rounded-lg text-xs font-bold transition-all bg-emerald-600 text-white shadow-2xs">@all</button>
          <button type="button" onclick="selectSingleReplyTag('family')" data-tag="family" class="talkback-tag-chip px-2.5 py-1 rounded-lg text-xs font-bold transition-all bg-slate-100 text-slate-700 hover:bg-slate-200">@family</button>
          <button type="button" onclick="selectSingleReplyTag('friends')" data-tag="friends" class="talkback-tag-chip px-2.5 py-1 rounded-lg text-xs font-bold transition-all bg-slate-100 text-slate-700 hover:bg-slate-200">@friends</button>
          <button type="button" onclick="selectSingleReplyTag('bffs')" data-tag="bffs" class="talkback-tag-chip px-2.5 py-1 rounded-lg text-xs font-bold transition-all bg-slate-100 text-slate-700 hover:bg-slate-200">@bffs</button>
          <button type="button" onclick="selectSingleReplyTag('work')" data-tag="work" class="talkback-tag-chip px-2.5 py-1 rounded-lg text-xs font-bold transition-all bg-slate-100 text-slate-700 hover:bg-slate-200">@work</button>
          <button type="button" onclick="selectSingleReplyTag('marcus')" data-tag="marcus" class="talkback-tag-chip px-2.5 py-1 rounded-lg text-xs font-mono font-bold transition-all bg-slate-100 text-slate-700 hover:bg-slate-200">@marcus</button>
          <button type="button" onclick="selectSingleReplyTag('sarah')" data-tag="sarah" class="talkback-tag-chip px-2.5 py-1 rounded-lg text-xs font-mono font-bold transition-all bg-slate-100 text-slate-700 hover:bg-slate-200">@sarah</button>
          <button type="button" onclick="selectSingleReplyTag('priya')" data-tag="priya" class="talkback-tag-chip px-2.5 py-1 rounded-lg text-xs font-mono font-bold transition-all bg-slate-100 text-slate-700 hover:bg-slate-200">@priya</button>
          <button type="button" onclick="selectSingleReplyTag('mom')" data-tag="mom" class="talkback-tag-chip px-2.5 py-1 rounded-lg text-xs font-mono font-bold transition-all bg-slate-100 text-slate-700 hover:bg-slate-200">@mom</button>
        </div>
        <input type="hidden" id="single-reply-active-tag" value="all" />
      </div>

      <!-- Single Reply Error Banner -->
      <div id="single-reply-error-banner" class="hidden p-2 rounded-lg bg-rose-50 border border-rose-200 text-xs text-rose-800 font-medium flex items-center gap-1.5">
        <span>⚠️</span> <span id="single-reply-error-text">Only one @list tagging allowed per reply.</span>
      </div>

      <!-- Textarea & Send Button -->
      <div class="flex items-center gap-2">
        <textarea id="talkback-answers-editor" rows="2" oninput="onSingleReplyTextInput(this.value)"
                  placeholder="Type one answer for the selected @tag..."
                  class="flex-1 bg-slate-50 border border-slate-200 rounded-xl p-2.5 text-sm text-slate-900 placeholder:text-slate-400 focus:outline-none focus:border-[#008069] font-sans leading-snug shadow-2xs resize-none"></textarea>
        <button type="button" id="submit-single-reply-btn" onclick="submitTalkBackAnswer()" 
                class="px-3.5 py-2.5 rounded-xl bg-[#008069] hover:bg-[#075e54] text-white font-bold text-sm shadow-xs transition-all active:scale-95 flex items-center gap-1 shrink-0">
          <span>💬</span> <span>Send</span>
        </button>
      </div>
    </div>

    <!-- Hidden answered directory container for test parity -->
    <div id="talkback-answered-directory" class="hidden"></div>
  </div>

</div>
`, p.CloneName, p.Subtitle, p.Greeting, chipsSb.String())

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
    /* =====================================================================
       Flutter-Native Touch Physics & Zero Scrollbars Everywhere
       ===================================================================== */
    *, *::before, *::after {
      scrollbar-width: none !important;
      -ms-overflow-style: none !important;
      -webkit-tap-highlight-color: transparent;
      box-sizing: border-box;
    }
    *::-webkit-scrollbar {
      display: none !important;
      width: 0 !important;
      height: 0 !important;
    }
    html, body {
      overflow-x: hidden;
      overscroll-behavior-y: none;
      -webkit-font-smoothing: antialiased;
      -moz-osx-font-smoothing: grayscale;
      touch-action: pan-y;
      padding-top: env(safe-area-inset-top);
      padding-bottom: env(safe-area-inset-bottom);
      padding-left: env(safe-area-inset-left);
      padding-right: env(safe-area-inset-right);
    }
    .custom-scroll, [class*="overflow-"] {
      scrollbar-width: none !important;
      -ms-overflow-style: none !important;
      -webkit-overflow-scrolling: touch;
      overscroll-behavior: contain;
    }
    .custom-scroll::-webkit-scrollbar, [class*="overflow-"]::-webkit-scrollbar {
      display: none !important;
      width: 0 !important;
      height: 0 !important;
    }
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
