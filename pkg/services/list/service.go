package list

import (
	"context"

	"github.com/opoccomaxao/myownranking/pkg/services/list/repo"
	"github.com/opoccomaxao/myownranking/pkg/services/list/structs"
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

func (s *Service) GetLists(
	ctx context.Context,
	params structs.ListParams,
) (*structs.ListResult, error) {
	res, err := s.repo.GetLists(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}
