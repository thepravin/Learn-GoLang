package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("pravinNalawade2003")

func main() {
	claims := jwt.MapClaims{
		"sub":   "12345",                 //subject
		"name":  "Pravin",                // public claims or custom claims
		"email": "pravin@pre-scient.com", //cusom claims
		"exp":   time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // header + playload

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		panic(err)
	}

	fmt.Println("Signed token is given below")
	fmt.Println(signedToken)
}
