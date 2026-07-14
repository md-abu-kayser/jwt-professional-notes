# JWT Structure

`xxxxx.yyyyy.zzzzz`

- **Header**: algorithm & token type (e.g., `{"alg":"HS256","typ":"JWT"}`)
- **Payload**: claims (iss, sub, exp, iat, custom)
- **Signature**: `HMACSHA256(base64UrlEncode(header) + "." + base64UrlEncode(payload), secret)`

> 📘 Next: [Claims & Registered Claim Names](04-claims-and-registered-claim-names.md)
