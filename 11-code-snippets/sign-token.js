const jwt = require("jsonwebtoken");
const secret = process.env.JWT_SECRET;
function generateAccessToken(user) {
  return jwt.sign({ sub: user.id, role: user.role }, secret, {
    expiresIn: "15m",
    algorithm: "HS256",
  });
}
function generateRefreshToken(user) {
  return jwt.sign({ sub: user.id, jti: uuidv4() }, secret, { expiresIn: "7d" });
}
