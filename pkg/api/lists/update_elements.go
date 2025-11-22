package lists

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/pkg/api/values"
	"github.com/opoccomaxao/myownranking/pkg/models"
	"github.com/opoccomaxao/myownranking/pkg/services/list/structs"
	"github.com/samber/lo"
)

type UpdateListElementsRequest struct {
	Items []*models.ListElementAPI `json:"items"`
}

type UpdateListElementsResponse struct {
	Items []*models.ListElementAPI `json:"items"`
}

// UpdateListElements godoc
//
//	@Summary		Update all elements of a list.
//	@Description	Update all elements of a list.
//	@Description	Replace all elements entirely.
//	@Tags			lists
//	@Accept			json
//	@Produce		json
//	@Param			list_id	path		int								true	"List ID"
//	@Param			body	body		lists.UpdateListElementsRequest	true	"Update elements request"
//	@Success		200		{object}	lists.UpdateListElementsResponse
//	@Failure		400,401	{object}	models.ErrorResponse
//	@Failure		404		{object}	models.ErrorResponse
//	@Failure		500		"Internal server error"
//	@Router			/api/lists/{list_id}/items [POST]
//	@Security		StdAuth
func (s *Service) UpdateListElements(ctx *gin.Context) {
	var req UpdateListElementsRequest

	if err := ctx.BindJSON(&req); err != nil {
		return
	}

	updated, err := s.list.UpdateListElements(ctx, structs.UpdateListElementsParams{
		ID:       GetListID(ctx),
		UserID:   values.UserID.Get(ctx),
		Elements: lo.Map(req.Items, (*models.ListElementAPI).ToModel),
	})
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusOK, &UpdateListElementsResponse{
		Items: models.ListElementAPI{}.FromModels(updated),
	})
}
