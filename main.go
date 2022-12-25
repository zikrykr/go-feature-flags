package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-feature-flag/config"
	"github.com/go-feature-flag/featureflag"
	"github.com/go-feature-flag/rest"
	"github.com/labstack/echo/v4"

	ffclient "github.com/thomaspoignant/go-feature-flag"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := New(cfg)

	app.Start()
}

type App struct {
	E      *echo.Echo
	config *config.Config
}

func New(config *config.Config) *App {
	app := &App{
		config: config,
		E:      echo.New(),
	}

	app.initFeatureFlag(config)
	app.initRoutes(config)

	return app
}

func (app *App) initRoutes(config *config.Config) {
	restFeatureFlagSvc := rest.NewRestFeatureFlag(config)

	restFeatureFlagController := rest.NewRestFeatureFlagController(app.E, restFeatureFlagSvc)

	users := app.E.Group("/users")
	users.GET("", getUser)

	featureFlag := app.E.Group("/feature-flags")
	featureFlag.POST("", restFeatureFlagController.CreateFeatureFlag)
	featureFlag.PUT("", restFeatureFlagController.UpdateFeatureFlag)
	featureFlag.GET("", restFeatureFlagController.GetFeatureFlags)
	featureFlag.DELETE("/:flagName", restFeatureFlagController.DeleteFeatureFlag)
}

func (app *App) initFeatureFlag(config *config.Config) {
	if err := featureflag.InitFeatureFlag(config); err != nil {
		app.E.Logger.Fatal(err)
	}
	defer ffclient.Close()
}

// Start the server and handle graceful shutdown
func (app *App) Start() {
	// Start server
	go func() {
		if err := app.E.Start(":8089"); err != nil {
			app.E.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.E.Shutdown(ctx); err != nil {
		app.E.Logger.Fatal(err)
	}
}

type User struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

func getUser(c echo.Context) error {
	result := User{}
	result.Name = "John Doe"
	if featureflag.IsFeatureEnabled("enable-show-phone-number") {
		result.PhoneNumber = "081234556789"
	}
	return c.JSON(http.StatusOK, result)
}
