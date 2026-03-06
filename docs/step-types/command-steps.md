---
sidebar_position: 3
---

# Command Steps

Command steps ask the user to execute shell commands and validate the output.

## Purpose

Use command steps for:
- Hands-on practice with real commands
- Verifying tool installations
- Building and testing artifacts
- Running operations that produce verifiable output

## Basic Structure

```yaml
- type: command
  prompt: "What to do"
  example: "echo 'Hello'"
  validation:
    type: "substring"
    contains: "Hello"
  success_msg: "✅ It works!"
  fail_msg: "❌ Try again"
```

## Fields Reference

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | `"command"` | Yes | Identifies this as a command step |
| `id` | string | No | Unique ID for tracking |
| `prompt` | string | Yes | Short description shown in UI |
| `instruction` | string | No | Detailed instruction text |
| `example` | string | Yes | Example command shown to user |
| `pre_content` | string | No | Explanation before the command |
| `post_content` | string | No | Explanation after success |
| `accepted_commands` | list | No | Valid command variations |
| `validation` | object | Yes | How to verify success |
| `alternative_validations` | list | No | Additional validation methods |
| `success_msg` | string | Yes | Message on success |
| `fail_msg` | string | Yes | Message on failure |
| `hints` | list | No | Progressive hints (shown with `?`) |
| `allow_skip` | boolean | No | Can user skip? (default: false) |
| `timeout` | integer | No | Max execution time in seconds (default: 30) |

## Full Example

```yaml
- type: command
  id: "check-cosign"
  prompt: "Verify Cosign is installed"
  instruction: "Run the command to check your Cosign version:"

  pre_content: |
    First, let's make sure Cosign is available on your system.

  example: "cosign version"

  accepted_commands:
    - "cosign version"
    - "cosign --version"

  validation:
    type: "regex"
    pattern: "GitVersion|cosign.*version"
    case_insensitive: true

  alternative_validations:
    - type: "exit_code"
      expected: 0

  post_content: |
    Great! Cosign is installed and ready to use.

  success_msg: "✅ Cosign detected!"
  fail_msg: "❌ Cosign not found. Install it first."

  hints:
    - level: 1
      text: "Type: cosign version"
    - level: 2
      text: "If not installed: brew install cosign"
    - level: 3
      text: "Full docs: https://docs.sigstore.dev/cosign/installation"

  allow_skip: true
  timeout: 10
```

## Validation Types

### substring

Output must contain a specific string.

```yaml
validation:
  type: "substring"
  contains: "Docker version"
  case_insensitive: true     # optional
```

### regex

Output must match a regular expression.

```yaml
validation:
  type: "regex"
  pattern: "v\\d+\\.\\d+\\.\\d+"
  case_insensitive: false
```

### exit_code

Command must exit with a specific code.

```yaml
validation:
  type: "exit_code"
  expected: 0
```

### file_exists

Specified files must exist after the command runs.

```yaml
validation:
  type: "file_exists"
  files:
    - "cosign.key"
    - "cosign.pub"
```

### file_contains

A file must contain a pattern.

```yaml
validation:
  type: "file_contains"
  path: "output.txt"
  pattern: "success"
```

### output_contains

Output must contain one or more patterns.

```yaml
validation:
  type: "output_contains"
  patterns:
    - "Successfully signed"
    - "Pushing signature"
  any_match: true    # match ANY pattern (OR)
  all_match: false   # match ALL patterns (AND)
```

## Progressive Hints

Users press `?` to reveal hints one at a time:

```yaml
hints:
  - level: 1
    text: "Try typing: cosign version"
  - level: 2
    text: "Install with: brew install cosign"
  - level: 3
    text: "Full command: cosign version"
```

Start vague, get more specific with each level.

## Best Practices

- **Always provide an example command** — users should never be stuck
- **Use multiple validation methods** — primary + alternative for flexibility
- **Include 2-3 hints** — from vague to explicit
- **Set reasonable timeouts** — 10s for quick commands, 60s for downloads
- **Allow skip on non-essential steps** — don't block progress
- **Use pre/post content** — explain the "why" before and the "what happened" after
