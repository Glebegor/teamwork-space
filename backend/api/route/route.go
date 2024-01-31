package route

import (
	"team-work-space/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRoute(env *bootstrap.Env, timeout time.Duration, db any, gin *gin.Engine) {
	// publicRouter := gin.Group("api/v1")
	// protectedRouter := gin.Group("api/v2")
	// openRouter := gin.Group("api/v3")
}
