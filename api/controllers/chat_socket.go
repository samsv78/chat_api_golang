package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/samsv78/chat_api_golang/api/auth"
	"github.com/samsv78/chat_api_golang/api/helpers"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (server *Server) wsEnpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// log.Println("Client Connected")
	reader(ws, server)
	res, index := helpers.ContainsWSConnection(server.WSConnections, ws)
	log.Println("user " + server.ConnectedClientsIds[index] + " disconnected")
	if res{
		server.WSConnections = helpers.RemoveElementByIndex(server.WSConnections, index)
		server.ConnectedClientsIds = helpers.RemoveElementByIndex(server.ConnectedClientsIds, index)
	}
	ws.Close()
	

}

func reader(conn *websocket.Conn, server *Server) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		
		tokenID, err := auth.CalcTokenID(string(p))
		if err != nil {
			log.Println(err)
			return
		}
		tokenIDSTR := strconv.Itoa(int(tokenID))
		log.Println("user " + tokenIDSTR + " connected")
		server.ConnectedClientsIds = append(server.ConnectedClientsIds, tokenIDSTR)
		server.WSConnections = append(server.WSConnections, conn)
		err = conn.WriteMessage(messageType, []byte("hi " + tokenIDSTR))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
