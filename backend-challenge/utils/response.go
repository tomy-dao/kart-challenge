package utils

import (
	"encoding/json"
	"local/model"
	"net/http"
)

func HandlerResponse[T any](w http.ResponseWriter, response model.Response[T]) {
	if response.Error() {
		w.WriteHeader(response.StatusCode())
		return
	}
	w.WriteHeader(response.StatusCode())
	json.NewEncoder(w).Encode(response.Data)
}
