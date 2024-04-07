package main

import (

	//"os"
	//"github.com/gin-gonic/gin"

	"fmt"

	"github.com/gofrs/uuid/v5"
	"github.com/keepgoing/core"
	"github.com/keepgoing/global"
	"github.com/keepgoing/utils"
)

func main() {
	/* 设置读取config文件
	   os.Setenv("CONF", "config.test.yaml")
	   gin.SetMode(gin.TestMode)
	*/

	global.VP = core.Viper() // 初始化Viper 读取配置

	/*
		获取配置文件值
		fmt.Println(global.VP.GetString("jwt.signing-key"))
		fmt.Println(global.VP.GetStringMap("email"))
		或者
		fmt.Println(global.Conf.JWT.BufferTime)
	*/

	// 连接Mysql 数据库
	// global.DB = database.GormMysql()

	// if global.DB != nil {
	// 	database.RegisterTables() // 初始化表
	// 	// 程序结束前关闭数据库链接
	// 	db, _ := global.DB.DB()
	// 	defer db.Close()
	// }

	// core.RunServer()

	jwt := utils.NewJWT()

	a := jwt.CreateClaims(utils.BaseClaims{
		UUID:        uuid.Must(uuid.NewV4()),
		ID:          1,
		Username:    "keepgoing",
		NickName:    "keepgoing",
		AuthorityId: 1,
	})

	fmt.Println(1111, a)
	b, _ := jwt.CreateToken(a)
	fmt.Println(b)

}
