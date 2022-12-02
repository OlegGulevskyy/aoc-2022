package day2

import (
	"strings"
)

func p1(in string) int {
	var ans int
	lines := strings.Split(in, "\n")

	for _, l := range lines {
		if l == "" {
			continue
		}
		ln := strings.Split(l, " ")

		aIndex := strings.Index("ABC", ln[0]) + 1
		bIndex := strings.Index("XYZ", ln[1]) + 1

		ans += bIndex

		switch bIndex - aIndex {
		case 1, -2:
			ans += 6
		case 0:
			ans += 3
		}
	}

	return ans
}

func p2(in string) int {
	var ans int
	lines := strings.Split(in, "\n")

	for _, l := range lines {
		if l == "" {
			continue
		}

		ln := strings.Split(l, " ")
		aIndex := strings.Index("ABC", ln[0])
		switch ln[1] {
		case "X":
			if (aIndex+1)%4-1 == 0 {
				ans += 3
			} else {
				ans += (aIndex+1)%4 - 1
			}
		case "Y":
			ans += 3
			ans += aIndex + 1

		case "Z":
			ans += 6
			ans += ((aIndex + 1) % 3) + 1
		}
	}

	return ans
}
