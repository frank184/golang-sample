package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const ProfileCount = 100000

type Profile struct {
	Name    string
	Hobbies []string
}

func buildProfiles(n int) []Profile {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	profiles := []Profile{}
	for i := 0; i < n; i++ {
		profiles = append(profiles, profile)
	}
	return profiles
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := buildProfiles(ProfileCount)
	json.NewEncoder(w).Encode(p)
	w.Header().Set("Server", "A Go JSON Server")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Status", "200")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
