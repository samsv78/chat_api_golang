package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errObject struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	errObject := errObject{Error: err.Error()}
	if err != nil {
		JSON(w, statusCode, errObject)
	} else {
		JSON(w, http.StatusBadRequest, nil)
	}

}
