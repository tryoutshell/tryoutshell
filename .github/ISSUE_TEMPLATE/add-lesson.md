---
name: 🎓 Add New Lesson
about: Create a new interactive lesson from a blog post or tutorial
title: '[LESSON] Add lesson for: [Topic Name]'
labels: ['lesson', 'content', 'needs-review']
assignees: ['github-copilot']
---

## 📚 Lesson Information

### Source Material
**Blog/Tutorial URL:**
<!-- Paste the URL of the blog post, tutorial, or documentation you want to convert into a lesson -->


**Topic/Technology:**
<!-- e.g., Kubernetes, Cosign, Docker, Terraform, etc. -->


**Estimated Duration:**
<!-- e.g., 15 min, 30 min, 1 hour -->


---

## 🎯 Lesson Details

### Organization
**Organization Name:**
<!-- e.g., sigstore, kubernetes, hashicorp, chainguard -->


### Difficulty Level
<!-- Check one -->
- [ ] Beginner
- [ ] Intermediate
- [ ] Advanced

### Prerequisites
<!-- List any required knowledge, tools, or setup (one per line) -->
-
-
-

### Tags
<!-- Add relevant tags for searchability (comma-separated) -->
<!-- e.g., docker, containers, signing, security, CI/CD -->


---

## 📝 Content Preferences

### Step Types to Include
<!-- Check all that apply -->
- [ ] Informational steps (concepts and explanations)
- [ ] Command exercises (hands-on practice)
- [ ] Quiz questions (knowledge checks)
- [ ] Challenges (open-ended tasks)
- [ ] Interview prep questions

### Special Requirements
<!-- Any specific requirements or notes for the lesson -->
- [ ] Include code examples with syntax highlighting
- [ ] Add diagrams or ASCII art where helpful
- [ ] Include security best practices callouts
- [ ] Add progressive hints for difficult commands
- [ ] Include real-world use cases

---

## 🤖 Instructions for GitHub Copilot

**Please generate a complete lesson YAML file following the TryOutShell format.**

### 📖 Documentation Reference

The complete lesson format documentation is available in the `docs/` directory:

```
docs/
├── examples/
│   ├── complete-lesson.md          # Full lesson example
│   └── step-by-step-tutorial.md    # Tutorial on creating lessons
├── getting-started/
│   ├── index.md                    # Overview and quick start
│   ├── lesson-structure.md         # Core structure explained
│   ├── minimal-example.md          # Simplest valid lesson
│   └── quick-start.md             # Getting started guide
├── guides/
│   ├── best-practices.md          # Writing guidelines
│   ├── hints-and-callouts.md      # Using hints effectively
│   ├── markdown-support.md        # Markdown formatting
│   └── validation-types.md        # Command validation methods
├── metadata/
│   ├── index.md                   # Metadata overview
│   ├── optional-fields.md         # Optional metadata fields
│   └── required-fields.md         # Required metadata fields
├── sections/
│   ├── conclusion.md              # Conclusion section format
│   └── introduction.md            # Introduction section format
└── step-types/
    ├── challenge-steps.md         # Challenge step format
    ├── command-steps.md           # Command step format
    ├── index.md                   # Step types overview
    ├── info-steps.md              # Info step format
    ├── interview-prep-steps.md    # Interview prep format
    └── quiz-steps.md              # Quiz step format
```

**Please review these documentation files** to ensure the generated lesson follows all format specifications, best practices, and validation patterns.

### Requirements:

1. **Parse the provided blog/tutorial** and extract:
   - Main concepts and learning objectives
   - Step-by-step instructions
   - Commands to execute
   - Code examples and explanations

2. **Structure the lesson** with:
   - Complete metadata section (id, org, title, description, difficulty, duration, tags)
   - Introduction section with clear learning objectives
   - 8-15 well-structured steps including:
     - `info` steps for concepts (with highlights, code_blocks, callouts, diagrams)
     - `command` steps for hands-on practice (with validation, hints, success/fail messages)
     - `quiz` steps for knowledge checks (2-3 questions minimum)
     - At least one `challenge` step for open-ended practice
   - Conclusion section with next steps and resources

3. **Follow best practices**:
   - Each command step must have:
     - Realistic validation (prefer regex/output_contains over exit_code)
     - 3 progressive hints
     - Alternative command variations in `accepted_commands`
     - Pre-content and post-content explanations
   - Include security warnings where relevant
   - Add emojis sparingly for visual interest
   - Use proper Markdown formatting in content fields
   - Include ASCII diagrams for complex concepts
   - Add badges to the conclusion section

4. **File naming**:
   - Create file as: `lessons/<org>/<topic-name>.yaml`
   - Use kebab-case for the filename
   - Example: `lessons/sigstore/cosign-sign-verify.yaml`

5. **Validation rules**:
   - Test all regex patterns for command validation
   - Ensure all file paths are relative
   - Verify all URLs are accessible
   - Check that quiz answers match option indices (0-based)

6. **Reference the example**:
   - Use `lessons/sigstore/cosign-sign-verify.yaml` as a quality benchmark
   - Match the tone and depth of explanations
   - Include similar levels of context and hints

---

## ✅ Acceptance Criteria

The generated lesson YAML must:
- [ ] Follow the exact format specified in the documentation
- [ ] Include all required metadata fields
- [ ] Have clear, actionable steps with proper validation
- [ ] Include at least one quiz and one challenge
- [ ] Contain helpful hints and explanations
- [ ] Be tested for syntax errors
- [ ] Include proper badges and resources in conclusion
- [ ] Use realistic examples (no placeholders)

---

## 🔗 Additional Context

<!-- Add any additional context, screenshots, or notes that might help -->


---

**Note:** Once this issue is created, GitHub Copilot will automatically generate the lesson YAML file and create a pull request for review. The maintainer will review and merge after testing.
