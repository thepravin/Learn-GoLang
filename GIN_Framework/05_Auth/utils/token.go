package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secretKey string = "secretKey"

func GenerateToken(email string, id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func ParseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil

}

func TokenCheck(jwtToken string) (jwt.MapClaims, error) {
	token, err := ParseToken(jwtToken)
	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("bad jwt token")
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, errors.New("jwt token expired")
	}

	return claims, nil
}
