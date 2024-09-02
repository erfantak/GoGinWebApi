package api

import (
	"fmt"
	"github.com/erfantak/golang-web-api/api/routers"
	"github.com/erfantak/golang-web-api/config"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	//r1 := gin.Default()

	r.Use(gin.Logger(), gin.Recovery())

	api := r.Group("/api")
	v1 := api.Group("/api/v1")
	{
		routers.Health(v1)
		routers.TestRouter(v1.Group("test"))
	}
	err := r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		return
	}
}
