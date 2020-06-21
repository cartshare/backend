package main

import (
	"net/http"
)

type deleteNotificationRequest struct {
	ID string `json:"notificationId"`
}

func deleteNotificationHandler(u *user, w http.ResponseWriter, r *http.Request) {
	req := &deleteNotificationRequest{}
	err := parseReq(r, req)

	if err != nil {
		w.WriteHeader(400)

		return
	}

	for i, n := range u.notifications {
		if n.ID == req.ID {
			u.notifications = append(u.notifications[:i], u.notifications[i+1:]...)
		}
	}

	w.WriteHeader(200)
	writeResp(w, map[string]interface{}{
		"error": nil,
	})
}
