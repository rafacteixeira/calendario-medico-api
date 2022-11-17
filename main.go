package main

import (
	"github.com/rafacteixeira/calendario-medico-api/database"
	"github.com/rafacteixeira/calendario-medico-api/routes"
	"github.com/rafacteixeira/calendario-medico-api/settings"
)

func main() {
	settings.LoadEnv()
	database.DBStart()
	routes.StartServer()
}
