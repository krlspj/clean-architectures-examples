package controller

import (
	"websocket-sandbox/interfaces"

	"github.com/labstack/echo"
)

type NotificationController struct {
	notificationService interfaces.NotificationService
}

func NewNotificationContoroller(notificationService interfaces.NotificationService) interfaces.NotificationController {
	return &NotificationController{
		notificationService: notificationService,
	}
}

func (c *NotificationController) Register(contex echo.Context) error {
	return c.notificationService.Register(contex)
}
