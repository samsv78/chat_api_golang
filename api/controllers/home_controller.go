package controllers

import (
	"net/http"

	"github.com/samsv78/chat_api_golang/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To Chat API!")
}
