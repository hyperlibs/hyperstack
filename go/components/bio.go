package components

import (
	"fmt"
	"strings"
)

// =============================================================================
// ABC — A BIO CHAT: GO COMPONENT DSL
// Mobile-First Link-in-Bio & Real-Time Bio Chat Application
// Conforms to Hyperstack Component Standards & 1:1 Parity Rule
// =============================================================================

// --- 1. Bio Profile Header ---

type SocialLink struct {
	Platform string
	URL      string
	Icon     string
	Label    string
}

type BioProfileProps struct {
	AvatarURL   string
	Name        string
	Handle      string
	Verified    bool
	Tagline     string
	Bio         string
	Location    string
	StatusText  string
	Socials     []SocialLink
	Followers   string
	Views       string
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
	if p.Tagline == "" {
		p.Tagline = "Systems Architect & Hypermedia Alchemist"
	}
	if p.Location == "" {
		p.Location = "San Francisco, CA"
	}
	if p.StatusText == "" {
		p.StatusText = "Open for Advisory & Projects"
	}

	verifiedBadge := ""
	if p.Verified {
		verifiedBadge = `<span class="inline-flex items-center justify-center w-5 h-5 rounded-full bg-blue-500 text-white text-[11px] font-bold shadow-sm" title="Verified Creator">✓</span>`
	}

	var socialsSb strings.Builder
	for _, s := range p.Socials {
		socialsSb.WriteString(fmt.Sprintf(`
      <a href="%s" target="_blank" rel="noopener noreferrer" 
         class="w-9 h-9 rounded-full bg-slate-800/80 hover:bg-indigo-600/80 border border-slate-700/60 flex items-center justify-center text-sm text-slate-300 hover:text-white transition-all active:scale-90"
         title="%s" hx-action="$appz.haptic('selection')">
        %s
      </a>`, s.URL, s.Label, s.Icon))
	}

	html := fmt.Sprintf(`
<div class="bio-profile-card flex flex-col items-center text-center p-6 relative">
  <!-- Glow Halo -->
  <div class="absolute -top-10 inset-x-10 h-28 bg-gradient-to-r from-indigo-500/20 via-purple-500/20 to-pink-500/20 blur-2xl pointer-events-none rounded-full"></div>

  <!-- Avatar with Pulsing Status Ring -->
  <div class="relative mb-4 group cursor-pointer" hx-action="$appz.haptic('selection'); $toast('Hey! Thanks for visiting my Bio Chat.', 'info')">
    <div class="w-24 h-24 rounded-full p-1 bg-gradient-to-tr from-indigo-500 via-purple-500 to-pink-500 shadow-xl shadow-indigo-500/25">
      <img src="%s" alt="%s" class="w-full h-full rounded-full object-cover border-2 border-slate-950" />
    </div>
    <!-- Online status badge -->
    <div class="absolute bottom-1 right-1 w-5 h-5 rounded-full bg-slate-950 border-2 border-slate-950 flex items-center justify-center" title="Online now">
      <span class="w-3.5 h-3.5 rounded-full bg-emerald-500 animate-pulse"></span>
    </div>
  </div>

  <!-- Name & Handle -->
  <div class="flex items-center gap-1.5 justify-center mb-1">
    <h2 class="text-xl font-black text-white tracking-tight">%s</h2>
    %s
  </div>
  <p class="text-xs font-mono font-medium text-indigo-400 mb-2">%s</p>

  <!-- Tagline & Bio -->
  <p class="text-sm font-semibold text-slate-200 max-w-xs mb-2 leading-snug">%s</p>
  <p class="text-xs text-slate-400 max-w-xs mb-3 leading-relaxed">%s</p>

  <!-- Location & Status Pill -->
  <div class="flex items-center gap-2 mb-4">
    <span class="inline-flex items-center gap-1 text-[11px] font-medium text-slate-400 bg-slate-900/80 px-2.5 py-1 rounded-full border border-slate-800">
      <span>📍</span> %s
    </span>
    <span class="inline-flex items-center gap-1.5 text-[11px] font-semibold text-emerald-400 bg-emerald-500/10 px-2.5 py-1 rounded-full border border-emerald-500/20">
      <span class="w-1.5 h-1.5 rounded-full bg-emerald-400 animate-ping"></span>
      %s
    </span>
  </div>

  <!-- Social Links Strip -->
  <div class="flex items-center justify-center gap-2 mb-5">
    %s
  </div>

  <!-- Quick Action Pill Bar -->
  <div class="grid grid-cols-4 gap-2 w-full max-w-xs">
    <button hx-action="activeTab = 'chat'; $appz.haptic('selection'); $sound('click')" 
            class="py-2 px-1 rounded-xl bg-indigo-600/90 hover:bg-indigo-600 text-white font-bold text-xs flex flex-col items-center gap-1 shadow-lg shadow-indigo-600/30 active:scale-95 transition-all">
      <span class="text-sm">💬</span>
      <span class="text-[10px]">Bio Chat</span>
    </button>
    <button hx-action="$appz.sheet('#sheet-tip'); $appz.haptic('selection'); $sound('ping')" 
            class="py-2 px-1 rounded-xl bg-slate-800/90 hover:bg-slate-700 text-amber-400 border border-amber-500/20 font-bold text-xs flex flex-col items-center gap-1 active:scale-95 transition-all">
      <span class="text-sm">☕</span>
      <span class="text-[10px]">Tip Jar</span>
    </button>
    <button hx-action="$appz.sheet('#sheet-booking'); $appz.haptic('selection'); $sound('click')" 
            class="py-2 px-1 rounded-xl bg-slate-800/90 hover:bg-slate-700 text-slate-300 font-bold text-xs flex flex-col items-center gap-1 active:scale-95 transition-all">
      <span class="text-sm">📅</span>
      <span class="text-[10px]">Book 1:1</span>
    </button>
    <button hx-action="$copy(window.location.href); $toast('Bio link copied to clipboard!', 'success'); $sound('click')" 
            class="py-2 px-1 rounded-xl bg-slate-800/90 hover:bg-slate-700 text-slate-300 font-bold text-xs flex flex-col items-center gap-1 active:scale-95 transition-all">
      <span class="text-sm">🔗</span>
      <span class="text-[10px]">Share</span>
    </button>
  </div>
</div>`, p.AvatarURL, p.Name, p.Name, verifiedBadge, p.Handle, p.Tagline, p.Bio, p.Location, p.StatusText, socialsSb.String())

	return Component{html: html}
}

// --- 2. Bio Link Card ---

type BioLinkProps struct {
	ID          string
	Title       string
	Description string
	URL         string
	Icon        string
	Category    string // "projects" | "articles" | "media" | "store"
	Badge       string
	Highlighted bool
	Metrics     string
}

func BioLinkCard(p BioLinkProps) Component {
	highlightClass := "bg-slate-900/90 border-slate-800 hover:border-slate-700"
	if p.Highlighted {
		highlightClass = "bg-gradient-to-r from-indigo-950/60 via-slate-900 to-indigo-950/60 border-indigo-500/50 shadow-lg shadow-indigo-500/10 hover:border-indigo-400"
	}

	badgeHtml := ""
	if p.Badge != "" {
		badgeHtml = fmt.Sprintf(`<span class="text-[10px] font-bold px-2 py-0.5 rounded-full bg-indigo-500/20 text-indigo-400 border border-indigo-500/30">%s</span>`, p.Badge)
	}

	metricsHtml := ""
	if p.Metrics != "" {
		metricsHtml = fmt.Sprintf(`<span class="text-[10px] font-mono text-slate-400 flex items-center gap-1"><span class="w-1.5 h-1.5 rounded-full bg-indigo-400"></span>%s</span>`, p.Metrics)
	}

	html := fmt.Sprintf(`
<div class="bio-link-card group relative rounded-2xl border p-4 transition-all duration-200 active:scale-[0.98] %s"
     data-category="%s"
     hx-vibe-card="8">
  <a href="%s" target="_blank" rel="noopener noreferrer" 
     class="flex items-center gap-3.5"
     hx-action="stats.clicks++; $appz.haptic('selection')">
    <div class="w-11 h-11 rounded-xl bg-slate-800 flex items-center justify-center text-xl shrink-0 border border-slate-700/60 shadow-sm group-hover:scale-110 transition-transform">
      %s
    </div>
    <div class="flex-1 min-w-0 text-left">
      <div class="flex items-center gap-2 mb-0.5">
        <h3 class="font-bold text-sm text-white group-hover:text-indigo-400 transition-colors truncate">%s</h3>
        %s
      </div>
      <p class="text-xs text-slate-400 line-clamp-1 leading-relaxed">%s</p>
      %s
    </div>
    <div class="w-8 h-8 rounded-full bg-slate-800/60 flex items-center justify-center text-slate-400 group-hover:text-white group-hover:bg-indigo-600 transition-colors shrink-0">
      ↗
    </div>
  </a>
</div>`, highlightClass, p.Category, p.URL, p.Icon, p.Title, badgeHtml, p.Description, metricsHtml)

	return Component{html: html}
}

// --- 3. Bio Chat (Messenger & AMA) ---

type ChatMessage struct {
	ID        int
	Sender    string // "alex" | "visitor"
	Text      string
	Timestamp string
}

type BioChatProps struct {
	InitialMessages []ChatMessage
}

func BioChat(p BioChatProps) Component {
	var messagesSb strings.Builder
	for _, m := range p.InitialMessages {
		if m.Sender == "alex" {
			messagesSb.WriteString(fmt.Sprintf(`
        <div class="flex items-start gap-2.5 max-w-[85%%]">
          <div class="w-7 h-7 rounded-full bg-indigo-600 flex items-center justify-center text-white text-xs font-bold shrink-0">A</div>
          <div class="p-3 rounded-2xl rounded-tl-sm bg-slate-800 border border-slate-700/70 text-xs text-slate-200 leading-relaxed shadow-sm">
            %s
            <div class="text-[9px] text-slate-400 text-right mt-1 font-mono">%s</div>
          </div>
        </div>`, m.Text, m.Timestamp))
		} else {
			messagesSb.WriteString(fmt.Sprintf(`
        <div class="flex items-start gap-2.5 max-w-[85%%] ml-auto flex-row-reverse">
          <div class="w-7 h-7 rounded-full bg-emerald-600 flex items-center justify-center text-white text-xs font-bold shrink-0">You</div>
          <div class="p-3 rounded-2xl rounded-tr-sm bg-indigo-600 text-white text-xs leading-relaxed shadow-sm">
            %s
            <div class="text-[9px] text-indigo-200 text-right mt-1 font-mono">%s</div>
          </div>
        </div>`, m.Text, m.Timestamp))
		}
	}

	html := fmt.Sprintf(`
<div class="bio-chat-container flex flex-col h-[520px] rounded-2xl bg-slate-900/90 border border-slate-800 overflow-hidden shadow-2xl">
  <!-- Chat Header -->
  <div class="h-12 px-4 bg-slate-950/80 border-b border-slate-800/80 flex items-center justify-between shrink-0">
    <div class="flex items-center gap-2">
      <div class="relative">
        <span class="w-3 h-3 rounded-full bg-emerald-500 block"></span>
        <span class="w-3 h-3 rounded-full bg-emerald-500 block absolute inset-0 animate-ping opacity-75"></span>
      </div>
      <div>
        <h3 class="text-xs font-bold text-white leading-none">Alex Rivera (AMA)</h3>
        <p class="text-[9px] text-slate-400 font-medium">Instant Bio Assistant • Replies in ~2m</p>
      </div>
    </div>
    <span class="text-[10px] font-mono px-2 py-0.5 rounded-full bg-indigo-500/10 text-indigo-400 border border-indigo-500/20 font-semibold">Live Socket</span>
  </div>

  <!-- Messages Stream -->
  <div id="bio-chat-stream" class="flex-1 overflow-y-auto p-4 space-y-3.5">
    %s
  </div>

  <!-- Quick Suggestion Starter Chips -->
  <div class="px-3 py-2 bg-slate-950/50 border-t border-slate-800/60 overflow-x-auto flex items-center gap-1.5 scrollbar-none">
    <button class="px-2.5 py-1 rounded-full bg-slate-800/80 hover:bg-indigo-600 border border-slate-700/60 text-[10px] font-semibold text-slate-300 hover:text-white shrink-0 active:scale-95 transition-all"
            hx-action="askPrompt('What are your top projects right now?')">
      🚀 Top Projects?
    </button>
    <button class="px-2.5 py-1 rounded-full bg-slate-800/80 hover:bg-indigo-600 border border-slate-700/60 text-[10px] font-semibold text-slate-300 hover:text-white shrink-0 active:scale-95 transition-all"
            hx-action="askPrompt('Can I book you for architectural advisory?')">
      🤝 Hire for Advisory?
    </button>
    <button class="px-2.5 py-1 rounded-full bg-slate-800/80 hover:bg-indigo-600 border border-slate-700/60 text-[10px] font-semibold text-slate-300 hover:text-white shrink-0 active:scale-95 transition-all"
            hx-action="askPrompt('Why did you choose Go & Hyperstack?')">
      ⚡ Why Hyperstack?
    </button>
  </div>

  <!-- Input Bar -->
  <form id="bio-chat-form" class="p-3 bg-slate-950/90 border-t border-slate-800 flex items-center gap-2"
        hx-post="/api/bio/chat"
        hx-target="#bio-chat-stream"
        hx-swap="beforeend"
        hx-on::after-request="this.reset(); scrollChatToEnd(); $sound('ping')">
    <input type="text" name="message" id="bio-chat-input" placeholder="Type a message or question..." required
           class="flex-1 bg-slate-900 border border-slate-800 rounded-xl px-3 py-2 text-xs text-white placeholder:text-slate-500 focus:outline-none focus:border-indigo-500 shadow-inner" />
    <button type="submit" 
            class="px-4 py-2 rounded-xl bg-indigo-600 hover:bg-indigo-500 text-white font-bold text-xs shadow-md shadow-indigo-600/30 active:scale-95 transition-all flex items-center gap-1">
      <span>Send</span> <span>➤</span>
    </button>
  </form>
</div>`, messagesSb.String())

	return Component{html: html}
}

// --- 4. Bio Tip Jar ---

type BioTipJarProps struct {
	TotalTips int
	Amounts   []int
}

func BioTipJar(p BioTipJarProps) Component {
	if len(p.Amounts) == 0 {
		p.Amounts = []int{3, 5, 15, 50}
	}
	if p.TotalTips == 0 {
		p.TotalTips = 240
	}

	var chipsSb strings.Builder
	for _, a := range p.Amounts {
		label := fmt.Sprintf("$%d", a)
		subtitle := "Coffee"
		if a == 5 {
			subtitle = "Boba"
		} else if a == 15 {
			subtitle = "Pizza"
		} else if a == 50 {
			subtitle = "Sponsor"
		}

		chipsSb.WriteString(fmt.Sprintf(`
      <button hx-action="tipAmount = %d; $appz.haptic('selection'); $sound('click')"
              hx-class="{ 'border-amber-400 bg-amber-500/20 text-white': tipAmount === %d, 'border-slate-800 bg-slate-900 text-slate-300': tipAmount !== %d }"
              class="p-3 rounded-2xl border flex flex-col items-center justify-center transition-all active:scale-95">
        <span class="text-base font-black">%s</span>
        <span class="text-[10px] text-slate-400">%s</span>
      </button>`, a, a, a, label, subtitle))
	}

	html := fmt.Sprintf(`
<div class="bio-tip-jar p-6 rounded-3xl bg-slate-900/90 border border-slate-800 flex flex-col items-center text-center shadow-xl relative overflow-hidden">
  <div class="w-14 h-14 rounded-2xl bg-amber-500/10 border border-amber-500/30 flex items-center justify-center text-2xl mb-3 text-amber-400 shadow-inner">
    ☕
  </div>
  <h3 class="font-extrabold text-base text-white mb-1">Fuel the Open Source Journey</h3>
  <p class="text-xs text-slate-400 max-w-xs mb-4">Support development of free, hyperreactive tools like Hyperstack, Hypersheet, and AppzOS.</p>

  <!-- Tip preset amounts -->
  <div class="grid grid-cols-4 gap-2.5 w-full mb-4">
    %s
  </div>

  <!-- Custom note -->
  <input type="text" id="tip-note" placeholder="Add an optional thank-you note..." 
         class="w-full bg-slate-950 border border-slate-800 rounded-xl px-3 py-2 text-xs text-white placeholder:text-slate-500 mb-4 focus:outline-none focus:border-amber-500" />

  <!-- Send Tip Button with sound & shockwave blast -->
  <button hx-action="sendTip(); $sound('success'); $blast({ radius: 300, force: 40 }); $toast('Thank you so much for your support! ❤️', 'success')"
          class="w-full py-3 rounded-2xl bg-gradient-to-r from-amber-500 to-amber-600 hover:from-amber-400 hover:to-amber-500 font-extrabold text-xs text-slate-950 tracking-wide shadow-lg shadow-amber-500/25 active:scale-95 transition-all flex items-center justify-center gap-2">
    <span>❤️ Send Support (<span hx-text="'$' + tipAmount">$15</span>)</span>
  </button>
</div>`, chipsSb.String())

	return Component{html: html}
}

// --- 5. Full Bio Page Component ---

type BioPageProps struct {
	Title   string
	Profile BioProfileProps
	Links   []BioLinkProps
	Chat    BioChatProps
	TipJar  BioTipJarProps
}

func BioPage(p BioPageProps) Component {
	if p.Title == "" {
		p.Title = "Alex Rivera — A Bio Chat"
	}

	profileComp := BioProfile(p.Profile)
	var linksSb strings.Builder
	for _, l := range p.Links {
		linksSb.WriteString(BioLinkCard(l).Render())
	}
	chatComp := BioChat(p.Chat)
	tipComp := BioTipJar(p.TipJar)

	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en" class="h-full bg-slate-950 text-slate-100 antialiased select-none">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no, viewport-fit=cover">
  <title>%s</title>
  <script src="/vendor/htmx.min.js"></script>
  <script src="/htmx-bolt.js"></script>
  <script src="/htmx-flash.js"></script>
  <script src="/htmx-form.js"></script>
  <script src="/htmx-vibe.js"></script>
  <script src="/appz.js"></script>
  <script src="https://cdn.tailwindcss.com"></script>
  <link rel="stylesheet" href="/styles.css">
</head>
<body class="h-full flex flex-col justify-center items-center p-0 md:p-6 bg-slate-950">
  <div class="w-full max-w-md h-full md:h-[860px] md:rounded-[44px] md:border-[10px] md:border-slate-900 bg-slate-950 flex flex-col overflow-hidden shadow-2xl relative">
    %s
    <div class="flex-1 overflow-y-auto p-4 space-y-4">
      %s
      %s
      %s
    </div>
  </div>
</body>
</html>`, p.Title, profileComp.Render(), linksSb.String(), chatComp.Render(), tipComp.Render())

	return Component{html: html}
}
