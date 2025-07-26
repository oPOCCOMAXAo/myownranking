package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Shutdown godoc
//
//	@Summary		Shutdown endpoint
//	@Description	Shutdown the API server gracefully
//	@Tags			system
//	@Produce		text/plain
//	@Success		200	"OK"
//	@Router			/api/shutdown [PUT]
//

func (s *Service) Shutdown(ctx *gin.Context) {
	ctx.Status(http.StatusOK)

	s.cancel(nil)
}
