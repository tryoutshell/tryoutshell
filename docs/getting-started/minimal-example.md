---
sidebar_position: 3
---

# Minimal Lesson Example

The simplest possible TryOutShell lesson — just two files.

## Data-Only Lesson (New Format)

### lesson.yaml

```yaml
id: hello-world
title: "Hello World"
description: "Your first TryOutShell lesson"
difficulty: "beginner"
duration: "5 min"
version: "1.0"
tags: ["intro"]

quiz:
  - question: "What command prints text to the terminal?"
    options: ["ls", "echo", "cd", "pwd"]
    answer: 1
    explain: "echo prints text to standard output."
```

### slides.md

```markdown
# Hello World

Welcome to your first TryOutShell lesson!

---

## What You'll Learn

- How TryOutShell slides work
- Basic terminal navigation

```bash
echo "Hello from TryOutShell!"
```

---

## Summary

You've completed your first lesson!

Try running the quiz: `tryoutshell quiz my-org hello-world`
```

### Directory Layout

```
lessons/my-org/
  meta.yaml
  hello-world/
    lesson.yaml
    slides.md
```

### Run It

```bash
go run . start my-org --lesson hello-world
go run . quiz my-org hello-world
```

---

## Interactive Lesson (Legacy Format)

For lessons that need real command execution and validation, use the single-file YAML format:

```yaml
metadata:
  id: "hello-world"
  org: "tutorial"
  title: "Hello World"
  description: "Your first interactive lesson"
  difficulty: "beginner"
  duration: "5 min"
  tags: ["intro"]

steps:
  - type: info
    title: "Welcome!"
    content: "Hello! This is your first interactive lesson."

  - type: command
    prompt: "Say hello"
    example: "echo 'Hello from TryOutShell'"
    validation:
      type: "substring"
      contains: "Hello"
    success_msg: "✅ Perfect! You said hello!"
    fail_msg: "❌ Try running the example command"
```

Save as `lessons/tutorial/hello-world.yaml` and run:

```bash
go run . start tutorial --lesson hello-world
```

The interactive format supports command execution, output validation, progressive hints, challenges, and more. See the [Step Types](../step-types/) reference for details.

## What's Next?

- [Creating a Full Lesson](./creating-lessons) — Complete guide to data-only lessons
- [Lesson Structure](./lesson-structure) — Deep dive into interactive lesson format
- [Step Types](../step-types/) — All available step types
