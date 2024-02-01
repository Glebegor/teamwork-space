package route

import (
	"fmt"
	"team-work-space/bootstrap"
	"team-work-space/mongo"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	fmt.Print("PUBLIC ROUTES...\n")
	publicRouter := gin.Group("api/v1")
	NewAuthRouter(env, db, timeout, publicRouter)
	NewUserPublicRouter(env, db, timeout, publicRouter)

	fmt.Print("PROTECT ROUTES...\n")
	protectedRouter := gin.Group("api/v2")
	NewUserProtectRouter(env, db, timeout, protectedRouter)

	fmt.Print("OPEN ROUTES...\n")
	// openRouter := gin.Group("api/v3")
}
