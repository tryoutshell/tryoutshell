```md
┌─────────────────────────────────────────────────────────────────────────┐
│                                                                         │
│   ████████╗██████╗ ██╗   ██╗ ██████╗ ██╗   ██╗████████╗███████╗██╗  ██╗│
│   ╚══██╔══╝██╔══██╗╚██╗ ██╔╝██╔═══██╗██║   ██║╚══██╔══╝██╔════╝██║  ██║│
│      ██║   ██████╔╝ ╚████╔╝ ██║   ██║██║   ██║   ██║   ███████╗███████║│
│      ██║   ██╔══██╗  ╚██╔╝  ██║   ██║██║   ██║   ██║   ╚════██║██╔══██║│
│      ██║   ██║  ██║   ██║   ╚██████╔╝╚██████╔╝   ██║   ███████║██║  ██║│
│      ╚═╝   ╚═╝  ╚═╝   ╚═╝    ╚═════╝  ╚═════╝    ╚═╝   ╚══════╝╚═╝  ╚═╝│
│                                                                         │
│                 🚀 Interactive Learning in Your Terminal          │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘

  Select an organization:

  ▸ 🛡️  Chainguard          Secure software supply chain tools
    🔗 In-Toto              Supply chain integrity framework
    🦉 Witness              Attestation and verification
    🤖 Unstable AI          ML security and guardrails


  ↑/↓: Navigate  •  Enter: Select  •  q: Quit
```

---

### **📚 Lesson Selection Screen**
```md
╔═══════════════════════════════════════════════════════════════════════╗
║  🛡️  Chainguard Lessons                                               ║
╚═══════════════════════════════════════════════════════════════════════╝

  Available Lessons:

  ▸ Container Image Signing with Cosign
    ⭐ Beginner  •  ⏱ 20 min  •  🔐 cosign, signing, supply-chain
    Learn to sign and verify container images for supply chain security

    Prerequisites: Docker installed, Basic understanding of containers

  ────────────────────────────────────────────────────────────────────

    Keyless Signing with Fulcio
    ⭐⭐ Intermediate  •  ⏱ 30 min  •  🎫 keyless, OIDC, fulcio
    Sign images without managing keys using OIDC identity

  ────────────────────────────────────────────────────────────────────

    Melange: Build Alpine APKs
    ⭐ Beginner  •  ⏱ 25 min  •  📦 melange, apk, builds
    Create secure Alpine packages with Melange


  ↑/↓: Navigate  •  Enter: Start Lesson  •  Esc: Back  •  q: Quit
```

---

### **📖 Introduction Screen**
```md
╔═══════════════════════════════════════════════════════════════════════╗
║  Container Image Signing with Cosign                                  ║
║  by Chainguard Team  •  v1.0                                          ║
╚═══════════════════════════════════════════════════════════════════════╝

┌─────────────────────────────────────────────────────────────────────┐
│  What You'll Learn                                                  │
└─────────────────────────────────────────────────────────────────────┘

In this lesson, you will:
  • Understand what Cosign is and why image signing matters
  • Install and verify Cosign
  • Sign a container image with ephemeral keys
  • Verify signatures to ensure image integrity

Time: ~20 minutes
Tools: Cosign, Docker


┌─────────────────────────────────────────────────────────────────────┐
│  💡 Tip: Have Docker running before you start!                      │
└─────────────────────────────────────────────────────────────────────┘


Progress: [░░░░░░░░░░░░░░░░░░░░] 0% (0/10 steps)


  Press Enter to continue  •  q: Quit
```

---

### **📄 Step 1: Info - "What is Cosign?"**
```md
╔═══════════════════════════════════════════════════════════════════════╗
║  Step 1/10                                                            ║
╚═══════════════════════════════════════════════════════════════════════╝

┌─────────────────────────────────────────────────────────────────────┐
│  What is Cosign?                                                    │
└─────────────────────────────────────────────────────────────────────┘

Cosign is a tool for signing and verifying container images and other
artifacts.

Think of it like a digital signature for software packages - it proves:
  • Who created the image (authenticity)
  • What exact image you're running (integrity)
  • When it was signed (provenance)


┌─ Why Does This Matter? ─────────────────────────────────────────────┐
│                                                                      │
│  Without signing, you can't be sure that myapp:latest hasn't been   │
│  tampered with between the registry and your production cluster.    │
│  An attacker could:                                                 │
│    • Replace your image with malicious code                         │
│    • Inject backdoors during CI/CD                                  │
│    • Modify dependencies without detection                          │
│                                                                      │
│  Cosign prevents this by creating cryptographic signatures that     │
│  are nearly impossible to forge.                                    │
│                                                                      │
└──────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────┐
│  💡 Fun Fact: Cosign is part of the Sigstore project, which also   │
│     includes Fulcio (certificate authority) and Rekor (transparency │
│     log).                                                           │
└─────────────────────────────────────────────────────────────────────┘


    Registry              Cosign                Your Cluster
    ┌────────┐            ┌──────┐             ┌───────────┐
    │ Image  │───sign────▶│ .sig │────verify──▶│ ✓ Deploy  │
    └────────┘            └──────┘             └───────────┘


Progress: [██░░░░░░░░░░░░░░░░░░] 10% (1/10 steps)


  Press Enter to continue  •  Esc: Back  •  q: Quit
```

---

### **📄 Step 2: Info - "Installing Cosign"**
```
╔═══════════════════════════════════════════════════════════════════════╗
║  Step 2/10                                                            ║
╚═══════════════════════════════════════════════════════════════════════╝

┌─────────────────────────────────────────────────────────────────────┐
│  Installing Cosign                                                  │
└─────────────────────────────────────────────────────────────────────┘

First, let's check if Cosign is already installed on your system.

Try running: cosign version

If you see version information, you're all set! Otherwise, install it
using:


╭─ macOS (Homebrew) ────────────────────────────────────────────────────╮
│                                                                        │
│  $ brew install cosign                                                 │
│                                                                        │
╰────────────────────────────────────────────────────────────────────────╯

╭─ Linux (Binary) ──────────────────────────────────────────────────────╮
│                                                                        │
│  $ curl -LO https://github.com/sigstore/cosign/releases/latest/\      │
│      download/cosign-linux-amd64                                       │
│  $ sudo mv cosign-linux-amd64 /usr/local/bin/cosign                   │
│  $ sudo chmod +x /usr/local/bin/cosign                                │
│                                                                        │
╰────────────────────────────────────────────────────────────────────────╯

╭─ Go Install ──────────────────────────────────────────────────────────╮
│                                                                        │
│  $ go install github.com/sigstore/cosign/v2/cmd/cosign@latest         │
│                                                                        │
╰────────────────────────────────────────────────────────────────────────╯


Progress: [████░░░░░░░░░░░░░░░░] 20% (2/10 steps)


  Press Enter to continue  •  Esc: Back  •  q: Quit
```

---

### **⌨️ Step 3: Command - "Verify Cosign is installed"**
```
╔═══════════════════════════════════════════════════════════════════════╗
║  Step 3/10  •  Command Execution                                      ║
╚═══════════════════════════════════════════════════════════════════════╝

┌─────────────────────────────────────────────────────────────────────┐
│  Verify Cosign is installed                                         │
└─────────────────────────────────────────────────────────────────────┘

Run the command to check your Cosign version:


╭─ Example ─────────────────────────────────────────────────────────────╮
│                                                                        │
│  $ cosign version                                                      │
│                                                                        │
╰────────────────────────────────────────────────────────────────────────╯


┌─ Your turn: ──────────────────────────────────────────────────────────┐
│                                                                        │
│  $ █                                                                   │
│                                                                        │
└────────────────────────────────────────────────────────────────────────┘


Progress: [██████░░░░░░░░░░░░░░] 30% (3/10 steps)


  Type command and press Enter  •  ?: Hint  •  Esc: Back  •  q: Quit
```

---

### **⌨️ Step 3: Command - User Types Command**
```
╔═══════════════════════════════════════════════════════════════════════╗
║  Step 3/10  •  Command Execution                                      ║
╚═══════════════════════════════════════════════════════════════════════╝

┌─────────────────────────────────────────────────────────────────────┐
│  Verify Cosign is installed                                         │
└─────────────────────────────────────────────────────────────────────┘

Run the command to check your Cosign version:


╭─ Example ─────────────────────────────────────────────────────────────╮
│                                                                        │
│  $ cosign version                                                      │
│                                                                        │
╰────────────────────────────────────────────────────────────────────────╯


┌─ Your turn: ──────────────────────────────────────────────────────────┐
│                                                                        │
│  $ cosign version█                                                     │
│                                                                        │
└────────────────────────────────────────────────────────────────────────┘


Progress: [██████░░░░░░░░░░░░░░] 30% (3/10 steps)


  Type command and press Enter  •  ?: Hint  •  Esc: Back  •  q: Quit
```

---

### **✅ Step 3: Command - Successful Execution**
```
╔═══════════════════════════════════════════════════════════════════════╗
║  Step 3/10  •  Command Execution                                      ║
╚═══════════════════════════════════════════════════════════════════════╝

┌─────────────────────────────────────────────────────────────────────┐
│  Verify Cosign is installed                                         │
└─────────────────────────────────────────────────────────────────────┘

Run the command to check your Cosign version:


┌─ Your command: ───────────────────────────────────────────────────────┐
│                                                                        │
│  $ cosign version                                                      │
│                                                                        │
└────────────────────────────────────────────────────────────────────────┘


╭─ Output ──────────────────────────────────────────────────────────────╮
│                                                                        │
│  GitVersion:    2.2.3                                                  │
│  GitCommit:     b2fc3a7b4fc1e8ebef2c1f6e11e5e0d1a8d4e8f8              │
│  GitTreeState:  clean                                                  │
│  BuildDate:     2024-01-15T10:30:00Z                                   │
│  GoVersion:     go1.21.5                                               │
│  Compiler:      gc                                                     │
│  Platform:      darwin/arm64                                           │
│                                                                        │
│  ⏱ Completed in 0.2s                                                   │
│                                                                        │
╰────────────────────────────────────────────────────────────────────────╯


✅ Great! Cosign v2.2.3 is installed.


Progress: [██████░░░░░░░░░░░░░░] 30% (3/10 steps)


  Press Enter to continue  •  Esc: Back  •  q: Quit
```

---

### **❌ Step 3: Command - Failed Execution (Example)**
```
╔═══════════════════════════════════════════════════════════════════════╗
║  Step 3/10  •  Command Execution                                      ║
╚═══════════════════════════════════════════════════════════════════════╝

┌─────────────────────────────────────────────────────────────────────┐
│  Verify Cosign is installed                                         │
└─────────────────────────────────────────────────────────────────────┘

Run the command to check your Cosign version:


┌─ Your command: ───────────────────────────────────────────────────────┐
│                                                                        │
│  $ cosign version                                                      │
│                                                                        │
└────────────────────────────────────────────────────────────────────────┘


╭─ Output ──────────────────────────────────────────────────────────────╮
│                                                                        │
│  bash: cosign: command not found                                       │
│                                                                        │
│  Exit code: 127                                                        │
│  ⏱ Completed in 0.1s                                                   │
│                                                                        │
╰────────────────────────────────────────────────────────────────────────╯


❌ Cosign not found. Please install it using one of the methods above.


┌─────────────────────────────────────────────────────────────────────┐
│  💡 Need help? Press ? for hints                                    │
└─────────────────────────────────────────────────────────────────────┘


Progress: [██████░░░░░░░░░░░░░░] 30% (3/10 steps)


  Try again  •  ?: Show hint  •  Esc: Back  •  q: Quit
```

---

### **💡 Step 3: Showing Hints (Press ?)**
```
╔═══════════════════════════════════════════════════════════════════════╗
║  Step 3/10  •  Command Execution                                      ║
╚═══════════════════════════════════════════════════════════════════════╝

┌─────────────────────────────────────────────────────────────────────┐
│  Verify Cosign is installed                                         │
└─────────────────────────────────────────────────────────────────────┘


╭─ Hint (Level 1/3) ────────────────────────────────────────────────────╮
│                                                                        │
│  💡 Type: cosign version                                               │
│                                                                        │
╰────────────────────────────────────────────────────────────────────────╯


┌─ Your turn: ──────────────────────────────────────────────────────────┐
│                                                                        │
│  $ █                                                                   │
│                                                                        │
└────────────────────────────────────────────────────────────────────────┘


Progress: [██████░░░░░░░░░░░░░░] 30% (3/10 steps)


  Type command and press Enter  •  ?: Next hint  •  Esc: Back
```

**Press ? again:**
```
╭─ Hint (Level 2/3) ────────────────────────────────────────────────────╮
│                                                                        │
│  💡 If not installed, use: brew install cosign (macOS) or download    │
│     from GitHub                                                        │
│                                                                        │
╰────────────────────────────────────────────────────────────────────────╯
```

**Press ? once more:**
```
╭─ Hint (Level 3/3) ────────────────────────────────────────────────────╮
│                                                                        │
│  💡 Still stuck? Check https://docs.sigstore.dev/cosign/installation   │
│                                                                        │
╰────────────────────────────────────────────────────────────────────────╯
```

---

### **📄 Step 4: Info - "How Signing Works"**
```
╔═══════════════════════════════════════════════════════════════════════╗
║  Step 4/10                                                            ║
╚═══════════════════════════════════════════════════════════════════════╝

┌─────────────────────────────────────────────────────────────────────┐
│  How Signing Works                                                  │
└─────────────────────────────────────────────────────────────────────┘

Cosign uses public key cryptography to sign images:

1. Generate a key pair (public + private)
     • Private key (secret) → used to sign
     • Public key → used to verify

2. Sign the image with your private key
     • Cosign creates a signature file
     • Signature is stored alongside the image in the registry

3. Verify using the public key
     • Anyone can verify, but only you can sign


┌─ Two Signing Modes: ──────────────────────────────────────────────────┐
│                                                                        │
│  🔑 Key-based signing (what we'll use today)                          │
│     • You generate and manage keys                                    │
│     • Good for testing and private registries                         │
│                                                                        │
│  🎫 Keyless signing (recommended for production)                      │
│     • Uses OIDC identity (GitHub, Google, etc.)                       │
│     • No keys to manage!                                              │
│     • We'll cover this in the advanced lesson                         │
│                                                                        │
└────────────────────────────────────────────────────────────────────────┘


┌─────────────────────────────────────────────────────────────────────┐
│  💡 Tip: In production, always use keyless signing with Fulcio      │
└─────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────┐
│  ⚠️  Warning: Never commit private keys to Git!                     │
└─────────────────────────────────────────────────────────────────────┘


Progress: [████████░░░░░░░░░░░░] 40% (4/10 steps)


  Press Enter to continue  •  Esc: Back  •  q: Quit
```

---

### **⌨️ Step 5: Command - "Generate a key pair"**
```
╔═══════════════════════════════════════════════════════════════════════╗
║  Step 5/10  •  Command Execution                                      ║
╚═══════════════════════════════════════════════════════════════════════╝

┌─────────────────────────────────────────────────────────────────────┐
│  Generate a key pair                                                │
└─────────────────────────────────────────────────────────────────────┘

Let's generate a test key pair. Cosign will ask you for a password -
use something simple like test123 (this is just for learning).


╭─ Example ─────────────────────────────────────────────────────────────╮
│                                                                        │
│  $ cosign generate-key-pair                                            │
│                                                                        │
╰────────────────────────────────────────────────────────────────────────╯


┌─ Your turn: ──────────────────────────────────────────────────────────┐
│                                                                        │
│  $ █                                                                   │
│                                                                        │
└────────────────────────────────────────────────────────────────────────┘


Progress: [██████████░░░░░░░░░░] 50% (5/10 steps)


  Type command and press Enter  •  ?: Hint  •  Esc: Back  •  q: Quit
```

---

### **⌨️ Step 5: Command - After User Types (Showing Interactive Output)**
```
╔═══════════════════════════════════════════════════════════════════════╗
║  Step 5/10  •  Command Execution                                      ║
╚═══════════════════════════════════════════════════════════════════════╝

┌─────────────────────────────────────────────────────────────────────┐
│  Generate a key pair                                                │
└─────────────────────────────────────────────────────────────────────┘


┌─ Your command: ───────────────────────────────────────────────────────┐
│                                                                        │
│  $ cosign generate-key-pair                                            │
│                                                                        │
└────────────────────────────────────────────────────────────────────────┘


╭─ Output (Live) ───────────────────────────────────────────────────────╮
│                                                                        │
│  Enter password for private key:                                       │
│  Enter password for private key again:                                 │
│  Private key written to cosign.key                                     │
│  Public key written to cosign.pub                                      │
│                                                                        │
│  ⏱ Completed in 3.4s                                                   │
│                                                                        │
╰────────────────────────────────────────────────────────────────────────╯


✅ Key pair generated successfully!


You should now see two files:
  • cosign.key - Your private key (keep secret!)
  • cosign.pub - Your public key (share freely)

The private key is encrypted with the password you chose.


Progress: [██████████░░░░░░░░░░] 50% (5/10 steps)


  Press Enter to continue  •  Esc: Back  •  q: Quit
```

---

### **📄 Step 6: Info - "Time to Sign!"**
```
╔═══════════════════════════════════════════════════════════════════════╗
║  Step 6/10                                                            ║
╚═══════════════════════════════════════════════════════════════════════╝

┌─────────────────────────────────────────────────────────────────────┐
│  Time to Sign!                                                      │
└─────────────────────────────────────────────────────────────────────┘

Now we'll sign an actual container image. We'll use a public test image
from Google's container registry.

The command structure is:

  cosign sign --key <private-key> <image>


Cosign will:
  1. Hash the image layers
  2. Create a signature using your private key
  3. Upload the signature to the registry (as <image>.sig)


Progress: [████████████░░░░░░░░] 60% (6/10 steps)


  Press Enter to continue  •  Esc: Back  •  q: Quit
```

---

### **⌨️ Step 7: Command - "Sign a container image"**
```
╔═══════════════════════════════════════════════════════════════════════╗
║  Step 7/10  •  Command Execution                                      ║
╚═══════════════════════════════════════════════════════════════════════╝

┌─────────────────────────────────────────────────────────────────────┐
│  Sign a container image                                             │
└─────────────────────────────────────────────────────────────────────┘

Let's sign the gcr.io/distroless/static:latest image.

You'll be prompted for your key password (the one you set earlier).


╭─ Example ─────────────────────────────────────────────────────────────╮
│                                                                        │
│  $ cosign sign --key cosign.key gcr.io/distroless/static:latest       │
│                                                                        │
╰────────────────────────────────────────────────────────────────────────╯


┌─ Your turn: ──────────────────────────────────────────────────────────┐
│                                                                        │
│  $ █                                                                   │
│                                                                        │
└────────────────────────────────────────────────────────────────────────┘


Progress: [██████████████░░░░░░] 70% (7/10 steps)
Timeout: 60 seconds


  Type command and press Enter  •  ?: Hint  •  Esc: Back  •  q: Quit
```

---

### **⌨️ Step 7: Command - Successful Signing**
```
╔═══════════════════════════════════════════════════════════════════════╗
║  Step 7/10  •  Command Execution                                      ║
╚═══════════════════════════════════════════════════════════════════════╝

┌─────────────────────────────────────────────────────────────────────┐
│  Sign a container image                                             │
└─────────────────────────────────────────────────────────────────────┘


┌─ Your command: ───────────────────────────────────────────────────────┐
│                                                                        │
│  $ cosign sign --key cosign.key gcr.io/distroless/static:latest       │
│                                                                        │
└────────────────────────────────────────────────────────────────────────┘


╭─ Output ──────────────────────────────────────────────────────────────╮
│                                                                        │
│  Enter password for private key:                                       │
│  Pushing signature to: gcr.io/distroless/static:sha256-abc123...sig   │
│                                                                        │
│  tlog entry created with index: 48392847                               │
│  Pushing signature to: gcr.io/distroless/static:sha256-abc123...sig   │
│                                                                        │
│  ⏱ Completed in 4.8s                                                   │
│                                                                        │
╰────────────────────────────────────────────────────────────────────────╯


✅ Image signed successfully!


🎉 Image signed! The signature is now stored in the registry.

Notice the output shows where the signature was pushed. Cosign stores
signatures as OCI artifacts alongside the original image.


Progress: [██████████████░░░░░░] 70% (7/10 steps)


  Press Enter to continue  •  Esc: Back  •  q: Quit
```

---

### **📄 Step 8: Info - "Verifying Signatures"**
```
╔═══════════════════════════════════════════════════════════════════════╗
║  Step 8/10                                                            ║
╚═══════════════════════════════════════════════════════════════════════╝

┌─────────────────────────────────────────────────────────────────────┐
│  Verifying Signatures                                               │
└─────────────────────────────────────────────────────────────────────┘

Now let's verify that our signature is valid. Anyone with the public key
can verify the signature, but only you (with the private key) can create
it.

This is the magic of public key cryptography! 🔐


Progress: [████████████████░░░░] 80% (8/10 steps)


  Press Enter to continue  •  Esc: Back  •  q: Quit
```
