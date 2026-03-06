---
sidebar_position: 4
---

# Interactive Lesson Structure

Interactive lessons use the legacy single-file YAML format with step-by-step command execution and validation. This is more powerful than data-only lessons but requires more effort to write.

## Overview

```yaml
metadata:        # Required — lesson identification
introduction:    # Optional — shown once at start
steps:           # Required — interactive step sequence
conclusion:      # Optional — shown after completion
```

## The Four Sections

### 1. Metadata (Required)

```yaml
metadata:
  id: "cosign-sign-verify"
  org: "sigstore"
  title: "Container Image Signing with Cosign"
  description: "Learn to sign and verify container images"
  difficulty: "beginner"
  duration: "20 min"
  prerequisites:
    - "Docker installed"
  tags: ["cosign", "signing", "supply-chain"]
  author: "TryOutShell"
  version: "1.0"
  resources:
    - title: "Cosign Docs"
      url: "https://docs.sigstore.dev"
      type: "docs"
```

| Field | Required | Description |
|-------|----------|-------------|
| `id` | Yes | Unique identifier (kebab-case) |
| `org` | Yes | Organization/topic ID |
| `title` | Yes | Human-readable title |
| `description` | Yes | Brief summary |
| `difficulty` | Yes | `beginner`, `intermediate`, or `advanced` |
| `duration` | Yes | Estimated time |
| `tags` | Yes | Searchable keywords |
| `prerequisites` | No | Required tools or knowledge |
| `author` | No | Creator name |
| `version` | No | Lesson version |
| `resources` | No | External links |

### 2. Introduction (Optional)

Sets expectations before the lesson begins.

```yaml
introduction:
  title: "What You'll Learn"
  content: |
    In this lesson, you will:
    - Understand what Cosign is
    - Install and verify Cosign
    - Sign and verify container images

    **Time:** ~20 minutes
    **Tools:** Cosign, Docker
```

Supports full Markdown: bold, italic, code, lists, blockquotes, headers.

### 3. Steps (Required)

The interactive core. Each step has a `type` that determines its behavior.

| Type | Purpose |
|------|---------|
| `info` | Display educational content |
| `command` | Execute and validate shell commands |
| `quiz` | Multiple-choice knowledge check |
| `challenge` | Open-ended hands-on task |
| `interview_prep` | Practice interview questions |

See the [Step Types](../step-types/) reference for full details on each type.

### 4. Conclusion (Optional)

Summarizes learning and awards badges.

```yaml
conclusion:
  title: "Congratulations!"
  content: |
    You've learned:
    - How to sign container images
    - How to verify signatures
    - Supply chain security basics

    **Next:** Try the Keyless Signing lesson

  badges:
    - id: "cosign-basics"
      name: "Cosign Fundamentals"
      icon: "🔐"
```

## Complete Example

```yaml
metadata:
  id: "docker-intro"
  org: "docker"
  title: "Docker Introduction"
  description: "Learn basic Docker commands"
  difficulty: "beginner"
  duration: "15 min"
  tags: ["docker", "containers"]

introduction:
  title: "Welcome to Docker"
  content: |
    Docker packages applications in containers.
    You'll learn to run your first container.

steps:
  - type: info
    title: "What is Docker?"
    content: |
      **Docker** is a platform for running applications
      in lightweight, portable containers.

  - type: command
    prompt: "Verify Docker is installed"
    example: "docker --version"
    validation:
      type: "substring"
      contains: "Docker version"
    success_msg: "✅ Docker is installed!"
    fail_msg: "❌ Docker not found."

  - type: command
    prompt: "Run your first container"
    example: "docker run hello-world"
    validation:
      type: "substring"
      contains: "Hello from Docker"
    success_msg: "✅ Container ran successfully!"
    fail_msg: "❌ Container failed to run"
    hints:
      - level: 1
        text: "Try: docker run hello-world"

  - type: quiz
    title: "Quick Check"
    questions:
      - id: "q1"
        question: "What does docker ps show?"
        type: "multiple_choice"
        options:
          - "All images"
          - "Running containers"
          - "Docker version"
        answer: 1
        explanation: "docker ps lists running containers."

conclusion:
  title: "Well Done!"
  content: |
    🎉 You've learned Docker basics!
  badges:
    - id: "docker-beginner"
      name: "Docker Beginner"
      icon: "🐳"
```

## File Location

Interactive lessons are single YAML files placed at:

```
lessons/<org-id>/<lesson-id>.yaml
```

They are auto-discovered alongside the newer directory-based format.

## Best Practices

- **Keep it focused** — one main concept per lesson
- **Start simple** — info steps first, then commands
- **Progressive difficulty** — easy commands → complex tasks → challenges
- **Test everything** — run `go run . start <org> --lesson <id>` end-to-end
- **Provide hints** — 3 progressive hints per command step
- **Include quizzes** — reinforce learning with 3-5 questions

## What's Next?

- [Step Types Reference](../step-types/) — Deep dive into each step type
- [Creating Data-Only Lessons](./creating-lessons) — The simpler lesson format
