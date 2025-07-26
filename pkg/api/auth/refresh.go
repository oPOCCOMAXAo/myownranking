package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/pkg/models"
)

type RefreshRequest struct {
	RefreshToken string `binding:"required" json:"refresh_token"`
}

type RefreshResponse struct {
	AuthToken    string `json:"auth_token"`
	RefreshToken string `json:"refresh_token"`
}

// Refresh godoc
//
//	@Summary		User token refresh
//	@Description	User refreshes access and refresh tokens
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		RefreshRequest	true	"Body"
//	@Success		200		{object}	RefreshResponse
//	@Failure		400,401	{object}	models.ErrorResponse
//	@Failure		500		"Internal Server Error"
//	@Router			/api/auth/refresh [POST]
func (s *Service) Refresh(ctx *gin.Context) {
	var req RefreshRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		return
	}

	tokens, err := s.auth.RefreshTokens(ctx, req.RefreshToken)
	if err != nil {
		if errors.Is(err, models.ErrInvalidAuth) {
			ctx.Error(err)

			ctx.JSON(http.StatusUnauthorized, &models.ErrorResponse{
				Errors: []string{"Invalid refresh token"},
			})

			return
		}

		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusOK, &RefreshResponse{
		AuthToken:    tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}
