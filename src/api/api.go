package api

import (
	"fmt"

	"github.com/amirmalekian/Mashinchi/src/api/middleware"
	"github.com/amirmalekian/Mashinchi/src/api/router"
	"github.com/amirmalekian/Mashinchi/src/api/validation"
	"github.com/amirmalekian/Mashinchi/src/config"
	"github.com/amirmalekian/Mashinchi/src/docs"
	"github.com/amirmalekian/Mashinchi/src/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	r := gin.New()

	RegisterValidators()

	r.Use(middleware.DefaultStructuredLogger(cfg))
	r.Use(middleware.Cors(cfg))
	r.Use(gin.Logger(), gin.CustomRecovery(middleware.ErrorHandler) /*middleware.TestMiddleware()*/, middleware.LimitByRequest())

	RegisterRoutes(r, cfg)
	RegisterSwagger(r, cfg)
	logger := logging.NewLogger(cfg)
	logger.Info(logging.General, logging.Startup, "Started", nil)
	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		// Test
		health := v1.Group("/health")
		test_router := v1.Group("/test" /*middleware.Authentication(cfg), middleware.Authorization([]string{"admin"})*/)

		// User
		users := v1.Group("/users")

		// Base
		countries := v1.Group("/countries", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		cities := v1.Group("/cities", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		files := v1.Group("/files", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))

		// Test
		router.Health(health)
		router.TestRouter(test_router)

		// User
		router.User(users, cfg)

		// Base
		router.Country(countries, cfg)
		router.City(cities, cfg)
		router.File(files, cfg)
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		router.Health(health)
	}
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validation.IranianPhoneNumberValidator, true)
		val.RegisterValidation("password", validation.PasswordValidator, true)
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
