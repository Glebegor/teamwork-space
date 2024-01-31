package main

import (
	"team-work-space/api/route"
	"team-work-space/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.Mongo.Database(env.DBname)

	timeout := time.Duration(10) * time.Second

	gin := gin.Default()
	gin.Use(bootstrap.CORS())
	gin.Run(":" + env.SERVERport)

	route.SetupRoute(env, timeout, db, gin)
}
