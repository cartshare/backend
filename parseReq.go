package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func parseReq(r *http.Request, dest interface{}) error {
	all, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(all, dest)

	if err != nil {
		return err
	}

	return nil
}
