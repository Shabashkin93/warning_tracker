package http_gin

import (
	"context"

	"github.com/Shabashkin93/warning_tracker/internal/transport/http/http_gin/slog_gin"
	"github.com/Shabashkin93/warning_tracker/pkg/logging"

	"github.com/gin-gonic/gin"
)

func Init(ctx context.Context, logger logging.Logger) (handler *gin.Engine) {
	handler = gin.New()
	handler.Use(slog_gin.GinLogger(ctx, logger), gin.Recovery())
	return
}
