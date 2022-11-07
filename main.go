package main

import (
	"net/http"

	"github.com/go-feature-flag/featureflag"
	"github.com/labstack/echo/v4"
	ffclient "github.com/thomaspoignant/go-feature-flag"
)

func main() {

	e := echo.New()

	if err := featureflag.InitFeatureFlag(); err != nil {
		e.Logger.Fatal(err)
	}
	defer ffclient.Close()

	e.GET("/users", getUser)
	e.Logger.Fatal(e.Start(":8082"))
}

type User struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

func getUser(c echo.Context) error {
	result := User{}
	result.Name = "John Doe"
	if featureflag.IsFeatureEnabled("enable-show-phone-number") {
		result.PhoneNumber = "081234556789"
	}
	return c.JSON(http.StatusOK, result)
}
