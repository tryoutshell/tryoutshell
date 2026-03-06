---
sidebar_position: 4
---

# Hints and Callouts

Hints help users when they're stuck. Callouts add contextual notes to info steps.

## Hints (Command Steps)

Progressive hints are revealed one at a time when the user presses `?`.

```yaml
hints:
  - level: 1
    text: "Use the docker command to check the version"
  - level: 2
    text: "Try: docker --version"
  - level: 3
    text: "Full command: docker --version"
```

### Design Pattern

| Level | What to Include |
|-------|----------------|
| 1 | Vague direction — which tool or concept |
| 2 | More specific — the rough command shape |
| 3 | Complete answer — exact command to type |

### Tips

- **Always include 2-3 hints** per command step
- **Don't repeat the example** — that's already visible
- **Level 3 should give the full answer** — last resort for stuck users
- **Add context, not just commands** — "Install with brew if you're on macOS"

## Callouts (Info Steps)

Callouts add visual emphasis to tips, warnings, and other notes.

```yaml
callouts:
  - type: "tip"
    text: "Use keyless signing in production"
  - type: "warning"
    text: "Never commit private keys to Git!"
  - type: "danger"
    text: "This will delete all containers"
  - type: "info"
    text: "Signatures are stored as OCI artifacts"
```

### Callout Types

| Type | Use For | Emoji |
|------|---------|-------|
| `tip` | Helpful advice, shortcuts | 💡 |
| `info` | Additional context, facts | ℹ️ |
| `warning` | Potential issues, caution | ⚠️ |
| `danger` | Serious risks, destructive actions | 🚨 |

### Tips

- **1-2 callouts per step** — more than that is noise
- **Keep text concise** — one sentence per callout
- **Use `warning` for security advice** — "Never share private keys"
- **Use `tip` for best practices** — "Use multi-stage builds"
- **Use `info` for fun facts** — "Cosign is part of Sigstore"
