package lists

import (
	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/pkg/services/auth"
	"github.com/opoccomaxao/myownranking/pkg/services/list"
	"go.uber.org/fx"
)

func Invoke() fx.Option {
	return fx.Module("api/lists",
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

	optionalGroup.GET("/api/lists", service.GetLists)

	requiredGroup.POST("/api/lists")
	optionalGroup.GET("/api/lists/:list_id")
	requiredGroup.PATCH("/api/lists/:list_id")
	requiredGroup.DELETE("/api/lists/:list_id")

	requiredGroup.GET("/api/lists/:list_id/items")
	requiredGroup.POST("/api/lists/:list_id/items")
	requiredGroup.PATCH("/api/lists/:list_id/items/:item_id")
	requiredGroup.DELETE("/api/lists/:list_id/items/:item_id")
}

type Service struct {
	list *list.Service
}

func NewService(
	list *list.Service,
) *Service {
	return &Service{
		list: list,
	}
}
