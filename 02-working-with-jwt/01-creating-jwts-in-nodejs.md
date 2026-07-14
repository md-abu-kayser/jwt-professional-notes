# Creating JWTs in Node.js

```javascript
const jwt = require("jsonwebtoken");
const token = jwt.sign({ userId: 123 }, process.env.JWT_SECRET, {
  expiresIn: "1h",
  algorithm: "HS256",
  issuer: "myapp",
  audience: "myapp-api",
});
```
