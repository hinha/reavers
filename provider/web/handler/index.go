package handler

import (
	"net/http"

	"github.com/hinha/reavers/provider"
	"github.com/labstack/echo/v4"
)

// IndexPage ...
type IndexPage struct{}

// NewIndexPage ...
func NewIndexPage() *IndexPage {
	return &IndexPage{}
}

// Path return web path
func (p *IndexPage) Path() string {
	return "/"
}

// Method return web method
func (p *IndexPage) Method() string {
	return "GET"
}

// Middleware ...
func (p *IndexPage) Middleware(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return handlerFunc(ctx)
	}
}

// Handle health which always return 200
func (p *IndexPage) Handle(context provider.WebContext) {
	_ = context.Render(http.StatusOK, "index.html", map[string]interface{}{
		"title": "Test Twitter",
	})
}
