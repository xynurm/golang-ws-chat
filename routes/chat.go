package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/xynurm/golang-ws-chat/internal/adapters/handlers"
	"github.com/xynurm/golang-ws-chat/internal/core/services"
	"github.com/xynurm/golang-ws-chat/pkg/websockets"
)

func ChatRoutes(e *echo.Group) {
	webSocketManager := websockets.GetWebSocketManager()
	webSocketService := services.NewWebSocketService(webSocketManager)

	h := handlers.NewWebSocketHandler(webSocketService)

	e.GET("/", h.HomeHandler)
	e.GET("/ws", h.WebSocketHandler)

}
