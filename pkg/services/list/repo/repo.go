package repo

import (
	"context"
	stderr "errors"
	"slices"

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
		query = query.Where("is_public = true")
	}

	if params.OnlyDeleted {
		query = query.Where("deleted_at > 0")
	} else {
		query = query.Where("deleted_at = 0")
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

func (r *Repo) CreateList(
	ctx context.Context,
	value *models.List,
) error {
	err := r.db.
		WithContext(ctx).
		Create(value).
		Error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *Repo) GetListByID(
	ctx context.Context,
	id int64,
) (*models.List, error) {
	var res models.List

	err := r.db.WithContext(ctx).
		Where("id = ?", id).
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

func (r *Repo) UpdateList(
	ctx context.Context,
	value *models.List,
	fields []string,
) error {
	if len(fields) == 0 {
		return nil
	}

	err := r.db.
		WithContext(ctx).
		Select(fields).
		Updates(value).
		Error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *Repo) GetElementsByListID(
	ctx context.Context,
	listID int64,
) ([]*models.ListElement, error) {
	var res []*models.ListElement

	err := r.db.WithContext(ctx).
		Where("list_id = ?", listID).
		Order("id ASC").
		Find(&res).
		Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}

func (r *Repo) CreateElements(
	ctx context.Context,
	elems []*models.ListElement,
	batchSize int,
) error {
	if len(elems) == 0 {
		return nil
	}

	err := r.db.WithContext(ctx).
		CreateInBatches(&elems, batchSize).
		Error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *Repo) DeleteElementsByIDs(
	ctx context.Context,
	ids []int64,
	batchSize int,
) error {
	if len(ids) == 0 {
		return nil
	}

	var errs []error

	for chunk := range slices.Chunk(ids, batchSize) {
		err := r.db.WithContext(ctx).
			Where("id IN ?", chunk).
			Delete(&models.ListElement{}).
			Error
		if err != nil {
			errs = append(errs, errors.WithMessagef(err, "failed to delete list element with ID %d", chunk))
		}
	}

	return errors.WithStack(stderr.Join(errs...))
}

func (r *Repo) UpdateListElements(
	ctx context.Context,
	elems []*models.ListElement,
	batchSize int,
) error {
	if len(elems) == 0 {
		return nil
	}

	var errs []error

	for chunk := range slices.Chunk(elems, batchSize) {
		err := r.db.WithContext(ctx).
			Transaction(func(tx *gorm.DB) error {
				for _, elem := range chunk {
					err := tx.
						Model(elem).
						Where("id = ?", elem.ID).
						Select("*").
						Updates(elem).
						Error
					if err != nil {
						errs = append(errs, errors.WithMessagef(err, "failed to update list element with ID %d", elem.ID))
					}
				}

				return nil
			})
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errors.WithStack(stderr.Join(errs...))
}
