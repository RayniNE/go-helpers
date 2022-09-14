package httphelper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DecodeJson(r *http.Request, dto interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		return fmt.Errorf("Error while decoding data: %v", err.Error())
	}

	return nil
}

func EncodeResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
