package lists

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetListID(ctx *gin.Context) int64 {
	listID, _ := strconv.ParseInt(ctx.Param("list_id"), 10, 64)

	return listID
}
