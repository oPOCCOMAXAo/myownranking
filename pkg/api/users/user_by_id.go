package users

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/pkg/models"
)

// GetUserByID godoc
//
//	@Summary		Get user by ID
//	@Description	Retrieve a user by their unique ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		int	true	"User ID"
//	@Success		200		{object}	models.UserAPI
//	@Failure		404		{object}	models.ErrorResponse	"User not found"
//	@Failure		500		"Internal server error"
//	@Router			/api/users/id{user_id} [GET]
//	@Security		StdAuth
func (s *Service) GetUserByID(ctx *gin.Context) {
	userID, _ := strconv.ParseInt(ctx.Param("user_id"), 10, 64)

	user, err := s.user.GetUserByID(ctx.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, &models.ErrorResponse{
				Errors: []string{"User not found"},
			})

			return
		}

		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusOK, models.UserAPI{}.FromModel(user, 0))
}
