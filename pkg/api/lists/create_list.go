package lists

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/pkg/api/values"
	"github.com/opoccomaxao/myownranking/pkg/models"
)

type CreateListRequest struct {
	Name string `json:"name"`
}

type CreateListResponse struct {
	ID int64 `json:"id"`
}

// CreateList godoc
//
//	@Summary		Create a new list
//	@Description	Creates a new list for the authenticated user
//	@Tags			lists
//	@Accept			json
//	@Produce		json
//	@Param			body	body		lists.CreateListRequest	true	"List data"
//	@Success		201		{object}	lists.CreateListResponse
//	@Failure		400,401	{object}	models.ErrorResponse
//	@Failure		500		"Internal server error"
//	@Router			/api/lists [POST]
//	@Security		StdAuth
func (s *Service) CreateList(ctx *gin.Context) {
	var req CreateListRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		return
	}

	list, err := s.list.CreateList(ctx, &models.List{
		UserID: values.UserID.Get(ctx),
		Name:   req.Name,
	})
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusCreated, &CreateListResponse{
		ID: list.ID,
	})
}
