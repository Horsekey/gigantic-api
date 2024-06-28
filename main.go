package main

import (
		"log"
		"net/http"
		"encoding/json"
		"fmt"

		"github.com/gorilla/mux"
)

type Hero struct {
	ID	 string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Role string `json:"role"`
  Archetype string `json:"archetype"`
}

// Init heroes var as a slice of Hero 
var heroes []Hero

func getheroes(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v", heroes)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(heroes)
}

func getHeroByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range heroes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func main() {

	heroes = append(heroes, Hero{ID: "1", Name: "Aising", Description: "By taking up her father's sword, Aisling conjured more than just his memory", "Role": ["Support", "Melee DPS"], "Archetype": ["Summoner", "Utility"]})

	router := mux.NewRouter()

	// get hero information based on name or id
	router.HandleFunc("/sushi", getheroes).Methods("GET")
	router.HandleFunc("/heroes/{name}", getHeroByName).Methods("GET")
	router.HandleFunc("/heroes/{id}", getHeroById).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":5000", router))
}
