# 📚 TryOutShell Lesson Format Documentation

## Table of Contents

1. [Quick Start](#quick-start)
2. [Lesson Structure](#lesson-structure)
3. [Metadata Section](#metadata-section)
4. [Introduction Section](#introduction-section)
5. [Steps](#steps)
6. [Step Types Reference](#step-types-reference)
   - [Info Steps](#1-type-info)
   - [Command Steps](#2-type-command)
   - [Quiz Steps](#3-type-quiz)
   - [Challenge Steps](#4-type-challenge)
   - [Interview Prep Steps](#5-type-interview_prep)
7. [Conclusion Section](#conclusion-section)

---

## Quick Start

### Minimal Lesson Example

```yaml
metadata:
  id: "my-first-lesson"
  org: "my-org"
  title: "My First Lesson"
  description: "A simple introduction"
  difficulty: "beginner"
  duration: "10 min"
  tags: ["intro"]

steps:
  - type: info
    title: "Welcome"
    content: "Hello, world!"

  - type: command
    prompt: "Check your shell"
    example: "echo 'Hello from TryOutShell'"
    validation:
      type: "substring"
      contains: "Hello"
    success_msg: "✅ It works!"
    fail_msg: "❌ Try again"
```

---

## Lesson Structure

Every lesson YAML file has 4 main sections:

```yaml
metadata:        # Required - Lesson identification and metadata
introduction:    # Optional - Shown once at the start
steps:           # Required - Main lesson content
conclusion:      # Optional - Shown after completion
```

---

## Metadata Section

### Required Fields

| Field | Type | Description | Example |
|-------|------|-------------|---------|
| `id` | string | Unique identifier (kebab-case) | `"cosign-basics"` |
| `org` | string | Organization name | `"chainguard"` |
| `title` | string | Human-readable title | `"Introduction to Cosign"` |
| `description` | string | Brief description (1-2 sentences) | `"Learn to sign containers"` |
| `difficulty` | enum | `beginner`, `intermediate`, `advanced` | `"beginner"` |
| `duration` | string | Estimated time | `"15 min"` |
| `tags` | list | Keywords for search | `["cosign", "signing"]` |

### Optional Fields

| Field | Type | Description | Example |
|-------|------|-------------|---------|
| `prerequisites` | list | Tools/knowledge needed | `["Docker installed"]` |
| `author` | string | Creator name | `"Chainguard Team"` |
| `version` | string | Lesson version | `"1.0"` |
| `updated_at` | date | Last update | `"2025-01-15"` |

### Example

```yaml
metadata:
  id: "cosign-sign-verify"
  org: "chainguard"
  title: "Container Image Signing with Cosign"
  description: "Learn to sign and verify container images for supply chain security"
  difficulty: "beginner"
  duration: "20 min"
  prerequisites:
    - "Docker installed"
    - "Basic understanding of containers"
  tags: ["cosign", "signing", "supply-chain", "security"]
  author: "Chainguard Team"
  version: "1.0"
```

---

## Introduction Section

**Purpose:** Set expectations and provide context before diving into steps.

### Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `title` | string | Yes | Introduction header |
| `content` | string (multiline) | Yes | Overview text (supports Markdown) |

### Markdown Support

- `**bold**` → bold text
- `` `code` `` → inline code
- `### Heading` → subheadings
- `> Quote` → blockquotes
- `- List item` → bullet lists

### Example

```yaml
introduction:
  title: "What You'll Learn"
  content: |
    In this lesson, you will:
    - Understand what Cosign is and why image signing matters
    - Install and verify Cosign
    - Sign a container image with ephemeral keys
    - Verify signatures to ensure image integrity

    **Time:** ~20 minutes
    **Tools:** Cosign, Docker

    > 💡 **Tip:** Have Docker running before you start!
```

---

## Steps

### Structure

Each step is a YAML object in the `steps` list.

```yaml
steps:
  - type: info           # Step type (see below)
    title: "Step Title"
    # ... type-specific fields

  - type: command
    prompt: "Do something"
    # ... type-specific fields
```

---

## Step Types Reference

### 1. Type: `info`

**Purpose:** Display educational content (text, diagrams, explanations).

#### Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | `"info"` | Yes | Identifies this as an info step |
| `title` | string | Yes | Step heading |
| `content` | string (multiline) | Yes | Main text (Markdown supported) |
| `highlights` | list | No | Inline text highlights (see below) |
| `code_blocks` | list | No | Code examples (see below) |
| `callouts` | list | No | Tips/warnings (see below) |
| `diagram` | string | No | ASCII art diagram |
| `wait_for_continue` | boolean | No | Pause until user presses Enter (default: true) |

#### Highlights

```yaml
highlights:
  - text: "myapp:latest"      # Text to highlight
    style: "code"             # Style: "code", "bold", "highlight"
  - text: "Sigstore"
    style: "bold"
```

#### Code Blocks

```yaml
code_blocks:
  - label: "macOS (Homebrew)"     # Description
    code: "brew install cosign"   # The command
    language: "bash"              # Syntax highlighting

  - label: "Linux"
    code: |
      curl -LO https://...
      sudo mv cosign /usr/local/bin/
    language: "bash"
```

#### Callouts

```yaml
callouts:
  - type: "tip"       # Types: "tip", "warning", "danger", "info"
    text: "Always verify signatures in production"

  - type: "warning"
    text: "Never commit private keys to Git!"
```

#### Full Example

```yaml
- type: info
  title: "What is Cosign?"
  content: |
    **Cosign** is a tool for signing and verifying container images.

    Think of it like a digital signature for software packages - it proves:
    - **Who** created the image (authenticity)
    - **What** exact image you're running (integrity)

    > 💡 **Fun Fact**: Cosign is part of the Sigstore project.

  highlights:
    - text: "Cosign"
      style: "bold"

  code_blocks:
    - label: "Example"
      code: "cosign sign myimage:latest"
      language: "bash"

  callouts:
    - type: "tip"
      text: "Cosign integrates with Kubernetes admission controllers"

  diagram: |
    Registry → Cosign → Verified ✓

  wait_for_continue: true
```

---

### 2. Type: `command`

**Purpose:** Execute a shell command and verify the output.

#### Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | `"command"` | Yes | Identifies this as a command step |
| `id` | string | No | Unique ID for tracking progress |
| `prompt` | string | Yes | Short description (shown in UI) |
| `instruction` | string | No | Detailed instruction text |
| `example` | string | Yes | Example command to show user |
| `pre_content` | string | No | Explanation before command |
| `post_content` | string | No | Explanation after success |
| `validation` | object | Yes | How to verify success (see below) |
| `alternative_validations` | list | No | Additional validation methods |
| `accepted_commands` | list | No | List of valid command variations |
| `success_msg` | string | Yes | Message on success |
| `fail_msg` | string | Yes | Message on failure |
| `hints` | list | No | Progressive hints (see below) |
| `allow_skip` | boolean | No | Can user skip this step? (default: false) |
| `timeout` | integer | No | Max execution time in seconds (default: 30) |

#### Validation Types

```yaml
# 1. Substring match
validation:
  type: "substring"
  contains: "version"              # Output must contain this
  case_insensitive: true           # Optional

# 2. Regex match
validation:
  type: "regex"
  pattern: "cosign.*v\\d+\\.\\d+"  # Regex pattern
  case_insensitive: false

# 3. Exit code check
validation:
  type: "exit_code"
  expected: 0                      # Command must exit with 0

# 4. File exists
validation:
  type: "file_exists"
  files:
    - "cosign.key"
    - "cosign.pub"

# 5. File contains
validation:
  type: "file_contains"
  path: "output.txt"
  pattern: "success"

# 6. Output contains (multiple patterns)
validation:
  type: "output_contains"
  patterns:
    - "Successfully signed"
    - "Pushing signature"
  any_match: true                  # Match ANY pattern (OR logic)
  all_match: false                 # Match ALL patterns (AND logic)

# 7. JSON validation
validation:
  type: "json_path"
  path: "$.status"
  expected: "success"
```

#### Hints

```yaml
hints:
  - level: 1
    text: "Try typing: cosign version"

  - level: 2
    text: "Install with: brew install cosign"

  - level: 3
    text: "Full command: cosign version"
```

#### Full Example

```yaml
- type: command
  id: "check-cosign"
  prompt: "Verify Cosign is installed"
  instruction: "Run the command to check your Cosign version:"

  pre_content: |
    First, let's make sure Cosign is available on your system.

  example: "cosign version"

  accepted_commands:
    - "cosign version"
    - "cosign --version"
    - "/usr/local/bin/cosign version"

  validation:
    type: "regex"
    pattern: "GitVersion|cosign.*version"
    case_insensitive: true

  alternative_validations:
    - type: "exit_code"
      expected: 0

  post_content: |
    Great! Cosign is installed and working.

  success_msg: "✅ Cosign v{version} detected!"
  fail_msg: "❌ Cosign not found. Install it first."

  hints:
    - level: 1
      text: "Type: cosign version"
    - level: 2
      text: "If not installed: brew install cosign"

  allow_skip: false
  timeout: 10
```

---

### 3. Type: `quiz`

**Purpose:** Test knowledge with multiple-choice questions.

#### Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | `"quiz"` | Yes | Identifies this as a quiz step |
| `title` | string | No | Quiz section header |
| `questions` | list | Yes | List of questions (see below) |

#### Question Structure

```yaml
questions:
  - id: "q1"                              # Unique ID
    question: "What does Cosign sign?"    # Question text
    type: "multiple_choice"               # Currently only this type
    options:
      - "Docker images"
      - "Git commits"
      - "Kubernetes manifests"
      - "All of the above"
    answer: 0                             # Correct option index (0-based)
    explanation: |                        # Shown after answering
      Cosign primarily signs container images, but can also sign
      other artifacts like blobs and attestations.
```

#### Full Example

```yaml
- type: quiz
  title: "Knowledge Check"
  questions:
    - id: "q1"
      question: "What cryptographic method does Cosign use?"
      type: "multiple_choice"
      options:
        - "Symmetric encryption"
        - "Public key cryptography"
        - "Hash functions only"
        - "OAuth tokens"
      answer: 1
      explanation: |
        Cosign uses **public key cryptography** with key pairs or
        keyless OIDC signing.

    - id: "q2"
      question: "Where are Cosign signatures stored?"
      type: "multiple_choice"
      options:
        - "Local filesystem"
        - "Git repository"
        - "Container registry as OCI artifacts"
        - "Separate database"
      answer: 2
      explanation: |
        Signatures are stored in the same container registry as
        OCI artifacts, tagged with `.sig` suffix.
```

---

### 4. Type: `challenge`

**Purpose:** Open-ended task for users to complete independently.

#### Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | `"challenge"` | Yes | Identifies this as a challenge |
| `title` | string | Yes | Challenge name |
| `description` | string | Yes | Task instructions |
| `verification` | object | Yes | How to verify completion |
| `hints` | list | No | Progressive hints |
| `success_msg` | string | Yes | Success message |
| `allow_skip` | boolean | No | Can skip? (default: true) |

#### Verification

```yaml
verification:
  type: "custom"
  checks:
    - type: "file_exists"
      path: "my-signature.txt"

    - type: "file_contains"
      path: "my-signature.txt"
      pattern: "successfully verified"

    - type: "command_succeeds"
      command: "test -f cosign.key && test -f cosign.pub"
```

#### Full Example

```yaml
- type: challenge
  title: "🚀 Sign Your Own Image"
  description: |
    Now it's your turn! Complete these tasks:

    1. Pick any public image (e.g., `nginx:latest`)
    2. Sign it with your key pair
    3. Verify the signature
    4. Save verification output to `result.txt`

    **Bonus:** Try signing a local Docker image!

  verification:
    type: "custom"
    checks:
      - type: "file_exists"
        path: "result.txt"
      - type: "file_contains"
        path: "result.txt"
        pattern: "successfully verified"

  hints:
    - level: 1
      text: "Use the same commands as before"
    - level: 2
      text: "Redirect output: cosign verify ... > result.txt"

  success_msg: "🎉 Challenge complete! You're a Cosign pro!"
  allow_skip: true
```

---

### 5. Type: `interview_prep`

**Purpose:** Practice open-ended interview questions.

#### Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | `"interview_prep"` | Yes | Identifies this type |
| `title` | string | Yes | Section header |
| `description` | string | No | Instructions |
| `questions` | list | Yes | List of question strings |
| `record_answers` | boolean | No | Save user's answers? (default: false) |
| `export_format` | enum | No | `"json"` or `"text"` (default: `"json"`) |

#### Example

```yaml
- type: interview_prep
  title: "Interview Questions"
  description: |
    Practice these questions to reinforce your learning.
    Your answers will be saved for review.

  questions:
    - "Explain how Cosign ensures container image integrity."
    - "What's the difference between key-based and keyless signing?"
    - "How would you integrate Cosign into a CI/CD pipeline?"
    - "What are the security risks if you lose your private key?"

  record_answers: true
  export_format: "json"
```

---

## Conclusion Section

**Purpose:** Summarize learning and suggest next steps.

### Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `title` | string | Yes | Conclusion header |
| `content` | string | Yes | Summary text (Markdown supported) |
| `badges` | list | No | Badges earned (see below) |

### Badges

```yaml
badges:
  - id: "cosign-basics"
    name: "Cosign Fundamentals"
    icon: "🔐"

  - id: "first-signer"
    name: "First Signature"
    icon: "✍️"
```

### Example

```yaml
conclusion:
  title: "What's Next?"
  content: |
    🎓 **Congratulations!** You've learned:
    - What Cosign is and why signing matters
    - How to generate key pairs
    - How to sign and verify container images

    ### Next Steps:
    - Try the **Keyless Signing** lesson
    - Explore **Cosign with Kubernetes**

    ### Resources:
    - [Cosign Docs](https://docs.sigstore.dev)
    - [Chainguard Academy](https://edu.chainguard.dev)

  badges:
    - id: "cosign-basics"
      name: "Cosign Fundamentals"
      icon: "🔐"
```

---

## Best Practices

### General Guidelines

1. **Keep steps focused** - Each step should teach one concept
2. **Provide context** - Explain *why* before *how*
3. **Use real examples** - Avoid placeholder commands
4. **Test thoroughly** - Validate all commands and validations
5. **Progressive difficulty** - Start simple, build complexity

### Writing Tips

- Use active voice and direct language
- Include emojis sparingly for visual interest
- Break long content into digestible chunks
- Provide hints that guide without giving away answers
- Use callouts for important warnings or tips

### Validation Tips

- Prefer specific validations (regex) over generic ones (exit code)
- Test validations with multiple command variations
- Provide clear error messages that guide users
- Use `alternative_validations` for flexible checking

---

## Complete Example Lesson

```yaml
metadata:
  id: "docker-basics"
  org: "tutorial"
  title: "Docker Basics"
  description: "Learn fundamental Docker commands"
  difficulty: "beginner"
  duration: "15 min"
  tags: ["docker", "containers"]

introduction:
  title: "Welcome to Docker"
  content: |
    Learn the essential Docker commands you'll use every day.

steps:
  - type: info
    title: "What is Docker?"
    content: |
      Docker is a platform for developing, shipping, and running applications
      in containers.

  - type: command
    prompt: "Check Docker version"
    example: "docker --version"
    validation:
      type: "substring"
      contains: "Docker version"
    success_msg: "✅ Docker is installed!"
    fail_msg: "❌ Docker not found"

  - type: quiz
    title: "Quick Check"
    questions:
      - id: "q1"
        question: "What command lists running containers?"
        type: "multiple_choice"
        options:
          - "docker list"
          - "docker ps"
          - "docker show"
        answer: 1
        explanation: "`docker ps` lists running containers"

conclusion:
  title: "Great Job!"
  content: |
    You've learned Docker basics. Keep practicing!
  badges:
    - id: "docker-beginner"
      name: "Docker Beginner"
      icon: "🐳"
```

---

## Support

For questions or issues, please refer to the TryOutShell documentation or contact support.
