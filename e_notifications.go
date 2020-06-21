package main

import (
	"net/http"
)

func notificationsHandler(u *user, w http.ResponseWriter, r *http.Request) {
	writeResp(w, map[string]interface{}{
		"error":         nil,
		"notifications": u.notifications,
	})
}
