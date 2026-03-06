```
 _____                 _   ____  _          _ _
|_   _| __ _   _  ___ | |_/ ___|| |__   ___| | |
  | || '__| | | |/ _ \| __\___ \| '_ \ / _ \ | |
  | || |  | |_| | (_) | |_ ___) | | | |  __/ | |
  |_||_|   \__, |\___/ \__|____/|_| |_|\___|_|_|
           |___/
       🚀 Interactive Learning in Your Terminal
```

# TryOutShell

An interactive, terminal-based learning tool that lets you explore DevSecOps, containers, CI/CD, and security topics without leaving your shell.

> **TODO**: Record demo with [vhs](https://github.com/charmbracelet/vhs) and add `demo.gif` here.

## Features

- 📖 **Interactive Lessons** — step-by-step learning with slides, quizzes, and hands-on exercises
- 🤖 **AI Blog Reader** — read any article in a split-pane TUI with an AI assistant (`tryoutshell read <url>`)
- 📝 **Quizzes** — test your knowledge with multiple-choice quizzes after each lesson
- 📊 **Progress Tracking** — track completion, quiz scores, and time spent
- 🎨 **Beautiful TUI** — powered by Bubble Tea, Glamour, and Lip Gloss
- 🔌 **Data-Only Lessons** — add lessons with just YAML + Markdown, no Go code needed
- 🔄 **Remote Updates** — download new lessons without rebuilding
- 🐚 **Shell Completions** — bash, zsh, fish support

## Installation

**Requirements:** Go 1.24 or later.

```bash
git clone https://github.com/tryoutshell/tryoutshell.git
cd tryoutshell
go build -o tryoutshell .
sudo mv tryoutshell /usr/local/bin/
```

**Homebrew** (coming soon):
```bash
brew install tryoutshell/tap/tryoutshell
```

**Binary download**: Check the [Releases](https://github.com/tryoutshell/tryoutshell/releases) page.

## Quick Start

```bash
# Browse all lessons interactively
tryoutshell list

# Start a specific lesson
tryoutshell start docker --lesson docker-101

# Take a quiz
tryoutshell quiz docker docker-101

# Read a blog post with AI assistant
tryoutshell read https://blog.sigstore.dev/cosign-2-0

# Check your progress
tryoutshell progress
```

## Available Lessons

| Org | Lesson | Difficulty | Duration |
|-----|--------|-----------|----------|
| 🐳 Docker | Docker 101: Container Fundamentals | Beginner | 25 min |
| ☸️ Kubernetes | Kubernetes 101 | Beginner | 30 min |
| 📚 Git | Git Internals | Intermediate | 25 min |
| 🔒 Security | OWASP Top 10 (2021) | Intermediate | 30 min |
| ⚡ CI/CD | GitHub Actions CI/CD | Beginner | 25 min |
| 🔐 Supply Chain | SLSA (Supply-chain Levels) | Intermediate | 25 min |
| 🛡️ Sigstore | Container Image Signing with Cosign | Beginner | 20 min |
| 🛡️ Chainguard | Introduction to Rekor | Beginner | 15 min |
| 🔗 In-Toto | Supply Chain Security with in-toto | Beginner | 20 min |
| 🦉 Witness | Getting Started with Witness | Beginner | 15 min |

## Commands

| Command | Description |
|---------|-------------|
| `start [org]` | Start an interactive learning session |
| `list` | Browse available organizations and lessons |
| `present <file.md>` | Present a Markdown file as terminal slides |
| `quiz <org> <lesson>` | Launch quiz mode for a lesson |
| `read <url>` | Read a blog post in split-pane TUI with AI chat |
| `read <url> --save` | Save article for offline reading |
| `saved` | List and open saved articles |
| `progress` | Show learning progress summary |
| `update` | Check for and download lesson updates |
| `update --check` | Check for updates without downloading |
| `completion [shell]` | Generate shell completion scripts |

## `tryoutshell read` — AI Blog Reader

Read any article in a beautiful split-pane terminal UI with an AI assistant.

```bash
tryoutshell read https://blog.sigstore.dev/cosign-2-0

# With AI chat (set one of these):
export OPENAI_API_KEY=sk-...
export ANTHROPIC_API_KEY=sk-ant-...
export GEMINI_API_KEY=AI...
```

**Controls:**
- `Tab` — switch between article and chat panes
- `j/k` — scroll article
- `[/]` — resize split
- `t` — cycle themes (default, dark, light, dracula, nord)
- `q` — quit

## Shell Completions

```bash
# Bash
source <(tryoutshell completion bash)

# Zsh
source <(tryoutshell completion zsh)

# Fish
tryoutshell completion fish | source
```

## Contributing

Contributions are welcome! **Adding a lesson requires zero Go code** — just YAML and Markdown.

1. Copy `lessons/_template/` as a starting point
2. Write your `lesson.yaml` and `slides.md`
3. Test with `go run . start <your-org> -l <your-lesson>`
4. Open a PR

See [CONTRIBUTING.md](CONTRIBUTING.md) for the full guide.

## Roadmap

- [ ] Lesson bookmarks and notes
- [ ] Interactive exercise runner with inline output
- [ ] Lesson search and filtering
- [ ] Community lesson marketplace
- [ ] VHS-recorded demo GIF
- [ ] Homebrew tap
- [ ] Progress sync across devices

## License

This project is licensed under the terms of the [LICENSE](LICENSE) file.
