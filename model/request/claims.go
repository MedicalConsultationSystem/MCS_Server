package request

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserId string
	BufferTime int64
	jwt.StandardClaims
}