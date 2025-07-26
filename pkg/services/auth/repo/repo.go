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

func (r *Repo) CreateUser(
	ctx context.Context,
	user *models.User,
) error {
	err := r.db.WithContext(ctx).
		Create(user).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.WithStack(err)
		}

		return errors.WithStack(err)
	}

	return nil
}

//nolint:nilnil
func (r *Repo) GetUserByIDOrNil(
	ctx context.Context,
	id int64,
) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&user).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.WithStack(err)
	}

	return &user, nil
}

//nolint:nilnil
func (r *Repo) GetUserByEmailOrNil(
	ctx context.Context,
	email string,
) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.WithStack(err)
	}

	return &user, nil
}
