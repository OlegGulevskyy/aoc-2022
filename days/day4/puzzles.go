package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/OlegGulevskyy/aoc2022/input-utils"
)

func p1(in string) (int, int) {
	var ans1 int
	var ans2 int

	lines := inpututils.Lines(in)

	for _, l := range lines {
		f := ""
		s := ""
		var totalOverlap bool

		left := strings.Split(l, ",")[0]
		right := strings.Split(l, ",")[1]

		lStartStr := strings.Split(left, "-")[0]
		lStart, _ := strconv.Atoi(lStartStr)

		lEndStr := strings.Split(left, "-")[1]
		lEnd, _ := strconv.Atoi(lEndStr)

		rStartStr := strings.Split(right, "-")[0]
		rStart, _ := strconv.Atoi(rStartStr)

		rEndStr := strings.Split(right, "-")[1]
		rEnd, _ := strconv.Atoi(rEndStr)

		for i := lStart; i <= lEnd; i++ {
			f += fmt.Sprintf("-%d-", i)
		}

		for i := rStart; i <= rEnd; i++ {
			char := fmt.Sprintf("-%d-", i)
			s += char

			if strings.Index(f, char) != -1 {
				totalOverlap = true
			}
		}

		if strings.Index(f, s) != -1 || strings.Index(s, f) != -1 {
			ans1++
		}
		if totalOverlap {
			ans2++
		}

	}

	return ans1, ans2
}

func p2(in string) int {
	var ans int

	return ans
}
