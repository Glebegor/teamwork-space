package controller

import (
	"net/http"
	"team-work-space/domain"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func (uc *UserController) GetAll(c *gin.Context) {
	var data []domain.User
	c.JSON(http.StatusOK, data)
}
func (uc *UserController) GetById(c *gin.Context) {
	var data domain.User
	c.JSON(http.StatusOK, data)
}
func (uc *UserController) GetByEmail(c *gin.Context) {
	var data domain.User

	c.JSON(http.StatusOK, data)
}
func (uc *UserController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, domain.SuccessResponse{Status: "ok"})
}
func (uc *UserController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, domain.SuccessResponse{Status: "ok"})
}
