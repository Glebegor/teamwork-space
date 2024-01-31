package route

import (
	"team-work-space/bootstrap"
	"team-work-space/mongo"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("api/v1")
	NewUserPublicRouter(env, db, timeout, publicRouter)
	NewUserProtectRouter(env, db, timeout, publicRouter)
	// protectedRouter := gin.Group("api/v2")
	// openRouter := gin.Group("api/v3")
}
