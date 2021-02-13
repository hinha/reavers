package web

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"strconv"

	"github.com/google/uuid"
	"github.com/hinha/reavers/provider"
	"github.com/hinha/reavers/provider/web/command"
	"github.com/hinha/reavers/provider/web/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// WEB ...
type WEB struct {
	engine *echo.Echo
	port   int
}

// Fabricate ...
func Fabricate(givenPort int) *WEB {
	return &WEB{engine: echo.New(), port: givenPort}
}

// FabricateCommand insert web related command
func (w *WEB) FabricateCommand(cmd provider.Command) {
	cmd.InjectCommand(command.NewRun(w))
}

// Inject new Route into reavers
func (w *WEB) Inject(handler provider.WebHandler) {
	w.engine.Add(handler.Method(), handler.Path(), func(context echo.Context) error {
		req := context.Request()
		if reqID := req.Header.Get("X-Request-ID"); reqID != "" {
			context.Set("request-id", reqID)
		} else {
			context.Set("request-id", uuid.New().String())
		}

		if userID := req.Header.Get("Resource-Owner-ID"); userID != "" {
			convertedUserID, err := strconv.Atoi(userID)
			if err == nil {
				context.Set("user-id", convertedUserID)
			}
		}

		handler.Handle(context)
		return nil
	})
}

// Run ...
func (w *WEB) Run() error {
	w.engine.Use(middleware.Logger())
	w.engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderAccessControlAllowOrigin,
		},
	}))
	w.engine.Static("/", "assets")
	w.engine.Renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.tmpl")),
	}

	w.Inject(handler.NewIndexPage())
	return w.engine.Start(fmt.Sprintf(":%d", w.port))
}

// Shutdown web engine
func (w *WEB) Shutdown(ctx context.Context) error {
	return w.engine.Shutdown(ctx)
}

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
