package status

import (
	"context"
	"net/http"

	"github.com/Shabashkin93/warning_tracker/internal/domain/status"
	"github.com/Shabashkin93/warning_tracker/internal/transport/http/http_gin/http_err"

	"github.com/gin-gonic/gin"
)

type StatusService interface {
	GetAll() (dto *status.Status, code int, err error)
}

type transport struct {
	usecase StatusService
	handler *gin.Engine
	debug   bool
	url     string
}

func NewTransport(ctx context.Context, usecase StatusService, i interface{}, url string, debug bool) *transport {
	handler := i.(*gin.Engine)
	return &transport{usecase: usecase, handler: handler, debug: debug, url: url}
}

func (t *transport) Register(version string, i interface{}) {
	router := i.(*gin.Engine)

	api := router.Group(version + t.url)
	{
		api.GET("", t.GetAll)
	}
}

func (t *transport) GetAll(c *gin.Context) {

	statuss, code, err := t.usecase.GetAll()
	if err != nil {
		if code == 0 {
			code = http.StatusInternalServerError
		}
		http_err.NewErrorResponse(c, code, err.Error())
		return
	}

	c.JSON(http.StatusOK, statuss)
}
