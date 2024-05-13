package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrakhaf/halo-suster/domain/nurse/interfaces"
	"github.com/mrakhaf/halo-suster/shared/common"
)

type handlerNurse struct {
	usecase        interfaces.Usecase
	repository     interfaces.Repository
	formatResponse common.JSON
}

func HandlerNurse(privateRoute *echo.Group, usecase interfaces.Usecase, repository interfaces.Repository, formatResponse common.JSON) {
	handler := &handlerNurse{
		usecase:        usecase,
		repository:     repository,
		formatResponse: formatResponse,
	}

	privateRoute.POST("/user/nurse/register", handler.Register)

}

func (h *handlerNurse) Register(c echo.Context) error {
	return c.JSON(http.StatusOK, "register")
}
