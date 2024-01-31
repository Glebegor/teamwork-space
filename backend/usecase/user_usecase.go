package usecase

import (
	"context"
	"team-work-space/domain"
	"time"
)

type userUsecase struct {
	repo           domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(repo domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		repo:           repo,
		contextTimeout: timeout,
	}
}

func (uu *userUsecase) GetById(c context.Context, id string) (*domain.User, error) {
	return nil, nil
}
func (uu *userUsecase) GetByEmail(c context.Context, email string) (*domain.User, error) {
	return nil, nil
}
func (uu *userUsecase) Update(c context.Context, id string, input *domain.UserUpdate) error {
	return nil
}
func (uu *userUsecase) Delete(c context.Context, id string) error {
	return nil
}
