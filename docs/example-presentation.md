# Welcome to TryOutShell Slides

Learn, explore, and present – all from your terminal.

---

## What is TryOutShell?

TryOutShell is a **terminal-based learning platform** that brings
interactive lessons, quizzes, and now presentations right to your
command line.

> No browser needed. Just your terminal.

---

## Key Features

- 🎓 **Interactive lessons** – run real commands and get instant feedback
- 📝 **Quizzes** – test your knowledge without leaving the terminal
- 🎞  **Slide presentations** – share ideas or read blog posts as slides
- 🔍 **Search** – jump to any slide instantly with `/`

---

## Navigation

| Key | Action |
|-----|--------|
| `space` / `→` | Next slide |
| `←` | Previous slide |
| `gg` | First slide |
| `G` | Last slide |
| `<n>G` | Jump to slide n |
| `/` | Search |
| `q` | Quit |

---

## Code is a First-Class Citizen

Every slide can contain syntax-highlighted code blocks:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello from TryOutShell! 🐚")
}
```

```bash
# Run a lesson
tryoutshell start sigstore --lesson cosign-101

# Present a markdown file
tryoutshell present my-talk.md
```

---

## Inspired by `slides`

This presentation mode is inspired by the excellent
**maaslalani/slides** project.

Key improvements in TryOutShell:

1. Deeply integrated with interactive lessons
2. Syntax highlighting via **chroma**
3. Search across all slides
4. Navigate by slide number (`5G` → slide 5)

---

## Getting Started

Install TryOutShell and try presenting this file:

```bash
go install github.com/tryoutshell/tryoutshell@latest

tryoutshell present docs/example-presentation.md
```

Or pipe content from the web:

```bash
curl -s https://raw.githubusercontent.com/.../slides.md \
  | tryoutshell present /dev/stdin
```

---

# Thank You! 🎉

Explore the lessons:

```bash
tryoutshell start
```

Read the docs, contribute a lesson, or write your own slides.

> Everything from the terminal. No browser required.
