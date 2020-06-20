package main

import (
	"github.com/umahmood/haversine"
)

// userDist returns user distance in miles
func userDist(a, b *user) float64 {
	c1 := haversine.Coord{
		Lat: a.loc.Lat,
		Lon: a.loc.Lng,
	}

	c2 := haversine.Coord{
		Lat: b.loc.Lat,
		Lon: b.loc.Lng,
	}

	mi, _ := haversine.Distance(c1, c2)

	return mi
}

func getNeighbors(a *user) []*user {
	// Find neighbors within half a mile

	neighbors := []*user{}

	for _, u := range users {
		if userDist(a, u) <= 0.5 {
			neighbors = append(neighbors, u)
		}
	}

	return neighbors
}
