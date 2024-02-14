package route

import (
	"github.com/gin-gonic/gin"
	"team-work-space/api/controller"
	"team-work-space/bootstrap"
	"team-work-space/domain"
	"team-work-space/mongo"
	"team-work-space/repository"
	"time"
)

func NewTeamPublicRouter(env *bootstrap.Env, db mongo.Database, timeout time.Duration, group *gin.RouterGroup) {
	tr := repository.NewTeamRepository(db, domain.CollectionTeam)
	tc := &controller.TeamController{
		//TeamUsecase: usecase.NewTeamUsecase(tr, timeout),
	}
}
