package status

import (
	"net/http"

	"github.com/Shabashkin93/warning_tracker/internal/service"
	"github.com/Shabashkin93/warning_tracker/internal/transport/http/http_gin/http_err"

	"github.com/gin-gonic/gin"
)

type transport struct {
	service *service.Service
	handler *gin.Engine
	debug   bool
	url     string
}

func NewTransport(service *service.Service, i interface{}, url string, debug bool) *transport {
	handler := i.(*gin.Engine)
	return &transport{service: service, handler: handler, debug: debug, url: url}
}

func (t *transport) Register(version string, i interface{}) {
	router := i.(*gin.Engine)

	api := router.Group(version + t.url)
	{
		api.GET("", t.GetAll)
	}
}

func (t *transport) GetAll(c *gin.Context) {

	statuss, code, err := t.service.Status.GetAll()
	if err != nil {
		if code == 0 {
			code = http.StatusInternalServerError
		}
		http_err.NewErrorResponse(c, code, err.Error())
		return
	}

	c.JSON(http.StatusOK, statuss)
}
