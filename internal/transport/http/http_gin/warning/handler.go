package warning

import (
	"context"
	"net/http"

	"github.com/Shabashkin93/warning_tracker/internal/domain/warning"

	"github.com/gin-gonic/gin"
)

type transport struct {
	usecase WarningService
	handler *gin.Engine
	debug   bool
	url     string
}

type WarningService interface {
	Create(in *warning.WarningCreate) (result warning.WarningResponse, err error)
}

func NewTransport(ctx context.Context, usecase WarningService, i interface{}, url string, debug bool) *transport {
	handler := i.(*gin.Engine)
	return &transport{usecase: usecase, handler: handler, debug: debug, url: url}
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

	result, err := t.usecase.Create(&entry)
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
