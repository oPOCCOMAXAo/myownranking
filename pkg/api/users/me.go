package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/pkg/api/values"
	"github.com/opoccomaxao/myownranking/pkg/models"
)

// GetMe godoc
//
//	@Summary		Get current user
//	@Description	Retrieve the currently authenticated user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.UserAPI
//	@Failure		404	{object}	models.ErrorResponse	"User not found"
//	@Failure		500	"Internal server error"
//	@Router			/api/users/me [GET]
//	@Security		StdAuth
func (s *Service) GetMe(ctx *gin.Context) {
	userID := values.UserID.Get(ctx)

	user, err := s.user.GetUserByID(ctx.Request.Context(), userID)
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusOK, models.UserAPI{}.FromModel(user, 0))
}
