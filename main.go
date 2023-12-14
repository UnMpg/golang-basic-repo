package main

import (
	"go-project/config"
	"go-project/middleware"
	"go-project/models"
	"go-project/routes"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @title Your API Title
// @version 1.0
// @description Your API description
// @host localhost:8080
// @BasePath /v1
func main() {
	var isDev bool

	if config.AppEnv.Env == "PROD" {
		isDev = false
	} else {
		isDev = true
	}

	r := routes.CreateRoute(isDev)

	// @Summary Add a new pet to the store
	// @Description get string by ID
	// @ID get-string-by-int
	// @Accept  json
	// @Produce  json
	// @Param   some_id     path    int     true        "Some ID"
	// @Success 200 {string} string  "ok"
	// @Router /string/{some_id} [get]
	r.GET("/", func(c *gin.Context) {
		message := "Welcome to golang Project"
		res := models.Response{RespCode: http.StatusOK, RespMessage: message, Status: "success"}
		c.JSON(http.StatusOK, res)
	})
	r.GET("/coba", func(c *gin.Context) {
		message := "Welcome to golang Project"
		res := models.Response{RespCode: http.StatusOK, RespMessage: message, Status: "success"}
		c.JSON(http.StatusOK, res)
	})

	r.HandleMethodNotAllowed = true
	r.NoMethod(middleware.HandleNoMethod)
	r.NoRoute(middleware.HandleNoRoute)

	routes.RouteV1(r)

	timeout := time.Duration(config.AppEnv.Timeout) * time.Second
	newHandler := http.TimeoutHandler(r, timeout, "Timeout!")

	server := &http.Server{
		Addr:         config.AppEnv.Port,
		Handler:      newHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Printf("Server Ready To Handler On PORT  :%v", config.AppEnv.Port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("Could not listen on %v : %v", config.AppEnv.Port, err)
	}

}
