package repository

import (
	"context"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	_ = os.Chdir("..")
	timeout := time.Duration(10) * time.Second

	// Upping the mongo container
	pool, err := dockertest.NewPool("")
	if err != nil {
		panic("Could not connect to Docker: " + err.Error())
	}
	options := &dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "latest",
		Name:       "db-team-work-space-test2",
		Env: []string{
			"MONGO_INITDB_ROOT_USERNAME=admin",
			"MONGO_INITDB_ROOT_PASSWORD=123321",
		},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"27017/tcp": {
				{HostIP: "", HostPort: "7979"},
			},
		},
	}
	resource, err := pool.RunWithOptions(options)
	if err != nil {
		panic("Could not start resource: " + err.Error())
	}
	_, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := pool.Purge(resource); err != nil {
		panic("Could not purge resource: " + err.Error())
	}
}
