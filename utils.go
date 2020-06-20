package main

import (
	"encoding/json"
	"net/http"
)

func getUserByUsername(un string) *user {
	for _, u := range users {
		if u.username == un {
			return u
		}
	}

	return nil
}

func writeResp(d http.ResponseWriter, r interface{}) error {
	m, err := json.Marshal(r)

	if err != nil {
		return err
	}

	d.Write(m)

	return nil
}
