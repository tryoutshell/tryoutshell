Build a modern landing + labs page inspired by rootly.com/ai-labs,but customized for TryOutShell, a CLI-native learning platform.

## Design overview:
* The homepage (tryoutshell.lol) should look like git-fork.com(https://git-fork.com/)
 — with a minimal hero section showing a terminal demo video or live terminal animation that demonstrates TryOutShell in action.
 * Beneath that, have a short 2–3 line mission statement and a “Get Started” button leading to /labs.

## Labs page (/labs) should look and feel like Rootly AI Labs(https://rootly.com/ai-labs):
* Replace their "AI Labs" heading with "TryOutShell Orgs".
* Replace the Rootly background with a calm CLI-inspired gradient or dark grid terminal aesthetic.
* Remove the top navbar items (“Product”, “Resources”, “Book a Demo”) — keep only a minimal sticky header with logo + “Docs / GitHub / Discord” links.
* Below the hero section, have a grid of cards representing organizations or categories of lessons.

## Grid layout:
* Each card = one org (like sigstore, chainguard, cosign, witness, etc.)
* On hover → show short lesson description, e.g.
“Learn how to verify container images with Sigstore using Cosign CLI.”
* Each card includes:
  - Status badge (e.g. “Available”, “Coming soon”)
  - Maintainer avatar
  - Stack tags (e.g. Go, Kubernetes, Cosign, Reproducible Builds)

## Card click behavior:

* Clicking opens /labs/:org, showing:
  * Overview of what lessons exist in that org
  * A small “Try a Demo” lesson (an embedded terminal video or animated code block)
  * Links to full lessons (tryoutshell start <lesson-name>)

## Style:
* Typography similar to Rootly (serif headers, clean sans-serif body)
* Dark mode default
* Smooth hover and entry animations for cards
* Rounded corners and slight neon-glow border for CLI vibe

Tech stack suggestion:
* Next.js + Tailwind + Framer Motion
* Recharts / ShadCN UI for consistency

Sections summary:
1. / → Hero terminal + short intro + Get Started
2. /labs → Orgs grid (like Rootly AI Labs)
3. /labs/:org → Details + demo lesson

```
/
├── Hero (Terminal demo + CTA)
└── Highlights (Open source, CLI-native, Learn by doing)

/labs
├── Hero (Title + Subtitle)
└── Grid of org cards
    ├── sigstore
    ├── witness
    ├── chainguard
    ├── cosign
    └── kubernetes

/labs/:org
├── Header (Org name + status)
├── Description
├── Demo terminal/video section
└── Lesson commands + links

```
Copy examples:

Hero text:
“Building the future of interactive CLI learning.”
“Learn DevSecOps, Kubernetes, and Supply Chain tools — right in your terminal.”

Subtitle:
“Hands-on, open-source, and built for engineers who love the command line.”
