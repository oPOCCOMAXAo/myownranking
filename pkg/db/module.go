package db

import (
	"context"
	"log/slog"

	"github.com/opoccomaxao/myownranking/pkg/migrations"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func ModulePostgres() fx.Option {
	return fx.Options(
		fx.Provide(newModulePostgres),
	)
}

type moduleParams struct {
	fx.In
	fx.Lifecycle

	Config Config
	Logger *slog.Logger `optional:"true"`
}

type moduleResults struct {
	fx.Out

	DB *gorm.DB
}

func newModulePostgres(
	params moduleParams,
) (moduleResults, error) {
	var res moduleResults

	var err error

	res.DB, err = NewPostgres(
		params.Config,
		params.Logger,
	)
	if err != nil {
		return res, err
	}

	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return migrations.Migrate(ctx, res.DB)
		},
	})

	return res, nil
}
