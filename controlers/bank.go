package controlers

import (
	"echobank/models"
	"echobank/storages"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func Postbank(c echo.Context) error {
	u := new(models.Bank) //structure
	if err := c.Bind(u); err != nil {
		return err
	}
	cur := storages.Getcursor()
	cur.Createbank(u)
	return c.JSON(http.StatusCreated, u)
	//return c.String(http.StatusOK, "ok")
}

func Listbank(c echo.Context) error {
	cur := storages.Getcursor()
	result, _ := cur.Listbank()

	return c.JSON(http.StatusCreated, result)
}

func Listbyidbank(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	cur := storages.Getcursor()
	result, _ := cur.Listbyidbank(Id)

	return c.JSON(http.StatusOK, result)

}

func Putbank(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	cur := storages.Getcursor()
	u := new(models.Bank)
	if err := c.Bind(u); err != nil {
		return err
	}
	fmt.Println("u", u)
	cur.Putbank(Id, u)
	return c.JSON(http.StatusCreated, u)

}

func Deletebank(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	cur := storages.Getcursor()
	cur.Deletebank(Id)

	return c.JSON(http.StatusOK, "Deleted")

}
