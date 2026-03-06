# What is Docker and Why It Matters

Docker is a platform that packages applications and their dependencies into lightweight, portable **containers** that run consistently across any environment.

## The "Works on My Machine" Problem

Before containers, deploying software was notoriously fragile:

- A developer writes code on macOS with Python 3.11, OpenSSL 3.0, and specific C libraries
- The staging server runs Ubuntu 20.04 with Python 3.8 and OpenSSL 1.1
- Production is on Amazon Linux 2 with yet another set of library versions
- The result: **"It works on my machine!"** — the most dreaded phrase in software engineering

## How Docker Solves This

Docker bundles your application code **along with its entire runtime environment** into an image:

```
┌──────────────────────────┐
│     Your Application     │
├──────────────────────────┤
│  Runtime (Python, Node)  │
├──────────────────────────┤
│  System Libraries / Deps │
├──────────────────────────┤
│   Base OS (Alpine, etc.) │
└──────────────────────────┘
        Docker Image
```

The same image runs identically on a developer's laptop, in CI/CD, and in production. No more environment drift.

## Key Benefits

- **Consistency** — identical behavior from dev to prod
- **Isolation** — containers don't interfere with each other or the host
- **Speed** — containers start in milliseconds (unlike VMs which take minutes)
- **Efficiency** — containers share the host kernel, using far less RAM than VMs
- **Portability** — runs on any machine with Docker installed

---

# Images vs Containers

Understanding the relationship between images and containers is fundamental to working with Docker.

## The Blueprint Analogy

Think of it like architecture:

| Concept        | Analogy               | Description                                      |
|----------------|-----------------------|--------------------------------------------------|
| **Image**      | Blueprint / Recipe    | A read-only template with all files and config   |
| **Container**  | Building / Dish       | A live, running instance created from an image   |

You can create **many containers** from **one image**, just like you can build many houses from one blueprint.

## Images Are Layered

Docker images are built from **layers**. Each instruction in a Dockerfile creates a new layer:

```
Layer 4: CMD ["python", "app.py"]        (metadata)
Layer 3: COPY . /app                     (+2 MB)
Layer 2: RUN pip install flask           (+15 MB)
Layer 1: FROM python:3.12-slim          (+120 MB)
```

Layers are **cached and shared** between images. If two images both start with `python:3.12-slim`, that base layer is stored only once on disk.

## Container Lifecycle

```
   Image
     │
     ▼
  Created  ──►  Running  ──►  Stopped  ──►  Removed
 (docker       (docker       (docker       (docker
  create)       start)        stop)         rm)
```

A running container has a thin **writable layer** on top of the read-only image layers. When the container is removed, that writable layer is discarded — unless you use volumes to persist data.

## Quick Commands

```bash
# List local images
docker images

# List running containers
docker ps

# List all containers (including stopped)
docker ps -a
```

---

# Essential Docker Commands

These are the commands you will use every day when working with Docker.

## Running Containers

```bash
# Run a container from an image (pulls if not local)
docker run nginx

# Run in detached mode (background)
docker run -d nginx

# Run with a name for easy reference
docker run -d --name my-web nginx

# Run interactively with a shell
docker run -it ubuntu bash

# Map host port 8080 to container port 80
docker run -d -p 8080:80 nginx
```

## Building Images

```bash
# Build an image from a Dockerfile in the current directory
docker build -t my-app:1.0 .

# Build with a specific Dockerfile
docker build -f Dockerfile.prod -t my-app:prod .

# Build with build arguments
docker build --build-arg NODE_ENV=production -t my-app .
```

## Inspecting and Managing

```bash
# View container logs
docker logs my-web
docker logs -f my-web        # follow (stream) logs

# Execute a command inside a running container
docker exec -it my-web bash

# Inspect container details (JSON output)
docker inspect my-web

# View resource usage
docker stats
```

## Stopping and Cleaning Up

```bash
# Stop a running container (graceful SIGTERM)
docker stop my-web

# Force-stop (SIGKILL)
docker kill my-web

# Remove a stopped container
docker rm my-web

# Remove an image
docker rmi nginx

# Nuclear option: remove all stopped containers, unused images, networks
docker system prune -a
```

---

# Writing a Dockerfile

A Dockerfile is a text file with instructions that Docker uses to build an image, layer by layer.

## Dockerfile Instructions

| Instruction | Purpose                                          |
|-------------|--------------------------------------------------|
| `FROM`      | Set the base image                               |
| `RUN`       | Execute a command during build (install packages) |
| `COPY`      | Copy files from host into the image              |
| `ADD`       | Like COPY but also handles URLs and tar archives |
| `WORKDIR`   | Set the working directory inside the container   |
| `ENV`       | Set environment variables                        |
| `EXPOSE`    | Document which port the app listens on           |
| `CMD`       | Default command when container starts            |
| `ENTRYPOINT`| Like CMD but harder to override                  |

## A Real-World Example

Here is a production-ready Dockerfile for a Python Flask application:

```dockerfile
FROM python:3.12-slim

WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

EXPOSE 5000

ENV FLASK_APP=app.py
ENV FLASK_ENV=production

CMD ["gunicorn", "--bind", "0.0.0.0:5000", "app:app"]
```

## Understanding the Build Order

The order of instructions matters for **caching**:

1. `COPY requirements.txt .` and `RUN pip install ...` are placed **before** `COPY . .`
2. This way, dependencies are only reinstalled when `requirements.txt` changes
3. Code changes (which happen frequently) only invalidate the final `COPY` layer

```
✓  Good: COPY requirements → RUN pip install → COPY code
✗  Bad:  COPY everything → RUN pip install
```

## Multi-Stage Builds

Use multi-stage builds to keep final images small:

```dockerfile
# Stage 1: Build
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o server .

# Stage 2: Runtime (only ~5 MB!)
FROM alpine:3.19
COPY --from=builder /app/server /server
CMD ["/server"]
```

The final image contains only the compiled binary — no Go toolchain, no source code.

---

# Docker Volumes and Bind Mounts

Containers are **ephemeral** by default. When a container is removed, all data inside it is lost. Volumes solve this.

## The Data Problem

```
Container created  →  App writes to /data/db  →  Container removed
                                                        │
                                                   Data is GONE ❌
```

## Three Storage Options

### 1. Volumes (Recommended)

Docker manages the storage location. Best for databases and persistent application data.

```bash
# Create a named volume
docker volume create my-data

# Mount it into a container
docker run -d -v my-data:/var/lib/mysql mysql:8

# The data in /var/lib/mysql persists even if the container is removed
docker rm -f <container-id>
docker run -d -v my-data:/var/lib/mysql mysql:8  # data is still there!
```

### 2. Bind Mounts

Map a specific host directory into the container. Great for development — edit code on the host and see changes instantly.

```bash
# Mount current directory into /app in the container
docker run -d -v $(pwd):/app -p 3000:3000 node:20

# Now editing files on your host immediately reflects inside the container
```

### 3. tmpfs Mounts

Store data in memory only. Useful for sensitive data that should never hit disk.

```bash
docker run -d --tmpfs /tmp:rw,size=64m my-app
```

## Volume Management

```bash
docker volume ls                 # List all volumes
docker volume inspect my-data    # Show volume details
docker volume rm my-data         # Delete a volume
docker volume prune              # Remove all unused volumes
```

---

# Docker Networking Basics

Docker creates isolated networks so containers can communicate with each other and the outside world.

## Network Drivers

| Driver      | Use Case                                      |
|-------------|-----------------------------------------------|
| **bridge**  | Default. Containers on the same host.         |
| **host**    | Container shares the host's network stack.    |
| **none**    | No networking at all.                         |
| **overlay** | Multi-host networking (Docker Swarm / K8s).   |

## Default Bridge Network

When you run a container without specifying a network, it joins the default `bridge` network:

```bash
docker run -d --name web nginx
docker run -d --name api node:20

# These can reach each other by IP, but NOT by name on the default bridge
```

## User-Defined Bridge Networks (Recommended)

Create a custom network to enable **DNS-based service discovery**:

```bash
# Create a network
docker network create my-app-net

# Run containers on that network
docker run -d --name web --network my-app-net nginx
docker run -d --name api --network my-app-net node:20

# Now 'web' can reach 'api' by hostname!
# Inside the web container:  curl http://api:3000
```

## Port Mapping

Containers are isolated by default. To expose a service to the host:

```bash
# -p HOST_PORT:CONTAINER_PORT
docker run -d -p 8080:80 nginx

# Now accessible at http://localhost:8080
```

## Inspecting Networks

```bash
docker network ls                    # List networks
docker network inspect my-app-net    # See connected containers, subnet, etc.
docker network connect my-app-net web  # Attach a running container to a network
```

---

# Docker Compose

Docker Compose lets you define and run **multi-container applications** with a single YAML file.

## Why Compose?

Running a real application often requires multiple services:

```bash
# Without Compose — you have to remember all of this:
docker network create app-net
docker run -d --name db --network app-net -v db-data:/var/lib/postgresql/data -e POSTGRES_PASSWORD=secret postgres:16
docker run -d --name redis --network app-net redis:7-alpine
docker run -d --name api --network app-net -p 3000:3000 -e DATABASE_URL=postgres://... my-api
docker run -d --name web --network app-net -p 80:80 my-frontend
```

That's four commands, each with multiple flags. Now imagine onboarding a new developer...

## A Real docker-compose.yml

```yaml
version: "3.9"

services:
  api:
    build: ./api
    ports:
      - "3000:3000"
    environment:
      DATABASE_URL: postgres://app:secret@db:5432/myapp
      REDIS_URL: redis://redis:6379
    depends_on:
      - db
      - redis

  web:
    build: ./frontend
    ports:
      - "80:80"
    depends_on:
      - api

  db:
    image: postgres:16-alpine
    environment:
      POSTGRES_USER: app
      POSTGRES_PASSWORD: secret
      POSTGRES_DB: myapp
    volumes:
      - db-data:/var/lib/postgresql/data

  redis:
    image: redis:7-alpine

volumes:
  db-data:
```

## Essential Compose Commands

```bash
# Start all services (build images if needed)
docker compose up -d

# View logs across all services
docker compose logs -f

# Stop and remove containers, networks
docker compose down

# Stop and also remove volumes (destroys data!)
docker compose down -v

# Rebuild images after code changes
docker compose up -d --build

# Scale a service
docker compose up -d --scale api=3
```

## How It Works

Compose automatically:
- Creates a **dedicated network** for all services (they can reach each other by service name)
- Creates **named volumes** as declared
- Starts services in **dependency order** (`depends_on`)

---

# Best Practices and Common Pitfalls

Follow these guidelines to build secure, efficient, and maintainable Docker images.

## Use Small Base Images

```dockerfile
# ✗ Bad: 900+ MB, includes compilers and tools you don't need
FROM ubuntu:22.04

# ✓ Better: ~130 MB
FROM python:3.12-slim

# ✓ Best (if compatible): ~5 MB
FROM python:3.12-alpine
```

## Don't Run as Root

By default containers run as root. This is a security risk.

```dockerfile
RUN addgroup --system app && adduser --system --ingroup app app
USER app
```

## Use .dockerignore

Prevent unnecessary files from being sent to the build context:

```
# .dockerignore
.git
node_modules
*.md
.env
__pycache__
```

## One Process per Container

Each container should run **one concern**:

- ✗ Don't run nginx + python + postgres in one container
- ✓ Use separate containers and connect them via a network

## Tag Your Images

```bash
# ✗ Bad: "latest" is ambiguous and causes deployment surprises
docker build -t my-app .

# ✓ Good: use semantic versioning or git SHA
docker build -t my-app:1.2.3 .
docker build -t my-app:$(git rev-parse --short HEAD) .
```

## Common Pitfalls

| Pitfall | Problem | Fix |
|---------|---------|-----|
| Storing data in containers | Data lost on removal | Use volumes |
| Ignoring layer caching | Slow rebuilds | Order Dockerfile instructions properly |
| Huge images | Slow pulls, waste disk | Use slim/alpine bases, multi-stage builds |
| Hardcoding config | Can't change per environment | Use ENV vars or config files |
| Running as root | Security vulnerability | Add a non-root USER |
| Not using .dockerignore | Slow builds, secrets leaked | Create .dockerignore |

## What's Next?

- Explore **Docker Hub** for official images
- Learn **Docker Compose** for multi-container workflows
- Dive into **Docker Swarm** or **Kubernetes** for orchestration
- Set up **CI/CD pipelines** that build and push images automatically
