package main

import (
	"net/http"
	"time"
)

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "sess",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour * 24 * 4),
		// Set Expires to a time in the past
	})
}
