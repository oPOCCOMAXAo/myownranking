package app

import (
	"github.com/opoccomaxao/myownranking/pkg/config"
	"github.com/opoccomaxao/myownranking/pkg/db"
	"github.com/opoccomaxao/myownranking/pkg/logger"
	"go.uber.org/fx"
)

func Run() error {
	app := fx.New(
		fx.Provide(NewCancelCause),
		fx.WithLogger(NewFxLogger),
		config.Module(),
		logger.Module(),
		db.ModulePostgres(),
	)

	app.Run()

	return app.Err()
}
