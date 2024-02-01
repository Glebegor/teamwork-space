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
	group.GET("/users", uc.GetById)
	group.GET("/users/getById/", uc.GetById)
	group.GET("/users/getByEmail/", uc.GetByEmail)
}

func NewUserProtectRouter(env *bootstrap.Env, db mongo.Database, timeout time.Duration, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	uc := &controller.UserController{
		UserUsecase: usecase.NewUserUsecase(ur, timeout),
	}
	group.DELETE("/users", uc.Delete)
	group.PUT("/users", uc.Update)
}
