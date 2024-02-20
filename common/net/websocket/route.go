package websocket

import (
	"easy-game/config"
	"easy-game/game/client"

	"github.com/gin-gonic/gin"
)

func (w *WebsocketService) Route(route *gin.Engine) {
	route.GET(config.Http.WSRoute, client.WebsocketUpgrader)
}
