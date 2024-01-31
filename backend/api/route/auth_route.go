package route

import (
	"team-work-space/bootstrap"
	"team-work-space/mongo"
	"time"

	"github.com/gin-gonic/gin"
)

func NewAuthRouter(env *bootstrap.Env, db mongo.Database, timeout time.Duration, group *gin.RouterGroup) {
	// ar := reposi
}
