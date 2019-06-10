package main

import (
	"github.com/labstack/echo"
	"github.com/td_rule_engine/router"
)

func main() {
	e := echo.New()
	router.InitRouter(e)
	e.Logger.Fatal(e.Start(":9090"))
}
