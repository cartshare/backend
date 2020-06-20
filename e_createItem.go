package main

import (
	"fmt"
	"net/http"
)

type createItemRequest struct {
	Desc string `json:"itemDesc"`
	Qty  int    `json:"itemQty"`
}

type createItemResponse struct {
	Error *string `json:"error"`
}

func createItemHandler(u *user, w http.ResponseWriter, r *http.Request) {
	req := &createItemRequest{}
	err := parseReq(r, req)

	if err != nil {
		w.WriteHeader(400)

		return
	}

	itemTok, err := genToken(16)

	if err != nil {
		w.WriteHeader(500)

		return
	}

	pushItem := item{
		ID:         itemTok,
		Desc:       req.Desc,
		Qty:        req.Qty,
		OnWishlist: false,
		owner:      u,
	}

	items = append(items, &pushItem)

	fmt.Println("Pushed item", pushItem)

	writeResp(w, &createItemResponse{
		Error: nil,
	})
}
