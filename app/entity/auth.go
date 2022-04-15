package entity

import "github.com/golang-jwt/jwt"

type JWTClaims struct {
	jwt.StandardClaims
	Fullname string
	Email    string
}
