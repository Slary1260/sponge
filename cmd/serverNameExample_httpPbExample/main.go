// Package main is the http server of the application.
package main

import (
	"github.com/zhufuyi/sponge/pkg/app"

	"github.com/zhufuyi/sponge/cmd/serverNameExample_httpPbExample/initial"
)

func main() {
	initial.InitApp()
	services := initial.CreateServices()
	closes := initial.Close(services)

	a := app.New(services, closes)
	a.Run()
}
