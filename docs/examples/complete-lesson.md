---
sidebar_position: 1
---

# Complete Lesson Example

A full data-only lesson with metadata, slides, quiz, and exercises.

## Directory Structure

```
lessons/tutorial/docker-basics/
  lesson.yaml
  slides.md
  exercises.sh
```

## lesson.yaml

```yaml
id: docker-basics
title: "Docker Basics: Your First Container"
description: "Learn to run, build, and manage Docker containers"
author: "TryOutShell"
tags: ["docker", "containers", "devops"]
difficulty: "beginner"
duration: "20 min"
version: "1.0"

quiz:
  - question: "What is a Docker container?"
    options:
      - "A virtual machine"
      - "A lightweight, isolated process environment"
      - "A programming language"
      - "A cloud service"
    answer: 1
    explain: "Containers are lightweight process environments, not full VMs."

  - question: "What command lists running containers?"
    options:
      - "docker images"
      - "docker ps"
      - "docker list"
      - "docker containers"
    answer: 1
    explain: "docker ps lists running containers. Add -a for all containers."

  - question: "What file defines a Docker image build?"
    options:
      - "Makefile"
      - "docker-compose.yml"
      - "Dockerfile"
      - "config.yaml"
    answer: 2
    explain: "A Dockerfile contains instructions for building a Docker image."

  - question: "What does docker run --rm do?"
    options:
      - "Runs in background"
      - "Restarts on failure"
      - "Removes the container after it exits"
      - "Runs as root"
    answer: 2
    explain: "--rm automatically removes the container when it exits."

  - question: "Which command builds an image from a Dockerfile?"
    options:
      - "docker create"
      - "docker build"
      - "docker make"
      - "docker compile"
    answer: 1
    explain: "docker build reads a Dockerfile and creates an image."
```

## slides.md

```markdown
# Docker Basics

Learn Docker fundamentals — images, containers, and Dockerfiles.

---

## What is Docker?

Docker packages applications with all dependencies into **containers**:

- Lightweight (not full VMs)
- Portable across environments
- Reproducible builds
- Fast startup (milliseconds)

```bash
docker --version
```

---

## Images vs Containers

| Concept | Analogy | Description |
|---------|---------|-------------|
| **Image** | Recipe | Read-only template |
| **Container** | Dish | Running instance of an image |

An image is a blueprint. A container is a running process from that blueprint.

---

## Running Your First Container

```bash
docker run hello-world
```

This will:
1. Pull the `hello-world` image from Docker Hub
2. Create a container from it
3. Run it and print a message
4. Exit

---

## Useful Docker Commands

```bash
# List running containers
docker ps

# List all containers (including stopped)
docker ps -a

# List images
docker images

# Remove a container
docker rm <container-id>

# Remove an image
docker rmi <image-name>
```

---

## Building a Custom Image

Create a `Dockerfile`:

```dockerfile
FROM alpine:latest
RUN echo "Hello from my custom image" > /greeting.txt
CMD ["cat", "/greeting.txt"]
```

Build and run:

```bash
docker build -t my-app .
docker run --rm my-app
```

---

## Summary

Key takeaways:

1. Docker packages apps in lightweight containers
2. Images are blueprints, containers are running instances
3. `docker run` creates and starts a container
4. `Dockerfile` defines how to build an image
5. `docker build` creates images from Dockerfiles

**Next**: Try the quiz to test your knowledge!
```

## exercises.sh

```bash
#!/bin/bash
# EXERCISE 1: Check Docker version
echo "=== Exercise 1: Check Docker version ==="
docker --version

# EXERCISE 2: Pull and run hello-world
echo "=== Exercise 2: Run hello-world ==="
docker run --rm hello-world

# EXERCISE 3: List images
echo "=== Exercise 3: List Docker images ==="
docker images
```

## Running This Lesson

```bash
# View the lesson slides
go run . start tutorial --lesson docker-basics

# Take the quiz
go run . quiz tutorial docker-basics
```
