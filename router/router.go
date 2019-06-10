package router

import (
	"github.com/labstack/echo"
	"github.com/td_rule_engine/controller"
)

func InitRouter(e *echo.Echo) {
	e.GET("/", controller.Test2)
}
