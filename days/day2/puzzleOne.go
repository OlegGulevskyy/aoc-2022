package day2

import "strings"

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
