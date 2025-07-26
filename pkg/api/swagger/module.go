package swagger

import (
	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/pkg/ginutils"
	"go.uber.org/fx"
)

func Invoke() fx.Option {
	return fx.Module("api/swagger",
		fx.Invoke(RegisterHandlers),
	)
}

func RegisterHandlers(
	router gin.IRouter,
) {
	router.GET("/api/swagger/*any", fullInstance())
	router.GET("/api/swagger", ginutils.StaticRedirect("/api/swagger/index.html"))
}
