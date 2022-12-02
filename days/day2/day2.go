package day2

import (
	"embed"
	"fmt"
	"strings"
)

const (
	losePoints = 0
	drawPoints = 3
	winPoints  = 6
	rockName = "rock"
	scissorsName = "scissors"
	paperName = "paper"
)

type item struct {
	id      string
	letters []string
	points  int
	looseFrom string
}

func (i *item) matchItemByLetter(l string) bool {
	found := false
	for _, letter := range i.letters {
		if letter == l {
			found = true
			break
		}
	}
	return found
}

var rock item = item{
	letters: []string{"A", "X"},
	points:  1,
	id:      rockName,
	looseFrom: paperName,
	
}

var paper item = item{
	letters: []string{"B", "Y"},
	points:  2,
	id:      paperName,
	looseFrom: scissorsName,
}

var scissors item = item{
	letters: []string{"C", "Z"},
	points:  3,
	id:      scissorsName,
	looseFrom: rockName,
}

type items []item

var gameItems items = items{rock, scissors, paper}


//go:embed input.txt
var input embed.FS

func Solution() error {
	d, err := input.ReadFile("input.txt")
	if err != nil {
		return err
	}

	resOne := puzzleOne(string(d))
	fmt.Printf("Total points of %d", resOne)

	return nil
}

func puzzleOne(in string) int {
	pairs := strings.Split(in, "\n")
	// remove last item that is always empty
	pairs = pairs[:len(pairs)-1]
	var totalPoints int

	for _, p := range pairs {
		letters := strings.Split(p, " ")
		matchedItems := leftAndRightItems(letters[0], letters[1])
		playerPoints := playerPointsPerPair(&matchedItems)
		totalPoints += playerPoints
	}

	return totalPoints
}

func leftAndRightItems(letters ...string) items {
	var leftAndRight items
	for _, l := range letters {

		for _, gItem := range gameItems {
			found := gItem.matchItemByLetter(l)
			if found {
				leftAndRight = append(leftAndRight, gItem)
				break
			}
		}
	}

	return leftAndRight
}

func playerPointsPerPair(pair *items) int {
	left := (*pair)[0]
	right := (*pair)[1]

	var pairPoints int
	// it's a draw
	if left.id == right.id {
		pairPoints = drawPoints + right.points
	}

	// left side wins, it's a loose
	if left.id == right.looseFrom {
		pairPoints = losePoints + right.points
	}

	// right side wins, it's a win for a player
	if left.looseFrom == right.id {
		pairPoints = winPoints + right.points
	}

	return pairPoints
}
