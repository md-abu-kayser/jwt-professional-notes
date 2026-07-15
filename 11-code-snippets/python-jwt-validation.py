import jwt
decoded = jwt.decode(token, 'secret', algorithms=['HS256'], audience='myapp', issuer='authservice')