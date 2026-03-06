#!/bin/bash
# Docker 101 Exercises
# Work through these exercises to practice Docker fundamentals.

# =============================================================================
# EXERCISE: Pull and run your first container
# Pull the official nginx image and run it, mapping port 8080 on your host
# to port 80 in the container. Then verify it's running with docker ps.
# =============================================================================

docker pull nginx:alpine
docker run -d --name my-nginx -p 8080:80 nginx:alpine
docker ps

# Verify the web server is responding
curl -s http://localhost:8080 | head -5

# Clean up
docker stop my-nginx && docker rm my-nginx

# =============================================================================
# EXERCISE: Build a custom image from a Dockerfile
# Create a simple Dockerfile that serves a custom HTML page with nginx.
# Build the image and run a container from it.
# =============================================================================

mkdir -p /tmp/docker-exercise && cd /tmp/docker-exercise

cat > index.html <<'HTMLEOF'
<!DOCTYPE html>
<html>
<body><h1>Hello from my custom Docker image!</h1></body>
</html>
HTMLEOF

cat > Dockerfile <<'DEOF'
FROM nginx:alpine
COPY index.html /usr/share/nginx/html/index.html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
DEOF

docker build -t my-custom-nginx:1.0 .
docker run -d --name custom-web -p 8080:80 my-custom-nginx:1.0

# Verify custom content is served
curl -s http://localhost:8080

# Clean up
docker stop custom-web && docker rm custom-web
docker rmi my-custom-nginx:1.0

# =============================================================================
# EXERCISE: Work with Docker volumes
# Create a named volume, mount it to a container, write data, remove the
# container, then start a new container with the same volume to prove
# the data persists.
# =============================================================================

docker volume create exercise-data

# Write data to the volume
docker run --rm -v exercise-data:/data alpine sh -c 'echo "Persistent data written at $(date)" > /data/message.txt'

# Read data back from a NEW container (proving persistence)
docker run --rm -v exercise-data:/data alpine cat /data/message.txt

# Clean up
docker volume rm exercise-data

# =============================================================================
# EXERCISE: Create a user-defined bridge network
# Create a custom network, run two containers on it, and verify they
# can communicate by hostname.
# =============================================================================

docker network create my-exercise-net

docker run -d --name server --network my-exercise-net nginx:alpine
docker run --rm --network my-exercise-net alpine sh -c 'wget -qO- http://server:80 | head -3'

# Clean up
docker stop server && docker rm server
docker network rm my-exercise-net

# =============================================================================
# EXERCISE: Inspect and debug a container
# Run a container, then use docker exec, docker logs, and docker inspect
# to explore its internals.
# =============================================================================

docker run -d --name debug-target nginx:alpine

# View the container's logs
docker logs debug-target

# Execute a command inside the running container
docker exec debug-target cat /etc/os-release

# Get the container's IP address
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' debug-target

# Check resource usage
docker stats debug-target --no-stream

# Clean up
docker stop debug-target && docker rm debug-target

# =============================================================================
# EXERCISE: Multi-stage build
# Create a Go application, build it with a multi-stage Dockerfile, and
# compare the image size to a single-stage build.
# =============================================================================

mkdir -p /tmp/multistage-exercise && cd /tmp/multistage-exercise

cat > main.go <<'GOEOF'
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from a tiny Go container!\n")
    })
    fmt.Println("Listening on :8080")
    http.ListenAndServe(":8080", nil)
}
GOEOF

cat > go.mod <<'MODEOF'
module multistage-demo
go 1.22
MODEOF

cat > Dockerfile <<'DEOF'
# Stage 1: Build
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -o server .

# Stage 2: Minimal runtime
FROM scratch
COPY --from=builder /app/server /server
EXPOSE 8080
CMD ["/server"]
DEOF

docker build -t go-multistage:1.0 .

# Check the image size — should be only a few MB!
docker images go-multistage:1.0

# Run it
docker run -d --name go-app -p 8080:8080 go-multistage:1.0
curl -s http://localhost:8080

# Clean up
docker stop go-app && docker rm go-app
docker rmi go-multistage:1.0
rm -rf /tmp/multistage-exercise /tmp/docker-exercise
