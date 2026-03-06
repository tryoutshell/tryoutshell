# What is CI/CD and Why Does It Matter?

**Continuous Integration (CI)** is the practice of automatically building
and testing code every time a developer pushes changes. **Continuous Delivery
(CD)** extends this by automatically deploying validated code to production
or staging environments.

**Before CI/CD:**
- "It works on my machine" syndrome
- Manual testing before every release
- Merging code was painful (integration hell)
- Deployments were risky, infrequent, and stressful

**With CI/CD:**
- Every commit is built and tested automatically
- Bugs are caught within minutes, not days
- Deployments become routine and low-risk
- Team moves faster with higher confidence

**The CI/CD pipeline:**

```
Code Push → Build → Test → Security Scan → Deploy to Staging → Deploy to Prod
   ↑                                                                    |
   └────────────── Feedback Loop (minutes, not days) ──────────────────┘
```

**GitHub Actions** is GitHub's built-in CI/CD platform that lets you automate
workflows directly in your repository — no external service needed.

---

# GitHub Actions Core Concepts

Understanding these five building blocks is key to working with GitHub Actions:

**Workflow** — An automated process defined in a YAML file. Lives in
`.github/workflows/`. A repository can have multiple workflows.

**Event (Trigger)** — Something that causes a workflow to run: a push,
pull request, schedule, manual dispatch, or external webhook.

**Job** — A set of steps that execute on the same runner. Jobs run in
parallel by default but can be configured to run sequentially.

**Step** — A single task within a job. Can be a shell command (`run:`)
or a reusable action (`uses:`).

**Runner** — The server that executes your jobs. GitHub provides hosted
runners (Ubuntu, Windows, macOS) or you can use self-hosted runners.

```
Workflow (ci.yaml)
├── Event: on push to main
├── Job: build
│   ├── Step 1: Checkout code
│   ├── Step 2: Setup Node.js
│   ├── Step 3: Install dependencies
│   └── Step 4: Run tests
└── Job: deploy (needs: build)
    ├── Step 1: Checkout code
    └── Step 2: Deploy to production
```

Each job gets a fresh runner — jobs don't share filesystem state
unless you explicitly use artifacts or caches.

---

# Anatomy of a Workflow YAML

Every workflow file lives in `.github/workflows/` and follows this structure:

```yaml
name: CI Pipeline

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

permissions:
  contents: read

jobs:
  build-and-test:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout repository
        uses: actions/checkout@v4

      - name: Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: "20"
          cache: "npm"

      - name: Install dependencies
        run: npm ci

      - name: Run linter
        run: npm run lint

      - name: Run tests
        run: npm test

      - name: Build
        run: npm run build
```

**Key points:**
- `name:` — Human-readable label shown in the Actions tab
- `on:` — Events that trigger the workflow
- `permissions:` — Principle of least privilege for the GITHUB_TOKEN
- `runs-on:` — Which runner OS to use
- `uses:` — Reference to a published action (org/repo@version)
- `run:` — Shell command to execute
- `with:` — Input parameters for an action
- `npm ci` — Clean install from lockfile (preferred over `npm install` in CI)

---

# Workflow Triggers

GitHub Actions supports a rich set of event triggers:

**Code events:**
```yaml
on:
  push:
    branches: [main, "release/**"]
    paths: ["src/**", "package.json"]
  pull_request:
    branches: [main]
    types: [opened, synchronize, reopened]
```

**Scheduled runs (cron):**
```yaml
on:
  schedule:
    - cron: "0 6 * * 1"    # Every Monday at 6:00 UTC
```

**Manual dispatch (run from the UI):**
```yaml
on:
  workflow_dispatch:
    inputs:
      environment:
        description: "Deploy target"
        required: true
        default: "staging"
        type: choice
        options: [staging, production]
```

**Other useful triggers:**
- `release:` — When a GitHub Release is published
- `workflow_run:` — After another workflow completes
- `repository_dispatch:` — External webhook events
- `issue_comment:` — When someone comments on an issue/PR

**Path filters** are powerful — you can skip running the entire CI pipeline
if only documentation files changed:

```yaml
on:
  push:
    paths-ignore:
      - "docs/**"
      - "*.md"
      - "LICENSE"
```

---

# The Actions Marketplace

The **GitHub Actions Marketplace** has 20,000+ community-built actions
you can use as building blocks in your workflows.

**Essential actions you'll use constantly:**

| Action | Purpose |
|--------|---------|
| `actions/checkout@v4` | Clone your repo |
| `actions/setup-node@v4` | Install Node.js |
| `actions/setup-python@v5` | Install Python |
| `actions/setup-go@v5` | Install Go |
| `actions/cache@v4` | Cache dependencies |
| `actions/upload-artifact@v4` | Save build output |
| `actions/download-artifact@v4` | Retrieve artifacts |
| `github/codeql-action/analyze@v3` | Security scanning |

**Using a marketplace action:**

```yaml
- name: Setup Go environment
  uses: actions/setup-go@v5
  with:
    go-version: "1.22"
    cache: true

- name: Run Go vulnerability check
  uses: golang/govulncheck-action@v1
```

**Security tip — always pin actions to a full SHA:**

```yaml
# Risky: tag can be moved to point at malicious code
- uses: actions/checkout@v4

# Safer: pinned to an immutable commit SHA
- uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11
```

Pinning to SHAs prevents supply chain attacks where an action maintainer
(or attacker who compromised their account) pushes malicious code to an
existing tag.

---

# Secrets and Environment Variables

**Secrets** store sensitive values (API keys, tokens, passwords) encrypted
at rest. They are never exposed in logs.

**Setting secrets:**
1. Go to repo **Settings → Secrets and variables → Actions**
2. Click **New repository secret**
3. Name: `DEPLOY_TOKEN`, Value: your token

**Using secrets in workflows:**

```yaml
jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - name: Deploy to production
        env:
          API_KEY: ${{ secrets.DEPLOY_TOKEN }}
        run: |
          curl -X POST https://api.example.com/deploy \
            -H "Authorization: Bearer $API_KEY" \
            -d '{"ref": "${{ github.sha }}"}'
```

**Environment-level secrets** let you have different values per environment
(staging vs. production) with optional approval gates:

```yaml
jobs:
  deploy-prod:
    runs-on: ubuntu-latest
    environment: production    # Requires approval + uses prod secrets
    steps:
      - run: echo "Deploying ${{ github.sha }}"
```

**Built-in environment variables:**
- `GITHUB_SHA` — The commit SHA that triggered the workflow
- `GITHUB_REF` — The branch or tag ref
- `GITHUB_REPOSITORY` — Owner/repo name
- `GITHUB_TOKEN` — Auto-generated token with scoped permissions
- `RUNNER_OS` — The runner OS (Linux, Windows, macOS)

**Important:** Never `echo` secrets or pass them in URLs — they'll be
redacted in logs, but it's still bad practice.

---

# Matrix Builds

Matrix strategy lets you run the same job across multiple configurations
in parallel — different OS versions, language versions, or any variable.

```yaml
jobs:
  test:
    runs-on: ${{ matrix.os }}
    strategy:
      matrix:
        os: [ubuntu-latest, macos-latest, windows-latest]
        node-version: [18, 20, 22]
      fail-fast: false

    steps:
      - uses: actions/checkout@v4

      - name: Setup Node.js ${{ matrix.node-version }}
        uses: actions/setup-node@v4
        with:
          node-version: ${{ matrix.node-version }}

      - run: npm ci
      - run: npm test
```

This creates **9 parallel jobs** (3 OS × 3 Node versions).

**Advanced: include/exclude specific combinations:**

```yaml
strategy:
  matrix:
    os: [ubuntu-latest, windows-latest]
    node-version: [18, 20]
    include:
      - os: ubuntu-latest
        node-version: 22
        experimental: true
    exclude:
      - os: windows-latest
        node-version: 18
```

**Options:**
- `fail-fast: false` — Don't cancel other jobs when one fails
- `max-parallel: 2` — Limit concurrency (useful for resource constraints)
- `include:` — Add extra combinations beyond the Cartesian product
- `exclude:` — Remove specific combinations

Matrix builds are ideal for library authors who need to ensure compatibility
across multiple runtime versions and operating systems.

---

# Caching Dependencies

Without caching, every workflow run downloads and installs all dependencies
from scratch. Caching can **cut build times by 50-80%**.

**Built-in caching with setup actions:**

```yaml
- uses: actions/setup-node@v4
  with:
    node-version: "20"
    cache: "npm"          # Automatically caches ~/.npm
```

**Explicit caching with actions/cache:**

```yaml
- name: Cache Go modules
  uses: actions/cache@v4
  with:
    path: |
      ~/.cache/go-build
      ~/go/pkg/mod
    key: go-${{ runner.os }}-${{ hashFiles('**/go.sum') }}
    restore-keys: |
      go-${{ runner.os }}-

- name: Install dependencies
  run: go mod download
```

**How the cache key works:**
1. Actions generates a key from `go-Linux-abc123` (OS + hash of go.sum)
2. On cache **hit**: restores files, skips download
3. On cache **miss**: runs install, saves new cache entry
4. `restore-keys` provides fallback — partial match restores the closest cache

**Cache limits:**
- 10 GB total per repository
- Entries not accessed in 7 days are evicted
- Exact key match always wins over restore-keys

**Docker layer caching (for container builds):**

```yaml
- name: Build Docker image
  uses: docker/build-push-action@v5
  with:
    context: .
    push: true
    tags: myapp:latest
    cache-from: type=gha
    cache-to: type=gha,mode=max
```

The `type=gha` cache backend stores Docker layers in GitHub Actions cache,
dramatically speeding up container builds.

---

# Putting It All Together

Here's a production-grade workflow combining everything we've learned:

```yaml
name: CI/CD Pipeline

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

permissions:
  contents: read
  packages: write

jobs:
  test:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        go-version: ["1.21", "1.22"]

    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version: ${{ matrix.go-version }}
          cache: true

      - run: go test -race -coverprofile=coverage.out ./...

      - name: Upload coverage
        if: matrix.go-version == '1.22'
        uses: actions/upload-artifact@v4
        with:
          name: coverage
          path: coverage.out

  deploy:
    needs: test
    if: github.ref == 'refs/heads/main' && github.event_name == 'push'
    runs-on: ubuntu-latest
    environment: production

    steps:
      - uses: actions/checkout@v4

      - name: Deploy
        env:
          DEPLOY_KEY: ${{ secrets.DEPLOY_KEY }}
        run: ./scripts/deploy.sh
```

**Key patterns demonstrated:**
- Matrix builds for Go version compatibility
- `needs:` for job dependencies (deploy waits for tests)
- `if:` conditions to deploy only on push to main
- Environment protection rules for production deploys
- Artifact upload for test coverage reports

**Resources:**
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Actions Marketplace](https://github.com/marketplace?type=actions)
- [Workflow Syntax Reference](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions)
- [Security Hardening Guide](https://docs.github.com/en/actions/security-guides/security-hardening-for-github-actions)
