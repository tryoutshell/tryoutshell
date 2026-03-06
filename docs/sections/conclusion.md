---
sidebar_position: 2
---

# Conclusion Section

The conclusion section is shown after all steps are completed. It summarizes learning and suggests next steps.

## Structure

```yaml
conclusion:
  title: "Congratulations!"
  content: |
    Summary and next steps.
  badges:
    - id: "badge-id"
      name: "Badge Name"
      icon: "🏆"
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `title` | string | Yes | Conclusion header |
| `content` | string | Yes | Summary text (Markdown supported) |
| `badges` | list | No | Achievement badges earned |

### Badge Fields

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique badge identifier |
| `name` | string | Display name |
| `icon` | string | Emoji icon |

## Example

```yaml
conclusion:
  title: "What's Next?"
  content: |
    🎓 **Congratulations!** You've learned:
    - What Cosign is and why signing matters
    - How to generate key pairs
    - How to sign and verify container images
    - Best practices for supply chain security

    ### Next Steps:
    - Try the **Keyless Signing** lesson
    - Explore **Cosign with Kubernetes**
    - Learn about **Rekor transparency logs**

    ### Resources:
    - [Cosign Documentation](https://docs.sigstore.dev)
    - [Sigstore Blog](https://blog.sigstore.dev)

  badges:
    - id: "cosign-basics"
      name: "Cosign Fundamentals"
      icon: "🔐"
    - id: "first-signer"
      name: "First Signature"
      icon: "✍️"
```

## Best Practices

- **Summarize key learnings** — list what was covered
- **Suggest next steps** — link to related lessons or external resources
- **Award badges** — recognition motivates continued learning
- **Keep it positive** — celebrate the completion
- **Include resource links** — documentation, repos, further reading
