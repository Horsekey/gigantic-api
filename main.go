package main

import (
		"net/http"
		"encoding/json"
		"fmt"
		"log"
		"github.com/gorilla/mux"
)

	type Test struct {
		Heroes []struct {
			ID              int      `json:"id"`
			Name            string   `json:"name"`
			Title           string   `json:"title"`
			Description     string   `json:"description"`
			Role            []string `json:"role"`
			Archetype       []string `json:"archetype"`
		}
	}

// Init heroes var as a slice of Hero 
var heroes Test

func getheroes(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v", heroes)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(heroes)
}

/*func getHeroByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range heroes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}*/

func main() {

	b := []byte(`{"heroes":[{"ID": 1, "Name": "Aising", "Description": "By taking up her father's sword, Aisling conjured more than just his memory", "Role": ["Support", "Melee DPS"], "Archetype": ["Summoner", "Utility"]},{"ID": 2, "Name": "Beckett", "Description": "An adventurer needs quick reflexes and quicker wits. A jetpack doesn't hurt either.", "Role": ["Ranged DPS"], "Archetype": ["Shooter"]}]}`)

	err := json.Unmarshal(b, &heroes)
	
	if err != nil {
		fmt.Printf("can't Unmarshal due to %s", err)
	}

	fmt.Println(heroes)

	router := mux.NewRouter()

	// get hero information based on name or id
	router.HandleFunc("/heroes", getheroes).Methods("GET")
	//router.HandleFunc("/heroes/{name}", getHeroByName).Methods("GET")
	//router.HandleFunc("/heroes/{id}", getHeroById).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":5000", router))
}
