package slog_gin

import (
	"context"
	"log/slog"
	"time"

	"github.com/Shabashkin93/warning_tracker/pkg/logging"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

type fields struct {
	ReqId           string
	StatusCode      int
	Bytes           int
	Duration        time.Duration
	DurationDisplay string
	Path            string
	Method          string
	MsgStr          string
	ClientIP        string
	Proto           string
	UserAgent       string
}

func GinLogger(ctx context.Context, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		msg := c.Errors.String()
		if msg == "" {
			msg = "Request"
		}
		cData := &fields{
			ReqId:           requestid.Get(c),
			StatusCode:      c.Writer.Status(),
			Bytes:           c.Writer.Size(),
			Duration:        time.Since(t),
			DurationDisplay: time.Since(t).String(),
			Path:            c.Request.URL.Path,
			Method:          c.Request.Method,
			MsgStr:          msg,
			ClientIP:        c.ClientIP(),
			Proto:           c.Request.Proto,
			UserAgent:       c.Request.UserAgent(),
		}

		logSwitch(ctx, logger, cData)
	}
}

func logSwitch(ctx context.Context, logger logging.Logger, data *fields) {
	switch {
	case data.StatusCode >= 400 && data.StatusCode < 500:
		{
			logger.Warn(ctx, data.MsgStr, slog.String("method", data.Method), slog.String("path", data.Path), slog.String("resp_time", data.Duration.String()), slog.Int("status", data.StatusCode), slog.String("client_ip", data.ClientIP))
		}
	case data.StatusCode >= 500:
		{
			logger.Error(ctx, data.MsgStr, slog.String("method", data.Method), slog.String("path", data.Path), slog.String("resp_time", data.Duration.String()), slog.Int("status", data.StatusCode), slog.String("client_ip", data.ClientIP))
		}
	default:
		logger.Info(ctx, data.MsgStr, slog.String("method", data.Method), slog.String("path", data.Path), slog.String("resp_time", data.Duration.String()), slog.Int("status", data.StatusCode), slog.String("client_ip", data.ClientIP))
	}
}
