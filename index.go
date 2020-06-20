package main

import (
	geo "github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/opencage"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Note: Unhashed password not to be saved to disk.
// Unhashed password is used in memory for demonstration.
type user struct {
	username string
	name     string
	password string
	loc      *geo.Location
}

var users []*user = []*user{}
var sessions map[string]string = map[string]string{}

var geocoder geo.Geocoder

func main() {
	// Init geocoder

	geocoder = opencage.Geocoder(os.Getenv("OPENCAGE_KEY"))

	// Init endpoints

	r := mux.NewRouter()

	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/signup", signupHandler)

	/*r.HandleFunc("/list", listHandler)
	r.HandleFunc("/createItem", createItemHandler)

	r.HandleFunc("/completeItem", completeItemHandler)
	r.HandleFunc("/setItemWishlisted", setItemWishlistedHandler)

	r.HandleFunc("/neighborList", neighborListHandler)*/

	err := http.ListenAndServe(":80", r)

	if err != nil {
		panic(err)
	}
}
