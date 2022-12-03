package day3

import (
	"fmt"
	"strings"
)

func letters() string {
	l := "abcdefghijklmnopqrstuvwxyz"
	return l + strings.ToUpper(l)
}

func p1(in string) int {
	var ans int

	lines := strings.Split(in, "\n")
	lines = lines[:len(lines)-1]

	for _, l := range lines {
		div := len(l) / 2

		for i, c := range l {
			char := fmt.Sprintf("%c", c)
			if i >= div {
				if strings.Index(l[:div], char) != -1 {
					ans += strings.Index(letters(), char) + 1
					break
				}
			}
		}
	}

	return ans
}

func p2(in string) int {
	var ans int

	lines := strings.Split(in, "\n")
	lines = lines[:len(lines)-1]
	var group []string
	var groups [][]string
	for i, l := range lines {
		group = append(group, l)
		ind := i + 1
		if ind%3 == 0 {
			groups = append(groups, group)
			group = []string{}
		}
	}

GroupsLoop:
	for _, grp := range groups {
		for i, line := range grp {
			for _, c := range line {
				char := fmt.Sprintf("%c", c)
				var firstHit bool
				var secondHit bool

				if strings.Index(grp[i+1], char) != -1 {
					firstHit = true
				}

				if strings.Index(grp[i+2], char) != -1 {
					secondHit = true
				}

				if firstHit && secondHit {
					ans += strings.Index(letters(), char) + 1
					continue GroupsLoop
				}
			}
		}
	}

	return ans
}
