package handlers

import (
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/xynurm/golang-ws-chat/internal/core/ports"
)

type webSocketHandlerImpl struct {
	webSocketService ports.WebSocketService
}

func NewWebSocketHandler(webSocketService ports.WebSocketService) ports.WebSocketHandlers {
	return &webSocketHandlerImpl{webSocketService}
}

func (h *webSocketHandlerImpl) HomeHandler(ctx echo.Context) error {
	content, err := os.ReadFile("view/index.html")
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Could not open requested file")
	}

	return ctx.HTMLBlob(http.StatusOK, content)
}

func (h *webSocketHandlerImpl) WebSocketHandler(ctx echo.Context) error {
	currentGorillaConn, err := websocket.Upgrade(ctx.Response(), ctx.Request(), ctx.Response().Header(), 1024, 1024)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Could not open WebSocket connection")
	}

	username := ctx.Request().URL.Query().Get("username")
	h.webSocketService.WebSocketConnection(currentGorillaConn, username)

	return nil

}
