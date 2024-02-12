package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"team-work-space/api/route"
	"team-work-space/bootstrap"
	"time"

	"github.com/gin-gonic/gin"

	_ "team-work-space/docs"
)

// @title Team work space Service API
func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	app := bootstrap.App("config.yml")
	env := app.Env
	db := app.Mongo.Database(env.DBname)

	timeout := time.Duration(10) * time.Second

	gin := gin.Default()
	gin.Use(bootstrap.CORS())

	route.SetupRoute(env, timeout, db, gin)
	if env.SERVERenv == "development" {
		fmt.Print("Running in development enviroment.\n	")
	}
	gin.Run(":" + env.SERVERport)

}
