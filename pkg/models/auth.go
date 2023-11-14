package models

import "github.com/dgrijalva/jwt-go"

type LoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var SecretKey = []byte("tu_clave_secreta")

type CustomClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}
