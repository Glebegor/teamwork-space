package usecase

import (
	"context"
	"team-work-space/domain"
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

func (au *authUsecase) Register(c context.Context, input domain.Reg) error {
	return au.repo.Create(c, input)
}
