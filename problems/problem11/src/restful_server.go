package main

import (
	"bufio"
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"sort"
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
	ListOfMatchDetails *list.List
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
	ListOfMatchDetails *list.List
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
var listOfTeam *list.List = list.New()
var listOfPointDistribution *list.List = list.New()
var listOfMatchesTpp *list.List = list.New()
var listOfMatchesFpp *list.List = list.New()
var resultListTpp []ResultItem
var resultListFpp []ResultItem

const dbName = "pgi.txt"

func MakeTeam(id int, name string, regional string) Team {
	return Team{id, name, regional}
}

func MakePointDistribution(placement int, point int) PointDistribution {
	return PointDistribution{placement, point}
}

func SearchTeam(id int) Team {
	for index := listOfTeam.Front(); index != nil; index = index.Next() {
		if id == index.Value.(Team).Id {
			return index.Value.(Team)
		}
	}

	return Team{0, "", ""}
}

func SearchPointDistribution(placement int) PointDistribution {
	for index := listOfPointDistribution.Front(); index != nil; index = index.Next() {
		if placement == index.Value.(PointDistribution).Placement {
			return index.Value.(PointDistribution)
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

					listOfTeam.PushBack(MakeTeam(id,
						strings.TrimSpace(tokens[1]),
						strings.TrimSpace(tokens[2])))

				case 2:
					var tokens = strings.Split(str, ",")
					placement, err := strconv.Atoi(strings.TrimSpace(tokens[0]))
					point, err1 := strconv.Atoi(strings.TrimSpace(tokens[1]))

					if err != nil || err1 != nil {
						fmt.Println(err)
					}

					listOfPointDistribution.PushBack(MakePointDistribution(placement, point))

				case 3:
					var tokens = strings.Split(str, ",")
					var tokensLen = len(tokens)
					//fmt.Println(tokensLen)

					id, err := strconv.Atoi(strings.TrimSpace(tokens[0]))
					if err != nil {
						fmt.Println(err)
					}

					match := Match{id, tokens[1], tokens[2], list.New()}

					for i := 0; i < (tokensLen-3)/3; i++ {
						teamid, err := strconv.Atoi(strings.TrimSpace(tokens[3+(i*3)]))
						placement, err := strconv.Atoi(strings.TrimSpace(tokens[3+(i*3)+1]))
						kill, err := strconv.Atoi(strings.TrimSpace(tokens[3+(i*3)+2]))
						if err != nil {
							fmt.Println(err)
						}
						matchdetail := MatchDetails{SearchTeam(teamid), SearchPointDistribution(placement), kill, 0}
						matchdetail.caculatePoint()

						match.ListOfMatchDetails.PushBack(matchdetail)
					}

					listOfMatchesTpp.PushBack(match)

				case 4:
					var tokens = strings.Split(str, ",")
					var tokensLen = len(tokens)
					//fmt.Println(tokensLen)

					id, err := strconv.Atoi(strings.TrimSpace(tokens[0]))
					if err != nil {
						fmt.Println(err)
					}

					match := Match{id, tokens[1], tokens[2], list.New()}

					for i := 0; i < (tokensLen-3)/3; i++ {
						teamid, err := strconv.Atoi(strings.TrimSpace(tokens[3+(i*3)]))
						placement, err := strconv.Atoi(strings.TrimSpace(tokens[3+(i*3)+1]))
						kill, err := strconv.Atoi(strings.TrimSpace(tokens[3+(i*3)+2]))
						if err != nil {
							fmt.Println(err)
						}
						matchdetail := MatchDetails{SearchTeam(teamid), SearchPointDistribution(placement), kill, 0}
						matchdetail.caculatePoint()

						match.ListOfMatchDetails.PushBack(matchdetail)
					}

					listOfMatchesFpp.PushBack(match)

				default:
					fmt.Printf("Unknown section")
				}
			} else {
				//fmt.Println("We have a blank line....")
			}
		}
	}
}

//2. Calculate table
func CalculatePGI() {
	resultListTpp = make([]ResultItem, 0)

	//Filter team result TPP
	for item := listOfTeam.Front(); item != nil; item = item.Next() {
		var data = item.Value.(Team)
		resultItem := ResultItem{data, list.New(), 0}

		//Each team we get all match details of each
		for inner := listOfMatchesTpp.Front(); inner != nil; inner = inner.Next() {
			var list = inner.Value.(Match).ListOfMatchDetails
			//Add detail of specific match
			for i := list.Front(); i != nil; i = i.Next() {
				var detail = i.Value.(MatchDetails)
				if detail.Team.Id == data.Id {
					resultItem.ListOfMatchDetails.PushBack(detail)
					resultItem.TotalPoint += detail.Detail.Point + detail.Kill*15
				}
			}
		}
		resultListTpp = append(resultListTpp, resultItem)
	}

	for item := listOfTeam.Front(); item != nil; item = item.Next() {
		var data = item.Value.(Team)
		resultItem := ResultItem{data, list.New(), 0}

		//Each team we get all match details of each
		for inner := listOfMatchesFpp.Front(); inner != nil; inner = inner.Next() {
			var list = inner.Value.(Match).ListOfMatchDetails
			//Add detail of specific match
			for i := list.Front(); i != nil; i = i.Next() {
				var detail = i.Value.(MatchDetails)
				if detail.Team.Id == data.Id {
					resultItem.ListOfMatchDetails.PushBack(detail)
					resultItem.TotalPoint += detail.Detail.Point + detail.Kill*15
				}
			}
		}
		resultListFpp = append(resultListFpp, resultItem)
	}

	//Sort result list
	//1. Point
	//2. if two point's team is equal, we compare kill
	sort.Sort(ByPoint(resultListTpp))
	sort.Sort(ByPoint(resultListFpp))
}

//3. Display section
func DisplayMatch(match Match) string {
	var result = ""

	result = "ID:         " + strconv.Itoa(match.Id) + "\n" +
		"Match Name: " + match.MatchName + "\n" +
		"Map:        " + match.MapName

	result += "\n\tDetails: \n"

	for inner := match.ListOfMatchDetails.Front(); inner != nil; inner = inner.Next() {
		result += inner.Value.(MatchDetails).Team.Name + ", " +
			inner.Value.(MatchDetails).Team.Regional + ", " +
			strconv.Itoa(inner.Value.(MatchDetails).Kill) + ", " +
			strconv.Itoa(inner.Value.(MatchDetails).TotalPoint) + "\n"
	}

	return result
}

func DisplayResultItems(list []ResultItem) string {
	var result = ""
	var start = 1
	for _, item := range list {
		result += strconv.Itoa(start) + ". " + item.Team.Name + "\t\t"
		for inner := item.ListOfMatchDetails.Front(); inner != nil; inner = inner.Next() {
			var data = inner.Value.(MatchDetails)
			result += "[" + strconv.Itoa(data.Detail.Placement) + ", " + strconv.Itoa(data.Kill) + "]" + " - "

		}
		result += strconv.Itoa(item.TotalPoint) + "\n"

		start++
	}
	return result
}

func DisplayAllTeams() {
	if listOfTeam != nil {
		fmt.Println("Total team: ", listOfTeam.Len())

		for index := listOfTeam.Front(); index != nil; index = index.Next() {
			fmt.Println(index.Value.(Team).Id, " - ", index.Value.(Team).Name, " - ", index.Value.(Team).Regional)
		}

		fmt.Print("\n")
	}
}

func DisplayPointDistribution() {
	if listOfPointDistribution != nil {
		fmt.Println("Total point distribution: ", listOfPointDistribution.Len())

		for index := listOfPointDistribution.Front(); index != nil; index = index.Next() {
			fmt.Println(index.Value.(PointDistribution).Placement, " - ", index.Value.(PointDistribution).Point)
		}
		fmt.Print("\n")
	}
}

func DisplayMatches() {
	if listOfMatchesTpp != nil {
		fmt.Println("Total matches: ", listOfMatchesTpp.Len())

		for index := listOfMatchesTpp.Front(); index != nil; index = index.Next() {
			fmt.Print(DisplayMatch(index.Value.(Match)))
			fmt.Print("\n")
		}
	}
}

func DisplayAll() {
	//1. Print teams
	//DisplayAllTeams()

	//2. Print point distributions
	//DisplayPointDistribution()

	//3. Print matches
	//DisplayMatches()

	//4. Print result items
	fmt.Println("PUBG Global Invitational 2018")
	fmt.Println("=============================\n")
	fmt.Println("TPP:")
	fmt.Println(DisplayResultItems(resultListTpp))
	fmt.Println("FPP:")
	fmt.Println(DisplayResultItems(resultListFpp))
}

func InitServer() {
	InitSettings()
	LoadData()
	CalculatePGI()
	DisplayAll()
}

func InitSettings() {

}

func GetTeams(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(listOfTeam)
}

func GetTeam(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)

	// for _, item := range listOfTeam {
	// 	if strconv.Itoa(item.Id) == params["id"] {
	// 		json.NewEncoder(w).Encode(item)
	// 		return
	// 	}
	// }
	// json.NewEncoder(w).Encode(Team{})

}

func main() {
	InitServer()

	fmt.Println("Hello problem 10")
	router := mux.NewRouter()
	router.HandleFunc("/teams", GetTeams).Methods("GET")
	router.HandleFunc("/teams/{id}", GetTeam).Methods("GET")
	// router.HandleFunc("/points", GetPoints).Methods("GET")
	// router.HandleFunc("/points/{placement}", GetPoint).Methods("GET")
	// router.HandleFunc("/matches/{id}", GetMatches).Methods("GET")
	// router.HandleFunc("/matches", GetMatch).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
