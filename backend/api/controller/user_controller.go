package controller

import (
	"team-work-space/domain"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func (uc *UserController) GetById(c *gin.Context)    {}
func (uc *UserController) GetByEmail(c *gin.Context) {}
func (uc *UserController) Update(c *gin.Context)     {}
