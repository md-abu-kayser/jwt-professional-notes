# Verifying JWTs

```javascript
try {
  const decoded = jwt.verify(token, process.env.JWT_SECRET, {
    algorithms: ["HS256"],
    issuer: "myapp",
    audience: "myapp-api",
  });
} catch (err) {
  // token invalid
}
```
