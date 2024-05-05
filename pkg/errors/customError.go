package errors

import (
	"encoding/json"
	"net/http"
)

func Error(w http.ResponseWriter, statusCode int, errMessage error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	jsonErr, err := json.Marshal(map[string]string{"error": errMessage.Error()})
	if err != nil {
		w.Write([]byte(errMessage.Error()))
	}

	w.Write(jsonErr)
}
