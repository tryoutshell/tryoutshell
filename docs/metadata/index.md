---
sidebar_position: 1
---

# Metadata

Lesson metadata controls how lessons appear in listings and search.

See the [full fields reference](./required-fields) for details on all available fields.

## Quick Reference

### Data-Only Lessons (lesson.yaml)

```yaml
id: "docker-101"
title: "Docker 101"
description: "Learn Docker fundamentals"
difficulty: "beginner"
duration: "25 min"
version: "1.0"
author: "TryOutShell"
tags: ["docker", "containers"]
```

### Interactive Lessons (metadata section)

```yaml
metadata:
  id: "cosign-sign-verify"
  org: "sigstore"
  title: "Container Image Signing"
  description: "Sign and verify container images"
  difficulty: "beginner"
  duration: "20 min"
  tags: ["cosign", "signing"]
  author: "TryOutShell"
  version: "1.0"
```

### Organization (meta.yaml)

```yaml
id: docker
name: "Docker"
description: "Container runtime and tooling"
logo: "🐳"
```
