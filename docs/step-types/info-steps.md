---
sidebar_position: 2
---

# Info Steps

Info steps display educational content without requiring command execution. The user reads the content and presses Enter to continue.

## Purpose

Use info steps to:
- Explain concepts and theory
- Provide context before hands-on steps
- Show diagrams and examples
- Give tips and warnings

## Basic Structure

```yaml
- type: info
  title: "Step Title"
  content: |
    Your content here with **Markdown** support.
```

## Fields Reference

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | `"info"` | Yes | Identifies this as an info step |
| `title` | string | Yes | Step heading |
| `content` | string | Yes | Main text (Markdown supported) |
| `highlights` | list | No | Inline text highlights |
| `code_blocks` | list | No | Code examples with labels |
| `callouts` | list | No | Tips, warnings, info boxes |
| `diagram` | string | No | ASCII art diagram |
| `wait_for_continue` | boolean | No | Pause until Enter (default: true) |

## Simple Example

```yaml
- type: info
  title: "What is Docker?"
  content: |
    **Docker** is a platform for developing and running
    applications in lightweight, portable containers.

    Containers package an application with all its dependencies,
    ensuring it runs the same everywhere.
```

## Full Example with All Fields

```yaml
- type: info
  title: "How Signing Works"
  content: |
    Cosign uses **public key cryptography** to sign images:

    1. **Generate a key pair** (public + private)
    2. **Sign the image** with your private key
    3. **Verify** using the public key

    > 💡 Anyone can verify, but only you can sign.

  highlights:
    - text: "public key cryptography"
      style: "bold"
    - text: "cosign.key"
      style: "code"

  code_blocks:
    - label: "macOS (Homebrew)"
      code: "brew install cosign"
      language: "bash"
    - label: "Linux"
      code: |
        curl -LO https://github.com/sigstore/cosign/releases/latest/download/cosign-linux-amd64
        sudo mv cosign-linux-amd64 /usr/local/bin/cosign
      language: "bash"

  callouts:
    - type: "tip"
      text: "Use keyless signing in production"
    - type: "warning"
      text: "Never commit private keys to Git!"

  diagram: |
    Registry              Cosign                Your Cluster
    ┌────────┐            ┌──────┐             ┌───────────┐
    │ Image  │───sign────▶│ .sig │────verify──▶│ ✓ Deploy  │
    └────────┘            └──────┘             └───────────┘

  wait_for_continue: true
```

## Highlights

Style specific text in the content:

```yaml
highlights:
  - text: "myapp:latest"
    style: "code"         # rendered as inline code
  - text: "Sigstore"
    style: "bold"         # rendered bold
  - text: "important"
    style: "highlight"    # rendered with background color
```

## Code Blocks

Show labeled code snippets:

```yaml
code_blocks:
  - label: "macOS"
    code: "brew install cosign"
    language: "bash"
  - label: "Go Install"
    code: "go install github.com/sigstore/cosign/v2/cmd/cosign@latest"
    language: "bash"
```

## Callouts

Add contextual notes:

```yaml
callouts:
  - type: "tip"
    text: "Helpful advice"
  - type: "warning"
    text: "Be careful about this"
  - type: "danger"
    text: "Serious risk"
  - type: "info"
    text: "Additional information"
```

## Best Practices

- Keep content focused on **one concept** per info step
- Use Markdown formatting for readability
- Add diagrams for complex workflows
- Use callouts sparingly — one or two per step
- Set `wait_for_continue: true` for important steps
- Follow info steps with command steps for hands-on practice
