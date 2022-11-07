package utils

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

type MetaToken struct {
	UserName      string
	FullName      string
	Email         string
	Role          string
	ExpiredAt     time.Time
	Authorization bool
}

type AccessToken struct {
	Claims MetaToken
}

func Sign(Data map[string]interface{}, SecretPublicKeyEnvName string, ExpiredAt time.Duration) (string, error) {

	expiredAt := time.Now().Add(time.Minute * ExpiredAt).Unix()

	jwtSecretKey := os.Getenv(SecretPublicKeyEnvName)

	claims := jwt.MapClaims{}
	claims["exp"] = expiredAt
	claims["authorization"] = true

	for i, v := range Data {
		claims[i] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(jwtSecretKey))

	if err != nil {
		logrus.Error(err.Error())
		return accessToken, err
	}

	return accessToken, nil
}

func VerifyTokenHeader(ctx *gin.Context, SecretPublicKeyEnvName string) (*jwt.Token, error) {
	tokenHeader := ctx.GetHeader("Authorization")
	accessToken := strings.SplitAfter(tokenHeader, "Bearer")[1]
	jwtSecretKey := os.Getenv(SecretPublicKeyEnvName)

	token, err := jwt.Parse(strings.Trim(accessToken, " "), func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

func ExtractClaimValueFromToken(ctx *gin.Context, claim string) string {
	claims, _ := ctx.Get("user")
	values := claims.(jwt.MapClaims)
	username := values[claim].(string)

	return username
}

func VerifyToken(accessToken, SecretPublicKeyEnvName string) (*jwt.Token, error) {
	jwtSecretKey := os.Getenv(SecretPublicKeyEnvName)

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

func DecodeToken(accessToken *jwt.Token) AccessToken {
	var token AccessToken
	stringify, _ := json.Marshal(&accessToken)

	json.Unmarshal(stringify, &token)

	return token
}
