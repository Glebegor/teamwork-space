package usecase

import (
	"context"
	"fmt"
	"team-work-space/domain"
	"team-work-space/internal/tokenutil"
	"time"
)

type authUsecase struct {
	repo           domain.AuthRepository
	contextTimeout time.Duration
}

func NewAuthUsecase(repo domain.AuthRepository, timeout time.Duration) domain.AuthUsecase {
	return &authUsecase{
		repo:           repo,
		contextTimeout: timeout,
	}
}

func (au *authUsecase) Register(c context.Context, input domain.User) error {
	return au.repo.Create(c, input)
}
func (au *authUsecase) GetByEmail(c context.Context, email string) (domain.User, error) {
	return au.repo.GetByEmail(c, email)
}
func (au *authUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	fmt.Print(user)
	return tokenutil.CreateJwtToken(user, secret, expiry)
}
