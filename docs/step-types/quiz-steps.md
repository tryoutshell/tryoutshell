---
sidebar_position: 4
---

# Quiz Steps

Quiz steps test knowledge with multiple-choice questions within interactive lessons.

> **Note**: For data-only lessons, quizzes are defined in `lesson.yaml` and launched with `tryoutshell quiz <org> <lesson>`.

## Basic Structure

```yaml
- type: quiz
  title: "Knowledge Check"
  questions:
    - id: "q1"
      question: "Question text?"
      type: "multiple_choice"
      options: ["Option A", "Option B", "Option C"]
      answer: 1
      explanation: "Why B is correct."
```

## Fields Reference

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | `"quiz"` | Yes | Identifies this as a quiz step |
| `title` | string | No | Section header |
| `questions` | list | Yes | List of questions |

### Question Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | Yes | Unique question ID |
| `question` | string | Yes | The question text |
| `type` | string | Yes | `"multiple_choice"` |
| `options` | list | Yes | Answer options (3-4 recommended) |
| `answer` | integer | Yes | Correct option index (0-based) |
| `explanation` | string | No | Shown after answering |

## Full Example

```yaml
- type: quiz
  title: "Container Security Quiz"
  questions:
    - id: "q1"
      question: "What does Cosign use to sign container images?"
      type: "multiple_choice"
      options:
        - "Passwords"
        - "Public key cryptography"
        - "OAuth tokens"
        - "SSH keys"
      answer: 1
      explanation: |
        Cosign uses **public key cryptography** (key pairs
        or keyless OIDC signing). The private key signs,
        and the public key verifies.

    - id: "q2"
      question: "Where are Cosign signatures stored?"
      type: "multiple_choice"
      options:
        - "Local filesystem"
        - "Git repository"
        - "Container registry as OCI artifacts"
        - "Separate database"
      answer: 2
      explanation: |
        Signatures are stored in the same container registry
        as OCI artifacts, tagged with a `.sig` suffix.

    - id: "q3"
      question: "Which key should you NEVER share?"
      type: "multiple_choice"
      options:
        - "cosign.pub (public key)"
        - "cosign.key (private key)"
        - "Both"
        - "Neither"
      answer: 1
      explanation: |
        The **private key** (cosign.key) must remain secret.
        The public key is meant to be shared for verification.
```

## Data-Only Lesson Quiz Format

In `lesson.yaml`, quizzes use a simpler format:

```yaml
quiz:
  - question: "What does docker ps show?"
    options:
      - "All images"
      - "Running containers"
      - "Docker version"
      - "Network settings"
    answer: 1
    explain: "docker ps lists running containers."
```

| Field | Type | Description |
|-------|------|-------------|
| `question` | string | Question text |
| `options` | list | Answer options |
| `answer` | integer | Correct index (0-based) |
| `explain` | string | Explanation |

Launch standalone: `tryoutshell quiz <org> <lesson>`

## Best Practices

- **3-5 questions** per quiz step
- **4 options** per question (all plausible)
- **Always include explanation** — teach even when the answer is wrong
- **Place quizzes after related content** — test what was just taught
- **Cover key concepts** — not trivia or trick questions
- **0-indexed answers** — first option is 0, second is 1, etc.
