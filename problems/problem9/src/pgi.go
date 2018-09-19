package main

import (
	"bufio"
	"container/list"
	"fmt"
	//"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
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

type ByPoint []ResultItem

func (this ByPoint) Len() int {
	return len(this)
}

func (this ByPoint) Less(i, j int) bool {
	return this[i].totalPoint > this[j].totalPoint
}

func (this ByPoint) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

//Global data
var listOfTeam *list.List
var listOfPointDistribution *list.List
var listOfMatches *list.List
var resultList []ResultItem

const dbName = "pgi.txt"

//

//Fake data
func CreateFakeData() {
	//1. Create teams
	listOfTeam = list.New()
	team1 := Team{1, "Team Liquid", "European Qualifier"}
	team2 := Team{2, "Welcome to South Georgo", "European Qualifier"}
	team3 := Team{9, "Gen.G Gold", "Korean Qualifier"}

	//listOfTeam.PushBack(team1)
	//listOfTeam.PushBack(team2)
	//listOfTeam.PushBack(team3)

	//2. Create point distributions
	listOfPointDistribution = list.New()
	// listOfPointDistribution.PushBack(PointDistribution{0, 15})
	// listOfPointDistribution.PushBack(PointDistribution{1, 500})
	// listOfPointDistribution.PushBack(PointDistribution{2, 410})
	// listOfPointDistribution.PushBack(PointDistribution{3, 345})
	// listOfPointDistribution.PushBack(PointDistribution{4, 295})
	// listOfPointDistribution.PushBack(PointDistribution{5, 250})
	// listOfPointDistribution.PushBack(PointDistribution{6, 210})
	// listOfPointDistribution.PushBack(PointDistribution{7, 175})
	// listOfPointDistribution.PushBack(PointDistribution{8, 145})
	// listOfPointDistribution.PushBack(PointDistribution{9, 120})
	// listOfPointDistribution.PushBack(PointDistribution{10, 100})
	// listOfPointDistribution.PushBack(PointDistribution{11, 80})
	// listOfPointDistribution.PushBack(PointDistribution{12, 65})
	// listOfPointDistribution.PushBack(PointDistribution{13, 50})
	// listOfPointDistribution.PushBack(PointDistribution{14, 40})
	// listOfPointDistribution.PushBack(PointDistribution{15, 30})
	// listOfPointDistribution.PushBack(PointDistribution{16, 20})
	// listOfPointDistribution.PushBack(PointDistribution{17, 15})
	// listOfPointDistribution.PushBack(PointDistribution{18, 10})
	// listOfPointDistribution.PushBack(PointDistribution{19, 5})
	// listOfPointDistribution.PushBack(PointDistribution{20, 0})
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
	matchdetail8 := MatchDetails{team2, PointDistribution{1, 500}, 7, 0}
	matchdetail8.caculatePoint()
	matchdetail9 := MatchDetails{team3, PointDistribution{2, 410}, 1, 0}
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

func ParseTeams() {

}

func ParsePointDistributions() {

}

func ParseTPP() {

}

func ParseFPP() {

}

func MakeTeam(id int, name string, regional string) Team {
	return Team{id, name, regional}
}

func MakePointDistribution(placement int, point int) PointDistribution {
	return PointDistribution{placement, point}
}

//1. Load database
func LoadData() {
	CreateFakeData()

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
		//var lineLen = len(str)

		if str == "team" {
			section = 1
			fmt.Printf("%s - len: %d\n", str, len(str))
		} else if str == "points distribution" {
			section = 2
			fmt.Printf("%s - len: %d\n", str, len(str))
		} else if str == "tpp" {
			section = 3
			fmt.Printf("%s - len: %d\n", str, len(str))
		} else if str == "fpp" {
			section = 4
			fmt.Printf("%s - len: %d\n", str, len(str))
		} else {
			//fmt.Printf("other text - len: %d\n", len(str))
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
					//fmt.Printf("%s - len: %d\n", str, len(str))
				case 3:
					//fmt.Printf("%s - len: %d\n", str, len(str))
					//9, 8, 1,1, 6, 6,2, 18, 2,11, 5, 6,3, 12, 3,10, 14, 3,17, 15, 2,5, 1, 12,18, 7, 4,7, 13, 6,19, 4, 4,15, 17, 1,12, 3, 9,13, 2, 4,6, 9, 3,8, 11, 2,20, 10, 4,14, 20, 0,16, 19, 0,4, 16, 1
					var tokens = strings.Split(str, ",")
					var tokensLen = len(tokens)

				case 4:
					//fmt.Printf("%s - len: %d\n", str, len(str))
				default:
					fmt.Printf("Unknown section")
				}
			} else {
				fmt.Println("We have a blank line....")
			}
		}

		//fmt.Printf("%s\n", str)

	}

	// var s = "jkk   jk   kj   "
	// fmt.Printf("%q", s[0])
	// fmt.Println("len: ", len(strings.TrimSpace(s)), " - ", strings.TrimSpace(s))

}

//2. Calculate table
func CalculatePGI() {
	resultList = make([]ResultItem, 0)

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
		resultList = append(resultList, resultItem)
	}

	//Sort result list
	//1. Point
	//2. if two point's team is equal, we compare kill
	sort.Sort(ByPoint(resultList))

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

	for _, item := range resultList {
		result += item.team.name + "\t"
		for inner := item.listOfMatchDetails.Front(); inner != nil; inner = inner.Next() {
			var data = inner.Value.(MatchDetails)
			result += "[" + strconv.Itoa(data.detail.placement) + ", " + strconv.Itoa(data.kill) + "]" + " - "

		}
		result += strconv.Itoa(item.totalPoint) + "\n"
	}
	return result
}

func DisplayAllTeams() {
	if listOfTeam != nil {
		fmt.Println("Total team: ", listOfTeam.Len())

		for index := listOfTeam.Front(); index != nil; index = index.Next() {
			fmt.Println(index.Value.(Team).id, " - ", index.Value.(Team).name, " - ", index.Value.(Team).regional)
		}

		fmt.Print("\n")
	}
}

func DisplayPointDistribution() {
	if listOfPointDistribution != nil {
		fmt.Println("Total point distribution: ", listOfPointDistribution.Len())

		for index := listOfPointDistribution.Front(); index != nil; index = index.Next() {
			fmt.Println(index.Value.(PointDistribution).placement, " - ", index.Value.(PointDistribution).point)
		}
		fmt.Print("\n")
	}
}

func DisplayMatches() {
	if listOfMatches != nil {
		fmt.Println("Total matches: ", listOfMatches.Len())

		for index := listOfMatches.Front(); index != nil; index = index.Next() {
			fmt.Print(DisplayMatch(index.Value.(Match)))
			fmt.Print("\n")
		}
	}
}

func DisplayAll() {
	//1. Print teams
	DisplayAllTeams()

	//2. Print point distributions
	DisplayPointDistribution()

	//3. Print matches
	//DisplayMatches()

	//4. Print result items
	fmt.Println(DisplayResultItems())
}

func main() {
	LoadData()
	CalculatePGI()
	DisplayAll()
}
