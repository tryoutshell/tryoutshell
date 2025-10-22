# tryoutshell 🦦

> first command "tryout list" - we might change it to make it default to show on first page or something
> second command "tryout start [lesson]" : i don't think so i need to add any flag to select ?? (start looks good i tried to 
> go with "select first" but does not look good to me)

```md
$ tryoutshell list (might be using their logo here not possible ?? or possible)
> in-toto
> witness
> chainguard
> unstable-ai

$ tryoutshell select in-toto (ok!! but start looks good)
$ tryoutshell start in-toto --theme bubble --color blue
```
```md
tryoutshell (main) $ ./bin/tryoutshell list
0 
[]
┌────────────┬─────────┬────────┐
│   LESSON   │ VERSION │ STATUS │
├────────────┼─────────┼────────┤
│ in-toto    │ v0.0.5  │ stable │
│ chainguard │ v1.1.0  │ latest │
│ witness    │ v1.1.0  │ latest │
└────────────┴─────────┴────────┘
tryoutshell (main) $ ./bin/tryoutshell list in-toto
1 
[in-toto]
┌────────┬───────────────────────────────────┬────────┐
│ S . NO │             DEMO NAME             │ STATUS │
├────────┼───────────────────────────────────┼────────┤
│ 1      │ in-toto demo                      │ stable │
│ 2      │ in-toto attestation sign & verify │ stable │
│ 3      │ in-toto dsse                      │ stabel │
└────────┴───────────────────────────────────┴────────┘

> we are doing this hardcoded but we need to handle this all from the `manifest.json` 
> so the manifest.json will have into the lesson repo or we can also have that manifest.json in ours main repo (fallback) 

> we will read everything from the `manifest.json` will look at orgs, lesson present under `orgs`

```json
{
  "organizations": [
    {
      "id": "in-toto",
      "name": "In-Toto",
      "description": "Supply chain integrity",
      "logo": "🔗",
      "lessons": ["getting-started", "advanced-policies"]
    }
  ]
}
```
```

---
```md
tryoutshell/
├── cmd/
│   ├── root.go          # Entry point, TTY detection
│   ├── start.go         # Start command
│   ├── list.go          # List command (TUI)
│   ├── config.go        # Config management
│   └── create.go        # Lesson creation wizard
├── internal/
│   ├── ui/
│   │   ├── models.go    # Bubble Tea state machine
│   │   ├── views.go     # Rendering logic
│   │   ├── themes.go    # Theme definitions
│   │   └── styles.go    # Lipgloss styles
│   ├── runner/
│   │   ├── executor.go  # Command execution
│   │   ├── safety.go    # Security checks
│   │   └── verifier.go  # Output validation
│   ├── lessons/
│   │   ├── type.go    # YAML → Go structs (all type)
│   │   ├── loader.go    # Load from filesystem
│   │   └── validator.go # Validate lesson format
│   ├── progress/
│   │   ├── db.go        # SQLite wrapper
│   │   └── badges.go    # Badge logic
│   └── config/
│       └── loader.go    # Config management
├── pkg/
│   └── types/           # Shared types
└── lessons/
    ├── manifest.json
    └── [org]/
        └── [lesson].yaml
```

```md
# Main interactive mode (default)
$ tryoutshell
$ tryoutshell start

# Direct organization selection
$ tryoutshell start in-toto --theme bubble --color blue
$ tryoutshell start chainguard --lesson cosign-basics

# List organizations/lessons
$ tryoutshell list
$ tryoutshell list in-toto

# Resume from checkpoint
$ tryoutshell resume

# Config & management
$ tryoutshell config --set theme=bubble
$ tryoutshell update  # Updates lesson repo
$ tryoutshell version
```
```md
rootCmd (interactive UI if no args)
├── start [org] [--theme] [--color] [--lesson]
├── list [org]
├── resume [org/lesson]
├── config [--set key=value]
├── update
└── version
```

## thinking...
`tryoutshell start in-toto --theme bubble --color blue`
1. we will except one args
2. we will look at some flags `--theme`, `--color`, `--lesson`