package service

import (
	"lets_bid/utils"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserID int
	*jwt.StandardClaims
}

// GenerateNewToken will generate jwt token for given user id
func GenerateNewToken(userID int) string {
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	token := &Token{
		UserID: userID,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	signedToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), token)

	tokenString, err := signedToken.SignedString([]byte("secret"))

	if err != nil {
		panic("Cannot sign token")
	}

	conn := utils.ConnectDb()
	conn.Query("insert into tokens values ($1, $2)", userID, tokenString)
	defer conn.Close()

	return tokenString
}
