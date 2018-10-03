package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//Data model for PGI
type Team struct {
	Id       int
	Name     string
	Regional string
}

type PointDistribution struct {
	Placement int
	Point     int
}

type Match struct {
	Id                 int
	MatchName          string
	MapName            string
	ListOfMatchDetails []MatchDetails
}

/////////////////////////////////////////////////////////////
type MatchDetails struct {
	Team       Team
	Detail     PointDistribution
	Kill       int
	TotalPoint int
}

func (dt *MatchDetails) caculatePoint() {
	dt.TotalPoint = dt.Detail.Point + dt.Kill*15
}

/////////////////////////////////////////////////////////////

type ResultItem struct {
	Team               Team
	ListOfMatchDetails []MatchDetails
	TotalPoint         int
}

type ByPoint []ResultItem

func (this ByPoint) Len() int {
	return len(this)
}

func (this ByPoint) Less(i, j int) bool {
	return this[i].TotalPoint > this[j].TotalPoint
}

func (this ByPoint) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

/////////////////////////////////////////////////////////////

//Global data
var listOfTeam []Team = make([]Team, 0)
var listOfPointDistribution []PointDistribution = make([]PointDistribution, 0)
var listOfMatchesTpp []Match = make([]Match, 0)
var listOfMatchesFpp []Match = make([]Match, 0)

const dbName = "pgi.txt"

func MakeTeam(id int, name string, regional string) Team {
	return Team{id, name, regional}
}

func MakePointDistribution(placement int, point int) PointDistribution {
	return PointDistribution{placement, point}
}

func SearchTeam(id int) Team {
	for _, item := range listOfTeam {
		if id == item.Id {
			return item
		}
	}

	return Team{0, "", ""}
}

func SearchPointDistribution(placement int) PointDistribution {
	for _, item := range listOfPointDistribution {
		if placement == item.Placement {
			return item
		}
	}

	return PointDistribution{0, 15}
}

//

func LoadData() {

	file, err := os.Open(dbName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var str = ""
	var section = 0 //1 (team),2 (point distribution),3 (tpp),4 (fpp)

	for scanner.Scan() {
		str = strings.TrimSpace(scanner.Text())

		if str == "team" {
			section = 1
		} else if str == "points distribution" {
			section = 2
		} else if str == "tpp" {
			section = 3
		} else if str == "fpp" {
			section = 4
		} else {
			if len(str) > 0 {
				switch section {
				case 1:
					var tokens = strings.Split(str, ",")
					id, err := strconv.Atoi(tokens[0])

					if err != nil {
						fmt.Println(err)
					}

					listOfTeam = append(listOfTeam, MakeTeam(id,
						strings.TrimSpace(tokens[1]),
						strings.TrimSpace(tokens[2])))

				case 2:
					var tokens = strings.Split(str, ",")
					placement, err := strconv.Atoi(strings.TrimSpace(tokens[0]))
					point, err1 := strconv.Atoi(strings.TrimSpace(tokens[1]))

					if err != nil || err1 != nil {
						fmt.Println(err)
					}

					listOfPointDistribution = append(listOfPointDistribution, MakePointDistribution(placement, point))

				case 3:
					var tokens = strings.Split(str, ",")
					var tokensLen = len(tokens)
					//fmt.Println(tokensLen)

					id, err := strconv.Atoi(strings.TrimSpace(tokens[0]))
					if err != nil {
						fmt.Println(err)
					}

					match := Match{id, tokens[1], tokens[2], make([]MatchDetails, 0)}

					for i := 0; i < (tokensLen-3)/3; i++ {
						teamid, err := strconv.Atoi(strings.TrimSpace(tokens[3+(i*3)]))
						placement, err := strconv.Atoi(strings.TrimSpace(tokens[3+(i*3)+1]))
						kill, err := strconv.Atoi(strings.TrimSpace(tokens[3+(i*3)+2]))
						if err != nil {
							fmt.Println(err)
						}
						matchdetail := MatchDetails{SearchTeam(teamid), SearchPointDistribution(placement), kill, 0}
						matchdetail.caculatePoint()

						match.ListOfMatchDetails = append(match.ListOfMatchDetails, matchdetail)
					}

					listOfMatchesTpp = append(listOfMatchesTpp, match)

				case 4:
					var tokens = strings.Split(str, ",")
					var tokensLen = len(tokens)
					//fmt.Println(tokensLen)

					id, err := strconv.Atoi(strings.TrimSpace(tokens[0]))
					if err != nil {
						fmt.Println(err)
					}

					match := Match{id, tokens[1], tokens[2], make([]MatchDetails, 0)}

					for i := 0; i < (tokensLen-3)/3; i++ {
						teamid, err := strconv.Atoi(strings.TrimSpace(tokens[3+(i*3)]))
						placement, err := strconv.Atoi(strings.TrimSpace(tokens[3+(i*3)+1]))
						kill, err := strconv.Atoi(strings.TrimSpace(tokens[3+(i*3)+2]))
						if err != nil {
							fmt.Println(err)
						}
						matchdetail := MatchDetails{SearchTeam(teamid), SearchPointDistribution(placement), kill, 0}
						matchdetail.caculatePoint()

						match.ListOfMatchDetails = append(match.ListOfMatchDetails, matchdetail)
					}

					listOfMatchesFpp = append(listOfMatchesFpp, match)

				default:
					fmt.Printf("Unknown section")
				}
			} else {
				//fmt.Println("We have a blank line....")
			}
		}
	}
}

func InitSettings() {

}

func InitServer() {
	InitSettings()
	LoadData()
}

///////////////////////////////////////////////////////////////////
//List of API
/*
/teams				-> Get all teams in tournament
/teams/{id}			-> Get a team in all team
/points				-> Get all point distributions
/points/{placement}	-> Get a point distribution
/matches			-> Get all match with details
/matches/{id}		-> Get a detail match
*/

//1. Get all teams
func GetTeams(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request url: ", r.RequestURI)
	json.NewEncoder(w).Encode(listOfTeam)
}

//2. Get a teams
func GetTeam(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request url: ", r.RequestURI)
	params := mux.Vars(r)

	for _, item := range listOfTeam {
		if strconv.Itoa(item.Id) == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(Team{})

}

//3. Get all point distributions
func GetPoints(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request url: ", r.RequestURI)
	json.NewEncoder(w).Encode(listOfPointDistribution)
}

//4. Get a point distribution
func GetPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request url: ", r.RequestURI)
}

//5. Get all matches
func GetMatches(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request url: ", r.RequestURI)
	params := mux.Vars(r)

	if params["type"] == "tpp" {
		json.NewEncoder(w).Encode(listOfMatchesTpp)
	} else if params["type"] == "fpp" {
		json.NewEncoder(w).Encode(listOfMatchesFpp)
	}

}

//6. Get a match
func GetMatch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request url: ", r.RequestURI)
	params := mux.Vars(r)

	if params["type"] == "tpp" {
		for _, item := range listOfMatchesTpp {
			if strconv.Itoa(item.Id) == params["id"] {
				json.NewEncoder(w).Encode(item)
				break
			}
		}
	} else if params["type"] == "fpp" {
		for _, item := range listOfMatchesFpp {
			if strconv.Itoa(item.Id) == params["id"] {
				json.NewEncoder(w).Encode(item)
				break
			}
		}
	}
}

func main() {
	InitServer()

	fmt.Println("Starting server....")
	fmt.Println("Running at port 8000")
	router := mux.NewRouter()
	router.HandleFunc("/teams", GetTeams).Methods("GET")
	router.HandleFunc("/teams/{id}", GetTeam).Methods("GET")
	router.HandleFunc("/points", GetPoints).Methods("GET")
	router.HandleFunc("/points/{placement}", GetPoint).Methods("GET")
	router.HandleFunc("/matches/{type}", GetMatches).Methods("GET")
	router.HandleFunc("/matches/{type}/{id}", GetMatch).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
