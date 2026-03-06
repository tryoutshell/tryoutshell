---
sidebar_position: 3
---

# Validation Types

Validation determines whether a user's command succeeded. Used in `command` steps of interactive lessons.

## Overview

| Type | Description |
|------|-------------|
| `substring` | Output contains a specific string |
| `regex` | Output matches a regular expression |
| `exit_code` | Command exits with expected code |
| `file_exists` | Specified files exist |
| `file_contains` | File contains a pattern |
| `output_contains` | Output matches one or more patterns |

## substring

Output must contain a specific string.

```yaml
validation:
  type: "substring"
  contains: "Docker version"
  case_insensitive: true     # optional, default: false
```

## regex

Output must match a regular expression.

```yaml
validation:
  type: "regex"
  pattern: "v\\d+\\.\\d+\\.\\d+"
  case_insensitive: false
```

Note: Backslashes must be escaped in YAML (`\\d` for `\d`).

## exit_code

Command must exit with a specific code.

```yaml
validation:
  type: "exit_code"
  expected: 0
```

## file_exists

Specified files must exist after the command runs.

```yaml
validation:
  type: "file_exists"
  files:
    - "cosign.key"
    - "cosign.pub"
```

## file_contains

A file must contain content matching a pattern.

```yaml
validation:
  type: "file_contains"
  path: "output.txt"
  pattern: "success"
```

## output_contains

Output must match one or more patterns with OR/AND logic.

```yaml
# Match ANY pattern (OR logic)
validation:
  type: "output_contains"
  patterns:
    - "Successfully signed"
    - "Pushing signature"
  any_match: true

# Match ALL patterns (AND logic)
validation:
  type: "output_contains"
  patterns:
    - "verified"
    - "cosign.pub"
  all_match: true
```

## Alternative Validations

When the primary validation might miss valid outputs, add fallbacks:

```yaml
validation:
  type: "regex"
  pattern: "cosign.*version"

alternative_validations:
  - type: "exit_code"
    expected: 0
  - type: "substring"
    contains: "GitVersion"
```

If the primary fails, alternatives are tried in order. The step passes if any succeeds.

## Best Practices

- **Use `regex` for flexible matching** — handles version number variations
- **Always add `exit_code` as an alternative** — catches success even if output format changes
- **Use `file_exists` for creation steps** — verifies the right files were produced
- **Be specific with `substring`** — too generic and false positives occur
- **Test your validations** — run the command yourself and verify the patterns match
