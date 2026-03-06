# Contributing to TryOutShell

Thank you for your interest in contributing! TryOutShell is designed so that **anyone can add a lesson without writing Go code** — just YAML and Markdown.

## Adding a New Lesson

### 1. Create the Directory Structure

```
lessons/<your-org>/<your-lesson>/
  lesson.yaml     ← metadata + quiz questions
  slides.md       ← slide content (--- separated)
  exercises.sh    ← optional: runnable exercises
```

If your organization doesn't exist yet, also create `lessons/<your-org>/meta.yaml`:

```yaml
id: my-org
name: "My Organization"
description: "What this org/topic is about"
logo: "🚀"
```

### 2. Write lesson.yaml

```yaml
id: my-lesson
title: "My Lesson Title"
description: "Brief description of what this teaches"
author: "Your Name"
tags: ["tag1", "tag2", "tag3"]
difficulty: "beginner"   # beginner | intermediate | advanced
duration: "15 min"
version: "1.0"

quiz:
  - question: "What is X?"
    options:
      - "Answer A"
      - "Answer B (correct)"
      - "Answer C"
      - "Answer D"
    answer: 1            # 0-indexed
    explain: "Explanation of why B is correct."
```

**Fields:**

| Field | Required | Description |
|-------|----------|-------------|
| `id` | Yes | Unique identifier (lowercase, hyphens) |
| `title` | Yes | Human-readable title |
| `description` | Yes | One-line summary |
| `author` | No | Your name |
| `tags` | No | Searchable tags |
| `difficulty` | Yes | `beginner`, `intermediate`, or `advanced` |
| `duration` | Yes | Estimated time, e.g. `"20 min"` |
| `version` | Yes | Semantic version |
| `quiz` | No | Array of quiz questions (recommended) |

### 3. Write slides.md

Slides are separated by `---` (horizontal rule). Use standard Markdown:

```markdown
# Slide Title

Introduction to the topic.

---

## Key Concepts

- Point 1
- Point 2
- Point 3

```bash
# Code examples are syntax-highlighted
docker run -it ubuntu:latest /bin/bash
```

> 💡 **Tip**: Use callouts for important information.

---

## Summary

What you learned today.
```

**Tips for great slides:**

- Aim for **8-15 slides** per lesson
- Each slide should focus on **one concept**
- Use code blocks with language tags for syntax highlighting
- Use tables, lists, and blockquotes for variety
- Keep individual slides readable without scrolling
- Start with "why" before "how"

### 4. Write Quiz Questions

- Include **3-5 questions** per lesson
- Cover key concepts from the slides
- Use 3-4 options per question
- Always provide an `explain` field
- `answer` is **0-indexed** (first option = 0)

### 5. Write exercises.sh (Optional)

```bash
#!/bin/bash
# EXERCISE: Description of what to do
# command to run

# EXERCISE: Another exercise
# another command
```

### 6. Test Locally

```bash
# Run from the repo root
go run . list                           # verify your lesson appears
go run . start <your-org> -l <lesson>   # test the full lesson
go run . quiz <your-org> <lesson>       # test quiz questions
```

### 7. Submit a PR

1. Fork the repository
2. Create a feature branch: `git checkout -b add-lesson-my-topic`
3. Add your lesson files
4. Test locally
5. Open a pull request

**PR Checklist:**

- [ ] `lesson.yaml` has all required fields
- [ ] `slides.md` has at least 8 slides
- [ ] Quiz questions have correct `answer` indices
- [ ] All quiz options are plausible (no obvious wrong answers)
- [ ] Content is accurate and educational
- [ ] No copyrighted content copied verbatim
- [ ] Tested with `go run . start` and `go run . quiz`

## Contributing Code

For Go code changes:

1. Ensure `go build ./...` succeeds
2. Ensure `go vet ./...` has zero warnings
3. Handle all errors — no `_` discards on user-facing errors
4. Follow existing code style and patterns
5. Test with `go test ./...`

## Template

Copy `lessons/_template/` as a starting point for new lessons.

## Questions?

Open an issue or start a discussion on GitHub.
