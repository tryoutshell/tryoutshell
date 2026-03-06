---
sidebar_position: 2
---

# Creating a Lesson

Add a new lesson to TryOutShell with **zero Go code** — just YAML and Markdown.

## Overview

A data-only lesson lives in a directory under `lessons/<org>/<lesson-id>/` and contains:

| File | Required | Purpose |
|------|----------|---------|
| `lesson.yaml` | Yes | Metadata + quiz questions |
| `slides.md` | Yes | Slide content (separated by `---`) |
| `exercises.sh` | No | Runnable shell exercises |

If the org is new, also create `lessons/<org>/meta.yaml`.

## Step 1: Create the Directory

```bash
mkdir -p lessons/my-org/my-lesson
```

## Step 2: Create meta.yaml (if new org)

```yaml
# lessons/my-org/meta.yaml
id: my-org
name: "My Organization"
description: "What this topic covers"
logo: "🚀"
```

## Step 3: Write lesson.yaml

```yaml
# lessons/my-org/my-lesson/lesson.yaml
id: my-lesson
title: "My Lesson Title"
description: "A brief description of what this teaches"
author: "Your Name"
tags: ["tag1", "tag2"]
difficulty: "beginner"   # beginner | intermediate | advanced
duration: "15 min"
version: "1.0"

quiz:
  - question: "What is the answer?"
    options:
      - "Wrong answer A"
      - "Correct answer B"
      - "Wrong answer C"
      - "Wrong answer D"
    answer: 1              # 0-indexed (B = index 1)
    explain: "B is correct because..."

  - question: "Another question?"
    options: ["A", "B", "C", "D"]
    answer: 2
    explain: "C is correct because..."
```

### lesson.yaml Fields

| Field | Required | Type | Description |
|-------|----------|------|-------------|
| `id` | Yes | string | Unique ID (kebab-case) |
| `title` | Yes | string | Human-readable title |
| `description` | Yes | string | One-line summary |
| `author` | No | string | Creator name |
| `tags` | No | list | Searchable keywords |
| `difficulty` | Yes | string | `beginner`, `intermediate`, or `advanced` |
| `duration` | Yes | string | Estimated time, e.g. `"20 min"` |
| `version` | Yes | string | Semantic version |
| `quiz` | No | list | Quiz questions (recommended) |

### Quiz Question Fields

| Field | Required | Type | Description |
|-------|----------|------|-------------|
| `question` | Yes | string | The question text |
| `options` | Yes | list | 3-4 answer options |
| `answer` | Yes | int | Correct option index (0-indexed) |
| `explain` | No | string | Explanation shown after answering |

## Step 4: Write slides.md

Slides are separated by `---` on its own line. Use standard Markdown.

```markdown
# My Lesson Title

Welcome! Here's what you'll learn in this lesson.

- Topic 1
- Topic 2
- Topic 3

---

## Topic 1: The Basics

Explain the first concept here.

### Key Points

- Point A
- Point B
- Point C

```bash
# Code blocks get syntax highlighting
echo "Hello World"
docker run -it ubuntu:latest
```

> 💡 **Tip**: Callouts work great for important info.

---

## Topic 2: Going Deeper

More detailed content with examples.

| Column 1 | Column 2 | Column 3 |
|----------|----------|----------|
| Data | More data | Even more |

---

## Summary

What you learned:

1. Topic 1 key takeaway
2. Topic 2 key takeaway

**Next**: Try the quiz to test your knowledge!
```

### Slide Tips

- Aim for **8-15 slides** per lesson
- Each slide should focus on **one concept**
- Use headers (`##`) to title each slide
- Use code blocks with language tags (`bash`, `yaml`, `python`, etc.)
- Use tables, lists, and blockquotes for variety
- Start with "why" before "how"
- Keep slides readable without horizontal scrolling

## Step 5: Write exercises.sh (Optional)

```bash
#!/bin/bash
# EXERCISE: Check your environment
echo "Hello from TryOutShell exercises!"

# EXERCISE: List files in the current directory
ls -la

# EXERCISE: Create a test file
echo "test content" > myfile.txt && cat myfile.txt
```

## Step 6: Test Locally

```bash
# Verify lesson appears in the list
go run . list

# Start the lesson
go run . start my-org --lesson my-lesson

# Test the quiz
go run . quiz my-org my-lesson
```

## Using the Template

Copy the built-in template as a starting point:

```bash
cp -r lessons/_template/my-lesson lessons/my-org/my-lesson
```

Edit the copied files with your content.

## Using AI to Generate Lessons

Copy `prompt-create-lesson.md` into any AI assistant along with a blog URL to auto-generate lesson content. See the [AI Lesson Generator](../guides/ai-generator) guide.

## Submitting Your Lesson

1. Fork the repository
2. Create a branch: `git checkout -b add-lesson-my-topic`
3. Add your lesson files
4. Test locally
5. Open a pull request

**PR Checklist:**

- [ ] `lesson.yaml` has all required fields
- [ ] `slides.md` has at least 8 slides with real content
- [ ] Quiz has 3-5 questions with correct `answer` indices
- [ ] All quiz options are plausible
- [ ] Content is accurate and educational
- [ ] Tested with `go run . start` and `go run . quiz`
