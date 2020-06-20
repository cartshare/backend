package main

import (
	"net/http"
)

type setItemWishlistedRequest struct {
	ID         string `json:"itemId"`
	Wishlisted bool   `json:"wishlisted"`
}

func setItemWishlistedHandler(u *user, w http.ResponseWriter, r *http.Request) {
	req := &setItemWishlistedRequest{}
	err := parseReq(r, req)

	if err != nil {
		w.WriteHeader(400)

		return
	}

	// Find and delete item if exists

	for _, item := range items {
		if item.ID == req.ID {
			if item.owner != u {
				w.WriteHeader(401)
				writeResp(w, map[string]string{
					"error": "Permission denied",
				})

				return
			}

			// Edit wishlisted status of item

			item.OnWishlist = req.Wishlisted

			w.WriteHeader(200)
			writeResp(w, map[string]interface{}{
				"error": nil,
			})

			return
		}
	}

	w.WriteHeader(404)

	writeResp(w, map[string]interface{}{
		"error": "Item not found",
	})
}
