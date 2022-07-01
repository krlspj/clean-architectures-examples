package interfaces

import (
	"websocket-sandbox/models"
	types "websocket-sandbox/types"

	"github.com/labstack/echo"
)

type NotificationController interface {
	Register(echo echo.Context) error
}

type NotificationService interface {
	Register(c echo.Context) error
	SendMessage()
	AppendeClient(client types.Client)
	RemoveClient(client types.Client)
}

type NotificationRepository interface {
	UpdateSendAt(id int64) error
	UpdateReadAt(id int64) error
	Add(userId, project_id int64) error
	GetNonSend(userId, project_id int64) ([]models.Notifications, error)
	GetNonRead(userId, project_id int64) ([]models.Notifications, error)
}
