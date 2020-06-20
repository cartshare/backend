package main

import (
	"encoding/json"
	"net/http"
)

func authWrap(o func(*user, http.ResponseWriter, *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authCookie, err := r.Cookie("sess")

		endError := func(e string) {
			ser, err := json.Marshal(map[string]string{"error": e})

			if err != nil {
				w.Write([]byte(e))
			}

			w.Write([]byte(ser))
		}

		if err != nil {
			w.WriteHeader(401)

			errText := err.Error()

			if err == http.ErrNoCookie {
				errText = "Missing auth cookie"
			}

			endError(errText)

			return
		}

		uname, ok := sessions[authCookie.Value]

		if !ok {
			w.WriteHeader(401)

			endError("Session could not be located")

			return
		}

		var relevantUser *user = nil

		for _, u := range users {
			if u.username == uname {
				relevantUser = u

				break
			}
		}

		if relevantUser == nil {
			w.WriteHeader(404)
		}

		o(relevantUser, w, r)
	}
}
