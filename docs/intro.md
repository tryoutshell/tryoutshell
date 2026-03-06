---
sidebar_position: 1
slug: /
---

# Welcome to TryOutShell

**TryOutShell** is an interactive, terminal-based learning tool that brings hands-on lessons, quizzes, and an AI-powered blog reader directly into your terminal.

## What is TryOutShell?

TryOutShell is a CLI learning platform where you can:

- Learn DevSecOps, containers, CI/CD, Git, and security — step-by-step in your terminal
- Take quizzes to test your knowledge
- Read any blog post in a split-pane TUI with an AI assistant
- Track your progress across all lessons
- Add new lessons with just YAML + Markdown (no Go code needed)

**Think of it as:** *Interactive slides + hands-on labs + AI reader — all in your terminal*

## Core Features

- **Slide-Based Lessons** — Markdown slides with quizzes, no code needed to contribute
- **Interactive Lessons** — Step-by-step labs with real command execution and validation
- **AI Blog Reader** — Read any URL in a split-pane TUI with OpenAI / Anthropic / Gemini chat
- **Quiz Mode** — Multiple-choice quizzes with explanations and score tracking
- **Progress Tracking** — Track completion, quiz scores, and time spent
- **Beautiful TUI** — Built with Bubble Tea, Glamour, and Lip Gloss
- **Shell Completions** — bash, zsh, fish support
- **Remote Updates** — Download new lessons without rebuilding

## Quick Start

```bash
# Install
git clone https://github.com/tryoutshell/tryoutshell.git
cd tryoutshell
go build -o tryoutshell .
sudo mv tryoutshell /usr/local/bin/

# Browse lessons
tryoutshell list

# Start a lesson
tryoutshell start docker --lesson docker-101

# Take a quiz
tryoutshell quiz docker docker-101

# Read a blog post with AI chat
OPENAI_API_KEY=sk-... tryoutshell read https://blog.sigstore.dev/cosign-2-0

# Check your progress
tryoutshell progress
```

## Two Lesson Formats

TryOutShell supports two lesson formats:

### 1. Data-Only Lessons (Recommended for new content)

Just YAML metadata + Markdown slides. **No Go code needed.**

```
lessons/docker/docker-101/
  lesson.yaml     ← metadata + quiz questions
  slides.md       ← slide content (--- separated)
  exercises.sh    ← optional exercises
```

### 2. Interactive Lessons (Legacy format)

Rich YAML with step-by-step command execution, validation, and inline quizzes.

```
lessons/sigstore/cosign-sign-verify.yaml
```

Both formats are auto-discovered — just put files in the `lessons/` directory.

## What's Next?

- [Getting Started](./getting-started/) — Install and run your first lesson
- [Creating Lessons](./getting-started/creating-lessons) — Build a data-only lesson in 5 minutes
- [Step Types Reference](./step-types/) — All step types for interactive lessons
- [AI Blog Reader](./features/reader) — Using `tryoutshell read`

## Community

- [GitHub](https://github.com/tryoutshell/tryoutshell) — Star, fork, and contribute
- [Issues](https://github.com/tryoutshell/tryoutshell/issues) — Report bugs or request features
- [Contributing Guide](https://github.com/tryoutshell/tryoutshell/blob/main/CONTRIBUTING.md) — How to add lessons
