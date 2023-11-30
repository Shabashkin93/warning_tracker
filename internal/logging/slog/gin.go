package slog

import (
	"log/slog"
	"time"

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

func GinLogger(i interface{}) gin.HandlerFunc {
	logger := i.(*slog.Logger)
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

		logSwitch(logger, cData)
	}
}

func logSwitch(logger *slog.Logger, data *fields) {
	switch {
	case data.StatusCode >= 400 && data.StatusCode < 500:
		{
			logger.Warn(data.MsgStr, slog.String("method", data.Method), slog.String("path", data.Path), slog.String("resp_time", data.Duration.String()), slog.Int("status", data.StatusCode), slog.String("client_ip", data.ClientIP))
		}
	case data.StatusCode >= 500:
		{
			logger.Error(data.MsgStr, slog.String("method", data.Method), slog.String("path", data.Path), slog.String("resp_time", data.Duration.String()), slog.Int("status", data.StatusCode), slog.String("client_ip", data.ClientIP))
		}
	default:
		logger.Info(data.MsgStr, slog.String("method", data.Method), slog.String("path", data.Path), slog.String("resp_time", data.Duration.String()), slog.Int("status", data.StatusCode), slog.String("client_ip", data.ClientIP))
	}
}
