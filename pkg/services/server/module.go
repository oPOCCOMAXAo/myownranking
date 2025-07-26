package server

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("services/server",
		fx.Provide(
			fx.Annotate(
				New,
				fx.OnStart((*Server).OnStart),
				fx.OnStop((*Server).OnStop),
			),
			(*Server).GetEngine,
			(*Server).GetRouter,
		),
	)
}
