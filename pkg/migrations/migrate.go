package migrations

import (
	"context"
	_ "embed" // for embedded SQL files

	"github.com/opoccomaxao/myownranking/pkg/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type DefaultAutoModel struct {
	CreatedAt int64 `gorm:"column:created_at;not null;default:EXTRACT(EPOCH FROM now())::bigint"`
	UpdatedAt int64 `gorm:"column:updated_at;not null;default:EXTRACT(EPOCH FROM now())::bigint;autoUpdateTime"`
}

type User struct {
	models.User
	DefaultAutoModel
}

type List struct {
	models.List
	DefaultAutoModel

	User *models.User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ListElement struct {
	models.ListElement
	DefaultAutoModel

	List *List `gorm:"foreignKey:ListID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

//go:embed sql/updated_at.sql
var sqlUpdatedAt string

func Migrate(
	ctx context.Context,
	dbOrig *gorm.DB,
) error {
	db := dbOrig.WithContext(ctx)

	err := db.AutoMigrate(
		&User{},
		&List{},
		&ListElement{},
	)
	if err != nil {
		return errors.WithStack(err)
	}

	sqls := []string{
		sqlUpdatedAt,
	}
	for _, sql := range sqls {
		err = db.Exec(sql).Error
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
