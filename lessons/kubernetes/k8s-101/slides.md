# What Problem Does Kubernetes Solve?

Docker lets you run containers on a single machine. But real-world applications need much more.

## The Challenge of Running Containers at Scale

Imagine you have a web application with 50 containers across 10 servers:

- A container crashes at 3 AM — who restarts it?
- Traffic spikes on Black Friday — how do you add more containers?
- You need to deploy a new version — how do you update all 50 containers without downtime?
- A server dies — how do you redistribute its containers?

Doing all this manually is impossible at scale.

## What Kubernetes Does

Kubernetes (K8s) is a **container orchestration platform** that automates:

| Concern              | What K8s Does                                        |
|----------------------|------------------------------------------------------|
| **Self-healing**     | Restarts crashed containers automatically            |
| **Scaling**          | Adds/removes replicas based on CPU, memory, or custom metrics |
| **Rolling updates**  | Deploys new versions with zero downtime              |
| **Service discovery**| Gives each service a stable DNS name and IP          |
| **Load balancing**   | Distributes traffic across healthy pods              |
| **Storage**          | Mounts persistent volumes to stateful apps           |
| **Secret management**| Injects passwords and API keys securely              |

## You Declare, Kubernetes Delivers

Kubernetes is **declarative**. You write a YAML file saying *"I want 3 replicas of my API"* and Kubernetes continuously works to make that true — even if a node fails.

```
You: "I want 3 replicas of my-api"
K8s: "Got it. I'll create 3 pods, watch them, and replace any that fail."
```

---

# Pods, Nodes, and Clusters

These are the foundational building blocks of every Kubernetes deployment.

## Cluster

A Kubernetes **cluster** is a set of machines (nodes) that run containerized applications. Every cluster has:

```
┌─────────────────────────────────────────────────┐
│                  CLUSTER                         │
│                                                  │
│  ┌──────────────┐    ┌──────────────────────┐   │
│  │ Control Plane │    │     Worker Nodes     │   │
│  │              │    │                      │   │
│  │ • API Server │    │  Node 1: [Pod][Pod]  │   │
│  │ • Scheduler  │    │  Node 2: [Pod][Pod]  │   │
│  │ • etcd       │    │  Node 3: [Pod]       │   │
│  │ • Controller │    │                      │   │
│  └──────────────┘    └──────────────────────┘   │
└─────────────────────────────────────────────────┘
```

- **Control Plane** — the brain: API server, scheduler, controllers, etcd (state store)
- **Worker Nodes** — the muscle: run your application pods

## Node

A **node** is a single machine (VM or physical) in the cluster. Each node runs:

- **kubelet** — agent that receives instructions from the control plane and manages pods on that node
- **kube-proxy** — handles networking rules so pods can communicate
- **Container runtime** — Docker, containerd, or CRI-O to actually run containers

## Pod

A **pod** is the smallest deployable unit — a wrapper around one or more containers.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-app
spec:
  containers:
    - name: api
      image: my-api:1.0
      ports:
        - containerPort: 3000
```

Key facts about pods:
- Containers in the same pod share **network** (localhost) and **storage**
- Pods get their own **IP address** within the cluster
- Pods are **ephemeral** — they can be killed and recreated at any time
- You rarely create pods directly; use Deployments instead

---

# Deployments and ReplicaSets

In practice, you almost never create pods directly. You use **Deployments** to manage them.

## Why Not Create Pods Directly?

If you create a standalone pod and it crashes, nothing restarts it. A Deployment ensures your desired state is always maintained.

## Deployment

A Deployment is a higher-level controller that:

1. Creates a **ReplicaSet** (which manages the pods)
2. Ensures the desired number of **replicas** are running
3. Handles **rolling updates** when you change the image
4. Supports **rollback** to previous versions

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-api
spec:
  replicas: 3
  selector:
    matchLabels:
      app: my-api
  template:
    metadata:
      labels:
        app: my-api
    spec:
      containers:
        - name: api
          image: my-api:1.2.0
          ports:
            - containerPort: 3000
          resources:
            requests:
              memory: "128Mi"
              cpu: "250m"
            limits:
              memory: "256Mi"
              cpu: "500m"
```

## ReplicaSet

A ReplicaSet ensures that a specified number of pod replicas are running at all times. Deployments create ReplicaSets automatically — you rarely interact with them directly.

```
Deployment (my-api)
    └── ReplicaSet (my-api-7d9f8b6c4)
            ├── Pod (my-api-7d9f8b6c4-abc12)
            ├── Pod (my-api-7d9f8b6c4-def34)
            └── Pod (my-api-7d9f8b6c4-ghi56)
```

## Rolling Updates

When you update the image version, Kubernetes performs a rolling update:

```bash
kubectl set image deployment/my-api api=my-api:1.3.0
```

```
Old ReplicaSet (v1.2): 3 pods → 2 → 1 → 0  (scaling down)
New ReplicaSet (v1.3): 0 pods → 1 → 2 → 3  (scaling up)
```

At no point is the application completely down. If something goes wrong:

```bash
kubectl rollout undo deployment/my-api
```

---

# Services: Exposing Your Application

Pods are ephemeral — they get new IP addresses when recreated. **Services** provide a stable endpoint.

## The Problem

```
Pod my-api-abc12  →  IP: 10.1.0.15  →  Pod crashes
Pod my-api-def34  →  IP: 10.1.0.22  ←  New pod, NEW IP!
```

Other services can't rely on pod IPs. A **Service** gives you a single, stable DNS name and IP.

## Service Types

### ClusterIP (Default)

Internal-only. Other pods can reach the service by name.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-api
spec:
  type: ClusterIP
  selector:
    app: my-api
  ports:
    - port: 80
      targetPort: 3000
```

Now any pod in the cluster can call `http://my-api:80` and traffic is load-balanced across all matching pods.

### NodePort

Exposes the service on a static port (30000–32767) on every node.

```yaml
spec:
  type: NodePort
  ports:
    - port: 80
      targetPort: 3000
      nodePort: 30080
```

Access from outside: `http://<any-node-ip>:30080`

### LoadBalancer

Provisions a cloud load balancer (AWS ELB, GCP LB, etc.) that routes external traffic.

```yaml
spec:
  type: LoadBalancer
  ports:
    - port: 80
      targetPort: 3000
```

This is the standard way to expose applications to the internet in cloud environments.

## How Services Route Traffic

Services use **label selectors** to find target pods:

```
Service (selector: app=my-api)
    │
    ├──► Pod (labels: app=my-api)  ✓ matches
    ├──► Pod (labels: app=my-api)  ✓ matches
    └──► Pod (labels: app=my-web)  ✗ no match
```

---

# ConfigMaps and Secrets

Kubernetes provides first-class objects for separating configuration from code.

## ConfigMaps: Non-Sensitive Config

Store configuration data as key-value pairs:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
data:
  DATABASE_HOST: "postgres.default.svc.cluster.local"
  LOG_LEVEL: "info"
  MAX_CONNECTIONS: "100"
```

### Injecting ConfigMaps into Pods

**As environment variables:**

```yaml
spec:
  containers:
    - name: api
      image: my-api:1.0
      envFrom:
        - configMapRef:
            name: app-config
```

**As mounted files:**

```yaml
spec:
  containers:
    - name: api
      image: my-api:1.0
      volumeMounts:
        - name: config
          mountPath: /etc/config
  volumes:
    - name: config
      configMap:
        name: app-config
```

## Secrets: Sensitive Data

Secrets are like ConfigMaps but for sensitive data (passwords, tokens, TLS certs). Values are base64-encoded (not encrypted by default — enable encryption at rest in production).

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: db-credentials
type: Opaque
data:
  username: YWRtaW4=          # base64 of "admin"
  password: cEBzc3cwcmQxMjM=  # base64 of "p@ssw0rd123"
```

```bash
# Create secrets from the command line
kubectl create secret generic db-credentials \
  --from-literal=username=admin \
  --from-literal=password='p@ssw0rd123'
```

### Important Security Notes

- Base64 is **encoding**, not encryption — anyone with cluster access can decode it
- Enable **encryption at rest** for etcd in production
- Use **external secret managers** (Vault, AWS Secrets Manager) via operators for serious workloads
- Apply **RBAC** to restrict who can read secrets

---

# kubectl Essential Commands

`kubectl` is the command-line tool for interacting with Kubernetes clusters.

## Getting Information

```bash
# View cluster info
kubectl cluster-info

# List nodes in the cluster
kubectl get nodes

# List all pods (current namespace)
kubectl get pods

# List pods across ALL namespaces
kubectl get pods -A

# Get detailed info about a specific pod
kubectl describe pod my-api-7d9f8b6c4-abc12

# List services, deployments, configmaps
kubectl get svc
kubectl get deployments
kubectl get configmaps
```

## Creating and Updating Resources

```bash
# Apply a YAML manifest (create or update)
kubectl apply -f deployment.yaml

# Apply all files in a directory
kubectl apply -f ./k8s/

# Scale a deployment
kubectl scale deployment my-api --replicas=5

# Update an image (triggers rolling update)
kubectl set image deployment/my-api api=my-api:2.0
```

## Debugging

```bash
# View pod logs
kubectl logs my-api-7d9f8b6c4-abc12

# Stream logs in real-time
kubectl logs -f my-api-7d9f8b6c4-abc12

# Open a shell inside a running pod
kubectl exec -it my-api-7d9f8b6c4-abc12 -- /bin/sh

# See events (useful for debugging scheduling issues)
kubectl get events --sort-by='.lastTimestamp'

# Check why a pod is pending/failing
kubectl describe pod <pod-name>
```

## Deleting Resources

```bash
# Delete a specific resource
kubectl delete pod my-api-7d9f8b6c4-abc12
kubectl delete deployment my-api

# Delete everything defined in a manifest
kubectl delete -f deployment.yaml

# Delete all pods in a namespace (careful!)
kubectl delete pods --all -n my-namespace
```

## Useful Shortcuts

```bash
# Short names:  po=pods, svc=services, deploy=deployments, ns=namespaces
kubectl get po
kubectl get svc
kubectl get deploy

# Output as YAML (great for learning)
kubectl get deployment my-api -o yaml

# Watch for changes in real-time
kubectl get pods -w
```

---

# Namespaces

Namespaces provide logical isolation within a single Kubernetes cluster.

## Why Namespaces?

In a shared cluster, you need boundaries:

- **Team isolation** — team-a and team-b don't accidentally interfere
- **Environment separation** — dev, staging, prod in one cluster
- **Resource quotas** — limit CPU/memory per namespace
- **Access control** — RBAC policies scoped to namespaces

## Default Namespaces

Every cluster comes with:

| Namespace           | Purpose                                          |
|---------------------|--------------------------------------------------|
| `default`           | Where resources go if no namespace is specified  |
| `kube-system`       | Kubernetes system components (DNS, proxy, etc.)  |
| `kube-public`       | Publicly accessible data (rarely used)           |
| `kube-node-lease`   | Node heartbeat leases                            |

## Working with Namespaces

```bash
# Create a namespace
kubectl create namespace staging

# List namespaces
kubectl get namespaces

# Deploy to a specific namespace
kubectl apply -f deployment.yaml -n staging

# Set your default namespace (avoid typing -n every time)
kubectl config set-context --current --namespace=staging

# View resources across all namespaces
kubectl get pods -A
```

## Resource Quotas

Limit how much a namespace can consume:

```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: team-quota
  namespace: team-a
spec:
  hard:
    requests.cpu: "4"
    requests.memory: 8Gi
    limits.cpu: "8"
    limits.memory: 16Gi
    pods: "20"
```

## Network Policies

By default, all pods can talk to all other pods across namespaces. Use **NetworkPolicies** to restrict this:

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: deny-other-namespaces
  namespace: production
spec:
  podSelector: {}
  ingress:
    - from:
        - podSelector: {}
```

This policy allows only pods within the `production` namespace to communicate with each other.

---

# Helm: The Kubernetes Package Manager

Helm simplifies deploying complex applications on Kubernetes.

## The Problem with Raw YAML

A typical application might require 10+ YAML files: Deployment, Service, ConfigMap, Secret, Ingress, HPA, PDB, ServiceAccount, RBAC... Managing all of these manually is tedious and error-prone.

## What Helm Does

Helm packages Kubernetes manifests into **charts** — reusable, version-controlled bundles with configurable values.

```
my-app-chart/
├── Chart.yaml          # Chart metadata (name, version)
├── values.yaml         # Default configuration values
└── templates/
    ├── deployment.yaml # Templates with {{ .Values.xxx }} placeholders
    ├── service.yaml
    ├── configmap.yaml
    └── ingress.yaml
```

## Using Helm

```bash
# Add a chart repository
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update

# Search for a chart
helm search repo postgresql

# Install a chart (creates a "release")
helm install my-db bitnami/postgresql \
  --set auth.postgresPassword=secret \
  --set primary.persistence.size=10Gi

# List installed releases
helm list

# Upgrade a release with new values
helm upgrade my-db bitnami/postgresql \
  --set primary.persistence.size=20Gi

# Rollback to a previous revision
helm rollback my-db 1

# Uninstall
helm uninstall my-db
```

## Custom values.yaml

Override defaults without modifying templates:

```yaml
# my-values.yaml
replicaCount: 3
image:
  repository: my-api
  tag: "1.5.0"
resources:
  requests:
    memory: "256Mi"
    cpu: "250m"
  limits:
    memory: "512Mi"
    cpu: "500m"
ingress:
  enabled: true
  hostname: api.example.com
```

```bash
helm install my-api ./my-app-chart -f my-values.yaml
```

## Why Helm Matters

- **Reproducibility** — same chart, same config → same deployment
- **Versioning** — rollback to any previous release
- **Ecosystem** — thousands of community charts for databases, monitoring, CI/CD tools
- **Templating** — one chart handles dev, staging, and prod with different values files

## What's Next?

- Practice with **Minikube** or **kind** for local K8s clusters
- Learn **Ingress controllers** for routing HTTP traffic
- Explore **Horizontal Pod Autoscaler** for auto-scaling
- Study **RBAC** for production access control
- Try **GitOps** with ArgoCD or Flux for declarative deployments
