package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
)

const SECRET = "testKey"

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

func ParseTokenString(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		logger.Fatalln(err)
		return nil, err
	} else {
		return token.Claims, nil
	}
}

func GenerateTokenWithClaim(claims jwt.Claims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SECRET))
	if err != nil {
		logger.Fatalln(err)
	}
	return tokenString
}

func main() {
	claims := jwt.MapClaims{
		"name": "liufy47",
	}
	token := GenerateTokenWithClaim(claims)
	fmt.Println(token)

	parsedClaim, err := ParseTokenString(token)
	if err != nil {
		logger.Fatalln(err)
	}
	fmt.Println(parsedClaim)
}
