package configs

import (
	"fmt"
	"net/mail"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil //bila errornya nil, kembalikan true
}

var jwtKey = []byte("bebas")

func CreateToken(email string) (string, error) {
	// Buat payload token
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token kadaluarsa dalam 24 jam

	// Buat token JWT dengan menggunakan algoritma HMAC dan kunci yang sudah ditentukan
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	// Parse token dengan kunci yang sudah ditentukan
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validasi algoritma yang digunakan dalam token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	// Pastikan token sudah terverifikasi
	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	// Ambil claims dari token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}
