---
sidebar_position: 5
---

# Challenge Steps

Challenge steps present open-ended tasks for users to complete independently.

## Simple Example
`````yaml
- type: challenge
  title: "🚀 Build Your First Image"
  description: |
    Create a Docker image for a simple web app.

    **Requirements:**
    1. Create a `Dockerfile` in the current directory
    2. Use `nginx:alpine` as the base image
    3. Build the image with tag `my-webapp:v1`

    **Hint:** Remember `docker build -t <tag> .`

  verification:
    type: "custom"
    checks:
      - type: "file_exists"
        path: "Dockerfile"

      - type: "command_succeeds"
        command: "docker images | grep my-webapp"

  success_msg: "🎉 Image built successfully!"
  allow_skip: true
`````

## Complete Examples

### Example 1: Container Deployment
`````yaml
- type: challenge
  title: "🚢 Deploy an Application"
  description: |
    Deploy a web application using everything you've learned.

    ## Requirements

    1. **Pull** the `nginx:alpine` image
    2. **Run** the container with these specs:
       - Name: `my-nginx`
       - Port: Map 8080 → 80
       - Detached mode
    3. **Verify** it's accessible at http://localhost:8080

    ## Success Criteria

    - Container is running
    - Port 8080 responds with nginx welcome page

    ## Tips

    - Use `docker run -d -p <host>:<container> --name <name> <image>`
    - Test with: `curl http://localhost:8080`
    - Check status: `docker ps`

  verification:
    type: "custom"
    checks:
      - type: "docker_container_running"
        name: "my-nginx"

      - type: "command_succeeds"
        command: "curl -f http://localhost:8080"

      - type: "port_open"
        port: 8080

  hints:
    - level: 1
      text: "Start by pulling the nginx:alpine image"

    - level: 2
      text: "Run with: docker run -d -p 8080:80 --name my-nginx nginx:alpine"

    - level: 3
      text: |
        Full solution:
        1. docker pull nginx:alpine
        2. docker run -d -p 8080:80 --name my-nginx nginx:alpine
        3. curl http://localhost:8080

  success_msg: "🎉 Application deployed! You're a DevOps pro!"
  allow_skip: true
`````

### Example 2: Image Signing Challenge
`````yaml
- type: challenge
  title: "🔐 Sign Your Own Image"
  description: |
    Now it's your turn! Apply what you've learned about Cosign.

    ## Your Task

    1. Pick any public image (e.g., `nginx:latest`, `alpine:latest`)
    2. Sign it with your key pair (`cosign.key`)
    3. Verify the signature
    4. Save the verification output to `verification.txt`

    ## Bonus Points

    - Try signing a local Docker image you built
    - Sign multiple images
    - Experiment with keyless signing

    ## Need Help?

    Use the hints if you get stuck. Remember the commands from earlier!

  verification:
    type: "custom"
    checks:
      - type: "file_exists"
        path: "verification.txt"

      - type: "file_contains"
        path: "verification.txt"
        pattern: "Verification for"

  hints:
    - level: 1
      text: "Use the same sign and verify commands from earlier steps"

    - level: 2
      text: |
        Steps:
        1. docker pull <image>
        2. cosign sign --key cosign.key <image>
        3. cosign verify --key cosign.pub <image> > verification.txt

    - level: 3
      text: |
        Example with nginx:
        docker pull nginx:latest
        cosign sign --key cosign.key nginx:latest
        cosign verify --key cosign.pub nginx:latest > verification.txt

  success_msg: "🎉 Challenge complete! You're a Cosign expert!"
  allow_skip: true
`````

### Example 3: Kubernetes Deployment
`````yaml
- type: challenge
  title: "☸️ Deploy to Kubernetes"
  description: |
    Time to deploy a real application to Kubernetes!

    ## Objective

    Create and deploy a complete application stack:

    1. **Create** a deployment YAML (`deployment.yaml`)
       - Use `nginx:latest` image
       - 3 replicas
       - Label: `app=nginx`

    2. **Create** a service YAML (`service.yaml`)
       - Type: LoadBalancer
       - Port: 80
       - Selector: `app=nginx`

    3. **Apply** both to your cluster

    4. **Verify** pods are running:
       - All 3 replicas ready
       - Service has external IP

    ## Documentation

    - [Deployments](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
    - [Services](https://kubernetes.io/docs/concepts/services-networking/service/)

    ## Testing
````bash
    kubectl get pods
    kubectl get service
````

  verification:
    type: "custom"
    checks:
      - type: "file_exists"
        path: "deployment.yaml"

      - type: "file_exists"
        path: "service.yaml"

      - type: "command_succeeds"
        command: "kubectl get deployment nginx -o jsonpath='{.status.readyReplicas}' | grep 3"

      - type: "command_succeeds"
        command: "kubectl get service nginx"

  hints:
    - level: 1
      text: "Start with a basic deployment template from kubernetes.io"

    - level: 2
      text: |
        Deployment structure:
````yaml
        apiVersion: apps/v1
        kind: Deployment
        metadata:
          name: nginx
        spec:
          replicas: 3
          selector:
            matchLabels:
              app: nginx
          template:
            metadata:
              labels:
                app: nginx
            spec:
              containers:
              - name: nginx
                image: nginx:latest
````

    - level: 3
      text: |
        Service structure:
````yaml
        apiVersion: v1
        kind: Service
        metadata:
          name: nginx
        spec:
          type: LoadBalancer
          ports:
          - port: 80
          selector:
            app: nginx
````

        Apply with:
        kubectl apply -f deployment.yaml
        kubectl apply -f service.yaml

  success_msg: "🎉 Kubernetes deployment successful! You're ready for production!"
  allow_skip: true
`````

### Example 4: Troubleshooting Challenge
`````yaml
- type: challenge
  title: "🔍 Debug a Broken Deployment"
  description: |
    Uh oh! There's a broken deployment in your cluster.

    ## The Problem

    A deployment named `broken-app` exists but isn't working.
    Your job is to:

    1. **Investigate** what's wrong
    2. **Fix** the issue
    3. **Verify** the app is running
    4. **Document** what was wrong in `fix-report.txt`

    ## Investigation Tools
````bash
    kubectl get pods
    kubectl describe pod <pod-name>
    kubectl logs <pod-name>
    kubectl get events
````

    ## Common Issues

    - Image pull errors
    - Resource constraints
    - Configuration errors
    - Port conflicts

    ## Success Criteria

    - All pods running and ready
    - Report explains the issue and fix

  verification:
    type: "custom"
    checks:
      - type: "command_succeeds"
        command: "kubectl get deployment broken-app -o jsonpath='{.status.readyReplicas}' | grep -v 0"

      - type: "file_exists"
        path: "fix-report.txt"

      - type: "file_contains"
        path: "fix-report.txt"
        pattern: "issue|problem|fixed"

  hints:
    - level: 1
      text: "Start by checking pod status: kubectl get pods"

    - level: 2
      text: "Describe the pod to see events: kubectl describe pod <pod-name>"

    - level: 3
      text: "Check logs for errors: kubectl logs <pod-name>"

    - level: 4
      text: "Common fix: Edit deployment with: kubectl edit deployment broken-app"

  success_msg: "🎉 Bug fixed! Great troubleshooting skills!"
  allow_skip: true
`````
