package auth

import (
	"github.com/opoccomaxao/myownranking/pkg/services/auth/repo"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("services/auth",
		fx.Provide(
			NewService,
		),
		fx.Provide(fx.Private,
			repo.NewRepo,
		),
	)
}
