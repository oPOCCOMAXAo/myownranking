package list

import (
	"context"
	"time"

	"github.com/opoccomaxao/myownranking/pkg/models"
	"github.com/opoccomaxao/myownranking/pkg/services/list/repo"
	"github.com/opoccomaxao/myownranking/pkg/services/list/structs"
	"github.com/opoccomaxao/myownranking/pkg/utils/diff"
	"github.com/opoccomaxao/myownranking/pkg/utils/update"
	"github.com/samber/lo"
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

func (s *Service) CreateList(
	ctx context.Context,
	value *models.List,
) (*models.List, error) {
	value.DeletedAt = 0
	value.IsPublic = false

	return value, s.repo.CreateList(ctx, value)
}

func (s *Service) GetList(
	ctx context.Context,
	params structs.SingleListParams,
) (*models.List, error) {
	res, err := s.repo.GetListByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	err = s.CheckViewAccessToList(ctx, params.UserID, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) UpdateList(
	ctx context.Context,
	params structs.UpdateListParams,
) (*models.List, error) {
	list, err := s.repo.GetListByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	err = s.CheckUpdateAccessToList(ctx, params.UserID, list)
	if err != nil {
		return nil, err
	}

	state := update.NewState()
	update.Field(state, "name", &list.Name, params.Name)
	update.Field(state, "is_public", &list.IsPublic, params.IsPublic)
	update.Field(state, "deleted_at", &list.DeletedAt, params.DeletedAt)

	if state.HasChanges() {
		err = s.repo.UpdateList(ctx, list, state.Fields())
		if err != nil {
			return nil, err
		}
	}

	return list, nil
}

func (s *Service) DeleteList(
	ctx context.Context,
	params structs.SingleListParams,
) (*models.List, error) {
	return s.UpdateList(ctx, structs.UpdateListParams{
		ID:        params.ID,
		UserID:    params.UserID,
		DeletedAt: lo.ToPtr(time.Now().Unix()),
	})
}

func (s *Service) RestoreList(
	ctx context.Context,
	params structs.SingleListParams,
) (*models.List, error) {
	return s.UpdateList(ctx, structs.UpdateListParams{
		ID:        params.ID,
		UserID:    params.UserID,
		DeletedAt: lo.ToPtr(int64(0)),
	})
}

func (s *Service) UpdateListElements(
	ctx context.Context,
	params structs.UpdateListElementsParams,
) ([]*models.ListElement, error) {
	for _, element := range params.Elements {
		element.ListID = params.ID
	}

	list, err := s.repo.GetListByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	err = s.CheckUpdateAccessToList(ctx, params.UserID, list)
	if err != nil {
		return nil, err
	}

	oldElements, err := s.repo.GetElementsByListID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	diff := diff.OrderedSlice(
		oldElements,
		params.Elements,
		(*models.ListElement).PrepareForComparison,
		(*models.ListElement).Equals,
	)

	err = s.repo.CreateElements(ctx, diff.Created, 1000)
	if err != nil {
		return nil, err
	}

	err = s.repo.UpdateListElements(ctx, diff.Updated, 1000)
	if err != nil {
		return nil, err
	}

	ids := lo.Map(diff.Deleted, func(el *models.ListElement, _ int) int64 {
		return el.ID
	})

	err = s.repo.DeleteElementsByIDs(ctx, ids, 10000)
	if err != nil {
		return nil, err
	}

	updated, err := s.repo.GetElementsByListID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
