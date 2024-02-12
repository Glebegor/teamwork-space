package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"team-work-space/api/route"
	"team-work-space/bootstrap"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	_ "team-work-space/docs"
)

func TestMain(m *testing.M) {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	app := bootstrap.App("tests")
	env := app.Env
	db := app.Mongo.Database(env.DBname)

	timeout := time.Duration(10) * time.Second

	gin := gin.Default()
	gin.Use(bootstrap.CORS())

	route.SetupRoute(env, timeout, db, gin)
	fmt.Print("Running in test enviroment.\n	")
	gin.Run(":" + env.SERVERport)

}
