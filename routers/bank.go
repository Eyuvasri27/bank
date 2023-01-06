package routers

import (
	"echobank/controlers"

	"github.com/labstack/echo"
)

func Setbankrouter(e *echo.Echo) {

	e.POST("/bank", controlers.Postbank)
	e.GET("/bank", controlers.Listbank)
	e.GET("/bank/:id", controlers.Listbyidbank)
	e.PUT("/bank/:id", controlers.Putbank)
	e.DELETE("/bank/:id", controlers.Deletebank)
	e.POST("/createaccount", controlers.CreateAccount)
	e.POST("/verifyaccount", controlers.LoginAndToken)

}
