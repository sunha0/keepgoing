package main

import (
	"fmt"
	//"os"
	//"github.com/gin-gonic/gin"
	"github.com/go_demo/core"
	"github.com/go_demo/global"
)

func main() {
	// 设置读取config文件

	// os.Setenv("CONF", "config.test.yaml")
	// gin.SetMode(gin.TestMode)

	global.VP = core.Viper()

	/*
		获取配置文件值
		fmt.Println(global.VP.GetString("jwt.signing-key"))
		fmt.Println(global.VP.GetStringMap("email"))
		或者
		fmt.Println(global.Conf.JWT.BufferTime)
	*/
	fmt.Println(global.VP.GetString("jwt.signing-key"))
	fmt.Println(global.VP.GetStringMap("email"))

	fmt.Println(global.Conf.JWT.BufferTime)

}
