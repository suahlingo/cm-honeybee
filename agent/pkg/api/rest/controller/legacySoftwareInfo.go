package controller

import (
	"net/http"

	"github.com/cloud-barista/cm-honeybee/agent/driver/software"
	"github.com/cloud-barista/cm-honeybee/agent/pkg/api/rest/common"
	_ "github.com/cloud-barista/cm-honeybee/agent/pkg/api/rest/model/onprem/legacy" // Need for swag
	"github.com/labstack/echo/v4"
)

// GetLegacySoftwareInfo godoc
//
//	@ID				get-legacy-software-info
//	@Summary		Get a list of legacy software information
//	@Description	Collect and return legacy software information from the host (via /proc scanning).
//	@Tags			[Legacy] Get legacy software info
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	legacy.LegacySoftware	"Successfully retrieved legacy software information."
//	@Failure		400	{object}	common.ErrorResponse	"Sent bad request."
//	@Failure		500	{object}	common.ErrorResponse	"Failed to get information of legacy software."
//	@Router			/software/legacy [get]
func GetLegacySoftwareInfo(c echo.Context) error {
	legacyInfo, err := software.GetLegacySoftwareInfo()
	if err != nil {
		return common.ReturnInternalError(c, err, "Failed to get information of legacy software.")
	}

	return c.JSONPretty(http.StatusOK, legacyInfo, " ")
}
