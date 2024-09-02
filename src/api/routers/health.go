package routers

import (
	"github.com/erfantak/golang-web-api/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	r.GET("/health", handler.HealthHandler)
	r.GET("/health/:id", handler.HealthById)
}
