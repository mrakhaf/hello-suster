package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrakhaf/halo-suster/domain/medical-record/interfaces"
	"github.com/mrakhaf/halo-suster/models/request"
	"github.com/mrakhaf/halo-suster/shared/common"
	"github.com/mrakhaf/halo-suster/shared/common/jwt"
)

type handlerMedicalRecord struct {
	usecase        interfaces.Usecase
	repository     interfaces.Repository
	formatResponse common.JSON
	jwtAccess      *jwt.JWT
}

func HandlerMedicalRecord(privateRoute *echo.Group, publicRoute *echo.Group, usecase interfaces.Usecase, repository interfaces.Repository, formatResponse common.JSON, jwtAccess *jwt.JWT) {
	handler := &handlerMedicalRecord{
		usecase:        usecase,
		repository:     repository,
		formatResponse: formatResponse,
		jwtAccess:      jwtAccess,
	}

	privateRoute.POST("/medical/patient", handler.SaveMedicalRecord)
}

func (h *handlerMedicalRecord) SaveMedicalRecord(c echo.Context) error {
	_, err := h.jwtAccess.GetUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	var req request.SaveMedicalRecord

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	data, err := h.usecase.SaveMedicalRecord(req)

	if err != nil && err.Error() == "Identity number already exist" {
		return c.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return h.formatResponse.FormatJson(c, http.StatusOK, "success", data)

}
