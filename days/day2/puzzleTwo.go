package day2

var instructions []instruction = []instruction{win, draw, lose}

func puzzleTwo(in string) int {
	pairs :=  pairsFromInput(in)
	var totalPoints int

	for _, p := range pairs {
		letters := lettersFromPair(p)
		
		var left item
		for _, gItem := range gameItems {
			for _, l := range gItem.letters {
				if l == letters[0] {
					left = gItem
					break
				}
			}
		}

		var right instruction
		for _, instr := range instructions {
			if instr.letter == letters[1] {
				right = instr
			}
		}

		playersItems := leftAndRightFromInstructions(&left, &right)
		playerPoints := playerPointsByItems(&playersItems)
		totalPoints += playerPoints
	}

	return totalPoints
}

func leftAndRightFromInstructions(i *item, instr *instruction) items {
	res := items{*i}
	var right *item

	if instr.letter == "X" {
		right = instr.loseItem(i)
	} else if instr.letter == "Y" {
		right = instr.drawItem(i)
	} else {
		right = instr.winItem(i)
	}

	res = append(res, *right)
	return res

}
