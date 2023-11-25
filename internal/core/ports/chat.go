package ports

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/xynurm/golang-ws-chat/internal/core/models"
)

type WebSocketService interface {
	WebSocketConnection(conn *websocket.Conn, usernam string)
	WebSocketService(currentConn *models.WebSocketConnection, connection []*models.WebSocketConnection)
}

type WebSocketHandlers interface {
	HomeHandler(ctx echo.Context) error
	WebSocketHandler(ctx echo.Context) error
}
