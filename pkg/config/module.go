package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/opoccomaxao/myownranking/pkg/db"
	"github.com/opoccomaxao/myownranking/pkg/logger"
	"github.com/opoccomaxao/myownranking/pkg/server"
	"github.com/pkg/errors"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("config",
		fx.Provide(New),
	)
}

type Config struct {
	fx.Out

	Logger logger.Config `envPrefix:"LOGGER_"`
	DB     db.Config     `envPrefix:"DB_"`
	Server server.Config `envPrefix:"SERVER_"`
}

func New() (Config, error) {
	var res Config

	err := env.ParseWithOptions(&res, env.Options{
		UseFieldNameByDefault: false,
		RequiredIfNoDef:       false,
	})
	if err != nil {
		return res, errors.WithStack(err)
	}

	return res, nil
}
