package customerrors

import (
	"encoding/json"
	"log"
	"net/http"
)

func HttpError(w http.ResponseWriter, r *http.Request, errorCode int, errorMessage string, err error) {
	if err != nil {
		log.Println(err.Error())
	}
	errResponse := errorResponse{
		errorMessage,
	}

	w.WriteHeader(errorCode)
	errBytes, marshalErr := json.Marshal(errResponse)
	if marshalErr != nil {
		log.Println(err)
	}

	w.Write(errBytes)
}

type errorResponse struct {
	Message string `json:"message"`
}
