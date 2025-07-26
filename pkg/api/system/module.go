package system

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func Invoke() fx.Option {
	return fx.Module("api/system",
		fx.Provide(NewService),
		fx.Invoke(RegisterEndpoints),
	)
}

func RegisterEndpoints(
	router gin.IRouter,
	service *Service,
) {
	router.GET("/api/health", service.Health)
	router.PUT("/api/shutdown", service.Shutdown)
}

type Service struct {
	cancel context.CancelCauseFunc
}

func NewService(
	cancel context.CancelCauseFunc,
) *Service {
	return &Service{
		cancel: cancel,
	}
}
