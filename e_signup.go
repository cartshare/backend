package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"net/http"
	"strings"
	"time"
)

type signupRequest struct {
	Username string `json:"email"`
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

	emailErr := checkmail.ValidateFormat(req.Username)

	if emailErr != nil {
		w.WriteHeader(400)

		emsg := "Invalid email address"

		writeResp(w, &signupResponse{
			Error: &emsg,
		})

		return
	}

	// Lock usersMux mutex

	usersMux.Lock()
	defer usersMux.Unlock()

	// Check if username taken

	userWithSameUsername := getUserByUsername(req.Username)

	if userWithSameUsername != nil {
		w.WriteHeader(400)

		emsg := "Email already has an account"

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
		username:      req.Username,
		name:          req.Name,
		password:      req.Password,
		loc:           loc,
		notifications: []*notification{},
	}

	users = append(users, u)

	fmt.Println("Pushed user", u.username)

	// End signup, create session

	tokStr, err := genToken(64)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	sessions[tokStr] = u.username

	http.SetCookie(w, &http.Cookie{
		Name:     "sess",
		Value:    tokStr,
		Expires:  time.Now().Add(time.Hour * 24 * 4),
		SameSite: http.SameSiteNoneMode,
	})

	w.WriteHeader(200)

	writeResp(w, &loginResponse{
		Error: nil,
	})
}
