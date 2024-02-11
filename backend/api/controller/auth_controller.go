package controller

import (
	"net/http"
	"team-work-space/bootstrap"
	"team-work-space/domain"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthController struct {
	AuthUsecase domain.AuthUsecase
	Env         *bootstrap.Env
}

func (ac *AuthController) Login(c *gin.Context) {
	var input domain.Reg
	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	encodedPassword, err := ac.AuthUsecase.EncryptPassword(input.Password, ac.Env.SERVERsecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials"})
		return
	}
	input.Password = encodedPassword

	user, err := ac.AuthUsecase.GetByEmail(c, input.Email)
	if err != nil {
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

// @BasePath /api/v1

// Registration godoc
// @Summary Registration example
// @Schemes
// @Description Do Registration
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /auth/registration [post]
func (ac *AuthController) Reg(c *gin.Context) {
	var input domain.Reg
	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	if _, err := ac.AuthUsecase.GetByEmail(c, input.Email); err == nil {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "This is user already created with same email."})
		return
	}
	encodedPassword, err := ac.AuthUsecase.EncryptPassword(input.Password, ac.Env.SERVERsecret)
	if err != nil {
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
		c.JSON(http.StatusBadGateway, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Status: "ok"})
}

func (ac *AuthController) Refresh(c *gin.Context) {
	var input domain.Refresh

	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	id, err := ac.AuthUsecase.GetIdFromRefreshToken(input.RefreshToken, ac.Env.SERVERsecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User not found"})
		return
	}

	user, err := ac.AuthUsecase.GetUserById(c, id)
	if err != nil {
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
