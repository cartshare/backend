package main

import (
	"fmt"
	"net/http"
	"strings"
)

type signupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Name     string `json:"name"`
}

type signupResponse struct {
	Error *string `json:"error"`
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	req := &signupRequest{}
	err := parseReq(r, req)

	if err != nil {
		w.WriteHeader(400)

		return
	}

	req.Username = strings.ToLower(req.Username)

	userWithSameUsername := getUserByUsername(req.Username)

	if userWithSameUsername != nil {
		w.WriteHeader(400)

		emsg := "Username is taken"

		writeResp(w, &signupResponse{
			Error: &emsg,
		})

		return
	}

	// Perform geocoding

	loc, err := geocoder.Geocode(req.Address)

	if err != nil {
		w.WriteHeader(404)

		emsg := err.Error()

		writeResp(w, &signupResponse{
			Error: &emsg,
		})

		return
	}

	// Create user

	u := &user{
		username: req.Username,
		name:     req.Name,
		password: req.Password,
		loc:      loc,
	}

	users = append(users, u)

	fmt.Println("Pushed user", u)
	fmt.Println(*u.loc)

	// Hand it over to login

	loginHandler(w, r)
}
