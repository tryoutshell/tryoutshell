# The .git Directory Structure

When you run `git init`, Git creates a `.git` directory that **is** the repository. Everything Git knows lives here.

## Anatomy of .git

```
.git/
├── HEAD              # Points to the current branch (or commit)
├── config            # Repository-level configuration
├── description       # Used by GitWeb (rarely touched)
├── index             # The staging area (binary file)
├── hooks/            # Client-side and server-side hook scripts
├── info/
│   └── exclude       # Like .gitignore but not committed
├── objects/          # THE object database — all your data
│   ├── pack/         # Compressed (packed) objects
│   ├── info/
│   ├── ab/           # Loose objects stored by first 2 chars of SHA
│   └── ...
└── refs/             # Branch and tag pointers
    ├── heads/        # Local branch refs (e.g., refs/heads/main)
    ├── tags/         # Tag refs
    └── remotes/      # Remote-tracking branch refs
```

## The Key Insight

Git is fundamentally a **content-addressable filesystem** — a key-value store where:

- **Key** = SHA-1 hash of the content
- **Value** = the compressed content itself

Every file, directory snapshot, and commit is stored as an object identified by its SHA-1 hash. If two files have identical content, they share the same object — automatic deduplication.

## Try It Yourself

```bash
# Peek inside the objects directory after a commit
find .git/objects -type f | head -10

# Look at what HEAD contains
cat .git/HEAD
# output: ref: refs/heads/main

# See what branch 'main' points to
cat .git/refs/heads/main
# output: a3f5b8c9d2e1f4a6b7c8d9e0f1a2b3c4d5e6f7a8
```

---

# Blobs, Trees, Commits, and Tags

Git has exactly four object types. Understanding them unlocks everything.

## 1. Blob (Binary Large Object)

A blob stores the **raw content** of a single file. It does NOT store the filename or permissions — just the bytes.

```bash
# Create a blob manually
echo "Hello, Git!" | git hash-object -w --stdin
# Returns: e965047ad7c57865823c7d992b1d046ea66edf78

# Read the blob back
git cat-file -p e965047
# Output: Hello, Git!

# Check the type
git cat-file -t e965047
# Output: blob
```

Two files with the same content → same blob hash → stored only once.

## 2. Tree

A tree represents a **directory snapshot**. It maps filenames and permissions to blob (file) or tree (subdirectory) references.

```bash
git cat-file -p main^{tree}
# Output:
# 100644 blob a3f5b8...  README.md
# 100644 blob b2c4d6...  app.py
# 040000 tree c3d5e7...  src
```

Think of a tree as a directory listing: each entry has a mode, type, hash, and name.

## 3. Commit

A commit ties everything together. It contains:

```bash
git cat-file -p HEAD
# Output:
# tree     9a8b7c...              ← pointer to the root tree (snapshot)
# parent   f1e2d3...              ← pointer to parent commit(s)
# author   Jane <jane@x.com> 1709...  ← who wrote the change
# committer Jane <jane@x.com> 1709... ← who committed it
#
# Fix the login bug                ← commit message
```

A commit is a snapshot pointer (tree) + metadata + parent link(s) that form the history chain.

## 4. Annotated Tag

An annotated tag is a full object that points to a commit with additional metadata:

```bash
git cat-file -p v1.0
# Output:
# object  a1b2c3...    ← the commit being tagged
# type    commit
# tag     v1.0
# tagger  Jane <jane@x.com> 1709...
#
# Release version 1.0
```

Lightweight tags are just refs (pointers) — annotated tags are stored as objects with a message and signature.

## The Object Graph

```
          Tag (v1.0)
            │
            ▼
Commit ◄── Commit ◄── Commit  (HEAD)
  │          │           │
  ▼          ▼           ▼
 Tree       Tree        Tree
 / \        / \         / \
Blob Blob  Blob Tree   Blob Blob
                 |
                Blob
```

---

# How git add and git commit Work Under the Hood

Understanding the internal mechanics transforms Git from a magic black box into a predictable tool.

## What git add Really Does

When you run `git add README.md`, three things happen:

### Step 1: Create a blob object

Git reads the file content, prepends a header (`blob <size>\0`), computes the SHA-1 hash, compresses it with zlib, and writes it to `.git/objects/`.

```bash
# This is essentially what git add does internally:
git hash-object -w README.md
# Output: e69de29bb2d1d6434b8b29ae775ad8c2e48c5391
```

### Step 2: Update the index (staging area)

Git writes an entry to `.git/index` recording the filename, blob hash, file permissions, and timestamps.

```bash
# View the staging index
git ls-files --stage
# Output:
# 100644 e69de29bb2d1d6434b8b29ae775ad8c2e48c5391 0  README.md
```

### Step 3: Done

That's it. No commit yet. The blob is stored and the index is updated. The working directory is unchanged.

## What git commit Really Does

When you run `git commit -m "Initial commit"`:

### Step 1: Write tree objects

Git reads the index and creates tree objects that represent the directory structure. Subdirectories become nested tree objects.

```bash
# Manually create a tree from the current index
git write-tree
# Output: 4b825dc642cb6eb9a060e54bf899d15f3d70a0  (tree hash)
```

### Step 2: Create the commit object

Git creates a commit object containing the root tree hash, parent commit hash(es), author/committer info, and your message.

```bash
# Manually create a commit
echo "Initial commit" | git commit-tree 4b825dc -p HEAD
```

### Step 3: Update the branch ref

Git writes the new commit's SHA to `.git/refs/heads/main` (or whatever branch HEAD points to).

```
Before: refs/heads/main → abc123 (old commit)
After:  refs/heads/main → def456 (new commit, parent=abc123)
```

## The Full Picture

```
Working Dir  ──git add──►  Index (Staging)  ──git commit──►  Repository
   files                   .git/index                     .git/objects/
                           blob references                commits, trees,
                                                          blobs
```

---

# The Staging Index

The index (`.git/index`) is the staging area — the crucial middle ground between your working directory and the repository.

## Why Does the Index Exist?

The index lets you **craft commits precisely**. Instead of committing everything that changed, you choose exactly which changes to include.

```bash
# You modified 3 files but only want to commit 2:
git add file1.py file2.py    # Stage these two
git commit -m "Fix auth bug"  # Only file1 and file2 are committed

# file3.py remains modified but uncommitted
```

## What the Index Contains

The index is a binary file (`.git/index`) that stores a sorted list of entries:

```bash
git ls-files --stage
# 100644 a1b2c3d4... 0  src/app.py
# 100644 e5f6a7b8... 0  src/utils.py
# 100644 c9d0e1f2... 0  README.md
```

Each entry records:
- **File mode** (100644 = regular file, 100755 = executable, 120000 = symlink)
- **Blob SHA** — hash of the file content stored in the object database
- **Stage number** — 0 for normal, 1/2/3 during merge conflicts
- **File path**

## The Three-Tree Architecture

Git constantly compares three "trees":

```
    HEAD (last commit)      Index (staging)      Working Directory
    ──────────────────      ───────────────      ─────────────────
    app.py  (v1)            app.py  (v2)         app.py  (v3)
    utils.py (v1)           utils.py (v1)        utils.py (v1)
```

- `git diff` → compares **Index** vs **Working Directory**
- `git diff --staged` → compares **HEAD** vs **Index**
- `git diff HEAD` → compares **HEAD** vs **Working Directory**

## Stage Numbers During Merge Conflicts

During a merge conflict, the index holds three versions of the conflicted file:

```bash
git ls-files --stage
# 100644 aaa...  1  file.txt    ← Stage 1: common ancestor
# 100644 bbb...  2  file.txt    ← Stage 2: ours (current branch)
# 100644 ccc...  3  file.txt    ← Stage 3: theirs (incoming branch)
```

When you resolve and `git add`, it collapses back to stage 0.

---

# Branches as Pointers

Branches are the most lightweight concept in Git — and the most misunderstood.

## A Branch Is Just a File

A branch is a 41-byte file (40 hex chars + newline) in `.git/refs/heads/` containing a commit SHA:

```bash
cat .git/refs/heads/main
# a3f5b8c9d2e1f4a6b7c8d9e0f1a2b3c4d5e6f7a8

cat .git/refs/heads/feature-login
# b4c6d8e0f2a3b5c7d9e1f3a5b7c9d1e3f5a7b9c1
```

Creating a branch = writing a 40-character hash to a file. That's why `git branch feature-x` is instant, even in a repository with millions of commits.

## How Branch Pointers Move

When you make a commit on a branch, Git simply updates the file to point to the new commit:

```
Before commit:
  main → C3 → C2 → C1

After git commit:
  main → C4 → C3 → C2 → C1
         ↑
     main now points here
```

## Visualizing Branches

```
          feature
            │
            ▼
  C1 ← C2 ← C3 ← C5
               ↖
                C4 ← C6
                      ↑
                     main
```

Both `feature` and `main` are just pointers. The commit graph is the real structure.

## git branch Internals

```bash
# Create a branch (writes a new ref file)
git branch feature-x
# Equivalent to: echo $(git rev-parse HEAD) > .git/refs/heads/feature-x

# Switch branches (updates HEAD)
git checkout feature-x
# Equivalent to: echo "ref: refs/heads/feature-x" > .git/HEAD
# + updates working directory and index to match the commit

# Delete a branch (just removes the ref file)
git branch -d feature-x
# Equivalent to: rm .git/refs/heads/feature-x
```

Deleting a branch does NOT delete any commits. The commits still exist in the object database until garbage collection runs (and only if no other ref points to them).

## Packed Refs

In repositories with many branches, Git packs refs into a single file for efficiency:

```bash
cat .git/packed-refs
# # pack-refs with: peeled fully-peeled sorted
# a3f5b8c9... refs/heads/main
# b4c6d8e0... refs/heads/develop
# c5d7e9f1... refs/tags/v1.0
```

---

# HEAD and Detached HEAD

HEAD is the pointer that tells Git "where you are right now."

## Normal HEAD (Symbolic Reference)

In normal operation, HEAD is a **symbolic reference** — it points to a branch, which points to a commit:

```bash
cat .git/HEAD
# ref: refs/heads/main

# HEAD → main → a3f5b8c9...
```

When you commit, the branch pointer moves forward and HEAD follows because it points to the branch.

```
Before:  HEAD → main → C3
After:   HEAD → main → C4 → C3
```

## Detached HEAD

When you check out a specific commit (not a branch), HEAD points directly to a commit:

```bash
git checkout a3f5b8c9

cat .git/HEAD
# a3f5b8c9d2e1f4a6b7c8d9e0f1a2b3c4d5e6f7a8   (raw SHA, not a ref!)
```

```
Normal:    HEAD → main → C4
Detached:  HEAD → C4  (main still points to C4, but HEAD is detached)
```

## Common Causes of Detached HEAD

```bash
git checkout v1.0          # Checking out a tag
git checkout a3f5b8c9      # Checking out a specific commit
git checkout HEAD~3        # Going back 3 commits
```

## The Danger

You **can** make commits in detached HEAD state, but they don't belong to any branch:

```
              HEAD
               │
               ▼
  C3 ← C4 ← C5 (orphan commit!)
         ↑
        main
```

If you switch away (`git checkout main`), commit C5 becomes unreachable and will eventually be garbage collected.

## Rescuing Detached HEAD Commits

```bash
# Option 1: Create a branch before switching away
git checkout -b rescue-branch

# Option 2: If you already switched away, find it in the reflog
git reflog
# e5f6a7b HEAD@{1}: commit: my orphan commit
git checkout -b rescue-branch e5f6a7b
```

---

# Rebasing vs Merging Internals

Both integrate changes from one branch into another, but they do it very differently under the hood.

## Merge: Create a Join Point

`git merge feature` creates a **merge commit** with two parents:

```
Before:
  main:    C1 ← C2 ← C4
  feature:       C2 ← C3 ← C5

After git merge feature (from main):
  main:    C1 ← C2 ← C4 ← M6
                  ↖  C3 ← C5 ↗

  M6 is a merge commit with parents C4 and C5
```

Internally:
1. Git finds the **merge base** (common ancestor) — C2
2. Performs a **three-way merge** between C2, C4 (ours), and C5 (theirs)
3. Creates a new tree from the merged result
4. Creates commit M6 with **two parent pointers**

```bash
git cat-file -p M6
# tree    abc123...
# parent  C4...    ← first parent (main)
# parent  C5...    ← second parent (feature)
# author  ...
```

## Rebase: Replay Commits

`git rebase main` (from feature branch) **replays** feature's commits on top of main:

```
Before:
  main:    C1 ← C2 ← C4
  feature:       C2 ← C3 ← C5

After git rebase main (from feature):
  main:    C1 ← C2 ← C4
  feature:                ← C3' ← C5'

  C3' and C5' are NEW commits (new SHAs!) with same diffs as C3 and C5
```

Internally, for each commit being rebased:
1. Compute the **diff** (patch) introduced by that commit
2. Apply the patch onto the new base
3. Create a **new commit** with a new SHA (because the parent changed)

The original commits (C3, C5) still exist in the object database but are orphaned.

## When to Use Which

| Scenario | Recommended | Why |
|----------|-------------|-----|
| Integrating a shared branch | **Merge** | Preserves full history, safe for pushed branches |
| Cleaning up local feature branch | **Rebase** | Linear history, easier to read |
| Commit already pushed to remote | **Merge** | Rebase rewrites history, breaking collaborators |
| Long-running feature branch | **Merge** | Rebase on many commits is conflict-prone |

## The Golden Rule

**Never rebase commits that have been pushed and shared with others.** Rebase rewrites commit SHAs. If someone else has based work on the original commits, their history diverges and chaos ensues.

---

# The Reflog: Your Safety Net

The reflog is Git's local history of where HEAD and branch tips have been. It's your "undo" mechanism.

## What the Reflog Tracks

Every time HEAD moves — commit, checkout, rebase, reset, merge, pull — Git records it:

```bash
git reflog
# a3f5b8c HEAD@{0}: commit: Add login feature
# b4c6d8e HEAD@{1}: checkout: moving from main to feature
# c5d7e9f HEAD@{2}: commit: Update README
# d6e8f0a HEAD@{3}: rebase (finish): returning to refs/heads/main
# e7f9a1b HEAD@{4}: reset: moving to HEAD~2
```

Each entry shows:
- The commit SHA at that point
- A relative position (`HEAD@{N}`)
- What operation caused the movement

## Recovering from Mistakes

### Undo a bad reset

```bash
# Oops, you just did git reset --hard HEAD~3
git reflog
# a3f5b8c HEAD@{0}: reset: moving to HEAD~3
# f1e2d3c HEAD@{1}: commit: Important work     ← want this back!

git reset --hard f1e2d3c    # Restored!
```

### Find a commit lost during rebase

```bash
git reflog
# Look for "rebase (start)" entries to find pre-rebase state
# abc1234 HEAD@{5}: rebase (start): ...

git checkout -b recovery abc1234
```

### Recover a deleted branch

```bash
git branch -D feature-x    # Accidentally deleted!

git reflog
# Look for the last commit on that branch
# def5678 HEAD@{3}: commit: Last commit on feature-x

git checkout -b feature-x def5678    # Branch restored!
```

## Important Reflog Details

- The reflog is **local only** — it's never pushed to remotes
- Entries expire after **90 days** by default (30 days for unreachable commits)
- Stored in `.git/logs/`
- Each branch has its own reflog: `.git/logs/refs/heads/main`

```bash
# View reflog for a specific branch
git reflog show main

# View with timestamps
git reflog --date=relative

# View reflog for a specific ref with full diff
git log -g --patch refs/heads/main
```

## The Reflog Is Not a Substitute for Good Practices

The reflog protects you from local mistakes, but:
- It doesn't protect against `rm -rf .git`
- It doesn't exist on other people's machines
- Entries expire — don't wait months to recover
- **Always push important work to a remote** as the ultimate backup
