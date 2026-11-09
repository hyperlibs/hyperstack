package components

import (
	"fmt"
	"strings"
)

// =============================================================================
// LINK IN BIO: GO COMPONENT DSL
// Pure Mobile-First Link-in-Bio Application using Hyperstack (Light Theme)
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

// --- 3. Full Mobile-First Link-in-Bio Page Component ---

type BioPageProps struct {
	Title   string
	Profile BioProfileProps
	Links   []BioLinkProps
}

func BioPage(p BioPageProps) Component {
	if p.Title == "" {
		p.Title = "Alex Rivera • Link in Bio"
	}

	profileComp := BioProfile(p.Profile)
	var linksSb strings.Builder
	for _, l := range p.Links {
		linksSb.WriteString(BioLinkCard(l).Render())
	}

	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en" class="h-full bg-slate-50 text-slate-900 antialiased select-none">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no, viewport-fit=cover">
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
<body class="min-h-full flex flex-col items-center justify-start bg-slate-50 text-slate-900 py-6 px-4">
  <main class="w-full max-w-md flex flex-col space-y-4">
    %s
    <div class="space-y-3">
      %s
    </div>
    <footer class="pt-6 pb-4 text-center">
      <p class="text-xs text-slate-400 font-mono">Built with Hyperstack • Zero Virtual DOM</p>
    </footer>
  </main>
</body>
</html>`, p.Title, profileComp.Render(), linksSb.String())

	return Component{html: html}
}
