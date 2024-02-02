package controller

import (
	"net/http"
	"team-work-space/bootstrap"
	"team-work-space/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthController struct {
	AuthUsecase domain.AuthUsecase
	Env         *bootstrap.Env
}

func (ac *AuthController) Login(c *gin.Context) {
	var input domain.Login
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	user, err := ac.AuthUsecase.GetByEmail(c, input.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found with the given email"})
		return
	}

	tokenAccess, err := ac.AuthUsecase.CreateAccessToken(&user, ac.Env.SERVERsecret, 2)
	// tokenRefresh, err := ac.AuthUsecase.CreateAccessToken(&user, ac.Env.SECRET, 2)
	loginResponse := &domain.LoginResponse{
		AccessToken: tokenAccess,
		// RefreshToken: tokenRefresh,
	}

	c.JSON(http.StatusOK, loginResponse)
}

func (ac *AuthController) Reg(c *gin.Context) {
	var input domain.Reg

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
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
