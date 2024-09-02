package routers

import (
	"github.com/erfantak/golang-web-api/api/handlers"
	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.RouterGroup) {
	handler := handlers.NewTestHandler()

	r.GET("/user", handler.Test)
}
