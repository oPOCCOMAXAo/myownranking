package user

import (
	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/pkg/services/auth"
	"github.com/opoccomaxao/myownranking/pkg/services/user"
	"go.uber.org/fx"
)

func Invoke() fx.Option {
	return fx.Module("api/user",
		fx.Provide(NewService, fx.Private),
		fx.Invoke(RegisterHandlers),
	)
}

func RegisterHandlers(
	router gin.IRouter,
	auth *auth.Service,
	service *Service,
) {
	optionalGroup := router.Group("/",
		auth.MiddlewareAuthCache,
	)

	requiredGroup := router.Group("/",
		auth.MiddlewareAuthCache,
		auth.MiddlewareAuthRequired,
	)

	requiredGroup.GET("/api/users/me", service.GetMe)
	optionalGroup.GET("/api/users/id:user_id", service.GetUserByID)
}

type Service struct {
	user *user.Service
}

func NewService(
	user *user.Service,
) *Service {
	return &Service{
		user: user,
	}
}
