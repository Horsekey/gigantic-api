package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Hero struct {
	Heroes []struct {
		ID          int      `json:"id"`
		Name        string   `json:"name"`
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Role        []string `json:"role"`
		Archetype   []string `json:"archetype"`
	}
}

// Init heroes var as a slice of Hero
var hero Hero

func getHeroes(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v", hero)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hero)
}

func getHeroByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, hero := range hero.Heroes {

		fmt.Println(hero.Name)
		fmt.Println(params["name"])

		if strings.EqualFold(hero.Name, params["name"]) {
			w.Header().Set("Content-Type", "application/json")

			err := json.NewEncoder(w).Encode(hero)
			if err != nil {
				http.Error(w, "Failed to encode hero", http.StatusInternalServerError)
			}
			return
		}
	}
}

func getHeroById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid hero ID", http.StatusBadRequest)
	}

	for _, hero := range hero.Heroes {
		if hero.ID == id {
			w.Header().Set("Content-Type", "application/json")

			err := json.NewEncoder(w).Encode(hero)
			if err != nil {
				http.Error(w, "Failed to encode hero", http.StatusInternalServerError)
			}
			return
		}
	}
}

func main() {

	fileContent, err := os.ReadFile("gigantic-json.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(fileContent, &hero)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(hero)
	router := mux.NewRouter()

	// get hero information based on name or id
	router.HandleFunc("/heroes", getHeroes).Methods("GET")
	router.HandleFunc("/heroes/{name:[a-zA-Z]+}", getHeroByName).Methods("GET")
	router.HandleFunc("/heroes/{id:[0-9]+}", getHeroById).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", router))
}
