package middleware

import (
	"errors"
	"go-project/config"
	"go-project/db"
	"go-project/models"
	"go-project/utils/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	if config.AppEnv.Env == "PROD" {

		log.Log.Printf("Starting %s on Production Environment", config.AppEnv.AppName)
	} else {
		log.Log.Printf("Starting %s on Production Environment", config.AppEnv.AppName)
	}

	log.SetupLogger()

	if err := db.InitConnectionDB(); err != nil {
		log.Log.Error("Error Connection DB ", err.Error())
	}

}

func HandlePanic(c *gin.Context, err interface{}) {
	c.JSON(http.StatusInternalServerError, models.CreateResponse(http.StatusInternalServerError, "failed", "Internal server Error", nil))

}

func HandleNoMethod(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, models.CreateResponse(http.StatusMethodNotAllowed, "failed", "Undefined Method", nil))
}

func HandleNoRoute(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, models.CreateResponse(http.StatusMethodNotAllowed, "failed", "Undefined Method", nil))
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func GetUserId(c *gin.Context) (string, error) {
	GetUserId, valid := c.Get("UserId")
	if !valid {
		return "", errors.New("undifined User ID")
	}
	return GetUserId.(string), nil
}
