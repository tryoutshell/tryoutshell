---
sidebar_position: 2
---

# Info Steps

Info steps display educational content without requiring user interaction (except pressing Enter to continue).

## Purpose

Use info steps to:
- Explain concepts and theory
- Provide context before hands-on steps
- Show diagrams or examples
- Give tips and warnings

## Basic Structure
````yaml
- type: info
  title: "Step Title"
  content: |
    Your content here with markdown support.
````

## Fields Reference

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | `"info"` | Yes | Identifies this as an info step |
| `title` | string | Yes | Step heading |
| `content` | string (multiline) | Yes | Main text (Markdown supported) |
| `highlights` | list | No | Inline text highlights |
| `code_blocks` | list | No | Code examples with labels |
| `callouts` | list | No | Tips, warnings, info boxes |
| `diagram` | string | No | ASCII art diagram |
| `wait_for_continue` | boolean | No | Pause until Enter (default: true) |

## Simple Example
````yaml
- type: info
  title: "What is Docker?"
  content: |
    **Docker** is a platform for developing and running applications
    in lightweight, portable containers.
     "Author Name"
  version: "1.0"
  updated_at: "2025-01-15"
````

## Required vs Optional

| Category | Fields |
|----------|--------|
| **Required** | `id`, `org`, `title`, `description`, `difficulty`, `duration`, `tags` |
| **Optional** | `prerequisites`, `author`, `version`, `updated_at` |

## Quick Example
````yaml
metadata:
  id: "cosign-basics"
  org: "chainguard"
  title: "Introduction to Cosign"
  description: "Learn to sign and verify container images"
  difficulty: "beginner"
  duration: "20 min"
  tags: ["cosign", "signing", "security"]
````

## Deep Dive

For detailed information about each field:

- [Required Fields](./required-fields) - Essential fields every lesson needs
- [Optional Fields](./optional-fields) - Additional metadata for enhanced discovery

## Validation

TryOutShell validates metadata when you run:
````bash
tryoutshell validate my-lesson.yaml
````

Common validation errors:
- Missing required fields
- Invalid `difficulty` value (must be: beginner, intermediate, advanced)
- Invalid `id` format (must be kebab-case: `my-lesson-name`)
- Empty `tags` list

## Best Practices

### IDs Should Be Descriptive

✅ Good: `docker-networking-basics`
❌ Bad: `lesson1`, `test`, `my-lesson`

### Titles Should Be Clear

✅ Good: "Container Image Signing with Cosign"
❌ Bad: "Learn Stuff", "Tutorial", "Lesson 5"

### Descriptions Should Be Concise

Keep descriptions to 1-2 sentences that clearly explain what the lesson covers.

✅ Good: "Learn to sign container images with Cosign and verify signatures for supply chain security"
❌ Bad: "This lesson will teach you about Cosign which is a tool that..."

### Tags Should Be Specific

Use focused keywords that help users find your lesson.

✅ Good: `["kubernetes", "security", "rbac"]`
❌ Bad: `["stuff", "things", "tutorial"]`

## Next Steps

- [Required Fields Reference](./required-fields)
- [Optional Fields Reference](./optional-fields)
- [Back to Lesson Structure](../getting-started/lesson-structure)
