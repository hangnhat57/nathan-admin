package web

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()

	handler := NewHandler( /* pass your service dependencies here */ )
	handler.RegisterRoutes(router)

	return router
}
