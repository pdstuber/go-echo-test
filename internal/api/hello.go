package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HelloWorld struct {
	Hello string `json:"hello"`
}

func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, &HelloWorld{Hello: "world!"})
}
