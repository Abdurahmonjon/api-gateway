package api

import (
	v1 "github.com/Abdurahmonjon/api-gateway/api/handlers/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Abdurahmonjon/api-gateway/config"
	"github.com/Abdurahmonjon/api-gateway/pkg/logger"
	"github.com/Abdurahmonjon/api-gateway/services"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

func New(option Option) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            config.Config{},
	})

	api := router.Group("/v1")
	api.POST("/student", handlerV1.CreateStudent)
	api.GET("student/:username", handlerV1.GetStudent)
	api.GET("/students", handlerV1.ListStudents)
	api.PUT("/student/:username", handlerV1.UpdateStudent)
	api.DELETE("/student/:id", handlerV1.DeleteStudent)

	url := ginSwagger.URL("swagger/docs.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}