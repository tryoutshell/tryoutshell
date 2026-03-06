---
sidebar_position: 5
---

# Challenge Steps

Challenge steps present open-ended tasks where users apply what they've learned independently.

## Purpose

Use challenge steps to:
- Test practical application of concepts
- Encourage creative problem-solving
- Provide a capstone experience at the end of a lesson
- Verify users can complete multi-step tasks

## Basic Structure

```yaml
- type: challenge
  title: "Challenge Name"
  description: |
    Task instructions.
  verification:
    type: "custom"
    checks:
      - type: "file_exists"
        path: "result.txt"
  success_msg: "🎉 Challenge complete!"
  allow_skip: true
```

## Fields Reference

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | `"challenge"` | Yes | Identifies this as a challenge |
| `title` | string | Yes | Challenge name |
| `description` | string | Yes | Task instructions (Markdown) |
| `verification` | object | Yes | How to verify completion |
| `hints` | list | No | Progressive hints |
| `success_msg` | string | Yes | Shown on completion |
| `allow_skip` | boolean | No | Can user skip? (default: true) |

## Verification Checks

```yaml
verification:
  type: "custom"
  checks:
    - type: "file_exists"
      path: "output.txt"

    - type: "file_contains"
      path: "output.txt"
      pattern: "successfully verified"

    - type: "command_succeeds"
      command: "test -f cosign.key"
```

### Available Check Types

| Check Type | Fields | Description |
|-----------|--------|-------------|
| `file_exists` | `path` | File must exist |
| `file_contains` | `path`, `pattern` | File must contain pattern |
| `command_succeeds` | `command` | Command must exit with 0 |

## Full Example

```yaml
- type: challenge
  title: "🚀 Sign Your Own Image"
  description: |
    Now it's your turn! Complete these tasks:

    1. Pick any public image (e.g., `nginx:latest`)
    2. Sign it with your key pair
    3. Verify the signature
    4. Save the verification output to `result.txt`

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
      text: "Use the same commands as before, just change the image name"
    - level: 2
      text: "Redirect output: cosign verify --key cosign.pub <image> > result.txt"
    - level: 3
      text: |
        Full solution:
        cosign sign --key cosign.key nginx:latest
        cosign verify --key cosign.pub nginx:latest > result.txt

  success_msg: "🎉 Challenge complete! You're a Cosign pro!"
  allow_skip: true
```

## Best Practices

- **Place at the end** of a lesson, after all teaching steps
- **Be specific** about what needs to be achieved
- **Always provide hints** — users shouldn't get stuck permanently
- **Verify observable outcomes** — files, command output, not intent
- **Use `allow_skip: true`** — challenges should be optional
- **Number the tasks** — clear 1-2-3 structure helps users plan
