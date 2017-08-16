package main

import (
	"time"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
)

func main() {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "pthethanh",
		"expire":   time.Date(2017, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	if tokenString, err := token.SignedString([]byte("my-secret")); err == nil {
		log.Println("Token:", tokenString)
	} else {
		fmt.Println("Failed to sign", err)
	}
}
