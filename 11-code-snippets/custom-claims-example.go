token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "sub": "1234567890",
    "name": "John Doe",
    "admin": true,
    "iat": time.Now().Unix(),
})