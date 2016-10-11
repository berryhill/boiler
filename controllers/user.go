package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/boiler/models"
)

// this variable is a work around for not having a database
// vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
var (
	users = map[int]*models.User{}
	seq   = 1
)
// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

func CreateUser(c echo.Context) error {
	u := &models.User{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}

	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func UpdateUser(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}