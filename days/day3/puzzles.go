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
	lines = lines[:len(lines) - 1]

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

	return ans
}
