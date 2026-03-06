---
sidebar_position: 1
---

# Metadata Fields

Lesson metadata identifies the lesson and controls how it appears in listings.

## Data-Only Lessons (lesson.yaml)

These fields go at the top level of `lesson.yaml`:

| Field | Required | Type | Description | Example |
|-------|----------|------|-------------|---------|
| `id` | Yes | string | Unique ID (kebab-case) | `"docker-101"` |
| `title` | Yes | string | Human-readable title | `"Docker 101"` |
| `description` | Yes | string | Brief summary (1-2 sentences) | `"Learn Docker basics"` |
| `difficulty` | Yes | string | `beginner`, `intermediate`, `advanced` | `"beginner"` |
| `duration` | Yes | string | Estimated time | `"20 min"` |
| `version` | Yes | string | Lesson version | `"1.0"` |
| `author` | No | string | Creator name | `"TryOutShell"` |
| `tags` | No | list | Searchable keywords | `["docker", "containers"]` |
| `quiz` | No | list | Quiz questions | See [Quiz Steps](../step-types/quiz-steps) |

### Example

```yaml
id: docker-101
title: "Docker 101: Container Fundamentals"
description: "Learn Docker from scratch — images, containers, Dockerfiles"
author: "TryOutShell"
tags: ["docker", "containers", "devops"]
difficulty: "beginner"
duration: "25 min"
version: "1.0"
```

## Interactive Lessons (metadata section)

These fields go inside the `metadata:` section of interactive YAML lessons:

| Field | Required | Type | Description |
|-------|----------|------|-------------|
| `id` | Yes | string | Unique ID (kebab-case) |
| `org` | Yes | string | Organization/topic ID |
| `title` | Yes | string | Human-readable title |
| `description` | Yes | string | Brief summary |
| `difficulty` | Yes | string | `beginner`, `intermediate`, `advanced` |
| `duration` | Yes | string | Estimated time |
| `tags` | Yes | list | Searchable keywords |
| `prerequisites` | No | list | Required tools or knowledge |
| `author` | No | string | Creator name |
| `version` | No | string | Lesson version |
| `resources` | No | list | External links |

### Example

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
    - "Basic CLI experience"
  tags: ["cosign", "signing", "supply-chain"]
  author: "TryOutShell"
  version: "1.0"
  resources:
    - title: "Cosign Docs"
      url: "https://docs.sigstore.dev"
      type: "docs"
    - title: "Cosign GitHub"
      url: "https://github.com/sigstore/cosign"
      type: "github"
```

## Organization meta.yaml

Each org directory can have a `meta.yaml`:

```yaml
id: docker
name: "Docker"
description: "Container runtime and tooling"
logo: "🐳"
```

| Field | Required | Type | Description |
|-------|----------|------|-------------|
| `id` | Yes | string | Org identifier |
| `name` | Yes | string | Display name |
| `description` | No | string | Brief description |
| `logo` | No | string | Emoji or text icon |

## Best Practices

### IDs

- Use kebab-case: `docker-101`, `cosign-sign-verify`
- Be descriptive: `docker-networking-basics` not `lesson1`

### Titles

- Be specific: "Container Image Signing with Cosign"
- Not vague: "Learn Stuff", "Tutorial 5"

### Descriptions

- 1-2 sentences maximum
- Explain what the user will learn, not the tool's marketing pitch

### Tags

- Use specific keywords users would search for
- Include both the tool name and the concept
- Good: `["kubernetes", "rbac", "security"]`
- Bad: `["stuff", "things", "tutorial"]`

### Difficulty

| Level | Target Audience |
|-------|----------------|
| `beginner` | No prior experience with this tool |
| `intermediate` | Familiar with basics, learning advanced features |
| `advanced` | Production use cases, deep internals |
