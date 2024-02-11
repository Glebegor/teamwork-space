package controller

import (
	"net/http"
	"team-work-space/bootstrap"
	"team-work-space/domain"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthController struct {
	AuthUsecase domain.AuthUsecase
	Env         *bootstrap.Env
}

// Login
// @Summary Login
// @Description Do authorization with using email and password
// @Tags auth v1
// @ID authorization-user
// @Accept json
// @Produce json
// @Param input body domain.Login true "Login"
// @Success 200 {object} domain.LoginResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure 502 {object} domain.ErrorResponse
// @Failure 400, 404 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /api/v1/auth/login [post]
func (ac *AuthController) Login(c *gin.Context) {
	var input domain.Reg
	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		logrus.Errorf("Error while binding body: %v", err.Error())
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	encodedPassword, err := ac.AuthUsecase.EncryptPassword(input.Password, ac.Env.SERVERsecret)
	if err != nil {
		logrus.Errorf("Error while encrypting password: %v", err.Error())
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials"})
		return
	}
	input.Password = encodedPassword

	user, err := ac.AuthUsecase.GetByEmail(c, input.Email)
	if err != nil {
		logrus.Errorf("Error while getting user by email: %v", err.Error())
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found with the given email"})
		return
	}
	if user.Password != encodedPassword {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Password incorrect."})
		return
	}

	tokenAccess, err := ac.AuthUsecase.CreateAccessToken(&user, ac.Env.SERVERsecret, 2)
	tokenRefresh, err := ac.AuthUsecase.CreateRefreshToken(&user, ac.Env.SERVERsecret, 24)
	loginResponse := &domain.LoginResponse{
		AccessToken:  tokenAccess,
		RefreshToken: tokenRefresh,
	}

	c.JSON(http.StatusOK, loginResponse)
}

// Reg
// @Summary Registration
// @Description Do Registration with using username, email and password
// @Tags auth v1
// @ID create-user
// @Accept json
// @Produce json
// @Param input body domain.Reg true "Registration"
// @Success 200 {object} domain.SuccessResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure 502 {object} domain.ErrorResponse
// @Failure 400, 404, 409 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /api/v1/auth/registration [post]
func (ac *AuthController) Reg(c *gin.Context) {
	var input domain.Reg
	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		logrus.Errorf("Error while binding body: %v", err.Error())
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	if _, err := ac.AuthUsecase.GetByEmail(c, input.Email); err == nil {
		logrus.Errorf("Error while getting user by email: %v", err.Error())
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "This is user already created with same email."})
		return
	}
	encodedPassword, err := ac.AuthUsecase.EncryptPassword(input.Password, ac.Env.SERVERsecret)
	if err != nil {
		logrus.Errorf("Error while encrypting password: %v", err.Error())
		c.JSON(http.StatusBadGateway, domain.ErrorResponse{Message: err.Error()})
		return
	}
	input.Password = encodedPassword

	newUser := domain.User{
		ID:       primitive.NewObjectID(),
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
		Role:     "User",
	}

	if err := ac.AuthUsecase.Register(c, newUser); err != nil {
		logrus.Errorf("Error while registering new user: %v", err.Error())
		c.JSON(http.StatusBadGateway, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Status: "ok"})
}

// Refresh
// @Summary Refresh
// @Description Refresh token to get access and another refresh token
// @Tags auth v1
// @ID refresh-token
// @Accept json
// @Produce json
// @Param input body domain.Refresh true "Refresh"
// @Success 200 {object} domain.RefreshTokenResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure 502 {object} domain.ErrorResponse
// @Failure 400, 404 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /api/v1/auth/refresh [post]
func (ac *AuthController) Refresh(c *gin.Context) {
	var input domain.Refresh

	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		logrus.Errorf("Error while binding body: %v", err.Error())
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	id, err := ac.AuthUsecase.GetIdFromRefreshToken(input.RefreshToken, ac.Env.SERVERsecret)
	if err != nil {
		logrus.Errorf("Error while getting user from refresh token: %v", err.Error())
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User not found"})
		return
	}

	user, err := ac.AuthUsecase.GetUserById(c, id)
	if err != nil {
		logrus.Errorf("Error while getting user by id: %v", err.Error())
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User not found"})
		return
	}
	accessToken, err := ac.AuthUsecase.CreateAccessToken(&user, ac.Env.SERVERsecret, 2)
	refreshToken, err := ac.AuthUsecase.CreateRefreshToken(&user, ac.Env.SERVERsecret, 168)

	refreshResponse := &domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, refreshResponse)
}
