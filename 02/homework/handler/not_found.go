package handler

import (
	"fmt"
	"go-chapter/02/homework/response"
	"net/http"
)

// NewNotFound return not found handler
func NewNotFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response.WriteErrorString(w, http.StatusNotFound, fmt.Sprintf("%s %s not found", r.Method, r.RequestURI))
	})
}