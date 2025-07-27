package user

import (
	"github.com/opoccomaxao/myownranking/pkg/services/user/repo"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("services/user",
		fx.Provide(fx.Private,
			repo.NewRepo,
		),
		fx.Provide(NewService),
	)
}
