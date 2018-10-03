package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
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
var resultListTpp []ResultItem
var resultListFpp []ResultItem

const dbName = "pgi.txt"
const urlServer = "http://127.0.0.1:8000"

/////////////////////////////////////////////////////////////
//Helper function
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

/////////////////////////////////////////////////////////////

//1. Load database
func LoadTeams() {
	var url = urlServer + "/teams"

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
	}

	err1 := json.Unmarshal([]byte(data), &listOfTeam)

	if err1 != nil {
		fmt.Println(err)
	}
}

func LoadPointDistributions() {
	var url = urlServer + "/points"

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
	}

	err1 := json.Unmarshal([]byte(data), &listOfPointDistribution)

	if err1 != nil {
		fmt.Println(err)
	}

}

func LoadMatches() {
	//Load TPP matches
	var urlTpp = urlServer + "/matches/tpp"
	response, err := http.Get(urlTpp)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
	}

	err1 := json.Unmarshal([]byte(data), &listOfMatchesTpp)

	if err1 != nil {
		fmt.Println(err)
	}

	//Load FPP matches
	var urlFpp = urlServer + "/matches/fpp"
	response1, err1 := http.Get(urlFpp)

	if err1 != nil {
		fmt.Println(err)
	}

	defer response1.Body.Close()
	data, err2 := ioutil.ReadAll(response1.Body)

	if err2 != nil {
		fmt.Println(err)
	}

	err3 := json.Unmarshal([]byte(data), &listOfMatchesFpp)

	if err3 != nil {
		fmt.Println(err)
	}
}

func LoadFromRestAPI() {
	LoadTeams()
	LoadPointDistributions()
	LoadMatches()
}

func LoadData() {
	LoadFromRestAPI()
}

//2. Calculate table
func CalculatePGI() {
	resultListTpp = make([]ResultItem, 0)

	//Filter team result TPP
	for _, team := range listOfTeam {
		resultItem := ResultItem{team, make([]MatchDetails, 0), 0}

		//Each team we get all match details of each
		for _, item := range listOfMatchesTpp {
			//Add detail of specific match
			for _, detail := range item.ListOfMatchDetails {
				if detail.Team.Id == team.Id {
					resultItem.ListOfMatchDetails = append(resultItem.ListOfMatchDetails, detail)
					resultItem.TotalPoint += detail.Detail.Point + detail.Kill*15
				}
			}
		}
		resultListTpp = append(resultListTpp, resultItem)
	}

	for _, team := range listOfTeam {
		resultItem := ResultItem{team, make([]MatchDetails, 0), 0}

		//Each team we get all match details of each
		for _, item := range listOfMatchesFpp {
			//Add detail of specific match
			for _, detail := range item.ListOfMatchDetails {
				if detail.Team.Id == team.Id {
					resultItem.ListOfMatchDetails = append(resultItem.ListOfMatchDetails, detail)
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

	for _, item := range match.ListOfMatchDetails {
		result += item.Team.Name + ", " +
			item.Team.Regional + ", " +
			strconv.Itoa(item.Kill) + ", " +
			strconv.Itoa(item.TotalPoint) + "\n"
	}

	return result
}

func DisplayResultItems(list []ResultItem) string {
	var result = ""
	var start = 1
	for _, item := range list {
		result += strconv.Itoa(start) + ". " + item.Team.Name + "\t\t"
		for _, data := range item.ListOfMatchDetails {
			result += "[" + strconv.Itoa(data.Detail.Placement) + ", " + strconv.Itoa(data.Kill) + "]" + " - "

		}
		result += strconv.Itoa(item.TotalPoint) + "\n"

		start++
	}
	return result
}

func DisplayAllTeams() {
	if listOfTeam != nil {
		fmt.Println("Total team: ", len(listOfTeam))

		for _, item := range listOfTeam {
			fmt.Println(item.Id, " - ", item.Name, " - ", item.Regional)
		}

		fmt.Print("\n")
	}
}

func DisplayPointDistribution() {
	if listOfPointDistribution != nil {
		fmt.Println("Total point distribution: ", len(listOfPointDistribution))

		for _, item := range listOfPointDistribution {
			fmt.Println(item.Placement, " - ", item.Point)
		}
		fmt.Print("\n")
	}
}

func DisplayMatches() {
	if listOfMatchesTpp != nil {
		fmt.Println("Total matches: ", len(listOfMatchesTpp))

		for _, item := range listOfMatchesTpp {
			fmt.Print(DisplayMatch(item))
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

func main() {
	LoadData()
	CalculatePGI()
	DisplayAll()
}
