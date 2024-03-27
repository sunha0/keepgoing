package main

import (
	"fmt"
	//"os"
	//"github.com/gin-gonic/gin"
	"github.com/keepgoing/core"
	"github.com/keepgoing/global"
	"github.com/keepgoing/db"
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
	global.DB = db.GormMysql()


	if global.DB != nil {
		db.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}

	fmt.Println(global.VP.GetString("jwt.signing-key"))
	fmt.Println(global.VP.GetStringMap("email"))

	fmt.Println(global.Conf.JWT.BufferTime)
	fmt.Println(global.Conf.Mysql.GeneralDB.Prefix, global.Conf.Mysql.GeneralDB.Singular)

}
