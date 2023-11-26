package health

import (
	"net/http"

	"github.com/gin-gonic/gin"

	FlareState "github.com/luckylu/flare/internal/state"
)

func RegisterRouting(router *gin.Engine) {
	router.GET(FlareState.MiscPages.HealthCheck.Path, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}
