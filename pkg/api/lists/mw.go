package lists

import (
	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/pkg/models"
)

func RequiredListID(ctx *gin.Context) {
	listID := GetListID(ctx)
	if listID == 0 {
		ctx.Error(models.ErrNotFound)
		ctx.Abort()

		return
	}
}
