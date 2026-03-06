# 🤖 TryOutShell Lesson Generator Prompt

Copy this entire prompt and paste it into any AI assistant (Claude, ChatGPT, etc.) along with your blog/tutorial link to generate a complete lesson.

---

## PROMPT START

I need you to create a data-only lesson for TryOutShell, an interactive terminal-based learning platform. Lessons are pure YAML + Markdown — no Go code needed.

### 📚 SOURCE MATERIAL

**Blog/Tutorial URL:** [PASTE YOUR URL HERE]

**Additional Context:** [Add any specific requirements, focus areas, or notes]

---

### 📋 LESSON FORMAT

A TryOutShell lesson consists of 3 files in a directory:

```
lessons/<org-id>/<lesson-id>/
  lesson.yaml      ← metadata + quiz questions
  slides.md        ← lesson content (--- separated slides)
  exercises.sh     ← optional: runnable shell exercises
```

If the org doesn't exist yet, also create `lessons/<org-id>/meta.yaml`:

```yaml
id: my-org
name: "Organization Name"
description: "Brief description"
logo: "🚀"
```

#### lesson.yaml Schema

```yaml
id: "lesson-id"                    # kebab-case, unique identifier
title: "Human Readable Title"
description: "Brief 1-2 sentence description"
author: "Your Name"
tags: ["tag1", "tag2", "tag3"]     # searchable tags
difficulty: "beginner"             # beginner | intermediate | advanced
duration: "20 min"                 # estimated time
version: "1.0"

quiz:
  - question: "What is X?"
    options:
      - "Wrong answer A"
      - "Correct answer B"
      - "Wrong answer C"
      - "Wrong answer D"
    answer: 1                      # 0-indexed (B is index 1)
    explain: "Explanation of why B is correct and why other options are wrong."

  - question: "Another question?"
    options: ["A", "B", "C", "D"]
    answer: 2
    explain: "Explanation..."
```

#### slides.md Format

Slides are separated by `---` (horizontal rule on its own line). Use standard Markdown:

```markdown
# Lesson Title

Welcome and overview of what you'll learn.

---

## Concept 1

Detailed explanation with:

- Bullet points
- **Bold** and *italic* text
- `inline code`

```bash
# Code blocks with syntax highlighting
docker run -it ubuntu:latest /bin/bash
```

> 💡 **Tip**: Use blockquote callouts for important info.

---

## Concept 2

More content...

| Column 1 | Column 2 |
|----------|----------|
| Data     | Data     |

---

## Summary

Recap of what was covered.
```

#### exercises.sh (Optional)

```bash
#!/bin/bash
# EXERCISE: Description of exercise 1
command_to_run

# EXERCISE: Description of exercise 2
another_command
```

### 📝 REQUIREMENTS

**Slides:**
- Minimum **8 slides**, maximum 15
- Each slide should focus on **one concept**
- Mix of text, code blocks, tables, and diagrams
- Start with "why" before "how"
- Include real, working command examples
- Use proper markdown syntax highlighting (```bash, ```yaml, ```python, etc.)

**Quiz:**
- **5 questions** minimum
- Cover key concepts from the slides
- 4 options per question
- All options should be plausible (no obviously wrong answers)
- Always include `explain` field
- `answer` is 0-indexed

**Content Quality:**
- Must be accurate and up-to-date
- Should be educational, not just descriptive
- Include practical examples
- Mention common mistakes/pitfalls
- Provide references where appropriate

### 📦 OUTPUT FORMAT

Generate the complete contents of all files:

1. **lesson.yaml** — complete metadata + quiz
2. **slides.md** — complete slide content
3. **exercises.sh** — if applicable to the topic

### 🎯 EXAMPLE

Here's a complete working example for a Docker lesson:

**lesson.yaml:**
```yaml
id: docker-101
title: "Docker 101: Container Fundamentals"
description: "Learn Docker from scratch - images, containers, Dockerfiles, volumes, networking"
author: "TryOutShell"
tags: ["docker", "containers", "devops"]
difficulty: "beginner"
duration: "25 min"
version: "1.0"

quiz:
  - question: "What is the difference between a Docker image and a container?"
    options:
      - "They are the same thing"
      - "An image is a template; a container is a running instance of an image"
      - "A container is a template; an image is a running instance"
      - "Images run on the host; containers run in the cloud"
    answer: 1
    explain: "A Docker image is a read-only template. A container is a runnable instance of an image."
```

**slides.md:**
```markdown
# Docker 101: Container Fundamentals

Docker solves the classic "works on my machine" problem by packaging
applications with all their dependencies into standardized containers.

---

## Images vs Containers

A **Docker image** is a read-only template — like a class definition.
A **container** is a running instance — like an object created from that class.

```bash
# List images
docker images

# List running containers
docker ps
```

---

## Summary

What you learned:
- Docker solves environment consistency
- Images are templates, containers are instances
```

---

## PROMPT END

Now paste the URL of the blog/tutorial you want to convert into a TryOutShell lesson above, and the AI will generate all the files for you.
