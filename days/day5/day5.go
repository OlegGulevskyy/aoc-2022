package day5

import (
	"embed"
	"fmt"
	"log"
)

//go:embed *.txt
var input embed.FS

func Solution() error {
	d, err := input.ReadFile("prod-input.txt")
	if err != nil {
		log.Panicln(err)
		return err
	}
	input := string(d)

	p1Ans := p1(input)
	p2Ans := p2(input)

	fmt.Println("--------------")
	fmt.Println("Day 05: Puzzle one answer: ", p1Ans)
	fmt.Println("Day 05: Puzzle two answer: ", p2Ans)
	fmt.Println("--------------")

	return nil
}


