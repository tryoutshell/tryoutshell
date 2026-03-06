---
sidebar_position: 1
---

# Best Practices

Tips for writing effective TryOutShell lessons.

## Lesson Design

### Keep It Focused

Each lesson should teach **one main concept**. Break complex topics into multiple lessons.

- Good: "Docker Basics: Running Containers"
- Too broad: "Complete Docker and Kubernetes Guide"

### Start with "Why"

Before teaching "how", explain why it matters. Users learn better when they understand the motivation.

### Progressive Difficulty

Order content from simple to complex:

1. Introduce the concept (info/slides)
2. Simple verification commands
3. Basic operations
4. More complex tasks
5. Open-ended challenge
6. Quiz to reinforce

### 8-15 Slides or Steps

- Too few: lesson feels shallow
- Too many: user loses focus
- Sweet spot: 8-15 slides/steps

## Content Quality

### Use Real Examples

Show actual commands, real config files, genuine error messages. Avoid "foo/bar" placeholder content.

### Code Blocks Matter

Always specify the language for syntax highlighting:

````markdown
```bash
docker run -it ubuntu:latest /bin/bash
```

```yaml
apiVersion: apps/v1
kind: Deployment
```
````

### Explain the "What Happened"

After a command step, use `post_content` to explain the output. Don't leave users guessing.

### Include Diagrams

ASCII diagrams help visualize workflows:

```
Developer → Build → Sign → Push → Verify → Deploy
```

## Quiz Writing

- **4 options per question** — fewer is too easy, more is overwhelming
- **All options plausible** — no obviously wrong answers
- **Test concepts, not trivia** — "What does this command do?" not "What flag was introduced in v2.3.1?"
- **Always include explanations** — teach even when wrong
- **Place after related content** — don't quiz on what hasn't been taught

## Interactive Lessons

### Hints Are Essential

Every command step should have 2-3 progressive hints:

1. Vague direction: "Use the docker command"
2. More specific: "Try docker ps to list containers"
3. Full answer: "Run: docker ps -a"

### Timeouts

Set reasonable timeouts:
- Quick commands (version check): 10 seconds
- Downloads/installs: 60 seconds
- Complex operations: 30 seconds

### Validation

Prefer flexible validation:
- Use `regex` over `substring` when output varies
- Add `alternative_validations` for different valid outputs
- Use `exit_code` as a fallback

### Allow Skipping

Non-essential steps should have `allow_skip: true`. Don't block the entire lesson because a tool isn't installed.
