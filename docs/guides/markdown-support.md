---
sidebar_position: 2
---

# Markdown Support

TryOutShell renders Markdown in slides and lesson content using Glamour (terminal Markdown renderer).

## Supported Syntax

### Headers

```markdown
# H1 Title
## H2 Section
### H3 Subsection
```

### Emphasis

```markdown
**bold text**
*italic text*
***bold and italic***
~~strikethrough~~
```

### Lists

```markdown
- Unordered item
- Another item
  - Nested item

1. First item
2. Second item
3. Third item
```

### Code

Inline: `` `docker ps` ``

Block with syntax highlighting:

````markdown
```bash
docker run -it ubuntu:latest /bin/bash
```

```yaml
apiVersion: v1
kind: Pod
```

```python
def hello():
    print("Hello from TryOutShell!")
```
````

Supported languages: `bash`, `yaml`, `json`, `go`, `python`, `javascript`, `typescript`, `dockerfile`, `sql`, `html`, `css`, and many more.

### Tables

```markdown
| Command | Description |
|---------|-------------|
| `docker ps` | List containers |
| `docker images` | List images |
```

### Blockquotes

```markdown
> This is a blockquote.

> 💡 **Tip**: Use emoji for visual callouts.

> ⚠️ **Warning**: Be careful with this command.
```

### Links

```markdown
[Link text](https://example.com)
```

### Horizontal Rules

```markdown
---
```

In `slides.md`, `---` is also the slide separator.

## Tips for Terminal Rendering

- **Keep lines under 80 characters** — terminals vary in width
- **Use tables sparingly** — they can overflow in narrow terminals
- **Prefer code blocks over inline code** for multi-word commands
- **Use blockquotes with emoji** for callouts (💡 tips, ⚠️ warnings)
- **Test in your terminal** — rendering can differ from browser Markdown
