package usecase

import (
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

func (uu *userUsecase) Create(*domain.User) (*domain.User, error)     { return nil, nil }
func (uu *userUsecase) GetById(id string) (*domain.User, error)       { return nil, nil }
func (uu *userUsecase) GetByEmail(email string) (*domain.User, error) { return nil, nil }
func (uu *userUsecase) Update(*domain.UserUpdate) error               { return nil }
