package core

import (
	"github.com/keepgoing/router"
)

func RunServer() {

	router := router.Routers()
	router.Run(":8080")

}
