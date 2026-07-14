# Blacklisting & Refresh Token Rotation

- **Blacklist**: store revoked token IDs (jti) until they expire.
- **Rotation**: issue new refresh token each time; invalidate previous one. Detect reuse.

> 📘 Next: [Key Management & Rotation](06-key-management-and-rotation.md)
