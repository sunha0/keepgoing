package main

import (
	"fmt"

	// "github.com/gin-gonic/gin"
	"github.com/go_demo/core"
	"github.com/go_demo/global"
)

func main() {

	global.VP = core.Viper()

	fmt.Println(global.VP.GetString("jwt.signing-key"))
	fmt.Println(global.VP.GetStringMap("email"))

	fmt.Println(global.Conf.JWT.BufferTime)

}
