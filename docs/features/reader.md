---
sidebar_position: 1
---

# AI Blog Reader

Read any blog post or article in a split-pane terminal UI with an AI assistant.

## Usage

```bash
tryoutshell read <url>
```

### Examples

```bash
# Read a Sigstore blog post
tryoutshell read https://blog.sigstore.dev/cosign-2-0

# Read and save for offline
tryoutshell read --save https://kubernetes.io/blog/2022/01/07/kubernetes-is-moving-on-from-dockershim/

# Open saved articles
tryoutshell saved
```

## Layout

```
┌──────────────────────────────────┬─────────────────────────┐
│  📖 ARTICLE TITLE                │  🤖 AI ASSISTANT         │
│                                  │                          │
│  [Scrollable markdown content]   │  > Ask anything about    │
│                                  │    this article...       │
│  • Section 1                     │                          │
│  • Section 2                     │  You: What is cosign?    │
│    ...                           │                          │
│                                  │  AI: Cosign is a tool    │
│                                  │  for signing containers. │
│                                  │                          │
│                                  │  [type your question]    │
├──────────────────────────────────┴─────────────────────────┤
│  tab: switch │ [/]: resize │ t: theme │ j/k: scroll │ q    │
└─────────────────────────────────────────────────────────────┘
```

## Keybindings

### Article Pane (left)

| Key | Action |
|-----|--------|
| `j` / `↓` | Scroll down |
| `k` / `↑` | Scroll up |
| `d` / `ctrl+d` | Half-page down |
| `u` / `ctrl+u` | Half-page up |
| `G` | Jump to bottom |
| `g` | Jump to top |
| `Tab` | Switch to chat pane |
| `[` / `]` | Resize split left/right |
| `t` | Cycle through themes |
| `q` | Quit |

### Chat Pane (right)

| Key | Action |
|-----|--------|
| `Enter` | Send message |
| `Tab` | Switch to article pane |
| `Esc` | Back to article pane |

## AI Chat Setup

The reader checks for API keys in this order:

1. `OPENAI_API_KEY` — Uses GPT-4o-mini
2. `ANTHROPIC_API_KEY` — Uses Claude Sonnet
3. `GEMINI_API_KEY` — Uses Gemini 2.0 Flash

Set one of these environment variables to enable AI chat:

```bash
export OPENAI_API_KEY="sk-..."
# or
export ANTHROPIC_API_KEY="sk-ant-..."
# or
export GEMINI_API_KEY="AI..."
```

If no API key is set, the reader still works — you just can't use the chat panel.

## Themes

Press `t` to cycle through 5 built-in themes:

| Theme | Description |
|-------|-------------|
| `default` | System colors |
| `dark` | Dark background, light text |
| `light` | Light background, dark text |
| `dracula` | Dracula color scheme |
| `nord` | Nord color scheme |

## Saving Articles

Save articles for offline reading:

```bash
tryoutshell read --save https://example.com/article
```

Articles are saved as Markdown to `~/.local/share/tryoutshell/saved/`.

Browse and open saved articles:

```bash
tryoutshell saved
```
