package service

import (
	"log"
	"time"
	"websocket-sandbox/interfaces"
	types "websocket-sandbox/types"

	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

const messageInternal = 3

var Clients = make(map[types.Client]bool)

type NotificationService struct {
	notificationRepository interfaces.NotificationRepository
}

func NewNotificationService(notificationRepository interfaces.NotificationRepository) interfaces.NotificationService {
	return &NotificationService{
		notificationRepository: notificationRepository,
	}
}

func (s *NotificationService) Register(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			var client types.Client
			err := websocket.JSON.Receive(ws, &client)
			if err != nil {
				s.RemoveClient(client)
			} else {
				client.UserID = 1
				client.ProjectID = 1
				client.Ws = ws
				s.AppendeClient(client)
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func (s *NotificationService) SendMessage() {
	for range time.Tick(messageInternal * time.Second) {
		for client := range Clients {
			notSendsRecords, err := s.notificationRepository.GetNonSend(client.UserID, client.ProjectID)
			log.Print(notSendsRecords)
			if err != nil {
				// TODO : logger
			} else {
				noticeCount := len(notSendsRecords)
				if noticeCount > 0 {
					ws := client.Ws
					// TODO :
					ev := &types.Event{Type: "notice_sample", Message: "update notice count", Count: noticeCount}
					err = websocket.JSON.Send(ws, ev)
					if err != nil {
						s.RemoveClient(client)
						ws.Close()
					} else {
						for _, v := range notSendsRecords {
							s.notificationRepository.UpdateSendAt(v.ID)
						}
					}
				}
			}
		}
	}
}

func (s *NotificationService) AppendeClient(newClient types.Client) {
	for client := range Clients {
		if client.UserID == newClient.UserID && client.ProjectID != newClient.ProjectID {
			s.RemoveClient(client)
			Clients[newClient] = true
		}
	}
	if !Clients[newClient] {
		Clients[newClient] = true
	}
}

func (s *NotificationService) RemoveClient(client types.Client) {
	delete(Clients, client)
}
