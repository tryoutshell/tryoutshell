---
sidebar_position: 1
slug: /
---

# Welcome to TryOutShell

**TryOutShell** is a CLI-native interactive learning engine that brings guided hands-on labs directly into your terminal — no browser UI, no cloud VMs required.

## What is TryOutShell?

TryOutShell is a command-line learning environment where users can:

- Learn real-world DevOps & cybersecurity workflows
- Practice hands-on with tools like Cosign, Witness, In-Toto, and Chainguard
- Get interactive, step-by-step guidance right in their terminal
- Receive context-aware feedback on commands they execute

**Think of it as:** *Katacoda + Charm.sh + Learn-by-Doing → in your own terminal*

## Vision

TryOutShell brings interactive learning experiences where engineers are most comfortable: **the CLI**. Users type commands, get real output from their system, and receive beautifully formatted feedback — all inside a custom terminal UI built with Charm.sh's Bubble Tea and Lipgloss.

## Core Features

- ✅ **CLI-Native**: No web browser required
- ✅ **Interactive Learning**: Step-by-step guided lessons
- ✅ **Real Environment**: Practice on your actual system
- ✅ **Beautiful UI**: Built with Charm.sh Bubble Tea & Lipgloss
- ✅ **Open Source**: All lessons and binaries are transparent
- ✅ **Extensible**: Create your own lessons in YAML

## How It Works

1. **Install TryOutShell** via curl, brew, or GitHub releases
2. **Browse available lessons** in your terminal
3. **Launch a lesson** and follow interactive steps
4. **Execute commands** and get real-time validation
5. **Complete challenges** and earn badges

## Installation
```bash
# Quick install (recommended)
curl -sSL https://tryoutshell.lol | sh

# Or via Homebrew
brew install tryoutshell

# Or via GitHub Releases
# Download from https://github.com/your-org/tryoutshell/releases
```

This script will:
- Download and install the latest tryoutshell binary
- Clone or update the default lesson repository (`~/.tryoutshell/lessons/`)
- Verify SHA256 checksum for security
- Add tryoutshell to your `$PATH`

All binaries and lessons are **open-source** and **cryptographically signed** for transparency.

## Quick Start

After installation, launch your first lesson:
```bash
tryoutshell start docker-basics
```

## What's Next?

- 📚 [Getting Started Guide](./getting-started/) - Install and run your first lesson
- 📝 [Creating Lessons](./getting-started/minimal-example) - Build your own interactive lessons
- 🔧 [Step Types Reference](./step-types/) - Learn about all available step types
- 💡 [Best Practices](./guides/best-practices) - Tips for writing great lessons

## Community & Support

- 🐛 Report issues on [GitHub](https://github.com/your-org/tryoutshell/issues)
- 💬 Join our [Discord](https://discord.gg/tryoutshell)
- 📖 Browse [example lessons](./examples/)
- ⭐ Star us on [GitHub](https://github.com/your-org/tryoutshell)

---

Ready to start learning? Head to the [Getting Started](./getting-started/) guide!
