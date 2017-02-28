package regserver

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Run()  {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:"time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	wechatCtx := NewWeChatCtx()

	e.GET("/wx", wechatCtx.Handle)
	e.POST("/wx", wechatCtx.HandlePost)
	e.Logger.Fatal(e.Start(":80"))
}