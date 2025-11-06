package route

import (
	"strings"

	"github.com/cloud-barista/cm-honeybee/agent/common"
	"github.com/cloud-barista/cm-honeybee/agent/pkg/api/rest/controller"
	_ "github.com/cloud-barista/cm-honeybee/agent/pkg/api/rest/docs"
	"github.com/labstack/echo/v4"
)

// RegisterLegacy sets up routes for legacy software collection
func RegisterLegacy(e *echo.Echo) {
	e.GET("/"+strings.ToLower(common.ShortModuleName)+"/software/legacy", controller.GetLegacySoftwareInfo)
}
