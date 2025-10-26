package lists

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/pkg/api/values"
	"github.com/opoccomaxao/myownranking/pkg/services/list/structs"
)

// DeleteList godoc
//
//	@Summary		Delete a list.
//	@Description	Delete a list by ID.
//	@Tags			lists
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			list_id	path	int	true	"List ID"
//	@Success		204		"Deleted"
//	@Failure		400,404	{object}	models.ErrorResponse
//	@Failure		500		"Internal Server Error"
//	@Router			/api/lists/{list_id} [DELETE]
//	@Security		StdAuth
func (s *Service) DeleteList(ctx *gin.Context) {
	_, err := s.list.DeleteList(ctx, structs.SingleListParams{
		ID:     GetListID(ctx),
		UserID: values.UserID.Get(ctx),
	})
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.Status(http.StatusNoContent)
}
