```
 _____                 _   ____  _          _ _
|_   _| __ _   _  ___ | |_/ ___|| |__   ___| | |
  | || '__| | | |/ _ \| __\___ \| '_ \ / _ \ | |
  | || |  | |_| | (_) | |_ ___) | | | |  __/ | |
  |_||_|   \__, |\___/ \__|____/|_| |_|\___|_|_|
           |___/
       🚀 Interactive Learning in Your Terminal
```

# tryoutshell

An interactive, terminal-based learning tool that lets you explore security topics and tools without leaving your shell.

## Overview

`tryoutshell` provides hands-on lessons for open-source supply chain security projects — such as Sigstore, Cosign, In-Toto, and Witness — directly in your terminal. Each lesson walks you through concepts, commands, and exercises step by step.

## Installation

**Requirements:** Go 1.24 or later.

```bash
git clone https://github.com/tryoutshell/tryoutshell.git
cd tryoutshell
go build -o tryoutshell .
```

Move the binary to somewhere on your `$PATH`:

```bash
sudo mv tryoutshell /usr/local/bin/
```

## Usage

```
tryoutshell [command]
```

### Commands

| Command | Description |
|---------|-------------|
| `start [org]` | Start an interactive learning session. Omit `[org]` to pick from a list. |
| `list` | Browse available organizations and lessons interactively. |
| `present <file.md>` | Present a Markdown file as full-screen terminal slides. |

### Examples

```bash
# Launch the interactive org/lesson picker
tryoutshell start

# Jump straight into a specific org
tryoutshell start sigstore

# Start a specific lesson directly
tryoutshell start sigstore --lesson cosign-101

# Browse everything interactively
tryoutshell list

# Present a markdown file as slides
tryoutshell present docs/intro.md
```

### `present` navigation

| Key | Action |
|-----|--------|
| `space` / `→` / `↓` / `enter` / `n` / `j` / `l` | Next slide |
| `←` / `↑` / `p` / `h` / `k` / `Shift+N` | Previous slide |
| `gg` | First slide |
| `G` | Last slide |
| `<number> G` | Jump to slide number |
| `/` | Search slides |
| `ctrl+u` / `ctrl+d` | Scroll up / down within a slide |
| `q` / `ctrl+c` | Quit |

## Available Organizations

| ID | Name | Description |
|----|------|-------------|
| `sigstore` | Sigstore | Supply chain security tools |
| `chainguard` | Chainguard | Supply chain security |
| `in-toto` | In-Toto | Supply chain integrity framework |
| `witness` | Witness | Attestation and verification |

## Contributing

Contributions are welcome! To add a new lesson or organization:

1. Fork the repository and create a feature branch.
2. Add your lesson files under `lessons/<org-id>/<lesson-id>/`.
3. Register the new lesson in `manifest.json`.
4. Open a pull request describing your changes.

## License

This project is licensed under the terms of the [LICENSE](LICENSE) file.
