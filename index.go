package main

import (
	geo "github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/opencage"
	"github.com/gorilla/handlers"
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

type item struct {
	ID         string `json:"id"`
	Desc       string `json:"Desc"`
	Qty        int    `json:"Qty"`
	OnWishlist bool   `json:"onWishlist"`
	owner      *user
}

// Note: Mutexes should be used for this data, or even better, a database, in prod

var users []*user = []*user{}
var sessions map[string]string = map[string]string{}

var items []*item = []*item{}

var geocoder geo.Geocoder

func main() {
	// Init geocoder

	geocoder = opencage.Geocoder(os.Getenv("OPENCAGE_KEY"))

	// Init endpoints

	r := mux.NewRouter()

	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/signup", signupHandler)

	r.HandleFunc("/list", authWrap(listHandler))
	r.HandleFunc("/createItem", authWrap(createItemHandler))

	r.HandleFunc("/completeItem", authWrap(completeItemHandler))
	r.HandleFunc("/setItemWishlisted", authWrap(setItemWishlistedHandler))

	/*r.HandleFunc("/neighborList", neighborListHandler)*/

	// Note: CORS allows all origins with current configuration. Do not use this configuration in production.

	err := http.ListenAndServe(":80", handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(r))

	if err != nil {
		panic(err)
	}
}
