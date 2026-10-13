package handlers

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Echo upgrades the request to a WebSocket and sends back whatever it receives.
func Echo(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		slog.Warn("upgrade failed", "err", err)
		return
	}
	defer conn.Close()

	slog.Info("ws connected", "remote", c.Request.RemoteAddr)

	defer slog.Info("ws handler exited", "remote", c.Request.RemoteAddr)

	for {
		msgType, data, err := conn.ReadMessage()
		if err != nil {
			slog.Info("ws disconnected", "err", err)
			return
		}
		slog.Info("ws message", "type", msgType, "bytes", len(data), "text", string(data))
		if err := conn.WriteMessage(msgType, data); err != nil {
			slog.Warn("write failed", "err", err)
			return
		}
	}
}
