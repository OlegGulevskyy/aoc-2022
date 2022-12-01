package day1

import (
	"embed"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input embed.FS

func Solution() error {
	d, err := input.ReadFile("input.txt")
	if err != nil {
		return err
	}

	resOne := puzzleOne(string(d))
	resTwo := puzzleTwo(resOne[:3])

	log.Println("Result of puzzle one => ", resOne[0])
	log.Println("Result of puzzle two => ", resTwo)

	return nil
}

func puzzleOne(in string) []int {
	ln := strings.Split(in, "\n\n")
	var allCalories []int

	for _, pocket := range ln {
		var caloriesCount int

		for _, c := range strings.Split(pocket, "\n") {
			if c == "" {
				continue
			}

			i, err := strconv.Atoi(c)
			if err != nil {
				log.Panicln(fmt.Sprintf("cannot conver number %v from string", c))
			}
			caloriesCount += i
		}
		allCalories = append(allCalories, caloriesCount)
	}

	sort.Slice(allCalories, func (i, j int) bool {
		return allCalories[i] > allCalories[j]
	})

	return allCalories
}

func puzzleTwo(topByCalories []int) int {
	var total int
	for _, c := range topByCalories {
		total += c
	}
	return total
}
