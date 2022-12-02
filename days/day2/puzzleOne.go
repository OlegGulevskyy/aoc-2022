package day2

func puzzleOne(in string) int {
	pairs :=  pairsFromInput(in)
	var totalPoints int

	for _, p := range pairs {
		letters := lettersFromPair(p)
		matchedItems := leftAndRightItems(letters[0], letters[1])
		playerPoints := playerPointsByItems(&matchedItems)
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

func playerPointsByItems(pair *items) int {
	left := (*pair)[0]
	right := (*pair)[1]

	var pairPoints int
	// it's a draw
	if left.id == right.id {
		pairPoints = draw.points + right.points
	}

	// left side wins, it's a loose
	if left.id == right.looseFrom {
		pairPoints = lose.points + right.points
	}

	// right side wins, it's a win for a player
	if left.looseFrom == right.id {
		pairPoints = win.points + right.points
	}

	return pairPoints
}
