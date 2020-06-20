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
