package repo

import (
	"context"

	"github.com/opoccomaxao/myownranking/pkg/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) GetUserByID(
	ctx context.Context,
	userID int64,
) (*models.User, error) {
	var res models.User

	err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", userID).
		Take(&res).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(models.ErrNotFound)
		}

		return nil, errors.WithStack(err)
	}

	return &res, nil
}
