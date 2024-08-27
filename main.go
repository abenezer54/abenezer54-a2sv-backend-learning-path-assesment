package main

import (
	"loan-api/cmd"
	"loan-api/delivery/router"
	"loan-api/infrastructure/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	defer app.CloseDBConnection()
	env := app.Env

	userController, loanController, logController := cmd.InitializeDepenencies(app)
	r := gin.Default()
	router.SetRouter(r, userController, loanController, logController, env)
	r.Run(env.ServerAddress)
}
