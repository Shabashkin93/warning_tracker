package warning

import (
	"net/http"

	"github.com/Shabashkin93/warning_tracker/internal/domain/warning"
	"github.com/Shabashkin93/warning_tracker/internal/service"

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
		api.POST("", t.Create)
	}
}

func (t *transport) Create(c *gin.Context) {
	req := httpWarningCreate{}

	if err := c.ShouldBind(&req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	entry := warning.NewWarningCreate()
	err := warningCreateHttpToDomain(&req, &entry)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	result, err := t.service.Create(&entry)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	resultStruct := httpWarningResponse{}
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	warningCreateDomainToHttp(&result, &resultStruct)

	c.JSON(http.StatusOK, resultStruct)
}
