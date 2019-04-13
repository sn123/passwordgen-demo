package main

import (
	"net/http"
	"strconv"
	"strings"

	echotemplate "github.com/foolin/echo-template"
	passwordgen "github.com/goavega-software/passwordgenerator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = echotemplate.Default()

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", echo.Map{
			"passwords": "",
		})
	})

	e.POST("/", func(c echo.Context) error {
		i := 0
		num, err := strconv.Atoi(c.FormValue("num"))
		if err != nil {
			num = 8
		}
		var sb strings.Builder
		for i < num {
			sb.WriteString(passwordgen.Generate())
			sb.WriteString("\n")
			i++
		}

		return c.Render(http.StatusOK, "index.html", echo.Map{"passwords": sb.String()})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
