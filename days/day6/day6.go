package day6

import (
	"embed"
	"fmt"
	"log"
)

//go:embed *.txt
var input embed.FS

func Solution() error {
	day := 6

	d, err := input.ReadFile("prod_in.txt")
	if err != nil {
		log.Panicln(err)
		return err
	}
	input := string(d)

	p1Ans := p1(input)

	p1msg := fmt.Sprintf("Day %d: Puzzle one answer: %v", day, p1Ans)

	fmt.Println("--------------")
	fmt.Println(p1msg)
	fmt.Println("--------------")

	return nil
}


