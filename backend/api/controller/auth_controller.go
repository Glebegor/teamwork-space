package controller

import (
	"fmt"
	"net/http"
	"team-work-space/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	fmt.Print(input)
	if _, err := ac.AuthUsecase.GetByEmail(c, input.Email); err == nil {
		// if foundUser.Username == input.Username {
		// 	c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "This is user already created with same name."})
		// 	return
		// }
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "This is user already created with same email."})
		return
	}

	newUser := domain.User{
		ID:       primitive.NewObjectID(),
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
	}

	if err := ac.AuthUsecase.Register(c, newUser); err != nil {
		c.JSON(http.StatusBadGateway, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Status: "ok"})
}
