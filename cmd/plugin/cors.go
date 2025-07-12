package plugin

import (
	"net/http"

	m "github.com/labstack/echo/v4/middleware"
)

var CustomCors = m.CORSConfig{
	Skipper:      m.DefaultSkipper,
	AllowOrigins: []string{"*"},
	AllowMethods: []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPut,
		http.MethodPatch,
		http.MethodPost,
		http.MethodDelete,
	},
	AllowHeaders:     []string{},
	AllowCredentials: false,
	ExposeHeaders:    []string{},
	MaxAge:           0,
}
