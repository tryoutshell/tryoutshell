---
sidebar_position: 1
---

# Getting Started

Install TryOutShell and run your first lesson in under 2 minutes.

## Installation

### Build from Source (Recommended)

**Requirements:** Go 1.24 or later.

```bash
git clone https://github.com/tryoutshell/tryoutshell.git
cd tryoutshell
go build -o tryoutshell .
sudo mv tryoutshell /usr/local/bin/
```

### Via Go Install

```bash
go install github.com/tryoutshell/tryoutshell@latest
```

### Via Homebrew (Coming Soon)

```bash
brew install tryoutshell/tap/tryoutshell
```

### Binary Download

Download pre-built binaries from [GitHub Releases](https://github.com/tryoutshell/tryoutshell/releases).

## Verify Installation

```bash
tryoutshell --help
```

You should see the help output listing all available commands.

## Your First Lesson

### Browse available lessons

```bash
tryoutshell list
```

This opens an interactive picker. Select an organization, then a lesson.

### Start a specific lesson

```bash
tryoutshell start docker --lesson docker-101
```

### Take a quiz

```bash
tryoutshell quiz docker docker-101
```

### Check your progress

```bash
tryoutshell progress
```

## All Commands

| Command | Description |
|---------|-------------|
| `tryoutshell list` | Browse organizations and lessons interactively |
| `tryoutshell start [org]` | Start a learning session (optionally specify org) |
| `tryoutshell start [org] -l [lesson]` | Start a specific lesson directly |
| `tryoutshell quiz <org> <lesson>` | Launch quiz mode for a lesson |
| `tryoutshell read <url>` | Read a blog post in split-pane TUI with AI chat |
| `tryoutshell read <url> --save` | Save article for offline reading |
| `tryoutshell saved` | Open saved articles |
| `tryoutshell progress` | Show learning progress summary |
| `tryoutshell present <file.md>` | Present any markdown file as slides |
| `tryoutshell update` | Download new/updated lessons |
| `tryoutshell update --check` | Check for updates without downloading |
| `tryoutshell completion [bash\|zsh\|fish]` | Generate shell completions |

## Shell Completions

```bash
# Bash
source <(tryoutshell completion bash)

# Zsh
source <(tryoutshell completion zsh)

# Fish
tryoutshell completion fish | source
```

## Directory Layout

Lessons are stored in the repo under `lessons/`:

```
lessons/
  <org-id>/
    meta.yaml           ← organization metadata (name, logo)
    <lesson-id>/
      lesson.yaml       ← lesson metadata + quiz questions
      slides.md         ← slide content (--- separated)
      exercises.sh      ← optional exercises
    legacy-lesson.yaml  ← interactive lesson (old format, still works)
```

Progress is stored at `~/.config/tryoutshell/progress.json`.

Saved articles are at `~/.local/share/tryoutshell/saved/`.

## What's Next?

- [Creating a Lesson](./creating-lessons) — Add a lesson with just YAML + Markdown
- [Lesson Structure](./lesson-structure) — Understand the anatomy of interactive lessons
- [Step Types](../step-types/) — Deep dive into each step type
