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
	UserID      int64 `binding:"omitempty"               form:"user_id"`
	Limit       int   `binding:"omitempty,min=1,max=100" form:"limit"`
	Offset      int   `binding:"omitempty,min=0"         form:"offset"`
	OnlyDeleted bool  `binding:"omitempty"               form:"only_deleted"`
	OnlyPublic  bool  `binding:"omitempty"               form:"only_public"`
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
//	@Param			user_id			query		int		false	"User ID. If not provided, uses the authenticated user ID."
//	@Param			limit			query		int		false	"Limit"	Default(10)	Minimum(1)	Maximum(100)
//	@Param			offset			query		int		false	"Offset"
//	@Param			only_deleted	query		bool	false	"Only Deleted. For authorized users lists only"
//	@Param			only_public		query		bool	false	"Only Public. For authorized users lists only"
//	@Success		200				{object}	lists.GetListsResponse
//	@Failure		400				{object}	models.ErrorResponse
//	@Failure		404				{object}	models.ErrorResponse
//	@Failure		500				"Internal server error"
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
	isForAuthed := authedUserID == req.UserID

	params := structs.ListParams{
		UserID:    lo.CoalesceOrEmpty(req.UserID, authedUserID),
		Limit:     lo.CoalesceOrEmpty(req.Limit, 10),
		Offset:    req.Offset,
		WithTotal: true,
	}

	if isForAuthed {
		params.OnlyDeleted = req.OnlyDeleted
		params.OnlyPublic = req.OnlyPublic
	} else {
		params.OnlyDeleted = false
		params.OnlyPublic = true
	}

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
