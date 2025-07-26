package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/opoccomaxao/myownranking/pkg/models"
	"github.com/pkg/errors"
)

type TokenService struct {
	issuer string
	secret []byte
	method jwt.SigningMethod
}

func NewTokenService(
	issuer string,
	secret []byte,
) *TokenService {
	return &TokenService{
		issuer: issuer,
		secret: secret,
		method: jwt.SigningMethodHS256,
	}
}

func (s *TokenService) keyFunc(token *jwt.Token) (any, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, errors.WithStack(models.ErrInvalidAuth)
	}

	return s.secret, nil
}

type TokenData struct {
	EntityID   string
	Audience   string
	Expiration time.Duration // -1 means no expiration
}

func (s *TokenService) SignTokenData(
	data TokenData,
) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:  data.EntityID,
		Audience: []string{data.Audience},
		Issuer:   s.issuer,
		IssuedAt: jwt.NewNumericDate(time.Now()),
	}

	if data.Expiration >= 0 {
		claims.ExpiresAt = jwt.NewNumericDate(claims.IssuedAt.Add(data.Expiration))
	}

	jwtToken := jwt.NewWithClaims(s.method, claims)

	tokenString, err := jwtToken.SignedString(s.secret)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return tokenString, nil
}

type TokenValidationOption func(*jwt.RegisteredClaims) error

func (s *TokenService) ParseTokenWithValidation(
	tokenString string,
	options ...TokenValidationOption,
) (*jwt.RegisteredClaims, error) {
	var claims jwt.RegisteredClaims

	token, err := jwt.
		NewParser().
		ParseWithClaims(tokenString, &claims, s.keyFunc)
	if err != nil {
		var jwtErr *jwt.ValidationError
		if errors.As(err, &jwtErr) {
			return nil, errors.WithStack(models.ErrInvalidAuth)
		}

		return nil, errors.WithStack(err)
	}

	if !token.Valid {
		return nil, errors.WithStack(models.ErrInvalidAuth)
	}

	for _, option := range options {
		err := option(&claims)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return &claims, nil
}

func (s *TokenService) ValidateIssuerSelf() TokenValidationOption {
	return func(claims *jwt.RegisteredClaims) error {
		if !claims.VerifyIssuer(s.issuer, true) {
			return errors.WithStack(models.ErrInvalidAuth)
		}

		return nil
	}
}

func (s *TokenService) ValidateAudience(audience string) TokenValidationOption {
	return func(claims *jwt.RegisteredClaims) error {
		if !claims.VerifyAudience(audience, true) {
			return errors.WithStack(models.ErrInvalidAuth)
		}

		return nil
	}
}
