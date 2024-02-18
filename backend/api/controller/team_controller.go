package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"team-work-space/domain"
)

type TeamController struct {
	TeamUsecase domain.TeamUsecase
}

func (tc *TeamController) GetAll(c *gin.Context) {
	data, err := tc.TeamUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusBadGateway, domain.ErrorResponse{Message: err.Error()})
		c.Abort()
	}
	c.JSON(http.StatusOK, data)
}
func (tc *TeamController) GetById(c *gin.Context) {
	var data domain.Team
	c.JSON(http.StatusOK, data)
}
func (tc *TeamController) Create(c *gin.Context) {
	var data domain.Team
	err := tc.TeamUsecase.Create(c, data)
	if err != nil {
		c.JSON(http.StatusBadGateway, domain.ErrorResponse{Message: err.Error()})

		c.Abort()
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Status: "ok"})
}
func (tc *TeamController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, domain.SuccessResponse{Status: "ok"})
}
func (tc *TeamController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, domain.SuccessResponse{Status: "ok"})
}
