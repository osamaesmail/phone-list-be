package router

import (
	"phone-list/app/handler"
	"phone-list/app/repository"
	"phone-list/app/service"
	"phone-list/database"

	_ "phone-list/docs"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter initialize routing information
// @title phone list API
// @version 1.0
// @description phone list API example
// @BasePath /api/v1.
func InitRouter(r *gin.Engine) {
	customerRepository := repository.NewCustomerRepository(database.DB())
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := r.Group("/api/v1")
	apiV1.GET("/customers", customerHandler.List())
}
