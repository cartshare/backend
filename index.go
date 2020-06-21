package main

import (
	"fmt"
	geo "github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/opencage"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type notification struct {
	ID      string    `json:"id"`
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
}

// Note: Unhashed password not to be saved to disk.
// Unhashed password is used in memory for demonstration.
type user struct {
	username      string
	name          string
	password      string
	loc           *geo.Location
	notifications []*notification
}

type item struct {
	ID         string `json:"id"`
	Desc       string `json:"desc"`
	Qty        int    `json:"qty"`
	OnWishlist bool   `json:"onWishlist"`
	owner      *user
}

// Note: Mutexes should be used for this data, or even better, a database, in prod

var usersMux sync.Mutex = sync.Mutex{}
var users []*user = []*user{}
var sessions map[string]string = map[string]string{}

var items []*item = []*item{}

var geocoder geo.Geocoder

func main() {
	fmt.Println("Preparing...")

	// Init geocoder

	geocoder = opencage.Geocoder(os.Getenv("OPENCAGE_KEY"))

	// Init endpoints

	r := mux.NewRouter()

	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/signup", signupHandler)
	r.HandleFunc("/logout", logoutHandler)

	r.HandleFunc("/list", authWrap(listHandler))
	r.HandleFunc("/createItem", authWrap(createItemHandler))

	r.HandleFunc("/completeItem", authWrap(completeItemHandler))
	r.HandleFunc("/setItemWishlisted", authWrap(setItemWishlistedHandler))

	r.HandleFunc("/session", authWrap(sessionHandler))

	r.HandleFunc("/neighborList", authWrap(neighborListHandler))

	r.HandleFunc("/notifications", authWrap(notificationsHandler))
	r.HandleFunc("/deleteNotification", authWrap(deleteNotificationHandler))

	// Note: CORS allows all origins with current configuration. Do not use this configuration in production:

	fmt.Println("Listening on :80...")

	err := http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost", "http://localhost:8080", "http://35.192.187.248", "https://35.192.187.248"}),
	)(r)))

	if err != nil {
		panic(err)
	}
}
