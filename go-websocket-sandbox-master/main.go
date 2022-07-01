package main

import (
	"net/http"

	"websocket-sandbox/controller"
	"websocket-sandbox/service"

	"websocket-sandbox/config"
	"websocket-sandbox/db"

	"websocket-sandbox/repositories"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	var (
		notificationRepo       = repositories.NewNotificationRepository(db.DB())
		notificationService    = service.NewNotificationService(notificationRepo)
		notificationController = controller.NewNotificationContoroller(notificationService)
	)

	go notificationService.SendMessage()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowCredentials: true,
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	e.GET("/register", func(e echo.Context) error {
		return notificationController.Register(e)
	})

	e.Logger.Fatal(e.Start(config.GetPort()))
}
