package route

import (
	"fmt"
	"team-work-space/api/middleware"
	"team-work-space/bootstrap"
	"team-work-space/mongo"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	fmt.Print("PUBLIC ROUTES...\n")
	publicRouter := gin.Group("api/v1")
	NewAuthRouter(env, db, timeout, publicRouter)
	NewUserPublicRouter(env, db, timeout, publicRouter)

	fmt.Print("PROTECT ROUTES...\n")
	protectedRouter := gin.Group("api/v2")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.SERVERsecret))
	NewUserProtectRouter(env, db, timeout, protectedRouter)

	fmt.Print("OPEN ROUTES...\n")
	// openRouter := gin.Group("api/v3")

	// Add docs
	if env.SERVERenv == "development" {
		fmt.Print("DOCS ROUTES...\n")
		// To see main page write index.html
		gin.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
