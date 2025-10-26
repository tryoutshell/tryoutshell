---
sidebar_position: 1
---

# Introduction Section

The `introduction` section is **optional** but recommended. It sets context and expectations before users begin the lesson steps.

## Purpose

The introduction:
- Explains what the lesson covers
- Sets learning objectives
- Provides context and motivation
- Shows estimated time and prerequisites

## Structure
````yaml
introduction:
  title: "Introduction Title"
  content: |
    Multi-line content with markdown support.

    Can include lists, code, and formatting.
````

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `title` | string | Yes | Section header |
| `content` | string (multiline) | Yes | Introduction text (Markdown supported) |

## Basic Example
````yaml
introduction:
  title: "What You'll Learn"
  content: |
    In this lesson, you will:
    - Understand what Cosign is
    - Install and verify Cosign
    - Sign your first container image
    - Verify image signatures
````

## Markdown Support

The `content` field supports full Markdown:

### Formatting
````yaml
content: |
  **Bold text** for emphasis
  *Italic text* for subtle emphasis
  `inline code` for commands or variables
````

### Lists
````yaml
content: |
  You will learn:
  - Bullet point one
  - Bullet point two
  - Bullet point three

  Steps to follow:
  1. First step
  2. Second step
  3. Third step
````

### Code Blocks
````yaml
content: |
  Example command:
```bash
  docker run hello-world
```
````

### Headings
````yaml
content: |
  ### What is Docker?

  Docker is a containerization platform.

  ### Why Use Docker?

  Benefits include isolation, portability, and consistency.
````

### Blockquotes
````yaml
content: |
  > 💡 **Tip:** Have Docker running before you start!

  > ⚠️ **Warning:** This lesson requires root access.
````

### Links
````yaml
content: |
  Learn more in the [official documentation](https://docs.docker.com).

  See also: [Docker Hub](https://hub.docker.com)
````

## Complete Examples

### Example 1: Simple Welcome
````yaml
introduction:
  title: "Welcome to Docker Basics"
  content: |
    Learn the essential Docker commands you'll use every day.

    This lesson takes approximately **15 minutes** to complete.
````

### Example 2: Detailed Overview
````yaml
introduction:
  title: "Introduction to Cosign"
  content: |
    **Cosign** is a tool for signing and verifying container images,
    developed as part of the Sigstore project.

    ### What You'll Learn

    In this lesson, you will:
    - Understand the importance of image signing
    - Install and configure Cosign
    - Generate key pairs for signing
    - Sign and verify container images
    - Explore keyless signing with OIDC

    ### Prerequisites

    - Docker installed and running
    - Basic understanding of containers
    - Terminal/command-line access

    ### Estimated Time

    ⏱️ **20-25 minutes**

    > 💡 **Tip:** Open the [Cosign documentation](https://docs.sigstore.dev)
    > in another tab for reference.
````

### Example 3: With Motivation
````yaml
introduction:
  title: "Why Learn Kubernetes?"
  content: |
    Kubernetes has become the standard for container orchestration,
    powering applications at Google, Netflix, Spotify, and thousands
    of other companies.

    ### What Makes Kubernetes Special?

    - **Scalability**: Run 1 or 10,000 containers
    - **Self-healing**: Automatic recovery from failures
    - **Declarative**: Describe what you want, not how to do it

    ### This Lesson Covers

    1. Core Kubernetes concepts (Pods, Services, Deployments)
    2. Your first deployment
    3. Scaling and updating applications
    4. Troubleshooting common issues

    **Prerequisites:** Docker knowledge recommended but not required.

    Let's get started! 🚀
````

### Example 4: Security-Focused
````yaml
introduction:
  title: "Container Security Fundamentals"
  content: |
    > ⚠️ **Security First:** This lesson covers critical security practices
    > for container environments.

    ### The Challenge

    Container images can contain vulnerabilities, malware, or be tampered
    with during distribution. How do you know an image is safe?

    ### The Solution

    **Image signing and verification** ensures:
    - ✅ **Authenticity**: The image comes from a trusted source
    - ✅ **Integrity**: The image hasn't been modified
    - ✅ **Non-repudiation**: Cryptographic proof of origin

    ### Tools You'll Use

    - **Cosign**: Sign and verify images
    - **Docker**: Pull and run containers
    - **OpenSSL**: Inspect signatures (optional)

    ---

    **Time Required:** 30 minutes
    **Difficulty:** Intermediate
    **Prerequisites:** Basic Docker knowledge
````

## Best Practices

### Keep It Concise

Introduction should be scannable in 30-60 seconds.

✅ **Good:**
````yaml
content: |
  Learn to deploy applications with Kubernetes.

  You'll create deployments, expose services, and scale apps.
````

❌ **Too long:**
````yaml
content: |
  Kubernetes is a container orchestration system originally developed
  by Google based on their internal Borg system, which has been used
  to run Google's services for over a decade... (continues for 5 paragraphs)
````

### Start with the "Why"

Explain motivation before diving into details.

✅ **Good:**
````yaml
content: |
  Container images can be tampered with. Cosign solves this with
  cryptographic signatures.

  Learn to sign and verify images in this lesson.
````

❌ **Too technical immediately:**
````yaml
content: |
  Cosign uses ECDSA P-256 signatures with SHA-256 hashes stored
  as OCI artifacts... (loses readers immediately)
````

### Use Visual Hierarchy

Break content into sections with headers.

✅ **Good:**
````yaml
content: |
  ### What You'll Learn
  - Docker basics
  - Container networking

  ### Prerequisites
  - Docker installed
````

❌ **Wall of text:**
````yaml
content: |
  In this lesson you'll learn Docker basics and container networking
  but you need to have Docker installed first and also basic command
  line knowledge...
````

### Set Expectations

Tell users what they'll gain and how long it takes.

✅ **Good:**
````yaml
content: |
  **Duration:** 20 minutes
  **Outcome:** Deploy your first Kubernetes app
````

### Use Emojis Sparingly

1-2 emojis add personality. More is distracting.

✅ **Good:** `🚀 Let's get started!`
❌ **Too much:** `🎉🚀💻🔥 Let's get started! 🎊🎈✨`

## Common Patterns

### Pattern 1: Learning Objectives
````yaml
introduction:
  title: "What You'll Learn"
  content: |
    By the end of this lesson, you will be able to:
    1. Install and configure [Tool]
    2. Perform [Action]
    3. Troubleshoot common issues
````

### Pattern 2: Problem/Solution
````yaml
introduction:
  title: "The Challenge"
  content: |
    **Problem:** [Describe the pain point]

    **Solution:** [Introduce the tool/concept]

    **This Lesson:** [What they'll learn]
````

### Pattern 3: Quick Start
````yaml
introduction:
  title: "Quick Start"
  content: |
    Get hands-on with [Tool] in 15 minutes.

    No prior experience needed!
````

## When to Skip Introduction

You can omit the introduction if:
- The lesson is extremely simple (< 5 minutes)
- The title and description are self-explanatory
- You want to jump straight into action

Example where introduction isn't needed:
````yaml
metadata:
  title: "Check Docker Version"
  description: "Verify Docker is installed on your system"
  duration: "2 min"

# No introduction needed - title says it all

steps:
  - type: command
    prompt: "Check Docker version"
    example: "docker --version"
    # ...
````

## Next Steps

- [Conclusion Section](./conclusion) - End your lesson effectively
- [Step Types](../step-types/) - Build interactive lesson content
- [Markdown Support](../guides/markdown-support) - Advanced formatting
