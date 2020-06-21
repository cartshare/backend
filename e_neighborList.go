package main

import (
	"net/http"
)

type neighborItem struct {
	ID        string `json:"id"`
	Desc      string `json:"desc"`
	Qty       int    `json:"qty"`
	OwnerName string `json:"owner"`
}

func neighborListHandler(u *user, w http.ResponseWriter, r *http.Request) {
	neighborItems := []*neighborItem{}

	for _, item := range items {
		if item.owner != u && userDist(u, item.owner) < 0.5 && item.OnWishlist {
			neighborItems = append(neighborItems, &neighborItem{
				ID:        item.ID,
				Desc:      item.Desc,
				Qty:       item.Qty,
				OwnerName: item.owner.name,
			})
		}
	}

	writeResp(w, map[string]interface{}{
		"neighborRequests": neighborItems,
	})
}
