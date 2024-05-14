package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mrakhaf/halo-suster/domain/nurse/interfaces"
	"github.com/mrakhaf/halo-suster/models/request"
	"github.com/mrakhaf/halo-suster/shared/common"
	"github.com/mrakhaf/halo-suster/shared/common/jwt"
	"github.com/mrakhaf/halo-suster/shared/utils"
)

type handlerNurse struct {
	usecase        interfaces.Usecase
	repository     interfaces.Repository
	formatResponse common.JSON
	jwtAccess      *jwt.JWT
}

func HandlerNurse(privateRoute *echo.Group, publicRoute *echo.Group, usecase interfaces.Usecase, repository interfaces.Repository, formatResponse common.JSON, jwtAccess *jwt.JWT) {
	handler := &handlerNurse{
		usecase:        usecase,
		repository:     repository,
		formatResponse: formatResponse,
		jwtAccess:      jwtAccess,
	}

	privateRoute.POST("/user/nurse/register", handler.Register)
	privateRoute.POST("/user/nurse/:userId/access", handler.AccessNurse)
	publicRoute.POST("/user/nurse/login", handler.LoginNurse)

}

func (h *handlerNurse) Register(c echo.Context) error {

	userId, err := h.jwtAccess.GetUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	if userId[:3] != "615" {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	var req request.RegisterNurse

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	err = utils.ValidateNIP(req.NIP, "NURSE")

	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	data, err := h.usecase.Register(req)
	if err != nil && err.Error() == "NIP already exist" {
		return c.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
	}

	if err != nil && err.Error() != "NIP already exist" {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return h.formatResponse.FormatJson(c, http.StatusCreated, "Register success", data)
}

func (h *handlerNurse) AccessNurse(c echo.Context) error {

	userId, err := h.jwtAccess.GetUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	if userId[:3] != "615" {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	var accessNurse request.AccessNurse

	if err := c.Bind(&accessNurse); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	if err := c.Validate(accessNurse); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	nurseId := c.Param("userId")

	nurseIdInt, err := strconv.Atoi(nurseId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	err = utils.ValidateNIP(nurseIdInt, "NURSE")

	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	err = h.usecase.GiveAccessNurse(accessNurse.Password, nurseIdInt)

	if err != nil && err.Error() == "NIP not found" {
		return c.JSON(http.StatusNotFound, "NIP not found")
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "success")

}

func (h *handlerNurse) LoginNurse(c echo.Context) error {
	var req request.Login

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	err := utils.ValidateNIP(req.NIP, "NURSE")

	if err != nil {
		return c.JSON(http.StatusNotFound, "False NIP")
	}

	// 404 user is not found / user is not from nurse (nip not starts with 303)
	// 	- `400` password is wrong
	// - `400` user is not having access
	// - `400` request doesn’t pass validation
	// - `500` if server error

	data, err := h.usecase.LoginNurse(req)
	if err != nil && (err.Error() == "Wrong password" || err.Error() == "NIP not has access") {
		return c.JSON(http.StatusBadRequest, "Wrong password // Not have access")
	}

	if err != nil && err.Error() == "NIP not found" {
		return c.JSON(http.StatusNotFound, "NIP not found")
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return h.formatResponse.FormatJson(c, http.StatusOK, "Login success", data)
}
