---
sidebar_position: 2
---

# Step-by-Step: Creating a Lesson

A walkthrough of building a data-only lesson from scratch.

## Goal

Create a lesson teaching basic Git commands.

## 1. Plan Your Content

Before writing, outline:

- **Topic**: Git basics
- **Audience**: Complete beginners
- **Slides**: ~8 slides
- **Quiz**: 5 questions
- **Time**: 15 minutes

## 2. Create the Directory

```bash
mkdir -p lessons/git/git-basics
```

If a `git` org doesn't exist yet, create `lessons/git/meta.yaml`:

```yaml
id: git
name: "Git"
description: "Version control with Git"
logo: "🔀"
```

## 3. Write lesson.yaml

```yaml
# lessons/git/git-basics/lesson.yaml
id: git-basics
title: "Git Basics: Your First Repository"
description: "Learn git init, add, commit, and log"
author: "Your Name"
tags: ["git", "version-control", "basics"]
difficulty: "beginner"
duration: "15 min"
version: "1.0"

quiz:
  - question: "What command initializes a new Git repository?"
    options: ["git start", "git init", "git new", "git create"]
    answer: 1
    explain: "git init creates a new .git directory in the current folder."

  - question: "What does git add do?"
    options:
      - "Creates a new file"
      - "Stages changes for the next commit"
      - "Pushes to remote"
      - "Deletes a file"
    answer: 1
    explain: "git add stages changes, preparing them to be included in the next commit."

  - question: "What does git log show?"
    options:
      - "Current branch name"
      - "File differences"
      - "Commit history"
      - "Remote URLs"
    answer: 2
    explain: "git log displays the commit history for the current branch."
```

## 4. Write slides.md

```markdown
# Git Basics

Learn the fundamental Git commands to manage your code.

---

## What is Git?

Git is a **distributed version control system** that tracks
changes in your files.

Why use Git?

- Track every change with full history
- Collaborate with others without conflicts
- Revert to any previous state
- Branch and experiment safely

---

## Initializing a Repository

```bash
mkdir my-project
cd my-project
git init
```

This creates a `.git` directory that stores all version history.

---

## The Git Workflow

```
Working Directory → Staging Area → Repository
     (edit)           (add)         (commit)
```

1. **Edit** files in your working directory
2. **Stage** changes with `git add`
3. **Commit** staged changes with `git commit`

---

## Your First Commit

```bash
echo "Hello Git" > README.md
git add README.md
git commit -m "Initial commit"
```

Check the result:

```bash
git log --oneline
```

---

## Summary

| Command | Purpose |
|---------|---------|
| `git init` | Create a new repo |
| `git add <file>` | Stage changes |
| `git commit -m "msg"` | Save a snapshot |
| `git log` | View history |
| `git status` | Check current state |

**Next**: Take the quiz!
```

## 5. Test Locally

```bash
# List lessons — yours should appear
go run . list

# Start the lesson
go run . start git --lesson git-basics

# Run the quiz
go run . quiz git git-basics
```

## 6. Verify Quality

- All slides render correctly
- Quiz has correct answer indices
- Content is accurate
- Navigation works (arrows, gg, G)

## 7. Submit

```bash
git checkout -b add-lesson-git-basics
git add lessons/git/
git commit -m "Add Git basics lesson"
git push -u origin HEAD
# Open a PR
```
