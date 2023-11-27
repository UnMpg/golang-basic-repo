package routes

import (
	"go-project/api/handler"
	"go-project/api/repository"
	"go-project/api/usecase"
	"go-project/config"
	"go-project/db"
	"go-project/middleware"
	"go-project/utils/log"

	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CreateRoute(isDev bool) *gin.Engine {
	urlSwagger := ginSwagger.URL(config.AppEnv.Host + config.AppEnv.Port + "/swagger/doc.json")
	router := gin.New()
	router.Use(requestid.New())
	router.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		middleware.HandlePanic(c, err)
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, urlSwagger))

	router.Use(middleware.CORSMiddleware())
	router.Use(requestid.New())
	router.Use(healthcheck.Default())
	router.Use(log.RequestLoggerActivity())
	return router
}

func RouteV1(r *gin.Engine) {
	dbPg, err := db.GetConnectionDB()
	if err != nil {
		log.Log.Error("Error to Get Connection DB")
	}

	V1UserRoute := r.Group("/user/api")
	// V1UserRoute.Use(middleware.UserMiddleware())

	UserRepositoty := repository.NewUserRepository(dbPg)
	UserUsecase := usecase.NewUserUsecase(UserRepositoty)
	handler.NewUserHandler(V1UserRoute, UserUsecase)
}
