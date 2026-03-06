# What is SLSA?

**SLSA** (pronounced "salsa") stands for **Supply-chain Levels for Software
Artifacts**. It is a security framework created by the **Open Source Security
Foundation (OpenSSF)** — originally developed at Google — that provides a
checklist of standards and controls to prevent tampering, improve integrity,
and secure packages and infrastructure.

**The problem SLSA solves:**

Modern software depends on hundreds of components — source code, build systems,
package registries, and CI/CD pipelines. An attacker who compromises **any**
link in this chain can inject malicious code into your software.

**Real-world supply chain attacks:**
- **SolarWinds (2020):** Attackers compromised the build system and injected a
  backdoor into updates sent to 18,000+ organizations including US government
- **Codecov (2021):** Attackers modified a bash uploader script in CI, stealing
  environment variables (secrets, tokens) from thousands of repos
- **xz Utils (2024):** A social engineering campaign over 2+ years planted a
  backdoor in a critical Linux compression library

**SLSA addresses this by asking:**
- Can I trace this artifact back to its source?
- Was the build process tamper-resistant?
- Can I verify no one modified the artifact after building?

SLSA provides **graduated security levels** — you don't have to do everything
at once. Each level builds on the previous one.

---

# SLSA Levels Explained

SLSA defines a track system. The primary track is **Build**, which has three
levels (L1 through L3):

## Build L1 — Provenance Exists

The build process produces **provenance** — a document that says *what* was
built, *how* it was built, and *by whom*.

**Requirements:**
- Build runs via a scripted process (not manual)
- Provenance is generated (but doesn't need to be tamper-proof)
- Provenance identifies the build platform and top-level inputs

```json
{
  "buildType": "https://github.com/slsa-framework/slsa-github-generator",
  "builder": { "id": "https://github.com/slsa-framework/slsa-github-generator/.github/workflows/generator_generic_slsa3.yml" },
  "invocation": {
    "configSource": {
      "uri": "git+https://github.com/my-org/my-repo@refs/tags/v1.0.0",
      "entryPoint": ".github/workflows/release.yml"
    }
  }
}
```

## Build L2 — Hosted, Signed Provenance

The build runs on a **hosted** platform and the provenance is **signed**
by the build service, making it tamper-evident.

**Additional requirements over L1:**
- Build runs on a hosted, managed build service
- Provenance is cryptographically signed by the build service
- Consumers can verify the signature

## Build L3 — Hardened Builds

The build platform provides **hardened isolation** between builds, preventing
one build from influencing another.

**Additional requirements over L2:**
- Build runs in an isolated, ephemeral environment
- Builds cannot access secrets from other builds
- The build service's signing keys are protected (e.g., in an HSM)
- Provenance is unforgeable — even the project maintainer cannot fake it

---

# Build Provenance Deep Dive

**Build provenance** is the core artifact of SLSA. It's a signed attestation
that answers three questions:

1. **What** was built? (artifact hash, name)
2. **How** was it built? (build command, workflow, inputs)
3. **Where** was it built? (build platform identity)

Provenance follows the **in-toto attestation format** with a SLSA predicate:

```json
{
  "_type": "https://in-toto.io/Statement/v1",
  "subject": [
    {
      "name": "my-binary",
      "digest": {
        "sha256": "abc123..."
      }
    }
  ],
  "predicateType": "https://slsa.dev/provenance/v1",
  "predicate": {
    "buildDefinition": {
      "buildType": "https://slsa-framework.github.io/github-actions-buildtypes/workflow/v1",
      "externalParameters": {
        "workflow": {
          "ref": "refs/tags/v1.0.0",
          "repository": "https://github.com/my-org/my-repo"
        }
      },
      "resolvedDependencies": [
        {
          "uri": "git+https://github.com/my-org/my-repo@refs/tags/v1.0.0",
          "digest": { "gitCommit": "abc123..." }
        }
      ]
    },
    "runDetails": {
      "builder": {
        "id": "https://github.com/slsa-framework/slsa-github-generator/.github/workflows/generator_generic_slsa3.yml@refs/tags/v2.0.0"
      }
    }
  }
}
```

**Key fields:**
- `subject` — The artifact(s) that were produced, identified by cryptographic hash
- `buildType` — What kind of build process was used
- `externalParameters` — Inputs the user controlled (tag, branch)
- `resolvedDependencies` — Pinned inputs the build actually used
- `builder.id` — The identity of the build platform

Provenance lets anyone verify: "Was this binary really built from commit X
using workflow Y on builder Z?"

---

# SLSA and Sigstore

**Sigstore** is the signing and transparency infrastructure that makes
SLSA practical at scale. It provides:

**Cosign** — Sign and verify container images and blobs:

```bash
# Sign a container image (keyless, uses OIDC identity)
cosign sign ghcr.io/my-org/my-image@sha256:abc123...

# Verify the signature
cosign verify ghcr.io/my-org/my-image@sha256:abc123... \
  --certificate-identity=https://github.com/my-org/my-repo/.github/workflows/release.yml@refs/tags/v1.0.0 \
  --certificate-oidc-issuer=https://token.actions.githubusercontent.com
```

**Rekor** — An immutable transparency log that records signing events.
Once an artifact is signed, the event is publicly recorded and cannot
be deleted or modified:

```bash
# Search for entries about a specific artifact
rekor-cli search --sha abc123...

# Get details about a log entry
rekor-cli get --uuid 24296fb24b8ad77a...
```

**Fulcio** — A certificate authority that issues short-lived signing
certificates based on OIDC identity (GitHub, Google, etc.).

**How they work together in CI:**

```
GitHub Actions workflow runs
  → Fulcio issues a short-lived cert (identity = workflow)
  → Cosign signs the artifact with that cert
  → Rekor records the signing event in the transparency log
  → Consumers verify using cosign + Rekor
```

No long-lived keys to manage. The signer's identity is bound to their
OIDC identity (e.g., a specific GitHub Actions workflow).

---

# SLSA in Practice with GitHub Actions

The **SLSA GitHub Generator** is a set of reusable workflows that generate
SLSA Build L3 provenance on GitHub Actions.

**Generating provenance for a Go binary:**

```yaml
name: Release

on:
  push:
    tags: ["v*"]

permissions: read-all

jobs:
  build:
    runs-on: ubuntu-latest
    outputs:
      hashes: ${{ steps.hash.outputs.hashes }}
    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version: "1.22"

      - name: Build binary
        run: go build -trimpath -o my-binary ./cmd/app

      - name: Generate hash
        id: hash
        run: |
          sha256sum my-binary | base64 -w0 > hashes.txt
          echo "hashes=$(cat hashes.txt)" >> "$GITHUB_OUTPUT"

      - uses: actions/upload-artifact@v4
        with:
          name: my-binary
          path: my-binary

  provenance:
    needs: build
    permissions:
      actions: read
      id-token: write
      contents: write
    uses: slsa-framework/slsa-github-generator/.github/workflows/generator_generic_slsa3.yml@v2.0.0
    with:
      base64-subjects: "${{ needs.build.outputs.hashes }}"
      upload-assets: true
```

**What happens:**
1. Your `build` job builds the binary and computes its SHA-256 hash
2. The `provenance` job (a trusted reusable workflow) generates and signs
   provenance attesting that the hash came from your repo + workflow
3. The provenance is uploaded alongside your release artifacts
4. Because the SLSA generator runs in an **isolated reusable workflow**,
   your build job cannot tamper with the provenance — achieving Build L3

---

# Generating Provenance for Container Images

For Docker/OCI container images, use the SLSA container generator:

```yaml
name: Release Container

on:
  push:
    tags: ["v*"]

permissions: read-all

jobs:
  build:
    runs-on: ubuntu-latest
    outputs:
      image: ${{ steps.build.outputs.image }}
      digest: ${{ steps.build.outputs.digest }}
    permissions:
      packages: write
    steps:
      - uses: actions/checkout@v4

      - name: Login to GHCR
        uses: docker/login-action@v3
        with:
          registry: ghcr.io
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}

      - name: Build and push
        id: build
        uses: docker/build-push-action@v5
        with:
          context: .
          push: true
          tags: ghcr.io/${{ github.repository }}:${{ github.ref_name }}

  provenance:
    needs: build
    permissions:
      actions: read
      id-token: write
      packages: write
    uses: slsa-framework/slsa-github-generator/.github/workflows/generator_container_slsa3.yml@v2.0.0
    with:
      image: ${{ needs.build.outputs.image }}
      digest: ${{ needs.build.outputs.digest }}
      registry-username: ${{ github.actor }}
    secrets:
      registry-password: ${{ secrets.GITHUB_TOKEN }}
```

The provenance is attached to the container image as a **cosign attestation**
in the OCI registry, making it discoverable alongside the image itself.

**Consumers can then verify both the image signature and its provenance:**

```bash
cosign verify-attestation \
  --type slsaprovenance \
  --certificate-identity-regexp "^https://github.com/slsa-framework/slsa-github-generator/" \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com \
  ghcr.io/my-org/my-image@sha256:abc123...
```

---

# Verification with slsa-verifier

**slsa-verifier** is the official tool for verifying SLSA provenance generated
by trusted builders (GitHub Actions, Google Cloud Build).

**Install:**

```bash
go install github.com/slsa-framework/slsa-verifier/v2/cli/slsa-verifier@latest
```

**Verify a binary artifact:**

```bash
slsa-verifier verify-artifact my-binary \
  --provenance-path my-binary.intoto.jsonl \
  --source-uri github.com/my-org/my-repo \
  --source-tag v1.0.0
```

**Verify a container image:**

```bash
slsa-verifier verify-image \
  ghcr.io/my-org/my-image@sha256:abc123... \
  --source-uri github.com/my-org/my-repo \
  --source-tag v1.0.0
```

**What slsa-verifier checks:**
1. The provenance signature is valid and chains to a trusted builder
2. The artifact hash matches what's in the provenance
3. The source repository matches (prevents someone from building
   your project in their repo and claiming it's yours)
4. The builder identity matches a known trusted builder
5. Optionally: the source tag or branch matches

**Example output on success:**

```
Verified build using builder
  "https://github.com/slsa-framework/slsa-github-generator/.github/workflows/generator_generic_slsa3.yml@refs/tags/v2.0.0"
  at commit abc123...

PASSED: Verified SLSA provenance
```

**Example output on failure:**

```
FAILED: expected source 'github.com/my-org/my-repo'
  but got 'github.com/attacker/fork'
```

This is the core security guarantee: even if an attacker builds the same code,
the provenance will not verify because the source URI won't match.

---

# SLSA Adoption and Ecosystem

**Who uses SLSA today?**

- **Google** — All internal builds produce SLSA L3+ provenance
- **npm** — Package provenance for published packages (SLSA Build L3)
- **PyPI** — Trusted Publishers with provenance via GitHub Actions
- **Homebrew** — Verifies bottles with SLSA provenance
- **Kubernetes** — Release artifacts ship with SLSA provenance
- **Sigstore** — All Sigstore releases include SLSA L3 provenance

**npm package provenance example:**

```bash
# Publish with provenance (in GitHub Actions)
npm publish --provenance

# Check provenance on npmjs.com
# Look for the green "Provenance" badge on the package page
```

**Verifying npm provenance:**

```bash
npm audit signatures
```

**SLSA levels as a maturity ladder:**

| Level | What you get | Effort |
|-------|-------------|--------|
| L0 | No guarantees | None |
| L1 | Provenance exists, scripted build | Low |
| L2 | Signed provenance from hosted builder | Medium |
| L3 | Hardened, isolated builds | Medium-High |

**Start with L1** — just having provenance is a significant improvement.
The SLSA GitHub Generator makes jumping to L3 straightforward for
GitHub-hosted projects.

---

# Summary & Resources

## Key takeaways:

1. **SLSA is a framework**, not a tool — it defines security levels for your
   build process and artifacts
2. **Build provenance** is the core concept — cryptographic proof of how
   an artifact was produced
3. **Three levels** provide a graduation path: L1 (provenance exists),
   L2 (signed), L3 (hardened builds)
4. **Sigstore** (Cosign, Rekor, Fulcio) provides the signing and
   transparency infrastructure that makes SLSA practical
5. **slsa-verifier** lets consumers verify provenance against trusted builders
6. **GitHub Actions** has first-class support via the SLSA GitHub Generator
7. **Start small** — even L1 provenance dramatically improves your security
   posture compared to having nothing

## Resources:

- [SLSA Official Website](https://slsa.dev)
- [SLSA Specification](https://slsa.dev/spec/v1.0/)
- [SLSA GitHub Generator](https://github.com/slsa-framework/slsa-github-generator)
- [slsa-verifier](https://github.com/slsa-framework/slsa-verifier)
- [Sigstore Documentation](https://docs.sigstore.dev)
- [OpenSSF Scorecard](https://securityscorecards.dev) — Automated security
  health checks for open source projects
- [in-toto Attestation Spec](https://github.com/in-toto/attestation)
- [Google's Perspective on SLSA](https://security.googleblog.com/2021/06/introducing-slsa-end-to-end-framework.html)

**Next steps:** Try generating SLSA provenance for one of your own projects
using the SLSA GitHub Generator. Start with a simple Go binary or container
image and work through the verification process end to end.
