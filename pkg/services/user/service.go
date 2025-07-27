package user

import (
	"context"

	"github.com/opoccomaxao/myownranking/pkg/models"
	"github.com/opoccomaxao/myownranking/pkg/services/user/repo"
)

type Service struct {
	repo *repo.Repo
}

func NewService(
	repo *repo.Repo,
) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetUserByID(
	ctx context.Context,
	userID int64,
) (*models.User, error) {
	return s.repo.GetUserByID(ctx, userID)
}
