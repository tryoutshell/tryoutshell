# 🤖 TryOutShell Lesson Generator Prompt

Copy this entire prompt and paste it into any AI assistant (Claude, ChatGPT, etc.) along with your blog/tutorial link to generate a complete lesson YAML file.

---

## PROMPT START

I need you to create an interactive lesson YAML file for TryOutShell based on a blog post or tutorial. TryOutShell is an interactive CLI learning platform where users execute real commands and get instant feedback.

### 📚 SOURCE MATERIAL

**Blog/Tutorial URL:** [PASTE YOUR URL HERE]

**Additional Context:** [Add any specific requirements, focus areas, or notes]

---

### 📋 LESSON FORMAT SPECIFICATION

You MUST follow this exact YAML structure. Here's the complete format:

```yaml
metadata:
  id: "topic-name"                    # kebab-case, unique identifier
  org: "organization"                 # e.g., kubernetes, docker, sigstore
  title: "Human Readable Title"
  description: "Brief 1-2 sentence description"
  difficulty: "beginner"              # beginner | intermediate | advanced
  duration: "20 min"                  # Estimated time
  prerequisites:                      # Optional
    - "Tool or knowledge needed"
  tags: ["tag1", "tag2"]             # For search
  author: "Author Name"               # Optional
  version: "1.0"                      # Optional

introduction:
  title: "What You'll Learn"
  content: |
    Clear introduction with:
    - Learning objectives (bullet points)
    - What users will build/understand
    - Required tools
    - Time estimate

steps:
  # Step 1: Info step (concept explanation)
  - type: info
    title: "Concept Title"
    content: |
      Detailed explanation using Markdown:
      - Use **bold** for emphasis
      - Use `code` for commands/terms
      - Use > for blockquotes
      - Keep paragraphs short

    highlights:                       # Optional: highlight specific terms
      - text: "important-term"
        style: "code"                 # code | bold | highlight

    code_blocks:                      # Optional: show code examples
      - label: "Example Label"
        code: "command or code here"
        language: "bash"              # bash | python | yaml | etc

    callouts:                         # Optional: tips/warnings
      - type: "tip"                   # tip | warning | danger | info
        text: "Helpful information"

    diagram: |                        # Optional: ASCII diagram
      Simple ASCII art diagram

    wait_for_continue: true           # User must press Enter

  # Step 2: Command step (hands-on execution)
  - type: command
    id: "unique-step-id"
    prompt: "Short description"
    instruction: "Detailed instruction"  # Optional

    pre_content: |                    # Explain BEFORE command
      Context about what we're doing

    example: "actual command here"    # Show this to user

    accepted_commands:                # Alternative valid commands
      - "command variation 1"
      - "command variation 2"

    validation:
      type: "regex"                   # See validation types below
      pattern: "expected.*output"
      case_insensitive: true

    alternative_validations:          # Fallback validations
      - type: "exit_code"
        expected: 0

    post_content: |                   # Explain AFTER success
      What just happened and why

    success_msg: "✅ Success message!"
    fail_msg: "❌ Failure guidance"

    hints:                            # Progressive help
      - level: 1
        text: "Gentle nudge"
      - level: 2
        text: "More specific help"
      - level: 3
        text: "Almost the answer"

    allow_skip: false                 # Can user skip?
    timeout: 30                       # Max execution time

  # Step 3: Quiz step (knowledge check)
  - type: quiz
    title: "Knowledge Check"
    questions:
      - id: "q1"
        question: "Question text?"
        type: "multiple_choice"
        options:
          - "Option A"
          - "Option B"
          - "Option C"
        answer: 1                     # 0-based index
        explanation: |
          Why this is correct

  # Step 4: Challenge step (open-ended task)
  - type: challenge
    title: "🚀 Challenge Title"
    description: |
      Task description with:
      1. Clear objectives
      2. What to create/do
      3. Success criteria

    verification:
      type: "custom"
      checks:
        - type: "file_exists"
          path: "expected-file.txt"
        - type: "file_contains"
          path: "expected-file.txt"
          pattern: "expected content"

    hints:
      - level: 1
        text: "Hint for the challenge"

    success_msg: "🎉 Challenge complete!"
    allow_skip: true

conclusion:
  title: "What's Next?"
  content: |
    Summary of what was learned

    ### Next Steps:
    - Suggested next lessons
    - Advanced topics

    ### Resources:
    - [Link](url)

  badges:                             # Optional achievements
    - id: "badge-id"
      name: "Badge Name"
      icon: "🏆"
```

---

### 🎯 VALIDATION TYPES REFERENCE

Use these validation types in `command` steps:

1. **substring** - Check if output contains text
```yaml
validation:
  type: "substring"
  contains: "expected text"
  case_insensitive: true
```

2. **regex** - Match pattern (PREFERRED for most commands)
```yaml
validation:
  type: "regex"
  pattern: "version.*\\d+\\.\\d+"
  case_insensitive: false
```

3. **exit_code** - Check command exit status
```yaml
validation:
  type: "exit_code"
  expected: 0
```

4. **file_exists** - Verify files were created
```yaml
validation:
  type: "file_exists"
  files:
    - "file1.txt"
    - "file2.txt"
```

5. **file_contains** - Check file content
```yaml
validation:
  type: "file_contains"
  path: "output.txt"
  pattern: "success"
```

6. **output_contains** - Multiple pattern matching
```yaml
validation:
  type: "output_contains"
  patterns:
    - "pattern1"
    - "pattern2"
  any_match: true    # OR logic
  all_match: false   # AND logic
```

---

### ✅ REQUIREMENTS & BEST PRACTICES

**MUST HAVE:**
1. ✅ 8-15 well-structured steps
2. ✅ Mix of info, command, quiz, and challenge steps
3. ✅ At least 2-3 quiz questions
4. ✅ At least 1 challenge step
5. ✅ Every command step must have:
   - Realistic validation (prefer regex over exit_code)
   - 3 progressive hints
   - Pre and post content explanations
   - Clear success/fail messages
6. ✅ Include security warnings where relevant
7. ✅ Use emojis sparingly (✅ ❌ 🎯 💡 ⚠️)
8. ✅ Proper Markdown formatting
9. ✅ ASCII diagrams for complex concepts
10. ✅ All URLs must be real and accessible

**CONTENT STYLE:**
- Write conversationally but professionally
- Explain WHY before HOW
- Use analogies for complex concepts
- Keep paragraphs short (3-4 lines max)
- Use bullet points for lists
- Include real-world use cases
- Add context about when to use each tool/command

**VALIDATION RULES:**
- Test regex patterns for realistic output
- Allow command variations in `accepted_commands`
- Use `alternative_validations` for flexibility
- Set realistic timeouts (10-60 seconds)
- Provide helpful error messages

**SECURITY:**
- Add warnings for dangerous commands
- Never ask users to commit secrets
- Explain security implications
- Use callouts for important security notes

---

### 📖 EXAMPLE LESSON

Here's a complete example to follow as a quality benchmark:

```yaml
metadata:
  id: "docker-basics"
  org: "docker"
  title: "Docker Fundamentals"
  description: "Learn essential Docker commands and container management"
  difficulty: "beginner"
  duration: "25 min"
  prerequisites:
    - "Docker installed"
  tags: ["docker", "containers", "beginner"]
  author: "TryOutShell Team"
  version: "1.0"

introduction:
  title: "Welcome to Docker"
  content: |
    In this lesson, you'll learn:
    - What containers are and why they matter
    - How to run and manage Docker containers
    - Basic Docker commands you'll use daily

    **Prerequisites:** Docker Desktop installed
    **Time:** ~25 minutes

steps:
  - type: info
    title: "What are Containers?"
    content: |
      **Containers** are lightweight, isolated environments that package
      your application with all its dependencies.

      Think of them like shipping containers for software:
      - **Consistent** - Run anywhere (dev, test, prod)
      - **Isolated** - Don't interfere with other apps
      - **Lightweight** - Share OS kernel, start in seconds

      > 💡 **Fun fact:** Docker containers use Linux namespaces and cgroups
      > for isolation, making them much faster than VMs!

    diagram: |
      ┌─────────────────────────────────────┐
      │         Your Application            │
      │  ┌──────────┐  ┌──────────────┐    │
      │  │  Code    │  │ Dependencies │    │
      │  └──────────┘  └──────────────┘    │
      └─────────────────────────────────────┘
      ────────────── Container ──────────────
      ┌─────────────────────────────────────┐
      │         Docker Engine               │
      └─────────────────────────────────────┘
      ┌─────────────────────────────────────┐
      │         Host OS (Linux)             │
      └─────────────────────────────────────┘

    wait_for_continue: true

  - type: command
    id: "check-docker"
    prompt: "Verify Docker is running"

    pre_content: |
      Let's make sure Docker is installed and running on your system.

    example: "docker --version"

    accepted_commands:
      - "docker --version"
      - "docker -v"
      - "docker version"

    validation:
      type: "regex"
      pattern: "Docker version \\d+\\.\\d+"
      case_insensitive: true

    alternative_validations:
      - type: "exit_code"
        expected: 0

    post_content: |
      ✅ Docker is installed! You're ready to start working with containers.

    success_msg: "✅ Docker version detected!"
    fail_msg: "❌ Docker not found. Please install Docker Desktop."

    hints:
      - level: 1
        text: "Try: docker --version"
      - level: 2
        text: "If not installed, download from docker.com"
      - level: 3
        text: "Make sure Docker Desktop is running"

    timeout: 10

  - type: command
    id: "run-hello-world"
    prompt: "Run your first container"

    pre_content: |
      Let's run the classic "Hello World" container. Docker will:
      1. Download the image (if not cached)
      2. Create a container
      3. Run the command
      4. Show output

    example: "docker run hello-world"

    validation:
      type: "output_contains"
      patterns:
        - "Hello from Docker"
        - "This message shows that your installation"
      any_match: true

    post_content: |
      🎉 Congratulations! You just ran your first Docker container.

      Behind the scenes:
      - Docker pulled the `hello-world` image from Docker Hub
      - Created a new container from that image
      - Executed the container's command
      - Showed you the output

    success_msg: "✅ First container executed successfully!"
    fail_msg: "❌ Container failed to run. Check Docker daemon."

    hints:
      - level: 1
        text: "Use: docker run hello-world"
      - level: 2
        text: "Make sure Docker Desktop is running"
      - level: 3
        text: "Try: sudo docker run hello-world (Linux)"

    timeout: 60

  - type: quiz
    title: "Quick Check"
    questions:
      - id: "q1"
        question: "What makes containers faster than VMs?"
        type: "multiple_choice"
        options:
          - "They use more RAM"
          - "They share the host OS kernel"
          - "They don't need CPU"
          - "They run on cloud only"
        answer: 1
        explanation: |
          Containers share the host OS kernel, so they don't need to boot
          an entire OS like VMs do. This makes them start in seconds!

  - type: challenge
    title: "🚀 Run a Web Server"
    description: |
      Your challenge: Run an nginx web server container!

      Tasks:
      1. Run nginx in detached mode
      2. Map port 8080 to container port 80
      3. Verify it's running with `docker ps`
      4. Save the container ID to `nginx-container.txt`

      **Hint:** Use `docker run -d -p 8080:80 nginx`

    verification:
      type: "custom"
      checks:
        - type: "file_exists"
          path: "nginx-container.txt"
        - type: "command_succeeds"
          command: "docker ps | grep nginx"

    success_msg: "🎉 You're running a web server in Docker!"
    allow_skip: true

conclusion:
  title: "Next Steps"
  content: |
    🎓 You've learned Docker fundamentals!

    ### What's Next:
    - Build custom images with Dockerfile
    - Use Docker Compose for multi-container apps
    - Learn about volumes and networking

    ### Resources:
    - [Docker Docs](https://docs.docker.com)
    - [Docker Hub](https://hub.docker.com)

  badges:
    - id: "docker-basics"
      name: "Docker Fundamentals"
      icon: "🐳"
```

---

### 🎬 YOUR TASK

Based on the blog/tutorial URL I provided:

1. **Read and analyze** the entire content
2. **Extract key concepts** and learning objectives
3. **Identify practical commands** users should execute
4. **Create a complete YAML lesson** following the format above
5. **Ensure validation** is realistic and matches actual command output
6. **Add context** - explain WHY, not just HOW
7. **Include security notes** where relevant
8. **Test mentally** - would a beginner understand each step?

**Output Format:**
- Provide ONLY the complete YAML file
- No additional commentary
- Ready to save as `lessons/<org>/<topic>.yaml`
- Ensure proper YAML indentation (2 spaces)

**Quality Checklist:**
- [ ] 8-15 steps total
- [ ] Mix of info, command, quiz, challenge steps
- [ ] Every command has realistic validation
- [ ] 3 hints per command step
- [ ] Pre/post content explanations
- [ ] 2-3 quiz questions minimum
- [ ] At least 1 challenge
- [ ] Security warnings included
- [ ] ASCII diagrams for complex concepts
- [ ] Proper Markdown formatting
- [ ] Badges in conclusion

---

## PROMPT END

---

### 📝 HOW TO USE THIS PROMPT

1. **Copy everything** from "PROMPT START" to "PROMPT END"
2. **Replace** `[PASTE YOUR URL HERE]` with your blog/tutorial URL
3. **Add any additional context** in the "Additional Context" section
4. **Paste into** Claude, ChatGPT, or any AI assistant
5. **Review the output** YAML file
6. **Test commands** to ensure validation works
7. **Save as** `lessons/<org>/<topic-name>.yaml`

### 🎯 Example Usage

```
I need you to create an interactive lesson YAML file for TryOutShell...

📚 SOURCE MATERIAL

Blog/Tutorial URL: https://kubernetes.io/docs/tutorials/hello-minikube/

Additional Context: Focus on basics, make it beginner-friendly, include
troubleshooting hints
```

---

### ✨ Tips for Best Results

- **Provide complete URLs** - AI needs access to the content
- **Specify difficulty** - Mention if it should be beginner/intermediate/advanced
- **Note special requirements** - Security focus, specific tools, etc.
- **Review validation** - Always test the regex patterns match real output
- **Adjust tone** - Ask for more/less technical language if needed

---

**Questions?** Check the TryOutShell documentation at `docs/` or create an issue!
