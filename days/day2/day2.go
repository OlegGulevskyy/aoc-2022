package day2

import (
	"embed"
	"fmt"
	"strings"
)

type instruction struct {
	points int
	letter string
}

type item struct {
	id        string
	letters   []string
	points    int
	looseFrom string
}

var win instruction = instruction{
	letter: "Z",
	points: 6,
}

var draw instruction = instruction{
	letter: "Y",
	points: 3,
}

var lose instruction = instruction{
	letter: "X",
	points: 0,
}

func (instr *instruction) loseItem(i *item) *item {
	var found *item
	for _, gItem := range gameItems {
		if gItem.looseFrom == i.id {
			found = &gItem
			break
		}
	}
	return found
}

func (instr *instruction) drawItem(i *item) *item {
	var found *item
	for _, gItem := range gameItems {
		if gItem.id == i.id {
			found = &gItem
			break
		}
	}
	return found
}

func (instr *instruction) winItem(i *item) *item {
	var found *item
	for _, gItem := range gameItems {
		if gItem.id == i.looseFrom {
			found = &gItem
			break
		}
	}
	return found
}


const (
	rockName     = "rock"
	scissorsName = "scissors"
	paperName    = "paper"
)

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
	letters:   []string{"A", "X"},
	points:    1,
	id:        rockName,
	looseFrom: paperName,
}

var paper item = item{
	letters:   []string{"B", "Y"},
	points:    2,
	id:        paperName,
	looseFrom: scissorsName,
}

var scissors item = item{
	letters:   []string{"C", "Z"},
	points:    3,
	id:        scissorsName,
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
	input := string(d)

	// resOne := puzzleOne(input)
	// fmt.Printf("Total points of Puzzle One: %d", resOne)
	resTwo := puzzleTwo(input)
	fmt.Printf("Total points of Puzzle Two: %d", resTwo)

	return nil
}

func pairsFromInput(in string) []string {
	pairs := strings.Split(in, "\n")
	// remove last item that is always empty
	pairs = pairs[:len(pairs)-1]
	return pairs
}

func lettersFromPair(in string) []string {
	return strings.Split(in, " ")
}
