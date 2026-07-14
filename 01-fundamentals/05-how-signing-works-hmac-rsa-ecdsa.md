# How Signing Works

- **HMAC** (`HS256`, `HS384`, `HS512`): symmetric secret, shared between issuer and verifier.
- **RSA** (`RS256`…): asymmetric, private key signs, public key verifies.
- **ECDSA** (`ES256`…): like RSA but with elliptic curves, smaller keys.

> 📘 Next: [JOSE & Related Standards](06-jose-and-related-standards.md)
