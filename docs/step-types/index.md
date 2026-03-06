---
sidebar_position: 1
---

# Step Types Overview

Steps are the core of interactive TryOutShell lessons. Each step has a `type` that determines its behavior.

> **Note**: Step types apply to the interactive (legacy) lesson format. Data-only lessons use `slides.md` for content and `lesson.yaml` for quizzes.

## Available Step Types

| Type | Purpose | User Interaction |
|------|---------|-----------------|
| [`info`](./info-steps) | Display educational content | Read and press Enter |
| [`command`](./command-steps) | Execute shell commands | Type commands, get validation |
| [`quiz`](./quiz-steps) | Test knowledge | Answer multiple-choice questions |
| [`challenge`](./challenge-steps) | Open-ended tasks | Complete complex objectives |
| [`interview_prep`](./interview-prep-steps) | Practice questions | Reflect and write answers |

## When to Use Each Type

### info — Explain concepts

```yaml
- type: info
  title: "What is Docker?"
  content: |
    **Docker** is a platform for running applications
    in lightweight containers.
```

Best for: theory, context, diagrams, tips before hands-on steps.

### command — Hands-on practice

```yaml
- type: command
  prompt: "Check Docker version"
  example: "docker --version"
  validation:
    type: "substring"
    contains: "Docker version"
  success_msg: "✅ Docker is installed!"
  fail_msg: "❌ Docker not found"
```

Best for: executing real commands with output validation.

### quiz — Knowledge check

```yaml
- type: quiz
  title: "Quick Check"
  questions:
    - id: "q1"
      question: "What command lists Docker images?"
      type: "multiple_choice"
      options: ["docker ps", "docker images", "docker list"]
      answer: 1
      explanation: "docker images lists all local images."
```

Best for: reinforcing key concepts after info/command steps.

### challenge — Apply knowledge

```yaml
- type: challenge
  title: "Deploy Your App"
  description: |
    Create a Dockerfile, build it, and run the container.
  verification:
    type: "custom"
    checks:
      - type: "file_exists"
        path: "Dockerfile"
  success_msg: "🎉 Challenge complete!"
```

Best for: open-ended tasks where users demonstrate understanding.

### interview_prep — Deeper thinking

```yaml
- type: interview_prep
  title: "Think Deeper"
  questions:
    - "Explain the difference between images and containers"
    - "How would you secure a production container?"
  record_answers: true
```

Best for: reflection, preparing for technical interviews.

## Choosing the Right Type

| Scenario | Recommended Type |
|----------|-----------------|
| Explaining a concept | `info` |
| User needs to run a command | `command` |
| Testing if user understood | `quiz` |
| User proves mastery | `challenge` |
| Open-ended discussion | `interview_prep` |

## Recommended Lesson Flow

A well-structured lesson typically follows this pattern:

1. **info** — Introduce the concept
2. **command** — Practice with guided commands
3. **info** — Explain what happened
4. **command** — More complex commands
5. **quiz** — Check understanding
6. **challenge** — Apply knowledge independently

## Next Steps

- [Info Steps](./info-steps) — Displaying content
- [Command Steps](./command-steps) — Executing and validating commands
- [Quiz Steps](./quiz-steps) — Multiple-choice questions
- [Challenge Steps](./challenge-steps) — Open-ended tasks
- [Interview Prep Steps](./interview-prep-steps) — Practice questions
