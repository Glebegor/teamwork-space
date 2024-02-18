package usecase

import (
	"context"
	"team-work-space/domain"
	"time"
)

type teamUsecase struct {
	repo           domain.TeamRepository
	contextTimeout time.Duration
}

func NewTeamUsecase(repo domain.TeamRepository, timeout time.Duration) domain.TeamUsecase {
	return &teamUsecase{
		repo:           repo,
		contextTimeout: timeout,
	}
}
func (tu *teamUsecase) Create(c context.Context, input domain.Team) error {
	return nil
}
func (tu *teamUsecase) GetAll() ([]domain.Team, error) {
	return tu.repo.GetAll()
}
func (tu *teamUsecase) GetById(c context.Context, id string) (domain.Team, error) {
	return tu.repo.GetById(c, id)
}
func (tu *teamUsecase) Update(id string, input domain.TeamUpdate) error {
	return nil
}
func (tu *teamUsecase) Delete(id string) error {
	return nil
}
