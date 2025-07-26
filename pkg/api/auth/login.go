package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/pkg/models"
	"github.com/opoccomaxao/myownranking/pkg/services/auth"
)

type LoginRequest struct {
	Email    string `binding:"required,email" format:"email"  json:"email"`
	Password string `binding:"required"       json:"password"`
}

type LoginResponse struct {
	AuthToken    string `json:"auth_token"`
	RefreshToken string `json:"refresh_token"`
}

// Login godoc
//
//	@Summary		User login
//	@Description	User login with username and password
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		LoginRequest	true	"Body"
//	@Success		200		{object}	LoginResponse
//	@Failure		400,401	{object}	models.ErrorResponse
//	@Failure		500		"Internal Server Error"
//	@Router			/api/auth/login [POST]
func (s *Service) Login(ctx *gin.Context) {
	var req LoginRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		return
	}

	tokens, err := s.auth.Login(ctx, auth.AuthParams{
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

	if errors.Is(err, models.ErrInvalidAuth) {
		ctx.JSON(http.StatusUnauthorized, &models.ErrorResponse{
			Errors: []string{"Invalid email or password"},
		})

		return
	}

	ctx.Error(err)
}
