package controllers

import (
	"net/http"
	"strconv"

	"golang-crud-test/models"
	"github.com/labstack/echo"
)

func FetchAllUser(c echo.Context) error {
	result, err := models.FetchAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreUser(c echo.Context) (err error) {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)

	u := new(models.User)
	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreUser(conv_id, u.Username, u.Password, u.FullName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateUser(c echo.Context) (err error) {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)

	u := new(models.User)
	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateUser(conv_id, u.Username, u.Password, u.FullName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteUser(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
