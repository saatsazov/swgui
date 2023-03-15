//go:build swguicdn
// +build swguicdn

package v3

import (
	"net/http"

	"github.com/saatsazov/swgui/v3cdn"
)

var staticServer http.Handler

const (
	assetsBase  = v3cdn.AssetsBase
	faviconBase = v3cdn.FaviconBase
)
