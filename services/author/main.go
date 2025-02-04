package main

import (
	"github.com/zikrykr/library-management/services/author/cmd/rest"
	appSetup "github.com/zikrykr/library-management/services/author/cmd/setup"
	"github.com/zikrykr/library-management/services/author/config"
)

func main() {
	// config init
	config.InitConfig()

	// app setup init
	setup := appSetup.InitSetup()

	// starting REST server
	rest.StartServer(setup)
}
