package main

import (
	"net/http"
)

type sessionResponse struct {
	Error    *string `json:"error"`
	Username string  `json:"email"`
}

func sessionHandler(u *user, w http.ResponseWriter, r *http.Request) {
	writeResp(w, &sessionResponse{
		Error:    nil,
		Username: u.username,
	})
}
