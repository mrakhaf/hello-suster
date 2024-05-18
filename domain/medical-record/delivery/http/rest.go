package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrakhaf/halo-suster/domain/medical-record/interfaces"
	"github.com/mrakhaf/halo-suster/models/request"
	"github.com/mrakhaf/halo-suster/shared/common"
	"github.com/mrakhaf/halo-suster/shared/common/jwt"
	"github.com/mrakhaf/halo-suster/shared/utils"
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

	privateRoute.POST("/medical/patient", handler.SavePatient)
	privateRoute.GET("/medical/patient", handler.GetPatients)
	privateRoute.POST("/medical/record", handler.SaveMedicalRecord)
	privateRoute.GET("/medical/record", handler.GetMedicalRecords)
}

func (h *handlerMedicalRecord) SavePatient(c echo.Context) error {
	_, err := h.jwtAccess.GetUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	var req request.SavePatient

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	isImageTrue := utils.CheckImageType(req.IdentityCardScanImage)

	if !isImageTrue {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	data, err := h.usecase.SavePatient(req)

	if err != nil && err.Error() == "Identity number already exist" {
		return c.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
	}

	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return h.formatResponse.FormatJson(c, http.StatusOK, "success", data)

}

func (h *handlerMedicalRecord) GetPatients(c echo.Context) error {
	_, err := h.jwtAccess.GetUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	var req request.GetPatientsParam

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	data, err := h.usecase.GetPatients(req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return h.formatResponse.FormatJson(c, http.StatusOK, "success", data)
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

	if err != nil {
		if err.Error() == "Identity not found" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return h.formatResponse.FormatJson(c, http.StatusCreated, "success", data)
}

func (h *handlerMedicalRecord) GetMedicalRecords(c echo.Context) error {
	_, err := h.jwtAccess.GetUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	var req request.GetMedicalRecordsParam

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	data, err := h.usecase.GetMedicalRecords(req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return h.formatResponse.FormatJson(c, http.StatusOK, "success", data)
}
