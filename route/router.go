package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsUpgrader websocket.Upgrader

func init() {
	wsUpgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
}

type wsClients struct {
	Conn       *websocket.Conn
	RemoteAddr string
	UserName   string
	RoomID     int
}

type message struct {
	Type    int
	Content string
	User    int
}

func Run(c *gin.Context) {
	wsConn, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		// TODO: do something
		return
	}
	defer wsConn.Close()

	go read(wsConn)
	go write()

	select {} // blocked forever
}

func read(c *websocket.Conn) {
	defer func() {
		if err := recover(); err != nil { // catch exception
			// TODO: do something
		}
	}()
	for {
		msgType, msg, err := c.ReadMessage()
		if err != nil {
			// TODO: do something
			return
		}
	}
}

func write() {
	defer func() {
		if err := recover(); err != nil { // catch exception
			// TODO: do something
		}
	}()
	for {
		select { // looking for order from channel

		}

	}
}
