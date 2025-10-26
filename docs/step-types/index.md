---
sidebar_position: 1
---

# Step Types Overview

Steps are the core of TryOutShell lessons. Each step has a `type` that determines its behavior and interaction model.

## Available Step Types

| Type | Purpose | Interaction |
|------|---------|-------------|
| [`info`](./info-steps) | Display educational content | Read and continue |
| [`command`](./command-steps) | Execute and validate commands | Type and run commands |
| [`quiz`](./quiz-steps) | Test knowledge | Answer multiple-choice questions |
| [`challenge`](./challenge-steps) | Open-ended tasks | Complete complex objectives |
| [`interview_prep`](./interview-prep-steps) | Practice interview questions | Write and reflect |

## Quick Comparison

### Info Steps
**Best for:** Explanations, context, theory
````yaml
- type: info
  title: "What is Docker?"
  content: "Docker is a containerization platform..."
````

[→ Learn more about info steps](./info-steps)

---

### Command Steps
**Best for:** Hands-on practice, executing commands
````yaml
- type: command
  prompt: "Check Docker version"
  example: "docker --version"
  validation:
    type: "substring"
    contains: "Docker version"
  success_msg: "✅ Docker is installed!"
  fail_msg: "❌ Docker not found"
````

[→ Learn more about command steps](./command-steps)

---

### Quiz Steps
**Best for:** Knowledge checks, reinforcement
````yaml
- type: quiz
  title: "Quick Check"
  questions:
    - id: "q1"
      question: "What command lists Docker images?"
      type: "multiple_choice"
      options:
        - "docker ps"
        - "docker images"
        - "docker list"
      answer: 1
      explanation: "`docker images` lists all local images"
````

[→ Learn more about quiz steps](./quiz-steps)

---

### Challenge Steps
**Best for:** Applying knowledge, complex tasks
````yaml
- type: challenge
  title: "Deploy Your Own App"
  description: |
    Deploy a web application using what you've learned.

    Requirements:
    1. Create a Dockerfile
    2. Build the image
    3. Run the container on port 8080
  verification:
    type: "custom"
    checks:
      - type: "command_succeeds"
        command: "curl -f http://localhost:8080"
  success_msg: "🎉 Application deployed!"
````

[→ Learn more about challenge steps](./challenge-steps)

---

### Interview Prep Steps
**Best for:** Reflection, deeper understanding
````yaml
- type: interview_prep
  title: "Think Deeper"
  description: "Practice explaining concepts"
  questions:
    - "Explain the difference between Docker images and containers"
    - "How would you secure a production container?"
    - "When would you use Docker Compose vs Kubernetes?"
  record_answers: true
````

[→ Learn more about interview prep steps](./interview-prep-steps)

---

## Choosing the Right Step Type

### Decision Tree
