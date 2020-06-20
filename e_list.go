package main

import (
	"net/http"
)

func listHandler(u *user, w http.ResponseWriter, r *http.Request) {
	userItems := []*item{}

	for _, i := range items {
		if i.owner == u {
			userItems = append(userItems, i)
		}
	}

	writeResp(w, map[string]interface{}{
		"items": userItems,
	})
}
