package auth_tool

import (
	"errors"
	"github.com/Kambar-ZH/simple-service/pkg/conf"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const AccessTokenKey = "access_token"

func SetToken(ctx *gin.Context, token *jwt.Token) {
	ctx.Set(AccessTokenKey, token)
}

func GetAccessToken(ctx *gin.Context) *jwt.Token {
	token, ok := ctx.Value(AccessTokenKey).(*jwt.Token)
	if !ok {
		return nil
	}
	if !token.Valid {
		return nil
	}
	return token
}

func GetCurrenUserID(ctx *gin.Context) int64 {
	return GetCurrenUserIDByToken(GetAccessToken(ctx))
}

func GetCurrenUserIDByToken(token *jwt.Token) int64 {
	if token == nil {
		return 0
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0
	}
	return int64(userID)
}

var ErrUnauthorized = errors.New("unauthorized request")

func ParseToken(tokenString string) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (obj interface{}, err error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrUnauthorized
		}
		return []byte(conf.GlobalConfig.JWT.SecretKey), nil
	})
	if err != nil {
		return
	}
	return
}

func GenerateTokenPair(userID int64, email string) (access, refresh string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["sub"] = email
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	access, err = token.SignedString([]byte(conf.GlobalConfig.JWT.SecretKey))
	if err != nil {
		return
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["user_id"] = userID
	rtClaims["sub"] = email
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	refresh, err = refreshToken.SignedString([]byte(conf.GlobalConfig.JWT.SecretKey))
	if err != nil {
		return
	}

	return
}
