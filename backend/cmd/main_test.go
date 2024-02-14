package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"team-work-space/api/route"
	"team-work-space/bootstrap"
	_ "team-work-space/docs"
	"team-work-space/domain"
	"team-work-space/repository"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	_ = os.Chdir("..")
	logrus.SetFormatter(&logrus.JSONFormatter{})
	timeout := time.Duration(10) * time.Second

	// Upping the mongo container
	pool, err := dockertest.NewPool("")
	if err != nil {
		panic("Could not connect to Docker: " + err.Error())
	}
	options := &dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "latest",
		Name:       "db-team-work-space-test",
		Env: []string{
			"MONGO_INITDB_ROOT_USERNAME=admin",
			"MONGO_INITDB_ROOT_PASSWORD=123321",
		},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"27017/tcp": {
				{HostIP: "", HostPort: "6464"},
			},
		},
	}
	resource, err := pool.RunWithOptions(options)
	if err != nil {
		panic("Could not start resource: " + err.Error())
	}
	_, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Setup the app
	app := bootstrap.App("test")
	env := app.Env
	db := app.Mongo.Database(env.DBname)

	// New Repositories
	AuthRepository = repository.NewAuthRepository(db, "users")

	gin := gin.Default()
	gin.Use(bootstrap.CORS())

	route.SetupRoute(env, timeout, db, gin)
	fmt.Print("Running in test enviroment.")
	exitCode := m.Run()

	if err := pool.Purge(resource); err != nil {
		panic("Could not purge resource: " + err.Error())
	}

	os.Exit(exitCode)
}

var AuthRepository domain.AuthRepository

// Auth repo
func TestAuthRepositoryCreate(t *testing.T) {
	var user domain.User
	ctx := context.Background()

	user.Email = "test@gmail.com"
	user.Username = "Testusername"
	user.Password = "123321123321"
	user.Role = "User"
	user.ID = primitive.NewObjectID()
	err := AuthRepository.Create(ctx, user)
	assert.NoError(t, err)
}

var testuser domain.User

func TestAuthRepositoryGetByEmail(t *testing.T) {
	ctx := context.Background()
	email := "test@gmail.com"
	user, err := AuthRepository.GetByEmail(ctx, email)
	testuser = user
	assert.NoError(t, err)
	assert.Equal(t, user.Email, email)
}

func TestAuthRepositoryGetById(t *testing.T) {
	ctx := context.Background()
	user, err := AuthRepository.GetById(ctx, testuser.ID.Hex())
	assert.NoError(t, err)
	assert.Equal(t, user.ID, testuser.ID)
}
