package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Spot struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// making mock/ dummy data.
var spots = []Spot{
	{1, "Spot 1", 37.7749, -122.4194},
	{2, "Spot 2", 37.7749, -122.4194},
	{3, "Spot 3", 37.7895, -122.4023},
	{4, "Spot 4", 37.7895, -122.4023},
	{5, "Spot 5", 37.7739, -122.4312},
}

func main() {
	http.HandleFunc("/spots", getSpotsInArea)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getSpotsInArea(w http.ResponseWriter, r *http.Request) {
	latitudeStr := r.URL.Query().Get("latitude")
	longitudeStr := r.URL.Query().Get("longitude")
	radiusStr := r.URL.Query().Get("radius")

	// Convert latitude and longitude to float64
	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		return
	}

	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		http.Error(w, "Invalid longitude", http.StatusBadRequest)
		return
	}

	// Convert radius to float64
	radius, err := strconv.ParseFloat(radiusStr, 64)
	if err != nil {
		http.Error(w, "Invalid radius", http.StatusBadRequest)
		return
	}

	// Calculate the maximum latitude and longitude within the given radius
	latDiff := (radius / 6371000) * (180 / math.Pi)
	lonDiff := (radius / 6371000) * (180 / math.Pi) / math.Cos(latitude*math.Pi/180)

	maxLatitude := latitude + latDiff
	minLatitude := latitude - latDiff
	maxLongitude := longitude + lonDiff
	minLongitude := longitude - lonDiff

	// Filter spots within the area
	filteredSpots := filterSpotsByArea(minLatitude, maxLatitude, minLongitude, maxLongitude)

	// Return the filtered spots as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredSpots)
}

func filterSpotsByArea(minLat, maxLat, minLon, maxLon float64) []Spot {
	var filtered []Spot

	for _, spot := range spots {
		if spot.Latitude >= minLat && spot.Latitude <= maxLat &&
			spot.Longitude >= minLon && spot.Longitude <= maxLon {
			filtered = append(filtered, spot)
		}
	}

	return filtered
}
