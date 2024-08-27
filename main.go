package main

import (
	"loan-api/cmd"
	"loan-api/delivery/router"

	"github.com/gin-gonic/gin"
)

func main() {
	userController, loanController, env := cmd.InitializeDepenencies()
	r := gin.Default()
	router.SetRouter(r, userController, loanController, env)
	r.Run(env.ServerAddress)
}
