package list

import (
	"github.com/opoccomaxao/myownranking/pkg/services/list/repo"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("services/list",
		fx.Provide(fx.Private,
			repo.NewRepo,
		),
		fx.Provide(
			NewService,
		),
	)
}
