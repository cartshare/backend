package main

import (
	"net/http"
)

type completeItemRequest struct {
	ID string `json:"itemId"`
}

func completeItemHandler(u *user, w http.ResponseWriter, r *http.Request) {
	req := &completeItemRequest{}
	err := parseReq(r, req)

	if err != nil {
		w.WriteHeader(400)

		return
	}

	// Find and delete item if exists

	for i, item := range items {
		if item.ID == req.ID {
			if item.owner != u && (userDist(item.owner, u) > 0.5 || !item.OnWishlist) {
				// If requesting user does not have access to item (either because not owner
				// and item not wishlisted or because not owner and too distant from owner),
				// they will be denied access.

				w.WriteHeader(401)
				writeResp(w, map[string]string{
					"error": "Permission denied",
				})

				return
			}

			// Delete item

			items = append(items[:i], items[i+1:]...)

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
