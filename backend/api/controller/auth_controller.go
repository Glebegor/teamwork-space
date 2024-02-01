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
	c.JSON(http.StatusOK, domain.SuccessResponse{Status: "ok"})
}
