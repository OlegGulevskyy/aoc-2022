package day5

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	inpututils "github.com/OlegGulevskyy/aoc2022/input-utils"
)

func p1(in string) string {
	var ans string
	lines := inpututils.Lines(in)
	var instructions []string

	containers := map[string][]string{}

	// parse containers
	for lineIndex, l := range lines {
		if strings.Index(l, "1") != -1 {
			for charIndex, c := range l {
				char := string(c)
				if char == " " {
					continue
				}

				for i := 0; i < lineIndex; i++ {
					container := string(lines[i][charIndex])
					if container == "" || container == " " {
						continue
					}

					containers[char] = append(containers[char], container)
				}

			}
			instructions = lines[lineIndex + 2:]
			break
		}
	}

	fmt.Println("Containers parsed => ", containers)
	for _, instr := range instructions {
		amount, _ := strconv.Atoi(string(instr[5]))
		srcStr := string(instr[12])
		destStr := string(instr[17])

		for k, c := range containers {
			if k != srcStr {
				continue
			}

			for valueIndex, value := range c {
				containers[destStr] = append([]string{value}, containers[destStr]...)
				containers[srcStr] = containers[srcStr][1:]

				if valueIndex + 1 == amount {
					break
				}
				
			}
		}

	}

	keys := make([]string, 0, len(containers))
	for k := range containers {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if len(containers[k]) > 0 {
			ans += containers[k][0]
		}
	}

	fmt.Println("Containers after instructions => ", containers)

	return ans
}

func p2(in string) int {
	var ans int
	return ans
}
