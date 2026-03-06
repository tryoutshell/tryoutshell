---
sidebar_position: 2
---

# Optional Metadata Fields

These fields enhance lesson discoverability and provide additional context.

## Data-Only Lessons

| Field | Type | Description | Example |
|-------|------|-------------|---------|
| `author` | string | Creator name | `"TryOutShell"` |
| `tags` | list | Searchable keywords | `["docker", "containers"]` |
| `quiz` | list | Quiz questions | See [Quiz Steps](../step-types/quiz-steps) |

## Interactive Lessons

| Field | Type | Description | Example |
|-------|------|-------------|---------|
| `prerequisites` | list | Required tools/knowledge | `["Docker installed"]` |
| `author` | string | Creator name | `"TryOutShell"` |
| `version` | string | Lesson version | `"1.0"` |
| `resources` | list | External links | See below |

### Resources

```yaml
resources:
  - title: "Official Documentation"
    url: "https://docs.example.com"
    type: "docs"        # docs | video | tutorial | github | blog
  - title: "Source Code"
    url: "https://github.com/example/project"
    type: "github"
```

## Organization meta.yaml

| Field | Type | Description |
|-------|------|-------------|
| `description` | string | Brief description of the org/topic |
| `logo` | string | Emoji or text icon |
