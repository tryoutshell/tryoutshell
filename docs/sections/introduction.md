---
sidebar_position: 1
---

# Introduction Section

The introduction section is shown once when an interactive lesson starts. It sets context and expectations.

## Structure

```yaml
introduction:
  title: "What You'll Learn"
  content: |
    Markdown content describing the lesson.
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `title` | string | Yes | Introduction header |
| `content` | string | Yes | Overview text (Markdown supported) |

## Example

```yaml
introduction:
  title: "What You'll Learn"
  content: |
    In this lesson, you will:
    - Understand what Cosign is and why image signing matters
    - Install and verify Cosign
    - Sign a container image with ephemeral keys
    - Verify signatures to ensure image integrity

    **Time:** ~20 minutes
    **Tools:** Cosign, Docker

    > 💡 **Tip:** Have Docker running before you start!
```

## Markdown Support

The `content` field supports full Markdown:

- `**bold**` and `*italic*`
- `code` inline code
- `### Heading` — subheadings
- `> Quote` — blockquotes
- `- List item` — bullet lists
- `1. Numbered` — numbered lists
- Code blocks with syntax highlighting

## Best Practices

- **Set clear expectations** — list 3-5 things users will learn
- **Mention prerequisites** — tools, knowledge, time required
- **Keep it concise** — save details for the steps
- **Use bullet lists** — easier to scan than paragraphs
