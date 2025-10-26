package lists

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/pkg/api/values"
	"github.com/opoccomaxao/myownranking/pkg/models"
	"github.com/opoccomaxao/myownranking/pkg/services/list/structs"
)

type GetListResponse struct {
	List *models.ListAPI `json:"list"`
}

// GetList godoc
//
//	@Summary		Get a list.
//	@Description	Get a list by ID.
//	@Tags			lists
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			list_id	path		int	true	"List ID"
//	@Success		200		{object}	lists.GetListResponse
//	@Failure		400,404	{object}	models.ErrorResponse
//	@Failure		500		"Internal Server Error"
//	@Router			/api/lists/{list_id} [GET]
//	@Security		StdAuth
func (s *Service) GetList(ctx *gin.Context) {
	list, err := s.list.GetList(ctx, structs.SingleListParams{
		ID:     GetListID(ctx),
		UserID: values.UserID.Get(ctx),
	})
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusOK, &GetListResponse{
		List: models.ListAPI{}.FromModel(list, 0),
	})
}
