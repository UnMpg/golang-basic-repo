package middleware

import (
	"errors"
	"go-project/models"
	jwttoken "go-project/utils/jwt_token"
	"go-project/utils/log"
	"go-project/utils/redis"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		MiddlewareUserCek(ctx)
	}
}

func MiddlewareUserCek(c *gin.Context) {
	token, err := GetTokenHeader(c)
	if err != nil {
		c.Header("WWW-Authenticate", "JWT realm=jwt-token")
		c.Abort()
		c.JSON(http.StatusUnauthorized, models.CreateResponse(http.StatusUnauthorized, "Failed", "Token UnAuthorization", nil))
		return
	}

	getUserId, err := jwttoken.ValidateTokenHeader(token)
	if err != nil {
		c.Header("WWW-Authenticate", "JWT realm=jwt-token")
		c.Abort()
		c.JSON(http.StatusUnauthorized, models.CreateResponse(http.StatusUnauthorized, "failed", "UnAutorization", nil))
		return
	}
	client, err := redis.GetCacheConnection()
	if err != nil {
		log.Log.Error("Error to connection Redis Middleware", err)
	}

	_, err = client.GetValue(token)
	if err != nil {
		c.Header("WWW-Authenticate", "JWT realm=jwt-token")
		c.Abort()
		c.JSON(http.StatusUnauthorized, models.CreateResponse(http.StatusUnauthorized, "Failed", "Token UnAuthorization", nil))
		return
	}

	c.Set("userId", getUserId)
	c.Set("token", token)
}

func GetTokenHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", errors.New("unknow Auth Bearer Token")
	}
	return parts[1], nil
}
