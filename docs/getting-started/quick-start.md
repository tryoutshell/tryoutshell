---
sidebar_position: 2
---

# Quick Start

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

### Lesson Structure
Every lesson YAML file has 4 main sections:

```
metadata:        # Required - Lesson identification and metadata
introduction:    # Optional - Shown once at the start
steps:           # Required - Main lesson content
conclusion:      # Optional - Shown after completion
```

### Metadata Section

Required fields:
FieldTypeDescriptionExampleidstringUnique identifier (kebab-case)"cosign-basics"orgstringOrganization name"chainguard"titlestringHuman-readable title"Introduction to Cosign"descriptionstringBrief description (1-2 sentences)"Learn to sign containers"difficultyenumbeginner, intermediate, advanced"beginner"durationstringEstimated time"15 min"tagslistKeywords for search["cosign", "signing"]
Optional fields:
FieldTypeDescriptionExampleprerequisiteslistTools/knowledge needed["Docker installed"]authorstringCreator name"Chainguard Team"versionstringLesson version"1.0"updated_atdateLast update"2025-01-15"

```
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

### Introduction Section
Purpose: Set expectations and provide context before diving into steps.

Fields:
FieldTypeRequiredDescriptiontitlestringYesIntroduction headercontentstring (multiline)YesOverview text (supports Markdown)

```yaml
introduction:
  title: "Introduction"
  content: "Welcome to the Container Image Signing with Cosign lesson!"
```

###Markdown Support:

`**bold**` → bold text
`code` → inline code
`### Heading` → subheadings
`> Quote` → blockquotes
`- List item` → bullet lists

Example:
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

### Steps
Structure: Each step is a YAML object in the steps list.
```yaml
steps:
  - type: info           # Step type (see below)
    title: "Step Title"
    # ... type-specific fields

  - type: command
    prompt: "Do something"
    # ... type-specific fields
```

### Steps
Structure: Each step is a YAML object in the steps list.

1. Type: info
Purpose: Display educational content (text, diagrams, explanations).

Fields:
FieldTypeRequiredDescriptiontype"info"YesIdentifies this as an info steptitlestringYesStep headingcontentstring (multiline)YesMain text (Markdown supported)highlightslistNoInline text highlights (see below)code_blockslistNoCode examples (see below)calloutslistNoTips/warnings (see below)diagramstringNoASCII art diagramwait_for_continuebooleanNoPause until user presses Enter (default: true)

```yaml
highlights:
  - text: "myapp:latest"      # Text to highlight
    style: "code"             # Style: "code", "bold", "highlight"
  - text: "Sigstore"
    style: "bold"
```
## Code Blocks:

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

## Callouts:

```yaml
callouts:
  - type: "tip"       # Types: "tip", "warning", "danger", "info"
    text: "Always verify signatures in production"

  - type: "warning"
    text: "Never commit private keys to Git!"
```

## Full Example:
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

### 2. Type: command
Purpose: Execute a shell command and verify the output.

Fields:
FieldTypeRequiredDescriptiontype"command"YesIdentifies this as a command stepidstringNoUnique ID for tracking progresspromptstringYesShort description (shown in UI)instructionstringNoDetailed instruction textexamplestringYesExample command to show userpre_contentstringNoExplanation before commandpost_contentstringNoExplanation after successvalidationobjectYesHow to verify success (see below)alternative_validationslistNoAdditional validation methodsaccepted_commandslistNoList of valid command variationssuccess_msgstringYesMessage on successfail_msgstringYesMessage on failurehintslistNoProgressive hints (see below)allow_skipbooleanNoCan user skip this step? (default: false)timeoutintegerNoMax execution time in seconds (default: 30)


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

Hints:
```yaml
hints:
  - level: 1
    text: "Try typing: cosign version"

  - level: 2
    text: "Install with: brew install cosign"

  - level: 3
    text: "Full command: cosign version"
```

### Full Example
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

3. Type: quiz
Purpose: Test knowledge with multiple-choice questions.
Fields:
FieldTypeRequiredDescriptiontype"quiz"YesIdentifies this as a quiz steptitlestringNoQuiz section headerquestionslistYesList of questions (see below)

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

### Full Example
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

4. Type: challenge
Purpose: Open-ended task for users to complete independently.
Fields:
FieldTypeRequiredDescriptiontype"challenge"YesIdentifies this as a challengetitlestringYesChallenge namedescriptionstringYesTask instructionsverificationobjectYesHow to verify completionhintslistNoProgressive hintssuccess_msgstringYesSuccess messageallow_skipbooleanNoCan skip? (default: true)

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

### Full Example
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

5. Type: interview_prep
Fields:
FieldTypeRequiredDescriptiontype"interview_prep"YesIdentifies this typetitlestringYesSection headerdescriptionstringNoInstructionsquestionslistYesList of question stringsrecord_answersbooleanNoSave user's answers? (default: false)export_formatenumNo"json" or "text" (default: "json")

Purpose: Practice open-ended interview questions.
```
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

### Conclusion Section
Purpose: Summarize learning and suggest next steps.
Fields:
FieldTypeRequiredDescriptiontitlestringYesConclusion headercontentstringYesSummary text (Markdown supported)badgeslistNoBadges earned (see below)

```
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
