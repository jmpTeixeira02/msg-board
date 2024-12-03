package util

import (
	"encoding/json"
	"net/http"
)

func WriteReplyJson[T any](w http.ResponseWriter, resp T) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
