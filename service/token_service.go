package service

import (
	"encoding/json"
	"lets_bid/utils"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/go-redis/redis"
)

type Token struct {
	UserID int
	*jwt.StandardClaims
}

type AuthContext struct {
	Secret     string    `json:"secret"`
	ExpireTime time.Time `json:"expire_time"`
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

	// TODO: Move it to lib or service package
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	currentTime := time.Now()
	authContext := &AuthContext{
		Secret:     "secret",
		ExpireTime: currentTime.Add(time.Second * 3600),
	}

	authContextJSON, err := json.Marshal(authContext)

	if err != nil {
		panic(err)
	}

	err = redisClient.Set("cas:auth_context:"+strconv.Itoa(userID), authContextJSON, 0).Err()
	if err != nil {
		panic(err)
	}

	return tokenString
}
