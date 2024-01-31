package main

import (
	"team-work-space/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.Mongo.Database(env.DBname)

	timeout := time.Second * 10

	gin := gin.Default()
	gin.Use(bootstrap.CORS())
	gin.Run(":" + env.SERVERport)
}
