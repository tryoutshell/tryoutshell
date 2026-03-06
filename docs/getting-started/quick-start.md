---
sidebar_position: 5
---

# Quick Reference

A condensed reference for all lesson formats, step types, and validation methods.

## Data-Only Lesson Format

```
lessons/<org>/<lesson-id>/
  lesson.yaml       ← metadata + quiz
  slides.md         ← slides separated by ---
  exercises.sh      ← optional exercises
```

### lesson.yaml

```yaml
id: "my-lesson"
title: "Lesson Title"
description: "Brief description"
author: "Name"
tags: ["tag1", "tag2"]
difficulty: "beginner"
duration: "15 min"
version: "1.0"

quiz:
  - question: "Question text?"
    options: ["A", "B", "C", "D"]
    answer: 1                      # 0-indexed
    explain: "Why B is correct."
```

### slides.md

```markdown
# Title Slide

Content...

---

## Slide 2

More content...

---

## Final Slide

Summary.
```

---

## Interactive Lesson Format

Single YAML file at `lessons/<org>/<lesson-id>.yaml`:

```yaml
metadata:
  id: "lesson-id"
  org: "org-id"
  title: "Title"
  description: "Description"
  difficulty: "beginner"
  duration: "15 min"
  tags: ["tag"]

introduction:
  title: "Header"
  content: |
    Markdown content.

steps:
  - type: info|command|quiz|challenge|interview_prep
    # fields vary by type

conclusion:
  title: "Done!"
  content: "Summary"
  badges:
    - id: "badge-id"
      name: "Badge Name"
      icon: "🎯"
```

---

## Step Types Quick Reference

### info

```yaml
- type: info
  title: "Title"
  content: |
    Markdown content.
  highlights:
    - text: "term"
      style: "code"       # code | bold | highlight
  code_blocks:
    - label: "Example"
      code: "echo hello"
      language: "bash"
  callouts:
    - type: "tip"          # tip | warning | danger | info
      text: "Helpful note"
  diagram: |
    ASCII art diagram
  wait_for_continue: true
```

### command

```yaml
- type: command
  id: "step-id"
  prompt: "What to do"
  instruction: "Detailed instruction"
  example: "echo hello"
  pre_content: "Explanation before"
  post_content: "Explanation after"
  accepted_commands:
    - "echo hello"
    - "echo 'hello'"
  validation:
    type: "substring"
    contains: "hello"
  alternative_validations:
    - type: "exit_code"
      expected: 0
  success_msg: "✅ Done!"
  fail_msg: "❌ Try again"
  hints:
    - level: 1
      text: "First hint"
    - level: 2
      text: "More specific hint"
  allow_skip: true
  timeout: 30
```

### quiz

```yaml
- type: quiz
  title: "Knowledge Check"
  questions:
    - id: "q1"
      question: "Question?"
      type: "multiple_choice"
      options: ["A", "B", "C", "D"]
      answer: 1
      explanation: "Why B is correct."
```

### challenge

```yaml
- type: challenge
  title: "Challenge Name"
  description: |
    Task instructions.
  verification:
    type: "custom"
    checks:
      - type: "file_exists"
        path: "output.txt"
      - type: "file_contains"
        path: "output.txt"
        pattern: "success"
  hints:
    - level: 1
      text: "Hint text"
  success_msg: "🎉 Done!"
  allow_skip: true
```

### interview_prep

```yaml
- type: interview_prep
  title: "Interview Questions"
  description: "Practice these questions."
  questions:
    - "Explain concept X."
    - "How would you implement Y?"
  record_answers: true
  export_format: "json"
```

---

## Validation Types

| Type | Key Fields | Description |
|------|-----------|-------------|
| `substring` | `contains`, `case_insensitive` | Output contains string |
| `regex` | `pattern`, `case_insensitive` | Output matches regex |
| `exit_code` | `expected` | Command exits with code |
| `file_exists` | `files` | Files exist on disk |
| `file_contains` | `path`, `pattern` | File content matches |
| `output_contains` | `patterns`, `any_match`, `all_match` | Output matches patterns |

### Examples

```yaml
# Substring
validation:
  type: "substring"
  contains: "Docker version"

# Regex
validation:
  type: "regex"
  pattern: "v\\d+\\.\\d+\\.\\d+"

# Exit code
validation:
  type: "exit_code"
  expected: 0

# File exists
validation:
  type: "file_exists"
  files: ["output.txt", "config.yaml"]

# Output contains (any)
validation:
  type: "output_contains"
  patterns: ["success", "ok"]
  any_match: true
```

---

## Organization meta.yaml

```yaml
id: my-org
name: "Organization Name"
description: "Brief description"
logo: "🚀"
```
