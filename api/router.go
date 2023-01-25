package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/uzbekman2005/real_time_chatting/api/docs"
	v1 "github.com/uzbekman2005/real_time_chatting/api/handlers/v1"
	"github.com/uzbekman2005/real_time_chatting/api/middleware"
	token "github.com/uzbekman2005/real_time_chatting/api/tokens"
	"github.com/uzbekman2005/real_time_chatting/config"
	"github.com/uzbekman2005/real_time_chatting/pkg/logger"
	"github.com/uzbekman2005/real_time_chatting/storage"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	CasbinEnforcer *casbin.Enforcer
	Storage        storage.IStorage
}

// New ...
// @title           Student API
// @version         1.0
// @description     This is OPEN API that you get more information about universeties.
// @termsOfService  2 term adds uz

// @contact.name   Azizbek
// @contact.url    https://t.me/azizbek_dev_2005
// @contact.email  azizbekhojimurodov@gmail.com

// @host 13.229.54.192:8000
// @BasePath  /v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	jwtHandler := token.JWTHandler{
		SigninKey: option.Conf.SignInKey,
		Log:       option.Logger,
	}

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:     option.Logger,
		Cfg:        option.Conf,
		Storage:    option.Storage,
		Jwthandler: jwtHandler,
	})

	router.Use(middleware.NewAuth(option.CasbinEnforcer, jwtHandler, config.Load()))

	corConfig := cors.DefaultConfig()
	corConfig.AllowAllOrigins = true
	corConfig.AllowCredentials = true
	corConfig.AllowHeaders = []string{"*"}
	corConfig.AllowBrowserExtensions = true
	corConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corConfig))

	router.GET("/", handlerV1.AppIsRunning)
	api := router.Group("/v1")
	api.GET("/", handlerV1.AppIsRunning)
	// Admin
	api.GET("/admin/login/:admin_name/:password", handlerV1.AdminLogin)
	api.POST("/admin", handlerV1.CreateAdmin)
	api.GET("/admin/all", handlerV1.GetAllAdmins)
	api.PUT("/admin", handlerV1.UpdateAdmin)
	api.DELETE("/admin", handlerV1.DeleteAdmin)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
