package day2

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
	input := string(d)

	p1Ans := p1(input)
	p2Ans := p2(input)

	fmt.Println("--------------")
	fmt.Println("Day 02: Puzzle one answer: ", p1Ans)
	fmt.Println("Day 02: Puzzle two answer: ", p2Ans)
	fmt.Println("--------------")

	return nil
}
