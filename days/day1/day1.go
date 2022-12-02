package day1

import (
	"embed"
	"fmt"
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

	fmt.Println("--------------")
	fmt.Println("Day 01: Result of puzzle one: ", resOne[0])
	fmt.Println("Day 01: Result of puzzle two: ", resTwo)
	fmt.Println("--------------")

	return nil
}

