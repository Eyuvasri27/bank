package controlers

import (
	"echobank/models"
	"echobank/storages"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func CreateAccount(c echo.Context) error {
	bank := new(models.BankUserDetails)
	if err := c.Bind(bank); err != nil {
		return err
	}
	cur := storages.Getcursor()
	cur.AccountDetails(bank)
	return c.JSON(http.StatusCreated, bank)
}
func LoginAndToken(c echo.Context) error {
	BV := new(models.BankUserDetails)
	if err := c.Bind(BV); err != nil {
		return err
	}
	cur := storages.Getcursor()
	result, err := cur.Verification(BV.Username, BV.Password)
	fmt.Println("result:", result)
	fmt.Println("err:", err)
	if err == nil {
		expirationtime := time.Now().Add(time.Hour * 5)

		claims := &jwt.StandardClaims{
			Issuer:    "BANKDETAILS",
			Subject:   "VERIFY TOKEN",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expirationtime.Unix(),
		}
		fmt.Println(claims)

		var sign = "bankdetailsverification"

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		signed := []byte(sign)
		ss, err := token.SignedString(signed)
		if err != nil {
			fmt.Println(err)
			fmt.Println("tokenstring:  ", ss)

			return c.JSON(http.StatusOK, ss)
		}
		return c.JSON(http.StatusOK, ss)
	}
	/*if result != nil {
		expirationtime := time.Now().Add(time.Hour * 5)

		claims := &jwt.StandardClaims{
			Issuer:    "BANKDETAILS",
			Subject:   "VERIFY TOKEN",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expirationtime.Unix(),
		}
		fmt.Println(claims)

		var sign = "bankdetailsverification"

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		signed := []byte(sign)
		ss, err := token.SignedString(signed)
		if err != nil {
			fmt.Println(err)
			fmt.Println("tokenstring:  ", ss)

			return c.JSON(http.StatusOK, ss)
		}
		return c.JSON(http.StatusOK, ss)
	}*/

	return c.JSON(http.StatusBadRequest, "loginfailed")
}

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
