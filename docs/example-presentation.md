---
sidebar_position: 2
---

# Welcome to TryOutShell Slides

Learn, explore, and present — all from your terminal.

---

## What is TryOutShell?

TryOutShell is an **interactive terminal learning platform** that brings
lessons, quizzes, presentations, and an AI blog reader to your CLI.

> No browser needed. Just your terminal.

---

## Key Features

- **Slide lessons** — YAML + Markdown, no code needed to contribute
- **Interactive labs** — run real commands with validation
- **Quiz mode** — multiple-choice with explanations and scoring
- **AI blog reader** — read any URL in a split-pane TUI with chat
- **Progress tracking** — completion, scores, and time spent
- **Shell completions** — bash, zsh, fish

---

## Navigation

| Key | Action |
|-----|--------|
| `space` / `right` | Next slide |
| `left` | Previous slide |
| `gg` | First slide |
| `G` | Last slide |
| `<n>G` | Jump to slide n |
| `/` | Search |
| `?` | Help overlay |
| `q` | Quit |

---

## Code Is First-Class

Every slide renders syntax-highlighted code blocks:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello from TryOutShell!")
}
```

```bash
tryoutshell start docker --lesson docker-101
tryoutshell quiz docker docker-101
tryoutshell read https://blog.sigstore.dev/cosign-2-0
```

---

## Creating a Lesson

Just two files — no Go code needed:

```
lessons/my-org/my-lesson/
  lesson.yaml     # metadata + quiz questions
  slides.md       # slide content (--- separated)
```

Try it:

```bash
cp -r lessons/_template/my-lesson lessons/my-org/my-lesson
# Edit the files, then:
tryoutshell start my-org --lesson my-lesson
```

---

## Getting Started

```bash
git clone https://github.com/tryoutshell/tryoutshell.git
cd tryoutshell
go build -o tryoutshell .

# Browse lessons
./tryoutshell list

# Present this file
./tryoutshell present docs/example-presentation.md
```

---

# Thank You!

Explore the lessons:

```bash
tryoutshell list
```

Read the docs, contribute a lesson, or write your own slides.

> Everything from the terminal. No browser required.
