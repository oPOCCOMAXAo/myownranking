package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/pkg/services/auth"
	"go.uber.org/fx"
)

func Invoke() fx.Option {
	return fx.Module("api/auth",
		fx.Provide(NewService, fx.Private),
		fx.Invoke(RegisterHandlers),
	)
}

func RegisterHandlers(
	router gin.IRouter,
	service *Service,
) {
	router.POST("/api/auth/login", service.Login)
	router.POST("/api/auth/register", service.Register)
	router.POST("/api/auth/refresh", service.Refresh)
}

type Service struct {
	auth *auth.Service
}

func NewService(
	auth *auth.Service,
) *Service {
	return &Service{
		auth: auth,
	}
}
