---
sidebar_position: 6
---

# Interview Prep Steps

Interview prep steps provide open-ended questions for users to practice and reflect on concepts.

## Purpose

Use interview prep steps to:
- Reinforce learning through reflection
- Practice articulating technical concepts
- Prepare for real interview scenarios
- Encourage deeper understanding beyond commands

## Basic Structure

```yaml
- type: interview_prep
  title: "Interview Questions"
  questions:
    - "Explain concept X in simple terms."
    - "How would you implement Y in production?"
```

## Fields Reference

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | `"interview_prep"` | Yes | Identifies this type |
| `title` | string | Yes | Section heading |
| `description` | string | No | Instructions |
| `questions` | list | Yes | List of question strings |
| `record_answers` | boolean | No | Save user's answers (default: false) |
| `export_format` | string | No | `"json"` or `"text"` (default: `"json"`) |

## Examples

### Basic

```yaml
- type: interview_prep
  title: "Practice Questions"
  questions:
    - "What is Docker and why is it useful?"
    - "Explain the difference between images and containers."
    - "When would you use Docker Compose vs Kubernetes?"
```

### With Answer Recording

```yaml
- type: interview_prep
  title: "Interview Questions"
  description: |
    Practice these questions to reinforce your learning.
    Your answers will be saved for review.

  questions:
    - "Explain how Cosign ensures container image integrity."
    - "What's the difference between key-based and keyless signing?"
    - "How would you integrate Cosign into a CI/CD pipeline?"
    - "What are the security risks if you lose your private key?"
    - "How does Cosign integrate with Kubernetes admission controllers?"

  record_answers: true
  export_format: "json"
```

### Security-Focused

```yaml
- type: interview_prep
  title: "Security Deep Dive"
  description: |
    Think about these scenarios from a security perspective.
    Try to cover both the attack surface and mitigations.

  questions:
    - "How would you detect a compromised supply chain?"
    - "What is the role of transparency logs in software security?"
    - "Explain the principle of least privilege in container security."
    - "How do you handle secret rotation in Kubernetes?"
```

## Best Practices

- **5-8 questions** per prep step
- **Start broad, go specific** — "What is X?" → "How would you implement X in scenario Y?"
- **Mix question types** — concept explanations, scenarios, comparisons, troubleshooting
- **Place at the end** — after the user has context from earlier steps
- **Make questions realistic** — use scenarios they'd encounter in actual interviews
