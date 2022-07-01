package types

import "golang.org/x/net/websocket"

type Client struct {
	Ws        *websocket.Conn
	UserID    int64  `json:"userID"`
	ProjectID int64  `json:"projectID"`
	Token     string `json:"token"`
}

type Event struct {
	Type    string `json:"Type"`
	Message string `json:"Message"`
	Count   int    "json:count"
}
