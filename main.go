package main

import (

	//"os"
	//"github.com/gin-gonic/gin"
	"github.com/keepgoing/core"
	"github.com/keepgoing/database"
	"github.com/keepgoing/global"
)

func main() {
	// 设置读取config文件

	// os.Setenv("CONF", "config.test.yaml")
	// gin.SetMode(gin.TestMode)

	global.VP = core.Viper() // 初始化Viper

	/*
		获取配置文件值
		fmt.Println(global.VP.GetString("jwt.signing-key"))
		fmt.Println(global.VP.GetStringMap("email"))
		或者
		fmt.Println(global.Conf.JWT.BufferTime)
	*/

	// 连接Mysql 数据库
	global.DB = database.GormMysql()

	if global.DB != nil {
		database.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}

	core.RunServer()

}
