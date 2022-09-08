package router

import (
	"example.com/m/v2/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	Routers := gin.Default()
	Routers.Use(middlewares.Cors())

	Routers.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "基础完成",
		})
	})

	return Routers
}
