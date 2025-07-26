package server

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/pkg/models"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

func (s *Server) mwRecover(ctx *gin.Context) {
	defer func() {
		rec := recover()
		if rec != nil {
			err, ok := rec.(error)
			if !ok {
				err = errors.Wrapf(models.ErrPanic, "%+v", rec)
			}

			ctx.Error(err).SetType(gin.ErrorTypePrivate)
		}
	}()

	ctx.Next()
}

func (s *Server) mwErrors(ctx *gin.Context) {
	ctx.Next()

	if len(ctx.Errors) == 0 {
		return
	}

	s.captureErrors(ctx)

	if errs := ctx.Errors.ByType(gin.ErrorTypePrivate); len(errs) > 0 {
		ctx.JSON(http.StatusInternalServerError, &models.ErrorResponse{
			Errors: []string{http.StatusText(http.StatusInternalServerError)},
		})

		return
	}

	if errs := ctx.Errors.ByType(gin.ErrorTypeBind); len(errs) > 0 {
		ctx.JSON(http.StatusBadRequest, &models.ErrorResponse{
			Errors: lo.Map(errs, func(e *gin.Error, _ int) string {
				return e.Error()
			}),
		})

		return
	}

	if errs := ctx.Errors.ByType(gin.ErrorTypePublic); len(errs) > 0 {
		ctx.JSON(http.StatusForbidden, &models.ErrorResponse{
			Errors: lo.Map(errs, func(e *gin.Error, _ int) string {
				return e.Error()
			}),
		})

		return
	}

	ctx.JSON(http.StatusInternalServerError, &models.ErrorResponse{
		Errors: []string{http.StatusText(http.StatusInternalServerError)},
	})
}

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
