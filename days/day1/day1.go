package day1

import (
	"embed"
	"log"
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

