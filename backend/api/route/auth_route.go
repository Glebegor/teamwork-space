package route

import (
	"team-work-space/api/controller"
	"team-work-space/api/middleware/validators"
	"team-work-space/bootstrap"
	"team-work-space/domain"
	"team-work-space/mongo"
	"team-work-space/repository"
	"team-work-space/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewAuthRouter(env *bootstrap.Env, db mongo.Database, timeout time.Duration, group *gin.RouterGroup) {
	ar := repository.NewAuthRepository(db, domain.AuthCollection)
	ac := &controller.AuthController{
		AuthUsecase: usecase.NewAuthUsecase(ar, timeout),
		Env:         env,
	}
	group.POST("/auth/login", ac.Login)
	group.POST("/auth/registration", validators.RegValidator(), ac.Reg)
	group.POST("/auth/refresh", ac.Refresh)
}
