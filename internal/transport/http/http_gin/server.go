package http_gin

import (
	"github.com/Shabashkin93/warning_tracker/internal/logging"
	"github.com/Shabashkin93/warning_tracker/internal/logging/slog"

	"github.com/gin-gonic/gin"
)

func Init(logger *logging.LoggerEntry) (handler *gin.Engine) {
	handler = gin.New()
	handler.Use(slog.GinLogger(logger.Handler), gin.Recovery())
	return
}
