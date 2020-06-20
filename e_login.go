package main

import (
	"net/http"
	"strings"
	"time"
)

type loginRequest struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Error *string `json:"error"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	req := &loginRequest{}
	err := parseReq(r, req)

	if err != nil {
		w.WriteHeader(400)

		return
	}

	req.Username = strings.ToLower(req.Username)

	target := getUserByUsername(req.Username)

	if target == nil || target.password != req.Password {
		w.WriteHeader(401)

		resErr := "Incorrect username or password"

		writeResp(w, &loginResponse{
			Error: &resErr,
		})

		return
	}

	tokStr, err := genToken(64)

	if err != nil {
		w.WriteHeader(500)

		return
	}

	// Push session into sessions

	sessions[tokStr] = target.username

	http.SetCookie(w, &http.Cookie{
		Name:    "sess",
		Value:   tokStr,
		Expires: time.Now().Add(time.Hour * 24 * 4),
	})

	w.WriteHeader(200)

	writeResp(w, &loginResponse{
		Error: nil,
	})
}
