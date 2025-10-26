---
sidebar_position: 6
---

# Interview Prep Steps

Interview prep steps help users practice explaining concepts in their own words.

## Purpose

Use interview prep steps to:
- Encourage deeper thinking
- Practice articulating knowledge
- Prepare for technical interviews
- Reinforce understanding
- Identify knowledge gaps

## Basic Structure
`````yaml
- type: interview_prep
  title: "Think Deeper"
  description: "Practice explaining these concepts"
  questions:
    - "Question 1?"
    - "Question 2?"
    - "Question 3?"
  record_answers: true
  export_format: "json"
`````

## Fields Reference

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | `"interview_prep"` | Yes | Identifies this type |
| `title` | string | Yes | Section header |
| `description` | string | No | Instructions |
| `questions` | list | Yes | List of question strings |
| `record_answers` | boolean | No | Save answers? (default: false) |
| `export_format` | enum | No | `"json"` or `"text"` (default: `"json"`) |

## Simple Example
`````yaml
- type: interview_prep
  title: "Explain What You Learned"
  questions:
    - "What is Docker and why is it useful?"
    - "Explain the difference between an image and a container"
    - "When would you use Docker Compose?"
`````

## Complete Examples

### Example 1: Docker Concepts
`````yaml
- type: interview_prep
  title: "Interview Questions"
  description: |
    Practice explaining these concepts as if you're in a technical interview.

    Take your time and think through your answers. There's no "perfect" answer,
    but try to be clear and thorough.

  questions:
    - "Explain what Docker is to someone who has never heard of it"
    - "What's the difference between a Docker image and a container?"
    - "How do you troubleshoot a container that won't start?"
    - "When would you use Docker Compose vs Kubernetes?"
    - "What are some security best practices for containers?"

  record_answers: true
  export_format: "json"
`````

### Example 2: Security Focus
`````yaml
- type: interview_prep
  title: "Security Interview Prep"
  description: |
    Security is critical in DevOps. Practice answering these questions:

  questions:
    - "Explain how image signing improves supply chain security"
    - "What's the difference between key-based and keyless signing?"
    - "How would you integrate Cosign into a CI/CD pipeline?"
    - "What are the security risks if a private key is compromised?"
    - "How do you verify the authenticity of a third-party container image?"

  record_answers: true
`````

### Example 3: Troubleshooting Scenarios
`````yaml
- type: interview_prep
  title: "Troubleshooting Scenarios"
  description: |
    Imagine you're in an interview and asked to solve these problems.
    Explain your approach step-by-step.

  questions:
    - "A container keeps restarting. How do you diagnose the issue?"
    - "Users can't access your deployed application. What do you check?"
    - "A Docker build is failing. Walk through your debugging process"
    - "An image verification fails. What could be the causes?"
    - "Your Kubernetes pod is in 'CrashLoopBackOff'. What steps do you take?"

  record_answers: true
  export_format: "text"
`````

### Example 4: Architecture Decisions
`````yaml
- type: interview_prep
  title: "Architecture & Design"
  description: |
    Practice explaining architectural decisions and trade-offs.

  questions:
    - "When would you choose containers over virtual machines?"
    - "How do you decide between a monolith and microservices?"
    - "Explain the trade-offs between Docker Swarm and Kubernetes"
    - "How would you design a secure container build pipeline?"
    - "- [Best Practices](../guides/best-practices) - Write better lessons
`````
