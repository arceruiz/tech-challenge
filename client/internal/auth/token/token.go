package token

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(userId string) (string, error) {
	permissions := jwt.MapClaims{}

	if userId == "guest" {
		permissions["guest"] = true
	}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte("vuIXaOK4OpJWA9ySX1UTpIWshXPpP6neGKGA724FauY"))
}

func ValidateToken(r *http.Request) error {
	tokenString := getToken(r)

	token, err := jwt.Parse(tokenString, returnSecretKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

func getToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnSecretKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado %v", token.Header["alg"])
	}

	return []byte("vuIXaOK4OpJWA9ySX1UTpIWshXPpP6neGKGA724FauY"), nil
}

func GetUserId(r *http.Request) (string, error) {
	tokenString := getToken(r)

	token, err := jwt.Parse(tokenString, returnSecretKey)
	if err != nil {
		return "", err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := fmt.Sprintf("%.0f", permissions["userId"])
		if err != nil {
			return "", nil
		}

		return userId, nil
	}

	return "", errors.New("token inválido")
}
