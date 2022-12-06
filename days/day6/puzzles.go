package day6

func p1(input string) int {
	ans := 0
	measureBy := 14

InputLoop:
	for i, _ := range input {
		currentGroup := input[i:i+measureBy]
		chars := make(map[string]bool)

		for _, l := range currentGroup {
			ch := string(l)

			if chars[ch] == false {
				chars[ch] = true
			} else {
				continue InputLoop
			}
		}

		tempAns := input[:i+measureBy]
		ans = len(tempAns)
		break InputLoop
	}

	return ans
}
