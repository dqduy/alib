package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Team struct {
	Id       int
	Name     string
	Regional string
}

var listOfTeam []Team

func InitServer() {
	listOfTeam = append(listOfTeam, Team{9, "Gen.G Gold", "Korean Qualifier你好吗"})
	listOfTeam = append(listOfTeam, Team{2, "Welcome to South Georgo", "European Qualifier"})
	listOfTeam = append(listOfTeam, Team{1, "Team Liquid", "European Qualifier"})
}

func GetTeams(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(listOfTeam)
}

func GetTeam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range listOfTeam {
		if strconv.Itoa(item.Id) == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	json.NewEncoder(w).Encode(Team{})

}

func main() {
	InitServer()

	fmt.Println("Hello problem 10")
	router := mux.NewRouter()
	router.HandleFunc("/team", GetTeams).Methods("GET")
	router.HandleFunc("/team/{id}", GetTeam).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
