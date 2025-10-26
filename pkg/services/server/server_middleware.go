package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/myownranking/pkg/models"
	"github.com/pkg/errors"
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

	if s.handleErrors(ctx) {
		return
	}

	ctx.JSON(http.StatusInternalServerError, &models.ErrorResponse{
		Errors: []string{http.StatusText(http.StatusInternalServerError)},
	})
}
