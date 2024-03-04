package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countGoals(first []int, second []int, hoa int) int {
	var firstTeamHome, firstTeamAway int
	var secondTeamHome, secondTeamAway int

	if hoa == 1 {
		firstTeamHome = first[0]
		firstTeamAway = second[0]
		secondTeamHome = second[1]
		secondTeamAway = first[1]
	} else {
		firstTeamHome = second[0]
		firstTeamAway = first[0]
		secondTeamHome = first[1]
		secondTeamAway = second[1]
	}

	firstTeamGoals := firstTeamHome + firstTeamAway
	secondTeamGoals := secondTeamHome + secondTeamAway

	if firstTeamGoals > secondTeamGoals {
		return 0
	}

	if firstTeamGoals == 0 && secondTeamGoals == 0 {
		return 1
	}

	equalizeGoals := secondTeamGoals - firstTeamGoals

	if hoa == 1 {
		if equalizeGoals+firstTeamAway > secondTeamAway {
			return equalizeGoals
		} else {
			return equalizeGoals + 1
		}
	} else {
		if firstTeamAway > secondTeamAway {
			return equalizeGoals
		} else {
			return equalizeGoals + 1
		}
	}
}

func main() {
	var firstGame, secondGame string
	var homeOrAway int

	fmt.Scan(&firstGame, &secondGame, &homeOrAway)

	firstGameScores := strings.Split(firstGame, ":")
	secondGameScores := strings.Split(secondGame, ":")

	var firstGameScoresInt, secondGameScoresInt []int

	for _, value := range firstGameScores {
		num, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		firstGameScoresInt = append(firstGameScoresInt, num)
	}

	for _, value := range secondGameScores {
		num, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		secondGameScoresInt = append(secondGameScoresInt, num)
	}

	fmt.Println(countGoals(firstGameScoresInt, secondGameScoresInt, homeOrAway))
}
