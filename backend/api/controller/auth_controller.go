package controller

import (
	"net/http"
	"team-work-space/domain"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthUsecase domain.AuthUsecase
}

func (ac *AuthController) Login(c *gin.Context) {
	c.JSON(http.StatusOK, domain.SuccessResponse{Status: "ok"})
}

func (ac *AuthController) Reg(c *gin.Context) {
	var input domain.Reg
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	if err := ac.AuthUsecase.Register(c, input); err != nil {
		c.JSON(http.StatusBadGateway, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Status: "ok"})
}
