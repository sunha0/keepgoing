package router

import (

	"github.com/gin-gonic/gin"
	"github.com/keepgoing/global"
)


func InitRouters() *gin.Engine{
	router := gin.New()
	router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		router.Use(gin.Logger())
	}

	PublicGroup := router.Group(global.Conf.System.RouterPrefix)
	{
		PublicGroup.GET("/test",func(c *gin.Context) {
			c.JSON(200, "ok")
		})

	}

	 return router


	
}