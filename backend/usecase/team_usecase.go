package usecase

import (
	"team-work-space/domain"
	"time"
)

type teamUsecase struct {
	repo           domain.TeamRepository
	contextTimeout time.Duration
}
