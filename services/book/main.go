package main

import (
	"github.com/zikrykr/library-management/services/book/cmd/rest"
	appSetup "github.com/zikrykr/library-management/services/book/cmd/setup"
	"github.com/zikrykr/library-management/services/book/config"
)

func main() {
	// config init
	config.InitConfig()

	// app setup init
	setup := appSetup.InitSetup()

	// starting REST server
	rest.StartServer(setup)
}
