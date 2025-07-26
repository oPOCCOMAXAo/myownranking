package migrations

import (
	"context"

	"github.com/opoccomaxao/myownranking/pkg/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type User struct {
	models.User
}

func Migrate(
	ctx context.Context,
	dbOrig *gorm.DB,
) error {
	db := dbOrig.WithContext(ctx)

	err := db.AutoMigrate(
		&User{},
	)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
