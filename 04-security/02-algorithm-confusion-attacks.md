# Algorithm Confusion Attacks

Attacker changes `alg` to `none` or swaps `RS256` to `HS256` using the public key as HMAC secret. Mitigation: always explicitly specify allowed algorithms.

> 📘 Next: [Token Storage – XSS, CSRF](03-token-storage-xss-csrf.md)
