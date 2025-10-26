package list

import (
	"context"

	"github.com/opoccomaxao/myownranking/pkg/models"
	"github.com/pkg/errors"
)

func (s *Service) CheckViewAccessToList(
	_ context.Context,
	userID int64,
	list *models.List,
) error {
	if list.UserID == userID {
		return nil
	}

	if list.DeletedAt != 0 {
		return errors.WithStack(models.ErrAccessDenied)
	}

	if list.IsPublic {
		return nil
	}

	return errors.WithStack(models.ErrAccessDenied)
}

func (s *Service) CheckUpdateAccessToList(
	_ context.Context,
	userID int64,
	list *models.List,
) error {
	if list.UserID == userID {
		return nil
	}

	return errors.WithStack(models.ErrAccessDenied)
}
