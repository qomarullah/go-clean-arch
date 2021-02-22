package article

import (
	"github.com/bxcodec/go-clean-arch/pkg/pagination"
	"github.com/bxcodec/go-clean-arch/shared"
	"github.com/labstack/echo"
)

// Handler  represent the httphandler
type handler struct {
	service Service
}

// RegisterHandler ...
func RegisterHandler(e *echo.Echo, service Service) {
	h := &handler{
		service: service,
	}

	e.GET("/articles", h.Get)
	//e.POST("/articles", h.Post)
	//e.GET("/articles/:id", h.GetByID)
	//e.DELETE("/articles/:id", h.Delete)
}

type Request struct {
}

// Get will fetch the item based on given params
func (h *handler) Get(c echo.Context) error {

	var req Request
	var resp shared.Response
	err := c.Bind(&req)
	if err != nil {
		resp = shared.StatusUnprocessableEntity("")
		return c.JSON(resp.StatusCode(), resp.Error())

	}

	var ok bool
	if ok, err = shared.IsRequestValid(&req); !ok {
		resp = shared.StatusBadRequest("")
		return c.JSON(resp.StatusCode(), resp.Error())

	}

	ctx := c.Request().Context()

	count, err := h.service.Count(ctx)
	if err != nil {
		resp = shared.StatusNotFound("")
		return c.JSON(resp.StatusCode(), resp.Error())
	}
	pages := pagination.NewFromRequest(c.Request(), count)
	data, err := h.service.Query(ctx, pages.Offset(), pages.Limit())

	if err != nil {
		resp = shared.StatusInternalServerError(err.Error())
		return c.JSON(resp.StatusCode(), resp.Error())
	}
	pages.Items = data

	resp = shared.Success(pages)
	return c.JSON(resp.StatusCode(), resp)
}
