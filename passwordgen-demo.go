package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	echotemplate "github.com/foolin/echo-template"
	passwordgen "github.com/goavega-software/passwordgenerator"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Renderer = echotemplate.Default()
	e.Static("/js", "js")
	e.Static("/css", "css")
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", echo.Map{
			"passwords": "",
		})
	})

	e.POST("/api/generate", func(c echo.Context) error {

		var response struct {
			Data struct {
				Passwords []string `json:"passwords"`
			} `json:"data"`
			Errors []string `json:"errors"`
		}
		request := new(passwordgen.GenerateRequest)

		if reqErr := c.Bind(request); reqErr != nil {
			fmt.Print(reqErr)
			response.Errors = append(response.Errors, "Invalid request")
			return c.JSON(http.StatusBadRequest, response)
		}
		i := 0
		if request.NumPasswords > 8 {
			request.NumPasswords = 8
		}
		var (
			passwords []string
			password  string
		)
		for i < request.NumPasswords {
			password = passwordgen.Generate(*request)
			passwords = append(passwords, password)
			i++
		}

		response.Data.Passwords = passwords
		return c.JSON(http.StatusOK, response)
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), e))
}
