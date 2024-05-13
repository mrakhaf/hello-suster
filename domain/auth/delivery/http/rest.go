package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrakhaf/halo-suster/domain/auth/interfaces"
	"github.com/mrakhaf/halo-suster/models/request"
	"github.com/mrakhaf/halo-suster/shared/common"
	"github.com/mrakhaf/halo-suster/shared/utils"
)

type handlerAuth struct {
	usecase    interfaces.Usecase
	repository interfaces.Repository
	Json       common.JSON
}

func AuthHandler(authRoute *echo.Group, usecase interfaces.Usecase, repository interfaces.Repository, Json common.JSON) {
	handler := handlerAuth{
		usecase:    usecase,
		repository: repository,
		Json:       Json,
	}

	authRoute.POST("/user/it/register", handler.Register)
	authRoute.POST("/user/it/login", handler.Login)
}

func (h *handlerAuth) Register(c echo.Context) error {

	var req request.Register

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	//validate
	err := utils.ValidateNIP(req.NIP, "IT")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	data, err := h.usecase.Register(req)
	if err != nil && err.Error() == "NIP already exist" {
		return c.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
	}

	if err != nil && err.Error() != "NIP already exist" {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return h.Json.FormatJson(c, http.StatusCreated, "Register success", data)
}

func (h *handlerAuth) Login(c echo.Context) error {
	var req request.Login

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	//validate
	err := utils.ValidateNIP(req.NIP, "IT")
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	data, err := h.usecase.Login(req)
	if err != nil && err.Error() == "Wrong password" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err != nil && err.Error() == "NIP not found" {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return h.Json.FormatJson(c, http.StatusOK, "Login success", data)

}
