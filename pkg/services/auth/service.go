package auth

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/pkg/api/values"
	"github.com/opoccomaxao/myownranking/pkg/models"
	"github.com/opoccomaxao/myownranking/pkg/services/auth/repo"
	"github.com/opoccomaxao/myownranking/pkg/utils/texts"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

const (
	AudienceAuth    = "auth"
	AudienceRefresh = "refresh"
)

type Config struct {
	Issuer    string    `env:"ISSUER,required"`
	JWTSecret texts.Hex `env:"JWT_SECRET,required"`
}

type Service struct {
	cfg   Config
	repo  *repo.Repo
	token *TokenService
}

func NewService(
	cfg Config,
	repo *repo.Repo,
) *Service {
	return &Service{
		cfg:  cfg,
		repo: repo,

		token: NewTokenService(cfg.Issuer, cfg.JWTSecret),
	}
}

//nolint:revive
type AuthParams struct {
	Email    string `json:"-"`
	Password string `json:"-"`
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func (s *Service) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return string(hash), nil
}

func (s *Service) checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

//nolint:mnd
func (s *Service) createTokens(
	user *models.User,
) (*Tokens, error) {
	var (
		res Tokens
		err error
	)

	res.AccessToken, err = s.token.SignTokenData(TokenData{
		EntityID:   strconv.FormatInt(user.ID, 10),
		Audience:   AudienceAuth,
		Expiration: 24 * time.Hour,
	})
	if err != nil {
		return nil, err
	}

	res.RefreshToken, err = s.token.SignTokenData(TokenData{
		EntityID:   strconv.FormatInt(user.ID, 10),
		Audience:   AudienceRefresh,
		Expiration: 30 * 24 * time.Hour,
	})
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *Service) parseTokenData(
	tokenString string,
	audience string,
) (*TokenData, error) {
	claims, err := s.token.ParseTokenWithValidation(
		tokenString,
		s.token.ValidateIssuerSelf(),
		s.token.ValidateAudience(audience),
	)
	if err != nil {
		return nil, err
	}

	return &TokenData{
		EntityID: claims.Subject,
	}, nil
}

func (s *Service) Login(
	ctx context.Context,
	params AuthParams,
) (*Tokens, error) {
	user, err := s.repo.GetUserByEmailOrNil(ctx, params.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.WithStack(models.ErrInvalidAuth)
	}

	if !s.checkPassword(params.Password, user.Password) {
		return nil, errors.WithStack(models.ErrInvalidAuth)
	}

	return s.createTokens(user)
}

func (s *Service) Register(
	ctx context.Context,
	params AuthParams,
) (*Tokens, error) {
	user, err := s.repo.GetUserByEmailOrNil(ctx, params.Email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, errors.WithStack(models.ErrDuplicate)
	}

	newUser := &models.User{
		Email: params.Email,
	}

	newUser.Password, err = s.hashPassword(params.Password)
	if err != nil {
		return nil, err
	}

	err = s.repo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return s.createTokens(newUser)
}

func (s *Service) RefreshTokens(
	ctx context.Context,
	refreshToken string,
) (*Tokens, error) {
	data, err := s.parseTokenData(
		refreshToken,
		AudienceRefresh,
	)
	if err != nil {
		return nil, err
	}

	userID, err := strconv.ParseInt(data.EntityID, 10, 64)
	if err != nil {
		return nil, errors.WithStack(models.ErrInvalidAuth)
	}

	user, err := s.repo.GetUserByIDOrNil(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.WithStack(models.ErrInvalidAuth)
	}

	return s.createTokens(user)
}

func (s *Service) GetUserIDByAuthToken(
	_ context.Context,
	authToken string,
) (int64, error) {
	data, err := s.parseTokenData(
		authToken,
		AudienceAuth,
	)
	if err != nil {
		return 0, err
	}

	userID, err := strconv.ParseInt(data.EntityID, 10, 64)
	if err != nil {
		return 0, errors.WithStack(models.ErrInvalidAuth)
	}

	return userID, nil
}

func (s *Service) MiddlewareAuthCache(ctx *gin.Context) {
	authToken := ctx.GetHeader("Authorization")
	if authToken == "" {
		return
	}

	userID, err := s.GetUserIDByAuthToken(ctx, authToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, &models.ErrorResponse{
			Errors: []string{"Invalid authorization token"},
		})

		return
	}

	values.UserID.Set(ctx, userID)
}

// MiddlewareAuthRequired godoc
//
// Must be called after MiddlewareAuthCache to ensure UserID is available.
func (s *Service) MiddlewareAuthRequired(ctx *gin.Context) {
	if values.UserID.IsEmpty(ctx) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, &models.ErrorResponse{
			Errors: []string{"Authorization required"},
		})

		return
	}
}
