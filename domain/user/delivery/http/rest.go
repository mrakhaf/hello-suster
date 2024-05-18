package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mrakhaf/halo-suster/domain/user/interfaces"
	"github.com/mrakhaf/halo-suster/models/request"
	"github.com/mrakhaf/halo-suster/shared/common"
	"github.com/mrakhaf/halo-suster/shared/common/jwt"
	"github.com/mrakhaf/halo-suster/shared/utils"
)

type handlerUser struct {
	usecase        interfaces.Usecase
	repository     interfaces.Repository
	formatResponse common.JSON
	jwtAccess      *jwt.JWT
}

func HandlerUser(privateRoute *echo.Group, publicRoute *echo.Group, usecase interfaces.Usecase, repository interfaces.Repository, formatResponse common.JSON, jwtAccess *jwt.JWT) {
	handler := &handlerUser{
		usecase:        usecase,
		repository:     repository,
		formatResponse: formatResponse,
		jwtAccess:      jwtAccess,
	}

	privateRoute.POST("/user/nurse/register", handler.Register)
	privateRoute.POST("/user/nurse/:userId/access", handler.AccessNurse)
	publicRoute.POST("/user/nurse/login", handler.LoginNurse)
	privateRoute.GET("/user", handler.GetUsers)
	privateRoute.PUT("/user/nurse/:userId", handler.EditNurse)
	privateRoute.DELETE("/user/nurse/:userId", handler.DeleteNurse)

}

func (h *handlerUser) Register(c echo.Context) error {

	userId, err := h.jwtAccess.GetUserIdFromToken(c)

	fmt.Println(userId)

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

func (h *handlerUser) AccessNurse(c echo.Context) error {

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

	err = h.usecase.GiveAccessNurse(accessNurse.Password, nurseId)

	if err != nil && err.Error() == "UserId not found / nip not nurse" {
		return c.JSON(http.StatusNotFound, "UserId not found / nip not nurse")
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "success")

}

func (h *handlerUser) LoginNurse(c echo.Context) error {
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

func (h *handlerUser) GetUsers(c echo.Context) error {

	userId, err := h.jwtAccess.GetUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	if userId[:3] != "615" {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	var req request.GetUsers

	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	data, err := h.usecase.GetUsers(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return h.formatResponse.FormatJson(c, http.StatusOK, "success", data)
}

func (h *handlerUser) EditNurse(c echo.Context) error {

	userId, err := h.jwtAccess.GetUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	if userId[:3] != "615" {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	var req request.EditNurse

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	nurseNipString := strconv.Itoa(req.NIP)

	if nurseNipString[:3] != "303" {
		return c.JSON(http.StatusNotFound, "bad request")
	}

	err = utils.ValidateNIP(req.NIP, "NURSE")

	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	nurseId := c.Param("userId")

	err = h.usecase.UpdateNurse(req, nurseId)

	if err != nil && err.Error() == "NIP already exist" {
		return c.JSON(http.StatusConflict, "NIP already exist")
	}

	if err != nil && err.Error() == "UserId not found" {
		return c.JSON(http.StatusNotFound, "UserId not found")
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "success")
}

func (h *handlerUser) DeleteNurse(c echo.Context) error {

	userId, err := h.jwtAccess.GetUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	if userId[:3] != "615" {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	var userIdNurse = c.Param("userId")

	err = h.usecase.DeleteNurse(userIdNurse)

	if err != nil && err.Error() == "not delete anything" {
		return c.JSON(http.StatusBadRequest, "not delete anything")
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "success")

}
