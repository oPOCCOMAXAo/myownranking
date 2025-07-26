package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health godoc
//
//	@Summary		Health check endpoint
//	@Description	Check the health of the API server
//	@Tags			system
//	@Produce		text/plain
//	@Success		200	"OK"
//	@Router			/api/health [GET]
func (s *Service) Health(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
