package service

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"lets_bid/utils"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/go-redis/redis"
)

// Token is play load part of JWT token
type Token struct {
	AuthContextID string `json:"auth_context_id"`
	*jwt.StandardClaims
}

// AuthContext is authentication information
type AuthContext struct {
	Status string   `json:"status"`
	JWT    JWTToken `json:"jwt"`
}

type JWTToken struct {
	Secret string `json:"secret"`
	Token  string `json:"token"`
}

// GenerateNewToken will generate jwt token for given user id
func GenerateNewToken(userID int) string {
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()
	newAuthTokenID := generateAuthContextID(userID)

	token := &Token{
		AuthContextID: newAuthTokenID,
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
	defer redisClient.Close()

	authContext := &AuthContext{
		Status: "OK",
		JWT: JWTToken{
			Secret: "my-secret-code",
			Token:  tokenString,
		},
	}

	authContextJSON, err := json.Marshal(authContext)
	if err != nil {
		panic(err)
	}

	err = redisClient.Set("cas:auth_context:"+newAuthTokenID, authContextJSON, 0).Err()
	if err != nil {
		panic(err)
	}

	return tokenString
}

// ValidateToken will validate token string and return
// When token string is valid
func ValidateToken(tokenString string) (*jwt.Token, error) {
	var returnError error

	jwtToken, err := ParseTokenString(tokenString)
	if err != nil {
		return nil, err
	}

	authContext := GetAuthContextByID(jwtToken.AuthContextID)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(authContext.JWT.Secret), nil
	})

	if token.Valid {
		return token, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			returnError = errors.New("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			returnError = errors.New("Timing is everything")
		} else {
			returnError = errors.New("Couldn't handle this token")
		}
	} else {
		returnError = errors.New("Couldn't handle this token")
	}
	return nil, returnError
}

// ParseTokenString will parse given token and return Token struct
func ParseTokenString(tokenString string) (*Token, error) {
	// parse Claims
	var claimBytes []byte
	var err error
	var tokenClaim Token

	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, errors.New("Token contains an invalid number of segments")
	}

	if claimBytes, err = jwt.DecodeSegment(parts[1]); err != nil {
		return nil, errors.New("Malformed claim")
	}
	if err = json.Unmarshal(claimBytes, &tokenClaim); err != nil {
		return nil, errors.New("Malformed claim")
	}

	return &tokenClaim, nil
}

// GetAuthContextByID will check token existence in redis and valid
func GetAuthContextByID(authContextID string) *AuthContext {
	var authContext AuthContext

	// TODO: Move it to lib or service package
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer redisClient.Close()

	contextJSON, err := redisClient.Get("cas:auth_context:" + authContextID).Result()

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(contextJSON), &authContext)
	if err != nil {
		panic(err)
	}

	return &authContext
}

func generateAuthContextID(userID int) string {
	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(strconv.Itoa(userID) + strconv.FormatInt(time.Now().Unix(), 10)))
	hashString := hex.EncodeToString(sha1Hash.Sum(nil))
	return hashString
}
