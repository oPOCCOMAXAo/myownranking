package api

import (
	"github.com/opoccomaxao/myownranking/pkg/api/auth"
	"github.com/opoccomaxao/myownranking/pkg/api/lists"
	"github.com/opoccomaxao/myownranking/pkg/api/swagger"
	"github.com/opoccomaxao/myownranking/pkg/api/system"
	"github.com/opoccomaxao/myownranking/pkg/api/users"
	"go.uber.org/fx"
)

func Invoke() fx.Option {
	return fx.Module("api",
		system.Invoke(),
		swagger.Invoke(),
		auth.Invoke(),
		users.Invoke(),
		lists.Invoke(),
	)
}
