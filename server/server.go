package airdisk

import (
	"html/template"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io"
)

type Template struct{
	templates *template.Template
}
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error{
	return t.templates.ExecuteTemplate(w, name, data)
}

var (
	Opts *Options
)

func Run()  {

	t := &Template{
		templates: template.Must(template.ParseGlob("template/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.Static("/static","static")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	portalCtx := NewPortalCtx()


	//mgGroup := e.Group("")
	//mgGroup.GET("/index.html", portalCtx.Portal)
	//mgGroup.POST("/upgrade", portalCtx.Upgrade)
	//mgGroup.POST("/config", portalCtx.Config)
	//db := InitDB("./airdisk.db")
	//defer db.Close()

	e.GET("/", portalCtx.Portal)
	e.GET("/index.html", portalCtx.Portal)
	e.POST("/upgrade", portalCtx.Upgrade)
	e.POST("/control", portalCtx.Control)

	e.Logger.Fatal(e.Start(Opts.Port))
}