package util

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func CaptureStdOutput(f func()) string {
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()
	w.Close()
	bytes, _ := io.ReadAll(r)

	return string(bytes)
}

func WriteJsonResponse[T any](w http.ResponseWriter, statusCode int, resp T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(resp)
}

func EncodeJson(r io.Reader) string {
	bytes, _ := io.ReadAll(r)
	res, _ := json.Marshal(bytes)
	return string(res)
}
