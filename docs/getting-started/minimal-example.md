---
sidebar_position: 2
---

# Minimal Lesson Example

Let's create the simplest possible TryOutShell lesson to understand the basics.

## The Minimal Lesson

Create a file named `hello-world.yaml`:
```yaml
metadata:
  id: "hello-world"
  org: "tutorial"
  title: "Hello World"
  description: "Your first TryOutShell lesson"
  difficulty: "beginner"
  duration: "5 min"
  tags: ["intro", "basics"]

steps:
  - type: info
    title: "Welcome!"
    content: "Hello, world! This is your first lesson."

  - type: command
    prompt: "Say hello"
    example: "echo 'Hello from TryOutShell'"
    validation:
      type: "substring"
      contains: "Hello"
    success_msg: "✅ Perfect! You said hello!"
    fail_msg: "❌ Try running the example command"
```

## Breaking It Down

### Metadata Section
```yaml
metadata:
  id: "hello-world"           # Unique identifier
  org: "tutorial"             # Organization name
  title: "Hello World"        # Display name
  description: "Your first TryOutShell lesson"
  difficulty: "beginner"      # beginner | intermediate | advanced
  duration: "5 min"           # Estimated time
  tags: ["intro", "basics"]   # Search keywords
```

Every lesson **must** have metadata with these required fields.

### Steps Section

Lessons consist of a series of **steps**. Each step has a `type` that determines its behavior.

#### Info Step
```yaml
- type: info
  title: "Welcome!"
  content: "Hello, world! This is your first lesson."
```

Displays informational content to the user. Great for explanations and context.

#### Command Step
```yaml
- type: command
  prompt: "Say hello"                          # What to do
  example: "echo 'Hello from TryOutShell'"    # Example command
  validation:                                  # How to check success
    type: "substring"
    contains: "Hello"
  success_msg: "✅ Perfect! You said hello!"   # Success feedback
  fail_msg: "❌ Try running the example command"
```

Asks the user to run a command and validates the output.

## Running Your Lesson

1. Save the lesson to `~/.tryoutshell/lessons/hello-world.yaml`

2. Run it:
```bash
   tryoutshell start hello-world
```

3. Follow the interactive prompts in your terminal!

## What Happens?

1. TryOutShell displays the **info step** with the welcome message
2. User presses Enter to continue
3. TryOutShell shows the **command step** with the example
4. User types the command (or a similar one)
5. TryOutShell validates the output
6. Success or failure message is displayed

## Try It Yourself

Modify the lesson to:
- Change the welcome message
- Require a different string in the output
- Add another command step

Example:
```yaml
- type: command
  prompt: "Check the date"
  example: "date"
  validation:
    type: "exit_code"
    expected: 0
  success_msg: "✅ Date command works!"
  fail_msg: "❌ Something went wrong"
```

## Next Steps

- 📚 [Learn about Lesson Structure](./lesson-structure) - Understand all sections
- 🎯 [Explore Step Types](../step-types/) - Info, Command, Quiz, Challenge, and more
- 💡 [See Complete Examples](../examples/) - Full-featured lessons

## Common Questions

### Where do I save lessons?

Place them in:
- `~/.tryoutshell/lessons/` (user lessons)
- Or create a custom repository

### Can I use markdown?

Yes! The `content` field supports full markdown:
```yaml
content: |
  **Bold text**

  - Bullet points
  - Work great

  `Inline code` is supported too!
```

### How do I test my lesson?

Run it locally:
```bash
tryoutshell start my-lesson
```

Or validate the YAML:
```bash
tryoutshell validate my-lesson.yaml
```

---

Ready to dive deeper? Check out the [Lesson Structure](./lesson-structure) guide!
