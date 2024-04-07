package system

import (
	"github.com/gin-gonic/gin"
	"github.com/keepgoing/api/system"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	Router.GET("/test", system.Login)

	Router.GET("/register", system.RegisterUser)

}
