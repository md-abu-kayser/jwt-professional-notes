# Nested JWTs & JWE

- **JWE**: encrypt the entire token (header, encrypted key, IV, ciphertext, tag).
- **Nested JWT**: sign then encrypt (JWS inside JWE). Protects confidentiality and integrity.

> 📘 Next: **04-security** – [JWT Attacks & Common Vulnerabilities](../04-security/01-jwt-attacks-and-common-vulnerabilities.md)
