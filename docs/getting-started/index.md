---
sidebar_position: 1
---

# Getting Started

Learn how to install TryOutShell and create your first interactive lesson.

## Installation

### Quick Install (Recommended)
```bash
curl -sSL https://tryoutshell.lol | sh
```

This script will:
1. Download the latest `tryoutshell` binary from GitHub Releases
2. Verify the SHA256 checksum for security
3. Install the binary to `/usr/local/bin` (or `~/.local/bin`)
4. Clone the default lesson repository to `~/.tryoutshell/lessons/`
5. Add `tryoutshell` to your `$PATH`

### Alternative Installation Methods

#### Via Homebrew (macOS/Linux)
```bash
brew tap your-org/tryoutshell
brew install tryoutshell
```

#### Via GitHub Releases (Manual)

1. Download the latest release for your platform from [GitHub Releases](https://github.com/your-org/tryoutshell/releases)
2. Extract the archive:
```bash
   tar -xzf tryoutshell-linux-amd64.tar.gz
```
3. Move the binary to your PATH:
```bash
   sudo mv tryoutshell /usr/local/bin/
```
4. Make it executable:
```bash
   chmod +x /usr/local/bin/tryoutshell
```

#### Build from Source
```bash
git clone https://github.com/your-org/tryoutshell.git
cd tryoutshell
go build -o tryoutshell ./cmd/tryoutshell
sudo mv tryoutshell /usr/local/bin/
```

## Verify Installation

Check that TryOutShell is installed correctly:
```bash
tryoutshell version
```

You should see output like:
TryOutShell v1.0.0
Built with Go 1.21.0

## Your First Lesson

Let's run a simple lesson to see TryOutShell in action:
```bash
tryoutshell start hello-world
```

This will launch an interactive lesson that guides you through basic commands.

## What's Next?

- 📝 [Create a Minimal Lesson](./minimal-example) - Build your first lesson
- 📚 [Understand Lesson Structure](./lesson-structure) - Learn the anatomy of a lesson
- 🎯 [Browse Examples](../examples/) - See complete lesson examples

## Directory Structure

After installation, TryOutShell creates the following structure:
~/.tryoutshell/
├── lessons/              # Lesson repository
│   ├── docker-basics/
│   ├── cosign-intro/
│   └── ...
├── config.yaml           # User configuration
├── progress.db           # Lesson progress tracking
└── cache/                # Cached data

## Configuration

Edit `~/.tryoutshell/config.yaml` to customize TryOutShell:
```yaml
# Default lesson repository
lesson_repo: "https://github.com/your-org/tryoutshell-lessons.git"

# UI preferences
theme: "dark"  # Options: dark, light
animations: true

# Telemetry (optional, anonymized)
telemetry_enabled: false
```

## Updating TryOutShell

To update to the latest version:
```bash
tryoutshell update
```

Or reinstall via the installation script:
```bash
curl -sSL https://tryoutshell.lol | sh
```

## Troubleshooting

### Command not found

If you see `tryoutshell: command not found`, ensure it's in your PATH:
```bash
export PATH="$PATH:/usr/local/bin"
```

Add this to your `~/.bashrc` or `~/.zshrc` to make it permanent.

### Permission denied

If you get permission errors during installation:
```bash
sudo curl -sSL https://tryoutshell.lol | sh
```

Or install to a user directory:
```bash
mkdir -p ~/.local/bin
# Download and place binary in ~/.local/bin
export PATH="$PATH:~/.local/bin"
```

### Lessons not loading

Ensure the lesson repository is cloned:
```bash
git clone https://github.com/your-org/tryoutshell-lessons.git ~/.tryoutshell/lessons
```

## Support

Need help? Check out:

- 📖 [Documentation](/)
- 🐛 [GitHub Issues](https://github.com/your-org/tryoutshell/issues)
- 💬 [Discord Community](https://discord.gg/tryoutshell)
