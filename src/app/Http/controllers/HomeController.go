package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var HomeController *homeController

type homeController struct{}

func (t *homeController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, "nice")
}
