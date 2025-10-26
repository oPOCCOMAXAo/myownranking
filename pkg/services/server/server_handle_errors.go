package server

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/pkg/models"
	"github.com/pkg/errors"
)

func (s *Server) captureErrors(ctx *gin.Context) {
	attrs := []any{
		slog.String("method", ctx.Request.Method),
		slog.String("path", ctx.Request.URL.Path),
	}

	for _, e := range ctx.Errors {
		attrs = append(attrs, slog.Any("error", e.Err))
	}

	s.logger.ErrorContext(ctx, "request", attrs...)
}

type handledError struct {
	Error    error // error to match.
	Status   int16
	Text     string // text to override error message. if empty, don't override.
	Priority byte   // higher wins.
}

func (e *handledError) StatusOrDefault(status int) int {
	if e == nil {
		return status
	}

	return int(e.Status)
}

func (e *handledError) GetHigherError(current *handledError) *handledError {
	if current == nil || e.Priority > current.Priority {
		return e
	}

	return current
}

func (e *handledError) GetErrorText(err error) string {
	if e.Text != "" {
		return e.Text
	}

	return err.Error()
}

//nolint:gochecknoglobals,mnd
var handledErrors = []handledError{
	{
		Error:    models.ErrNotFound,
		Status:   http.StatusNotFound,
		Text:     http.StatusText(http.StatusNotFound),
		Priority: 1,
	},
	{
		Error:    models.ErrAccessDenied,
		Status:   http.StatusNotFound,
		Text:     http.StatusText(http.StatusNotFound),
		Priority: 2,
	},
	{
		Error:    models.ErrInvalidAuth,
		Status:   http.StatusUnauthorized,
		Text:     http.StatusText(http.StatusUnauthorized),
		Priority: 255,
	},
}

//nolint:gochecknoglobals
var defaultHandledError = handledError{
	Error:    nil, // unused.
	Status:   http.StatusBadRequest,
	Priority: 0,
}

func (s *Server) handleErrors(ctx *gin.Context) bool {
	var (
		finalError *handledError
		res        models.ErrorResponse
	)

	for _, err := range ctx.Errors {
		var currentError *handledError

		if err.Type&gin.ErrorTypePublic != 0 {
			currentError = &defaultHandledError
		}

		for _, defErr := range handledErrors {
			if errors.Is(err.Err, defErr.Error) {
				currentError = &defErr

				break
			}
		}

		if currentError == nil {
			continue
		}

		res.Errors = append(res.Errors, currentError.GetErrorText(err.Err))

		finalError = currentError.GetHigherError(finalError)
	}

	if len(res.Errors) == 0 {
		return false
	}

	ctx.JSON(finalError.StatusOrDefault(http.StatusBadRequest), &res)

	return true
}
