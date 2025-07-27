package lists

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/pkg/api/values"
	"github.com/opoccomaxao/myownranking/pkg/models"
	"github.com/opoccomaxao/myownranking/pkg/services/list/structs"
	"github.com/samber/lo"
)

type GetListsRequest struct {
	UserID int64 `binding:"omitempty"               form:"user_id"`
	Limit  int   `binding:"omitempty,min=1,max=100" form:"limit"`
	Offset int   `binding:"omitempty,min=0"         form:"offset"`
}

type GetListsResponse struct {
	Total int64             `json:"total"`
	Lists []*models.ListAPI `json:"lists"`
}

// GetLists godoc
//
//	@Summary		Get user lists
//	@Description	Retrieve lists created by a specific user
//	@Tags			lists
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query	int	false	"User ID. If not provided, uses the authenticated user ID."
//	@Param			limit	query	int	false	"Limit"	Default(10)	Minimum(1)	Maximum(100)
//	@Param			offset	query	int	false	"Offset"
//	@Success		200		{object}	GetListsResponse
//	@Failure		400		{object}	models.ErrorResponse
//	@Failure		404		{object}	models.ErrorResponse
//	@Failure		500		"Internal server error"
//	@Router			/api/lists [GET]
//	@Security		StdAuth
//
//nolint:mnd
func (s *Service) GetLists(ctx *gin.Context) {
	var req GetListsRequest

	err := ctx.Bind(&req)
	if err != nil {
		return
	}

	authedUserID := values.UserID.Get(ctx)

	params := structs.ListParams{
		UserID:    lo.CoalesceOrEmpty(req.UserID, authedUserID),
		Limit:     lo.CoalesceOrEmpty(req.Limit, 10),
		Offset:    req.Offset,
		WithTotal: true,
	}

	params.OnlyPublic = authedUserID != params.UserID

	res, err := s.list.GetLists(ctx.Request.Context(), params)
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusOK, &GetListsResponse{
		Total: res.Total,
		Lists: models.ListAPI{}.FromModels(res.Lists),
	})
}
