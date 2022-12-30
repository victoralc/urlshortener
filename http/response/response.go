package response

import (
	"fmt"
	"net/http"
)

type Headers map[string]string

func With(w http.ResponseWriter, statusCode int, headers Headers) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(statusCode)
}

func WithJson(w http.ResponseWriter, json string) {
	With(w, http.StatusOK, Headers{
		"Content-Type": "application/json",
	})
	fmt.Fprintf(w, json)
}
