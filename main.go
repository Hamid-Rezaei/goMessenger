package main

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"

	"github.com/Hamid-Rezaei/goMessenger/internal/infra/db"
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/http/handler"
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/repository"
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/router"
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/ws"
	"github.com/joho/godotenv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allow all origins during development. You should restrict it in production.
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// serveWs handles ws requests from the peer.
func serveWs(hub *ws.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &ws.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cannot find env file")
	}

	r := router.New()

	dbConnection, err := db.New()
	if err != nil {
		log.Fatal("Cannot connect to database.")
	}

	db.AutoMigrate(dbConnection)

	v1 := r.Group("/api")

	ur := repository.NewUserRepo(dbConnection)
	cr := repository.NewContactRepo(dbConnection)
	chr := repository.NewChatRepo(dbConnection)
	mr := repository.NewMessageRepo(dbConnection)

	h := handler.NewHandler(ur, cr, chr, mr)
	h.Register(v1)

	// Handle WebSocket connections using the serveWs function
	r.GET("/ws", func(c echo.Context) error {
		serveWs(ws.NewHub(), c.Response(), c.Request())
		return nil
	})

	if err := r.Start("0.0.0.0:8080"); err != nil {
		log.Fatalf("server failed to start %v", err)
	}
}
