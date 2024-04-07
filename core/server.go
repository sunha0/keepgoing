package core

import (
	"fmt"
	"github.com/keepgoing/router"
	"github.com/keepgoing/global"
)

func RunServer() {

	server := router.InitRouters()

	addr := fmt.Sprintf(":%d",global.Conf.System.Addr)
	server.Run(addr)
}
 