package lists

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/pkg/api/values"
	"github.com/opoccomaxao/myownranking/pkg/models"
	"github.com/opoccomaxao/myownranking/pkg/services/list/structs"
)

type UpdateListRequest struct {
	Name     *string `json:"name"`
	IsPublic *bool   `json:"is_public"`
}

// UpdateList godoc
//
//	@Summary		Update a list
//	@Description	Update a list by ID.
//	@Description	All fields are optional, only provided fields will be updated.
//	@Tags			lists
//	@Accept			json
//	@Produce		json
//	@Param			list_id	path		string					true	"List ID"
//	@Param			request	body		lists.UpdateListRequest	true	"Update list request"
//	@Success		200		{object}	lists.GetListResponse
//	@Failure		400,404	{object}	models.ErrorResponse
//	@Router			/api/lists/{list_id} [PATCH]
//	@Security		StdAuth
func (s *Service) UpdateList(ctx *gin.Context) {
	var req UpdateListRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		return
	}

	list, err := s.list.UpdateList(ctx, structs.UpdateListParams{
		ID:       GetListID(ctx),
		UserID:   values.UserID.Get(ctx),
		Name:     req.Name,
		IsPublic: req.IsPublic,
	})
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusOK, &GetListResponse{
		List: models.ListAPI{}.FromModel(list, 0),
	})
}
