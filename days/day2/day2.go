package day2

import (
	"embed"
	"fmt"
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
