package app

import (
	"github.com/opoccomaxao/myownranking/pkg/api"
	"github.com/opoccomaxao/myownranking/pkg/config"
	"github.com/opoccomaxao/myownranking/pkg/db"
	"github.com/opoccomaxao/myownranking/pkg/logger"
	"github.com/opoccomaxao/myownranking/pkg/server"
	"go.uber.org/fx"
)

func Run() error {
	app := fx.New(
		// Fx modules

		fx.Provide(NewCancelCause),
		fx.WithLogger(NewFxLogger),

		// Modules

		config.Module(),
		logger.Module(),
		db.ModulePostgres(),
		server.Module(),

		// Invoke

		api.Invoke(),
	)

	app.Run()

	//nolint:wrapcheck
	return app.Err()
}
