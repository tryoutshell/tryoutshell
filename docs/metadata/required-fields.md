---
sidebar_position: 2
---

# Required Metadata Fields

Every lesson **must** include these fields in the `metadata` section.

## Field Reference

### `id`

**Type:** `string`
**Format:** kebab-case (lowercase with hyphens)
**Example:** `"cosign-sign-verify"`

Unique identifier for the lesson. Used internally and in URLs.

**Rules:**
- Must be unique across all lessons
- Use kebab-case: `my-lesson-name`
- Only alphanumeric characters and hyphens
- No spaces or special characters
````yaml
id: "docker-networking-basics"  ✅
id: "Docker Networking"         ❌ (spaces, capitals)
id: "lesson_1"                  ❌ (underscores)
````

---

### `org`

**Type:** `string`
**Example:** `"chainguard"`

Organization or team that created the lesson.

**Purpose:**
- Groups lessons by creator
- Enables filtering by organization
- Shows authorship in UI
````yaml
org: "chainguard"     ✅
org: "my-team"        ✅
org: "acme-corp"      ✅
````

---

### `title`

**Type:** `string`
**Example:** `"Introduction to Cosign"`

Human-readable title displayed to users.

**Best Practices:**
- Clear and descriptive
- Capitalize major words
- Keep under 60 characters
- Avoid jargon in beginner lessons
````yaml
title: "Container Image Signing with Cosign"  ✅
title: "Cosign Stuff"                         ❌ (vague)
title: "How to Use Cosign to Sign Images and Verify Them in Production"  ❌ (too long)
````

---

### `description`

**Type:** `string`
**Length:** 1-2 sentences
**Example:** `"Learn to sign and verify container images for supply chain security"`

Brief description of what the lesson teaches.

**Best Practices:**
- Summarize the learning outcome
- Keep it concise (under 150 characters)
- Start with an action verb
- Mention key tools/concepts
````yaml
description: "Learn to sign container images with Cosign and verify signatures"  ✅
description: "This lesson teaches you about Cosign..."                           ❌ (wordy)
description: "Cosign"                                                             ❌ (too brief)
````

---

### `difficulty`

**Type:** `enum`
**Values:** `"beginner"`, `"intermediate"`, `"advanced"`
**Example:** `"beginner"`

Skill level required for the lesson.

**Guidelines:**

| Level | Criteria | Example Topics |
|-------|----------|----------------|
| **beginner** | No prior knowledge required | "Introduction to Docker", "First Steps with Git" |
| **intermediate** | Assumes basic familiarity | "Docker Networking", "Git Branching Strategies" |
| **advanced** | Requires deep knowledge | "Kubernetes Operators", "Custom Admission Controllers" |
````yaml
difficulty: "beginner"      ✅
difficulty: "intermediate"  ✅
difficulty: "advanced"      ✅
difficulty: "easy"          ❌ (invalid value)
````

---

### `duration`

**Type:** `string`
**Format:** `"X min"` or `"X hour"`
**Example:** `"20 min"`

Estimated time to complete the lesson.

**Best Practices:**
- Round to nearest 5 minutes
- Test the actual time with real users
- Include time for reading and thinking
- Be realistic, not aspirational
````yaml
duration: "15 min"    ✅
duration: "1 hour"    ✅
duration: "45 min"    ✅
duration: "15"        ❌ (no unit)
duration: "quick"     ❌ (not specific)
````

---

### `tags`

**Type:** `list of strings`
**Example:** `["cosign", "signing", "security"]`

Keywords for search and filtering.

**Best Practices:**
- Use 3-7 tags per lesson
- Include tool names: `"docker"`, `"kubernetes"`
- Include concepts: `"security"`, `"networking"`
- Include categories: `"devops"`, `"cicd"`
- Use lowercase
- Be specific, not generic
````yaml
tags: ["docker", "containers", "networking"]     ✅
tags: ["Docker", "Containers"]                   ❌ (capitals)
tags: ["stuff", "things", "tutorial"]            ❌ (too generic)
tags: []                                          ❌ (empty)
````

**Common Tags:**

| Category | Examples |
|----------|----------|
| Tools | `docker`, `kubernetes`, `terraform`, `cosign` |
| Concepts | `security`, `networking`, `cicd`, `monitoring` |
| Domains | `devops`, `sre`, `platform-engineering` |
| Skill Level | `fundamentals`, `advanced-topics` |

---

## Complete Example
````yaml
metadata:
  id: "kubernetes-rbac-basics"
  org: "platform-team"
  title: "Kubernetes RBAC Fundamentals"
  description: "Learn to secure Kubernetes clusters with role-based access control"
  difficulty: "intermediate"
  duration: "30 min"
  tags: ["kubernetes", "security", "rbac", "access-control"]
````

## Validation

Check your metadata with:
````bash
tryoutshell validate my-lesson.yaml
````

Common errors and fixes:

| Error | Fix |
|-------|-----|
| "Missing required field: id" | Add `id: "lesson-name"` |
| "Invalid difficulty value" | Use `beginner`, `intermediate`, or `advanced` |
| "ID must be kebab-case" | Change `My_Lesson` to `my-lesson` |
| "Tags cannot be empty" | Add at least one tag: `tags: ["docker"]` |

## Next Steps

- [Optional Fields](./optional-fields) - Enhance your lesson metadata
- [Lesson Structure](../getting-started/lesson-structure) - Build your lesson
- [Best Practices](../guides/best-practices) - Write better lessons
