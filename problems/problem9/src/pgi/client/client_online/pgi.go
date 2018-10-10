package main

import (
	"encoding/json"
	"fmt"
	"github.com/dqduy/pgi"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
)

/////////////////////////////////////////////////////////////
//Global var
const urlServer = "http://127.0.0.1:8000"

var resultListTpp []pgi.ResultItem
var resultListFpp []pgi.ResultItem

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

	err1 := json.Unmarshal([]byte(data), &pgi.ListOfTeam)

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

	err1 := json.Unmarshal([]byte(data), &pgi.ListOfPointDistribution)

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

	err1 := json.Unmarshal([]byte(data), &pgi.ListOfMatchesTpp)

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

	err3 := json.Unmarshal([]byte(data), &pgi.ListOfMatchesFpp)

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
	resultListTpp = make([]pgi.ResultItem, 0)

	//Filter team result TPP
	for _, team := range pgi.ListOfTeam {
		resultItem := pgi.ResultItem{team, make([]pgi.MatchDetails, 0), 0}

		//Each team we get all match details of each
		for _, item := range pgi.ListOfMatchesTpp {
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

	//Filter team result FPP
	for _, team := range pgi.ListOfTeam {
		resultItem := pgi.ResultItem{team, make([]pgi.MatchDetails, 0), 0}

		//Each team we get all match details of each
		for _, item := range pgi.ListOfMatchesFpp {
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
	sort.Sort(pgi.ByPoint(resultListTpp))
	sort.Sort(pgi.ByPoint(resultListFpp))
}

//3. Display section
func DisplayResultItems(list []pgi.ResultItem) string {
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

func DisplayAll() {
	//Print result items
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
