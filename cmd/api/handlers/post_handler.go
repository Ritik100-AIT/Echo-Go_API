package handlers

import (
	"net/http"
	"strconv"

	"github.com/RitikAIT100/cmd/api/services"
	"github.com/labstack/echo/v4"
)

func PostIndexHandler(c echo.Context) error {
	data, err := services.GetAll()
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process")
	}
	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}

func PostSingleHandler(c echo.Context) error {
	id := c.Param("id")
	indx, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process Id")
	}
	data, err := services.GetById(indx)
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process the data from Id")
	}
	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)

}
