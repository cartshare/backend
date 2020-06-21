package main

import (
	"encoding/json"
	"net/http"
	"time"
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

func pushNotification(u *user, n *notification) error {
	n.Created = time.Now()
	id, err := genToken(16)

	if err != nil {
		return err
	}

	n.ID = id

	u.notifications = append(u.notifications, n)

	return nil
}
