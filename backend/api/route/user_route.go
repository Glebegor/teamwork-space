package route

import (
	"team-work-space/api/controller"
	"team-work-space/bootstrap"
	"team-work-space/domain"
	"team-work-space/mongo"
	"team-work-space/repository"
	"team-work-space/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewUserPublicRouter(env *bootstrap.Env, db mongo.Database, timeout time.Duration, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	uc := &controller.UserController{
		UserUsecase: usecase.NewUserUsecase(ur, timeout),
	}
	group.POST("/users", uc.GetById)
	group.POST("/users", uc.GetByEmail)
}

func NewUserProtectRouter(env *bootstrap.Env, db mongo.Database, timeout time.Duration, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	uc := &controller.UserController{
		UserUsecase: usecase.NewUserUsecase(ur, timeout),
	}
	group.PUT("/users/MyInfo", uc.GetById)
	group.PUT("/users", uc.GetById)
}
