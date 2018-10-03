package main

import (
	"bufio"
	"fmt"
	"github.com/dqduy/pgi"
	"os"
	"sort"
	"strconv"
	"strings"
)

/////////////////////////////////////////////////////////////
//Global data
var resultListTpp []pgi.ResultItem
var resultListFpp []pgi.ResultItem

const dbName = "pgi.txt"

/////////////////////////////////////////////////////

//1. Load database
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

					pgi.ListOfTeam = append(pgi.ListOfTeam, pgi.MakeTeam(id,
						strings.TrimSpace(tokens[1]),
						strings.TrimSpace(tokens[2])))

				case 2:
					var tokens = strings.Split(str, ",")
					placement, err := strconv.Atoi(strings.TrimSpace(tokens[0]))
					point, err1 := strconv.Atoi(strings.TrimSpace(tokens[1]))

					if err != nil || err1 != nil {
						fmt.Println(err)
					}

					pgi.ListOfPointDistribution = append(pgi.ListOfPointDistribution, pgi.MakePointDistribution(placement, point))

				case 3:
					var tokens = strings.Split(str, ",")
					var tokensLen = len(tokens)
					//fmt.Println(tokensLen)

					id, err := strconv.Atoi(strings.TrimSpace(tokens[0]))
					if err != nil {
						fmt.Println(err)
					}

					match := pgi.Match{id, tokens[1], tokens[2], make([]pgi.MatchDetails, 0)}

					for i := 0; i < (tokensLen-3)/3; i++ {
						teamid, err := strconv.Atoi(strings.TrimSpace(tokens[3+(i*3)]))
						placement, err := strconv.Atoi(strings.TrimSpace(tokens[3+(i*3)+1]))
						kill, err := strconv.Atoi(strings.TrimSpace(tokens[3+(i*3)+2]))
						if err != nil {
							fmt.Println(err)
						}
						matchdetail := pgi.MatchDetails{pgi.SearchTeam(teamid), pgi.SearchPointDistribution(placement), kill, 0}
						matchdetail.CaculatePoint()

						match.ListOfMatchDetails = append(match.ListOfMatchDetails, matchdetail)
					}

					pgi.ListOfMatchesTpp = append(pgi.ListOfMatchesTpp, match)

				case 4:
					var tokens = strings.Split(str, ",")
					var tokensLen = len(tokens)
					//fmt.Println(tokensLen)

					id, err := strconv.Atoi(strings.TrimSpace(tokens[0]))
					if err != nil {
						fmt.Println(err)
					}

					match := pgi.Match{id, tokens[1], tokens[2], make([]pgi.MatchDetails, 0)}

					for i := 0; i < (tokensLen-3)/3; i++ {
						teamid, err := strconv.Atoi(strings.TrimSpace(tokens[3+(i*3)]))
						placement, err := strconv.Atoi(strings.TrimSpace(tokens[3+(i*3)+1]))
						kill, err := strconv.Atoi(strings.TrimSpace(tokens[3+(i*3)+2]))
						if err != nil {
							fmt.Println(err)
						}
						matchdetail := pgi.MatchDetails{pgi.SearchTeam(teamid), pgi.SearchPointDistribution(placement), kill, 0}
						matchdetail.CaculatePoint()

						match.ListOfMatchDetails = append(match.ListOfMatchDetails, matchdetail)
					}

					pgi.ListOfMatchesFpp = append(pgi.ListOfMatchesFpp, match)

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
