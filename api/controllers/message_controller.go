package controllers

import (
	"encoding/json"
	// "errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/samsv78/chat_api_golang/api/auth"
	"github.com/samsv78/chat_api_golang/api/dto"
	"github.com/samsv78/chat_api_golang/api/models"
	"github.com/samsv78/chat_api_golang/api/responses"
	"github.com/samsv78/chat_api_golang/api/services"
)

func (server *Server) SendMessage(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	request := dto.SendMessageRequest{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	message := models.Message{}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	message.Ctor(request, uid)
	messageInfo, err := services.GetMessageInfo(server.DB, message)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = message.SaveMessage(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	
	//signal via rabbitmq

	messageInfo.ID = message.ID
	responses.JSON(w, http.StatusCreated, messageInfo)
}

func (server *Server) GetChatRooms(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	chatrooms, err := services.GetChatRoomsInfo(server.DB, userID)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, chatrooms)
}

func (server *Server) GetChatRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	otherUserID, err := strconv.ParseUint(vars["otherUserID"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	chatRoom, err := services.GetChatRoomInfo(server.DB, userID, uint32(otherUserID))
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// if chatRoom.Messages == nil{
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("No Messages Found!"))
	// 	return
	// }
	responses.JSON(w, http.StatusOK, chatRoom)
}