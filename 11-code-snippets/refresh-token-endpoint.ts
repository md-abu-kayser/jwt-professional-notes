app.post("/auth/refresh", async (req, res) => {
  const refreshToken = req.cookies.refresh_token;
  // verify, check database, rotate, return new access token
  const newAccess = generateAccessToken(user);
  res.json({ accessToken: newAccess });
});
