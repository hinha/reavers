package provider

import (
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
)

// WebContext used by web handler to modify it's request
type WebContext interface {
	// Response returns `*Response`.
	Response() *echo.Response

	// SetResponse sets `*Response`.
	SetResponse(r *echo.Response)

	// SetRequest sets `*http.Request`.
	SetRequest(r *http.Request)

	// IsTLS returns true if HTTP connection is TLS otherwise false.
	IsTLS() bool

	// IsWebSocket returns true if HTTP connection is WebSocket otherwise false.
	IsWebSocket() bool

	// Scheme returns the HTTP protocol scheme, `http` or `https`.
	Scheme() string

	// Request returns `*http.Request`.
	Request() *http.Request

	// RealIP returns the client's network address based on `X-Forwarded-For`
	// or `X-Real-IP` request header.
	// The behavior can be configured using `Echo#IPExtractor`.
	RealIP() string

	// Path returns the registered path for the handler.
	Path() string

	// SetPath sets the registered path for the handler.
	SetPath(p string)

	// Param returns path parameter by name.
	Param(name string) string

	// SetParamNames sets path parameter names.
	SetParamNames(names ...string)

	// SetParamValues sets path parameter values.
	SetParamValues(values ...string)

	// ParamNames returns path parameter names.
	ParamNames() []string

	// ParamValues returns path parameter values.
	ParamValues() []string

	// QueryParam returns the query param for the provided name.
	QueryParam(name string) string

	// QueryParams returns the query parameters as `url.Values`.
	QueryParams() url.Values

	// QueryString returns the URL query string.
	QueryString() string

	// FormParams returns the form parameters as `url.Values`.
	FormParams() (url.Values, error)

	// FormValue returns the form field value for the provided name.
	FormValue(name string) string

	// FormFile returns the multipart form file for the provided name.
	FormFile(name string) (*multipart.FileHeader, error)

	// Cookie returns the named cookie provided in the request.
	Cookie(name string) (*http.Cookie, error)

	// SetCookie adds a `Set-Cookie` header in HTTP response.
	SetCookie(cookie *http.Cookie)

	// Cookies returns the HTTP cookies sent with the request.
	Cookies() []*http.Cookie

	// Get retrieves data from the context.
	Get(key string) interface{}

	// Set saves data in the context.
	Set(key string, val interface{})

	// Bind binds the request body into provided type `i`. The default binder
	// does it based on Content-Type header.
	Bind(i interface{}) error

	// JSON sends a JSON response with status code.
	JSON(code int, i interface{}) error

	// NoContent sends a response with no body and a status code.
	NoContent(code int) error

	// JSONPretty sends a pretty-print JSON with status code.
	JSONPretty(code int, i interface{}, indent string) error

	// JSONBlob sends a JSON blob response with status code.
	JSONBlob(code int, b []byte) error

	// JSONP sends a JSONP response with status code. It uses `callback` to construct
	// the JSONP payload.
	JSONP(code int, callback string, i interface{}) error

	MultipartForm() (*multipart.Form, error)

	Validate(i interface{}) error

	Render(code int, name string, data interface{}) error

	HTML(code int, html string) error
	HTMLBlob(code int, b []byte) error
	String(code int, s string) error

	JSONPBlob(code int, callback string, b []byte) error
	XML(code int, i interface{}) error
	XMLPretty(code int, i interface{}, indent string) error

	XMLBlob(code int, b []byte) error
	Blob(code int, contentType string, b []byte) error
	Stream(code int, contentType string, r io.Reader) error

	File(file string) error
	Attachment(file string, name string) error
	Inline(file string, name string) error

	Redirect(code int, url string) error
	Error(err error)
	Handler() echo.HandlerFunc

	SetHandler(h echo.HandlerFunc)
	Logger() echo.Logger
	SetLogger(l echo.Logger)

	Echo() *echo.Echo
	Reset(r *http.Request, w http.ResponseWriter)
}

// WebHandler handling route web request from client
type WebHandler interface {
	Handle(context WebContext)
	Method() string
	Path() string
	Middleware(handlerFunc echo.HandlerFunc) echo.HandlerFunc
}

// WebEngine ...
type WebEngine interface {
	Run() error
	Inject(handler WebHandler)
	Shutdown(ctx context.Context) error
}
