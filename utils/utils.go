package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

import "github.com/golang-jwt/jwt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 10).Unix(),
	})
	return token.SignedString([]byte("secret"))
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ParseJWT(tokenString string) string {
	tokenString = tokenString[7:]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if err != nil {
		return ""
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//fmt.Printf("%v", claims)
		return claims["username"].(string)
	} else {
		return ""
	}
}
