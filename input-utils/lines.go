package inpututils

import "strings"

func Lines(file string) []string {
	lines := strings.Split(file, "\n")
	lines = lines[:len(lines)-1]
	return lines
}
