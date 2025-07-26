package api

import (
	"github.com/opoccomaxao/myownranking/pkg/api/auth"
	"github.com/opoccomaxao/myownranking/pkg/api/swagger"
	"github.com/opoccomaxao/myownranking/pkg/api/system"
	"go.uber.org/fx"
)

func Invoke() fx.Option {
	return fx.Module("api",
		system.Invoke(),
		swagger.Invoke(),
		auth.Invoke(),
	)
}
