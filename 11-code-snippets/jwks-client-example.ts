import { createRemoteJWKSet } from "jose";
const JWKS = createRemoteJWKSet(
  new URL("https://auth.example.com/.well-known/jwks.json"),
);
const { payload } = await jwtVerify(token, JWKS, {
  issuer: "auth.example.com",
  audience: "api",
});
