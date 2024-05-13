package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	authHandler "github.com/mrakhaf/halo-suster/domain/auth/delivery/http"
	authRepository "github.com/mrakhaf/halo-suster/domain/auth/repository"
	authUsecase "github.com/mrakhaf/halo-suster/domain/auth/usecase"
	"github.com/mrakhaf/halo-suster/shared/common"
	formatJson "github.com/mrakhaf/halo-suster/shared/common/json"
	"github.com/mrakhaf/halo-suster/shared/common/jwt"
	"github.com/mrakhaf/halo-suster/shared/config/database"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = common.NewValidator()

	err := godotenv.Load(".env")
	if err != nil {
		e.Logger.Fatal(err)
	}

	//db config
	database, err := database.ConnectDB()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Enqilo Store!")
	})

	//create group
	publicGroup := e.Group("/v1")

	restrictedGroup := e.Group("/v1")
	restrictedGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte("secret"),
		ErrorHandler: func(err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		},
	}))

	//
	formatResponse := formatJson.NewResponse()
	jwtAccess := jwt.NewJWT()

	//auth
	authRepo := authRepository.NewRepository(database)
	authUsecase := authUsecase.NewUsecase(authRepo, jwtAccess)
	authHandler.AuthHandler(publicGroup, authUsecase, authRepo, formatResponse)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}
