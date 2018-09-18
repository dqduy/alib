package main

import (
	"container/list"
	"fmt"
	"strconv"
)

//Data model for PGI
type Team struct {
	id       int
	name     string
	regional string
}

type PointDistribution struct {
	placement int
	point     int
}

type Match struct {
	id                 int
	matchName          string
	mapName            string
	listOfMatchDetails *list.List
}

/////////////////////////////////////////////////////////////
type MatchDetails struct {
	team       Team
	detail     PointDistribution
	kill       int
	totalPoint int
}

func (dt *MatchDetails) caculatePoint() {
	dt.totalPoint = dt.detail.point + dt.kill*15
}

/////////////////////////////////////////////////////////////

type ResultItem struct {
	team               Team
	listOfMatchDetails *list.List
	totalPoint         int
}

//Global data
var listOfTeam *list.List
var listOfPointDistribution *list.List
var listOfMatches *list.List
var resultList *list.List

//

//Fake data
func CreateFakeData() {
	//1. Create teams
	listOfTeam = list.New()
	team1 := Team{1, "Team Liquid", "European Qualifier"}
	team2 := Team{2, "Welcome to South Georgo", "European Qualifier"}
	team3 := Team{9, "Gen.G Gold", "Korean Qualifier"}

	listOfTeam.PushBack(team1)
	listOfTeam.PushBack(team2)
	listOfTeam.PushBack(team3)

	//2. Create point distributions
	listOfPointDistribution = list.New()
	listOfPointDistribution.PushBack(PointDistribution{0, 15})
	listOfPointDistribution.PushBack(PointDistribution{1, 500})
	listOfPointDistribution.PushBack(PointDistribution{2, 410})
	listOfPointDistribution.PushBack(PointDistribution{3, 345})
	listOfPointDistribution.PushBack(PointDistribution{4, 295})
	listOfPointDistribution.PushBack(PointDistribution{5, 250})
	listOfPointDistribution.PushBack(PointDistribution{6, 210})
	listOfPointDistribution.PushBack(PointDistribution{7, 175})
	listOfPointDistribution.PushBack(PointDistribution{8, 145})
	listOfPointDistribution.PushBack(PointDistribution{9, 120})
	listOfPointDistribution.PushBack(PointDistribution{10, 100})
	listOfPointDistribution.PushBack(PointDistribution{11, 80})
	listOfPointDistribution.PushBack(PointDistribution{12, 65})
	listOfPointDistribution.PushBack(PointDistribution{13, 50})
	listOfPointDistribution.PushBack(PointDistribution{14, 40})
	listOfPointDistribution.PushBack(PointDistribution{15, 30})
	listOfPointDistribution.PushBack(PointDistribution{16, 20})
	listOfPointDistribution.PushBack(PointDistribution{17, 15})
	listOfPointDistribution.PushBack(PointDistribution{18, 10})
	listOfPointDistribution.PushBack(PointDistribution{19, 5})
	listOfPointDistribution.PushBack(PointDistribution{20, 0})
	//fmt.Println(*listOfPointDistribution)

	//3. Create matches & match details
	listOfMatches = list.New()

	//Round 1
	matchdetail1 := MatchDetails{team1, PointDistribution{3, 345}, 7, 0}
	matchdetail1.caculatePoint()
	matchdetail2 := MatchDetails{team2, PointDistribution{2, 410}, 2, 0}
	matchdetail2.caculatePoint()
	matchdetail3 := MatchDetails{team3, PointDistribution{1, 500}, 4, 0}
	matchdetail3.caculatePoint()
	//fmt.Println(matchdetail1)
	match1 := Match{1, "Round 1", "Erangel", nil}
	match1.listOfMatchDetails = list.New()
	match1.listOfMatchDetails.PushBack(matchdetail1)
	match1.listOfMatchDetails.PushBack(matchdetail2)
	match1.listOfMatchDetails.PushBack(matchdetail3)

	//Round 2
	matchdetail4 := MatchDetails{team1, PointDistribution{2, 410}, 3, 0}
	matchdetail4.caculatePoint()
	matchdetail5 := MatchDetails{team2, PointDistribution{1, 500}, 5, 0}
	matchdetail5.caculatePoint()
	matchdetail6 := MatchDetails{team3, PointDistribution{3, 345}, 2, 0}
	matchdetail6.caculatePoint()
	match2 := Match{2, "Round 2", "Erangel", nil}
	match2.listOfMatchDetails = list.New()
	match2.listOfMatchDetails.PushBack(matchdetail4)
	match2.listOfMatchDetails.PushBack(matchdetail5)
	match2.listOfMatchDetails.PushBack(matchdetail6)

	//Round 3
	matchdetail7 := MatchDetails{team1, PointDistribution{3, 345}, 5, 0}
	matchdetail7.caculatePoint()
	matchdetail8 := MatchDetails{team2, PointDistribution{2, 410}, 1, 0}
	matchdetail8.caculatePoint()
	matchdetail9 := MatchDetails{team3, PointDistribution{1, 500}, 4, 0}
	matchdetail9.caculatePoint()
	match3 := Match{3, "Round 3", "Erangel", nil}
	match3.listOfMatchDetails = list.New()
	match3.listOfMatchDetails.PushBack(matchdetail7)
	match3.listOfMatchDetails.PushBack(matchdetail8)
	match3.listOfMatchDetails.PushBack(matchdetail9)

	//Add matches
	listOfMatches.PushBack(match1)
	listOfMatches.PushBack(match2)
	listOfMatches.PushBack(match3)
}

//1. Load database
func LoadData() {
	CreateFakeData()
}

//2. Calculate table
func CalculatePGI() {
	resultList = list.New()

	//Filter team result
	for item := listOfTeam.Front(); item != nil; item = item.Next() {
		var data = item.Value.(Team)
		resultItem := ResultItem{data, list.New(), 0}

		//Each team we get all match details of each
		for inner := listOfMatches.Front(); inner != nil; inner = inner.Next() {
			var list = inner.Value.(Match).listOfMatchDetails
			//Add detail of specific match
			for i := list.Front(); i != nil; i = i.Next() {
				var detail = i.Value.(MatchDetails)
				if detail.team.id == data.id {
					resultItem.listOfMatchDetails.PushBack(detail)
					resultItem.totalPoint += detail.detail.point + detail.kill*15
				}
			}
		}

		resultList.PushBack(resultItem)
	}

	//Sort

}

//3. Display section
func DisplayMatch(match Match) string {
	var result = ""

	result = "ID:         " + strconv.Itoa(match.id) + "\n" +
		"Match Name: " + match.matchName + "\n" +
		"Map:        " + match.mapName

	result += "\n\tDetails: \n"

	for inner := match.listOfMatchDetails.Front(); inner != nil; inner = inner.Next() {
		result += inner.Value.(MatchDetails).team.name + ", " +
			inner.Value.(MatchDetails).team.regional + ", " +
			strconv.Itoa(inner.Value.(MatchDetails).kill) + ", " +
			strconv.Itoa(inner.Value.(MatchDetails).totalPoint) + "\n"
	}

	return result
}

func DisplayResultItems() string {
	var result = ""

	for item := resultList.Front(); item != nil; item = item.Next() {
		result += item.Value.(ResultItem).team.name + "\t"
		for inner := item.Value.(ResultItem).listOfMatchDetails.Front(); inner != nil; inner = inner.Next() {
			var data = inner.Value.(MatchDetails)
			result += "[" + strconv.Itoa(data.detail.placement) + ", " + strconv.Itoa(data.kill) + "]" + " - "

		}
		result += strconv.Itoa(item.Value.(ResultItem).totalPoint) + "\n"
	}

	return result
}

func DisplayAll() {
	//1. Print teams
	if listOfTeam != nil {
		fmt.Println("Total team: ", listOfTeam.Len())

		for index := listOfTeam.Front(); index != nil; index = index.Next() {
			fmt.Println(index.Value.(Team).name, " - ", index.Value.(Team).regional)
		}

		fmt.Print("\n")
	}

	//2. Print point distributions
	if listOfPointDistribution != nil {
		fmt.Println("Total point distribution: ", listOfPointDistribution.Len())

		for index := listOfPointDistribution.Front(); index != nil; index = index.Next() {
			fmt.Println(index.Value.(PointDistribution).placement, " - ", index.Value.(PointDistribution).point)
		}
		fmt.Print("\n")
	}

	//3. Print matches
	if listOfMatches != nil {
		fmt.Println("Total matches: ", listOfMatches.Len())

		for index := listOfMatches.Front(); index != nil; index = index.Next() {
			fmt.Print(DisplayMatch(index.Value.(Match)))
			fmt.Print("\n")
		}
	}

	//4. Print result items
	fmt.Println(DisplayResultItems())
}

func main() {
	LoadData()
	CalculatePGI()
	DisplayAll()
}
