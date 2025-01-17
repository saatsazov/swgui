// Package v4cdn provides Swagger UI v4 via CDN.
package v4cdn

import (
	"net/http"

	"github.com/saatsazov/swgui"
	"github.com/saatsazov/swgui/internal"
)

// Handler handle swagger UI request.
type Handler = internal.Handler

// NewWithConfig creates configurable handler constructor.
func NewWithConfig(config swgui.Config) func(title, swaggerJSONPath string, basePath string) http.Handler {
	return func(title, swaggerJSONPath string, basePath string) http.Handler {
		if config.Title == "" {
			config.Title = title
		}

		if config.SwaggerJSON == "" {
			config.SwaggerJSON = swaggerJSONPath
		}

		if config.BasePath == "" {
			config.BasePath = basePath
		}

		return NewHandlerWithConfig(config)
	}
}

// New creates HTTP handler for Swagger UI.
func New(title, swaggerJSONPath string, basePath string) http.Handler {
	return NewHandler(title, swaggerJSONPath, basePath)
}

// NewHandler creates HTTP handler for Swagger UI.
func NewHandler(title, swaggerJSONPath string, basePath string) *Handler {
	return NewHandlerWithConfig(swgui.Config{
		Title:       title,
		SwaggerJSON: swaggerJSONPath,
		BasePath:    basePath,
	})
}

// NewHandlerWithConfig creates HTTP handler for Swagger UI.
func NewHandlerWithConfig(config swgui.Config) *Handler {
	return internal.NewHandlerWithConfig(config, AssetsBase, FaviconBase, nil)
}
