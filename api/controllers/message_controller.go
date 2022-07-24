package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/samsv78/chat_api_golang/api/auth"
	"github.com/samsv78/chat_api_golang/api/dto"
	"github.com/samsv78/chat_api_golang/api/models"
	"github.com/samsv78/chat_api_golang/api/responses"
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
	messageInfo, err := message.GetMessageInfo(server.DB)
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