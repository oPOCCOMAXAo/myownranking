package repo

import (
	"context"

	"github.com/opoccomaxao/myownranking/pkg/models"
	"github.com/opoccomaxao/myownranking/pkg/services/list/structs"
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

func (r *Repo) GetLists(
	ctx context.Context,
	params structs.ListParams,
) (*structs.ListResult, error) {
	var res structs.ListResult

	query := r.db.WithContext(ctx).
		Model(&models.List{})

	if params.UserID != 0 {
		query = query.Where("user_id = ?", params.UserID)
	}

	if params.OnlyPublic {
		query = query.Where("is_private = ?", false)
	}

	if params.WithTotal {
		err := query.
			Select("COUNT(1)").
			Count(&res.Total).
			Error
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	if params.Limit > 0 {
		query = query.Limit(params.Limit)
	}

	if params.Offset > 0 {
		query = query.Offset(params.Offset)
	}

	err := query.
		Select("lists.*").
		Order("id DESC").
		Find(&res.Lists).
		Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &res, nil
}
