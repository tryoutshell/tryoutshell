# OWASP Top 10 (2021)

The **Open Web Application Security Project (OWASP)** maintains the most authoritative
list of critical web application security risks. Updated every few years, the
**OWASP Top 10** is the industry-standard awareness document for developers and
security teams worldwide.

**Why it matters:**
- Referenced by PCI DSS, NIST, and most compliance frameworks
- Used as a baseline for penetration testing and code review
- Helps development teams prioritize security efforts

The 2021 edition introduced three new categories and re-ranked several others
based on data from over 500,000 applications.

| Rank | Category |
|------|----------|
| A01  | Broken Access Control |
| A02  | Cryptographic Failures |
| A03  | Injection |
| A04  | Insecure Design |
| A05  | Security Misconfiguration |
| A06  | Vulnerable Components |
| A07  | Auth Failures |
| A08  | Data Integrity Failures |
| A09  | Logging Failures |
| A10  | SSRF |

---

# A01: Broken Access Control

**Moved from #5 to #1** in the 2021 edition. 94% of applications tested had
some form of broken access control.

Access control enforces that users cannot act outside their intended permissions.
Failures lead to unauthorized information disclosure, modification, or destruction
of data.

**Common vulnerabilities:**
- Bypassing access checks by modifying the URL or API request
- Viewing or editing someone else's account by changing the ID parameter
- Privilege escalation — acting as admin when logged in as a regular user
- Missing access controls on API endpoints (POST, PUT, DELETE)

**Real attack example — Insecure Direct Object Reference (IDOR):**

```
GET /api/invoices/1001        ← Your invoice
GET /api/invoices/1002        ← Someone else's invoice (no check!)
```

An attacker simply increments the invoice ID and accesses another user's data.

**Mitigations:**
- Deny by default — require explicit grants for every resource
- Implement server-side access control; never rely on client-side checks
- Use indirect references (UUIDs) instead of sequential IDs
- Log and alert on access control failures
- Rate-limit API requests to minimize automated abuse

---

# A02: Cryptographic Failures

Previously called "Sensitive Data Exposure," this category focuses on failures
related to cryptography that lead to exposure of sensitive data.

**What counts as sensitive data?**
- Passwords, credit card numbers, health records, PII
- API keys, session tokens, business-critical data

**Common vulnerabilities:**
- Transmitting data in cleartext (HTTP, SMTP, FTP)
- Using deprecated algorithms: MD5, SHA1, DES, RC4
- Storing passwords with reversible encryption or plain text
- Missing TLS or using outdated TLS versions (< 1.2)
- Weak or reused encryption keys

**Real attack — Plain-text password storage:**

```python
# VULNERABLE: storing password as-is
user.password = request.form["password"]
db.session.commit()

# SECURE: use bcrypt with a salt
import bcrypt
hashed = bcrypt.hashpw(
    request.form["password"].encode(), bcrypt.gensalt()
)
user.password_hash = hashed
db.session.commit()
```

**Mitigations:**
- Classify data by sensitivity; don't store sensitive data you don't need
- Encrypt all data in transit with TLS 1.2+
- Use strong adaptive hashing for passwords: bcrypt, scrypt, or Argon2
- Enforce HSTS headers and disable caching for sensitive responses
- Generate keys with cryptographically secure random functions

---

# A03: Injection

Injection flaws occur when untrusted data is sent to an interpreter as part of
a command or query. The attacker's hostile data tricks the interpreter into
executing unintended commands or accessing data without authorization.

**Types of injection:**
- **SQL Injection** — manipulating database queries
- **NoSQL Injection** — attacking MongoDB, CouchDB, etc.
- **OS Command Injection** — executing shell commands
- **LDAP Injection** — manipulating directory queries
- **Expression Language / Template Injection**

**Real attack — SQL Injection:**

```python
# VULNERABLE — string concatenation
query = "SELECT * FROM users WHERE name = '" + username + "'"
# Attacker input: ' OR '1'='1' --
# Resulting query: SELECT * FROM users WHERE name = '' OR '1'='1' --'
# Returns ALL users!

# SECURE — parameterized query
cursor.execute(
    "SELECT * FROM users WHERE name = %s", (username,)
)
```

**OS Command Injection:**

```python
# VULNERABLE
os.system("ping -c 1 " + user_input)
# Attacker input: 127.0.0.1; cat /etc/passwd

# SECURE — use subprocess with argument list
import subprocess
subprocess.run(["ping", "-c", "1", user_input], check=True)
```

**Mitigations:**
- Use parameterized queries / prepared statements everywhere
- Use ORMs that handle escaping (but still validate input)
- Validate and sanitize all user inputs with allowlists
- Escape special characters for the specific interpreter
- Run applications with least-privilege database accounts

---

# A04: Insecure Design

**New in the 2021 edition.** This category focuses on design-level flaws —
risks that cannot be fixed by a perfect implementation because the
architecture itself is insecure.

Insecure design is different from insecure implementation. A secure design
can still have implementation bugs, and an insecure design cannot be fixed
by perfect code.

**Common patterns:**
- No rate limiting on authentication or sensitive operations
- No account lockout after repeated failed login attempts
- Trusting client-side validation as the only check
- Missing business logic abuse protections
- Not using threat modeling during design phase

**Real scenario — Password recovery abuse:**

```
POST /api/password-reset
{ "email": "victim@example.com" }

→ Server responds: "Security question: What is your pet's name?"
→ Attacker can brute-force the answer with no rate limiting
→ No lockout after 1000 failed attempts
```

**Mitigations:**
- Integrate threat modeling into your SDLC (STRIDE, PASTA)
- Establish secure design patterns and reference architectures
- Use abuse cases / misuse cases alongside user stories
- Implement rate limiting and lockout at the design level
- Segregate tenant data at the architecture level
- Limit resource consumption by user or service

---

# A05: Security Misconfiguration

The most commonly seen issue. 90% of applications were tested for
misconfiguration, and 4.5% had issues.

This happens when security settings are not defined, implemented,
or maintained properly. Default configurations are often insecure.

**Common vulnerabilities:**
- Default credentials left unchanged (admin/admin)
- Unnecessary features enabled: ports, services, pages, accounts
- Error handling revealing stack traces or internal details
- Missing security headers (CSP, X-Frame-Options, HSTS)
- Cloud storage buckets with public access
- Directory listing enabled on web servers

**Real scenario — Exposed error pages:**

```
GET /api/users/undefined

HTTP 500 Internal Server Error
{
  "error": "TypeError: Cannot read property 'id' of undefined",
  "stack": "at UserController.getUser (/app/src/controllers/user.js:42:15)\n
            at Layer.handle [as handle_request] (/app/node_modules/express/...)",
  "database": "mongodb://admin:s3cret@db.internal:27017/production"
}
```

Stack traces reveal internal paths, libraries, and even database credentials.

**Mitigations:**
- Implement a repeatable hardening process for all environments
- Remove or don't install unused features and frameworks
- Review and update configurations as part of the patch process
- Use security headers: CSP, HSTS, X-Content-Type-Options
- Automate configuration verification (e.g., CIS Benchmarks)
- Use different credentials for every environment

---

# A06: Vulnerable and Outdated Components

Using components (libraries, frameworks, or software modules) with known
vulnerabilities. This is a supply chain risk that is extremely common and
difficult to test for.

**Why this is dangerous:**
- Attackers actively scan for CVEs in popular libraries
- Many projects have deep dependency trees (transitive deps)
- A single vulnerable package can compromise the entire app

**Real-world examples:**
- **Log4Shell (CVE-2021-44228):** Remote code execution in Log4j affected
  millions of Java applications worldwide
- **event-stream incident:** A malicious maintainer injected cryptocurrency-
  stealing code into a popular npm package
- **ua-parser-js hijack:** Attackers published malware versions of a package
  with 7M+ weekly downloads

**Check your dependencies:**

```bash
# Node.js
npm audit

# Python
pip-audit

# Go
govulncheck ./...

# Container images
grype my-image:latest
trivy image my-image:latest
```

**Mitigations:**
- Continuously inventory all components and their versions
- Monitor CVE databases: NVD, GitHub Advisory, OSV
- Enable automated dependency updates (Dependabot, Renovate)
- Remove unused dependencies — every dep is attack surface
- Only obtain components from official sources over secure links
- Pin dependency versions and use lock files

---

# A07: Identification and Authentication Failures

Previously "Broken Authentication." Covers weaknesses in session management,
credential handling, and identity verification.

**Common vulnerabilities:**
- Permitting brute-force or credential stuffing attacks
- Allowing weak or well-known passwords ("Password1", "admin123")
- Using weak credential recovery (knowledge-based questions)
- Storing passwords in plain text or with weak hashing (MD5)
- Missing or ineffective multi-factor authentication
- Exposing session IDs in URLs

**Real attack — Credential stuffing:**

```bash
# Attacker uses leaked credentials from another breach
while read email password; do
  curl -s -X POST https://target.com/login \
    -d "email=$email&password=$password" \
    | grep -q "Welcome" && echo "HIT: $email"
done < leaked_credentials.txt
```

Because users reuse passwords, breached lists from one site compromise others.

**Mitigations:**
- Implement multi-factor authentication (TOTP, WebAuthn, passkeys)
- Don't ship with default credentials
- Check passwords against known breached lists (Have I Been Pwned API)
- Enforce password length (12+ chars) over complexity rules
- Use rate limiting and account lockout on login endpoints
- Use secure session management — regenerate IDs after login
- Set session timeouts appropriate to the application's risk

---

# A08: Software and Data Integrity Failures

**New in the 2021 edition.** Focuses on code and infrastructure that does not
protect against integrity violations: untrusted plugins, CI/CD pipelines
without verification, or auto-update mechanisms without signature checks.

**Common vulnerabilities:**
- Using libraries from untrusted CDNs without integrity verification
- Insecure CI/CD pipelines that allow unauthorized code injection
- Auto-update functionality without signed packages
- Insecure deserialization of untrusted data

**Real scenario — Compromised CDN resource:**

```html
<!-- VULNERABLE: no integrity check -->
<script src="https://cdn.example.com/jquery-3.6.0.min.js"></script>

<!-- SECURE: Subresource Integrity (SRI) -->
<script
  src="https://cdn.example.com/jquery-3.6.0.min.js"
  integrity="sha384-vtXRMe3mGCbOeY7l30aIg8H9p3GdeSe4IFlP6G8JMa7o7lXvnz3GFKzPxzJdPfGK"
  crossorigin="anonymous">
</script>
```

If the CDN is compromised, SRI ensures the browser rejects tampered files.

**CI/CD Pipeline Attack:**

```yaml
# Dangerous: workflow uses unverified third-party action at latest
- uses: random-user/deploy-action@main

# Safer: pin to a specific commit SHA
- uses: random-user/deploy-action@a1b2c3d4e5f6
```

**Mitigations:**
- Verify digital signatures on software and updates
- Use SRI for all third-party CDN resources
- Pin CI/CD dependencies to specific versions or SHAs
- Ensure your CI/CD pipeline has proper segregation and access controls
- Don't send unsigned or unencrypted serialized data to untrusted clients
- Adopt frameworks like SLSA for build integrity

---

# A09: Security Logging and Monitoring Failures

Without proper logging and monitoring, breaches cannot be detected.
Studies show the average time to identify a breach is **287 days**.

**Common vulnerabilities:**
- Login failures, access control failures not logged
- Logs only stored locally (lost when server is compromised)
- No alerting for suspicious activity
- Logs lack sufficient detail (who, what, when, from where)
- Application unable to detect active attacks in real time

**What you should log:**

```json
{
  "timestamp": "2024-03-15T14:32:01Z",
  "level": "WARN",
  "event": "authentication_failure",
  "user": "admin",
  "source_ip": "203.0.113.42",
  "user_agent": "python-requests/2.28.0",
  "endpoint": "/api/login",
  "details": "Invalid password - attempt 47 in 60 seconds"
}
```

**Red flags that should trigger alerts:**
- Multiple failed logins from the same IP
- Access attempts to admin endpoints from non-admin users
- Unusual data export volumes
- API calls from unexpected geographies
- Repeated 403/401 responses to a single session

**Mitigations:**
- Log all authentication events (success and failure)
- Ensure logs have enough context for forensic analysis
- Use centralized log management (ELK, Splunk, Datadog)
- Set up real-time alerting for suspicious patterns
- Establish an incident response plan and test it regularly
- Protect log integrity — append-only, separate storage
- Follow OWASP Logging Cheat Sheet guidelines

---

# A10: Server-Side Request Forgery (SSRF)

**New in the 2021 edition.** SSRF occurs when a web application fetches a
remote resource without validating the user-supplied URL. This allows an
attacker to coerce the application into sending requests to unexpected
destinations, even when protected by a firewall or VPN.

**Why SSRF is devastating in cloud environments:**
Cloud providers expose metadata APIs on well-known internal addresses.
SSRF lets attackers reach them from the outside.

**Real attack — AWS metadata theft:**

```
# Application has a "preview URL" feature:
POST /api/preview
{ "url": "http://169.254.169.254/latest/meta-data/iam/security-credentials/my-role" }

# Server fetches the URL internally and returns:
{
  "AccessKeyId": "AKIA...",
  "SecretAccessKey": "wJalr...",
  "Token": "FwoGZX..."
}
```

The attacker now has temporary AWS credentials and can access S3 buckets,
databases, and other cloud resources.

**Capital One Breach (2019):** An SSRF vulnerability in a WAF allowed an
attacker to steal credentials via the EC2 metadata service, exfiltrating
100M+ customer records.

**Mitigations:**
- Validate and sanitize all user-supplied URLs on the server
- Use allowlists for permitted domains and IP ranges
- Block requests to private/internal IP ranges (10.x, 172.16.x, 169.254.x)
- Disable HTTP redirects or validate each redirect destination
- Use IMDSv2 (Instance Metadata Service v2) on AWS to require session tokens
- Don't return raw responses from server-side fetches to clients
- Segment remote resource access into separate networks

---

# Summary & Next Steps

## Key takeaways:

1. **Broken Access Control** is now the #1 risk — enforce least privilege
2. **Cryptographic Failures** require strong algorithms and proper key management
3. **Injection** remains critical — always use parameterized queries
4. **Insecure Design** means security must start at the architecture phase
5. **Misconfiguration** is the most common — automate your hardening
6. **Vulnerable Components** need continuous monitoring and patching
7. **Authentication Failures** are prevented by MFA and breach-aware passwords
8. **Data Integrity** requires signing, SRI, and secure CI/CD
9. **Logging Failures** let breaches go undetected for months
10. **SSRF** is devastating in cloud environments — validate all URLs

## Resources:

- [OWASP Top 10 Official Site](https://owasp.org/www-project-top-ten/)
- [OWASP Cheat Sheet Series](https://cheatsheetseries.owasp.org/)
- [OWASP Testing Guide](https://owasp.org/www-project-web-security-testing-guide/)
- [OWASP ASVS (Application Security Verification Standard)](https://owasp.org/www-project-application-security-verification-standard/)
- [PortSwigger Web Security Academy](https://portswigger.net/web-security) — free hands-on labs

**Practice:** Use tools like OWASP ZAP, Burp Suite, or `nikto` to scan
your own applications and identify these risks firsthand.
