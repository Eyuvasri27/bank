package main

import (
	"echobank/common"
	"echobank/routers"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	common.Database()
	routers.Setbankrouter(e)
	e.Logger.Fatal(e.Start(":1323"))
}
