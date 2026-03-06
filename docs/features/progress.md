---
sidebar_position: 2
---

# Progress Tracking

TryOutShell tracks your learning progress across all lessons.

## View Progress

```bash
tryoutshell progress
```

This shows a table with:
- Organization and lesson names
- Completion status (✓ completed, ⟳ in progress, ◌ not started)
- Quiz scores
- Time spent
- Overall completion percentage

## What's Tracked

| Data | When Updated |
|------|-------------|
| Completion status | When you finish all slides or steps |
| Quiz score | After completing a quiz |
| Time spent | While viewing lessons |
| Last accessed | Each time you open a lesson |

## Storage

Progress is stored locally at:

```
~/.config/tryoutshell/progress.json
```

This file is created automatically on first use.

## Reset Progress

To start fresh, delete the progress file:

```bash
rm ~/.config/tryoutshell/progress.json
```
