package route

import (
	"github.com/gin-gonic/gin"
	"team-work-space/api/controller"
	"team-work-space/bootstrap"
	"team-work-space/domain"
	"team-work-space/mongo"
	"team-work-space/repository"
	"team-work-space/usecase"
	"time"
)

func NewTeamPublicRouter(env *bootstrap.Env, db mongo.Database, timeout time.Duration, group *gin.RouterGroup) {
	tr := repository.NewTeamRepository(db, domain.CollectionTeam)
	tc := &controller.TeamController{
		TeamUsecase: usecase.NewTeamUsecase(tr, timeout),
	}
	group.GET("/teams", tc.GetAll)
	group.GET("/teams/:id", tc.GetById)
}
func NewTeamProtectedRouter(env *bootstrap.Env, db mongo.Database, timeout time.Duration, group *gin.RouterGroup) {
	tr := repository.NewTeamRepository(db, domain.CollectionTeam)
	tc := &controller.TeamController{
		TeamUsecase: usecase.NewTeamUsecase(tr, timeout),
	}
	group.POST("/teams", tc.Create)
	group.DELETE("/teams/:id", tc.Delete)
	group.PUT("/teams/:id", tc.Update)
}
