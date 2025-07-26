package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/pkg/models"
	"github.com/opoccomaxao/myownranking/pkg/services/auth"
)

type RegisterRequest struct {
	Email    string `binding:"required,email" format:"email"  json:"email"`
	Password string `binding:"required"       json:"password"`
}

type RegisterResponse struct {
	AuthToken    string `json:"auth_token"`
	RefreshToken string `json:"refresh_token"`
}

// Register godoc
//
//	@Summary		User registration
//	@Description	User registration with username and password
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		RegisterRequest	true	"Body"
//	@Success		200		{object}	RegisterResponse
//	@Failure		400,409	{object}	models.ErrorResponse
//	@Failure		500		"Internal Server Error"
//	@Router			/api/auth/register [POST]
func (s *Service) Register(ctx *gin.Context) {
	var req RegisterRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		return
	}

	tokens, err := s.auth.Register(ctx, auth.AuthParams{
		Email:    req.Email,
		Password: req.Password,
	})
	if err == nil {
		ctx.JSON(http.StatusOK, &LoginResponse{
			AuthToken:    tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		})

		return
	}

	if errors.Is(err, models.ErrDuplicate) {
		ctx.JSON(http.StatusConflict, &models.ErrorResponse{
			Errors: []string{"Email already exists"},
		})

		return
	}

	ctx.Error(err)
}
