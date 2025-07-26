package db

import (
	"context"

	"github.com/opoccomaxao/myownranking/pkg/migrations"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func ModulePostgres() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				NewPostgres,
				fx.OnStart(StartHook),
			),
		),
	)
}

func StartHook(
	ctx context.Context,
	db *gorm.DB,
) error {
	return migrations.Migrate(ctx, db)
}
