package router

import (
	"github.com/gin-gonic/gin"
	"github.com/keepgoing/global"
	"github.com/keepgoing/router/system"
)

// 路由组
type RouterGroup struct {
	system.UserRouter
}

func InitRouters() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		router.Use(gin.Logger())
	}

	// 用户路由
	systemRouter := new(RouterGroup)

	PublicGroup := router.Group(global.Conf.System.RouterPrefix)
	{

		systemRouter.InitUserRouter(PublicGroup)

	}

	return router

}
