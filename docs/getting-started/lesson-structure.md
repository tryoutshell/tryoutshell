---
sidebar_position: 3
---

# Lesson Structure

Every TryOutShell lesson follows a consistent YAML structure with four main sections.

## Overview
```yaml
metadata:        # Required - Lesson identification and metadata
introduction:    # Optional - Shown once at the start
steps:           # Required - Main lesson content
conclusion:      # Optional - Shown after completion
```

## The Four Sections

### 1. Metadata (Required)

Defines who created the lesson, what it teaches, and how to find it.
```yaml
metadata:
  id: "my-lesson"
  org: "my-org"
  title: "My Lesson"
  description: "What this lesson teaches"
  difficulty: "beginner"
  duration: "15 min"
  tags: ["docker", "containers"]
```

[→ See all metadata fields](../metadata/)

### 2. Introduction (Optional)

Sets context before the lesson begins. Shows once when the lesson starts.
```yaml
introduction:
  title: "What You'll Learn"
  content: |
    In this lesson, you will:
    - Install and verify a tool
    - Run basic commands
    - Complete a hands-on challenge
```

[→ Learn more about introductions](../sections/introduction)

### 3. Steps (Required)

The core of the lesson. A sequence of interactive steps the user completes.
```yaml
steps:
  - type: info
    title: "Welcome"
    content: "Let's get started!"

  - type: command
    prompt: "Run a command"
    example: "docker --version"
    validation:
      type: "substring"
      contains: "Docker version"
    success_msg: "✅ Great!"
    fail_msg: "❌ Try again"
```

[→ Explore all step types](../step-types/)

### 4. Conclusion (Optional)

Summarizes what was learned and suggests next steps.
```yaml
conclusion:
  title: "Congratulations!"
  content: |
    You've completed the lesson!

    Next steps:
    - Try the advanced lesson
    - Explore the documentation
  badges:
    - id: "docker-basics"
      name: "Docker Basics"
      icon: "🐳"
```

[→ Learn about conclusions](../sections/conclusion)

## Complete Example

Here's a complete lesson structure:
```yaml
metadata:
  id: "docker-intro"
  org: "tutorial"
  title: "Docker Introduction"
  description: "Learn basic Docker commands"
  difficulty: "beginner"
  duration: "15 min"
  tags: ["docker", "containers"]

introduction:
  title: "Welcome to Docker"
  content: |
    Docker lets you package applications in containers.

    In this lesson, you'll learn:
    - How to check Docker is installed
    - How to run your first container
    - Basic Docker commands

steps:
  - type: info
    title: "What is Docker?"
    content: |
      Docker is a platform for developing and running applications
      in lightweight, portable containers.

  - type: command
    prompt: "Verify Docker is installed"
    example: "docker --version"
    validation:
      type: "substring"
      contains: "Docker version"
    success_msg: "✅ Docker is installed!"
    fail_msg: "❌ Docker not found. Please install it first."

  - type: command
    prompt: "Run your first container"
    example: "docker run hello-world"
    validation:
      type: "substring"
      contains: "Hello from Docker"
    success_msg: "✅ Your first container ran successfully!"
    fail_msg: "❌ Container failed to run"

  - type: quiz
    title: "Quick Check"
    questions:
      - id: "q1"
        question: "What command checks Docker version?"
        type: "multiple_choice"
        options:
          - "docker version"
          - "docker --version"
          - "Both are correct"
        answer: 2
        explanation: "Both commands work to check Docker version!"

conclusion:
  title: "Well Done!"
  content: |
    🎉 You've learned Docker basics!

    **What's next?**
    - Explore the Docker CLI lesson
    - Learn about Dockerfiles
    - Build your own container
  badges:
    - id: "docker-beginner"
      name: "Docker Beginner"
      icon: "🐳"
```

## Validation

Before running your lesson, validate it:
```bash
tryoutshell validate my-lesson.yaml
```

This checks:
- YAML syntax
- Required fields are present
- Field types are correct
- Step types are valid

## Best Practices

### Keep It Focused

Each lesson should teach **one main concept**. Break complex topics into multiple lessons.

✅ Good: "Docker Basics - Running Containers"
❌ Too Broad: "Complete Docker and Kubernetes Guide"

### Start Simple

Begin with info steps to build context, then move to commands.
```yaml
steps:
  - type: info        # Explain first
    title: "What is X?"
    content: "..."

  - type: command     # Then practice
    prompt: "Try it yourself"
    example: "..."
```

### Progressive Difficulty

Order steps from easy to challenging:

1. Simple verification commands
2. Basic operations
3. More complex tasks
4. Final challenge

### Test Everything

Always test your lesson end-to-end:
```bash
tryoutshell start my-lesson
```

Verify:
- All commands work as expected
- Validations catch correct/incorrect outputs
- Success/failure messages are helpful
- The flow makes sense

## Next Steps

Now that you understand the structure:

- 📝 [Learn about Metadata](../metadata/) - All metadata fields explained
- 🎯 [Explore Step Types](../step-types/) - Deep dive into each step type
- 🔧 [Validation Methods](../guides/validation-types) - How to validate commands
- 💡 [Best Practices](../guides/best-practices) - Tips for great lessons

---

Ready to create your own lesson? Start with a [minimal example](./minimal-example)!
