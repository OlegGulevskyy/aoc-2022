package day7

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	inpututils "github.com/OlegGulevskyy/aoc2022/input-utils"
)

var TotalSizeThreshold = 100000

func p1(input string) int {
	var ans int
	lines := inpututils.Lines(input)
	fs2 := map[string][]string{}

	currDir := "/"
	for l := 2; l < len(lines); l++ {
		line := lines[l]

		if isDir(line) {
			dirName := strings.Split(line, " ")[1]
			if fs2[currDir + dirName + "/"] == nil {
				fs2[currDir + dirName + "/"] = []string{}
			} else {
				fs2[currDir + dirName + "/"] = fs2[currDir + dirName + "/"]
			}

		} else if isCmd(line) {
			cmd := strings.Split(line, " ")[1]
			if cmd == "cd" {
				toDir := strings.Split(line, " ")[2]
				if toDir == ".." {
					currDir = strings.Join(strings.Split(currDir, "/")[:len(strings.Split(currDir, "/"))-2], "/") + "/"
					if currDir == "" {
						currDir = "/"
					}
				} else {
					currDir = currDir + toDir + "/"
				}
			
			} 
		} else {
			// file
			fs2[currDir] = append(fs2[currDir], line)
		}
	}


	keys := make([]string, 0, len(fs2))
	for k := range fs2 {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println("Dir:", k, "Files:", fs2[k])
	}

	for _, k := range keys {
		dir := k
		files := fs2[k]
		if len(files) == 0 {
			continue
		}

		dirSize := folderSize(files)
		if dirSize > TotalSizeThreshold {
			continue
		}

		if dirSize < TotalSizeThreshold {
			fmt.Println("Smaller than threshold:", dir, dirSize)
			ans += dirSize
		}

		// find size of nested dirs
		for nestedDir, _ := range fs2 {
			if dir == "/" {
				continue
			}

			if strings.Index(nestedDir, dir) == 0 && nestedDir != dir {
				nestedDirSize := folderSize(fs2[nestedDir])
				dirSize += nestedDirSize
			}
		}
		if dirSize <= TotalSizeThreshold {
			// ans += dirSize
		}
	}

	return ans
}

func folderSize(dir []string) int {
	var ans int
	for _, file := range dir {
		ans += fileSize(file)
	}
	return ans
}

func fileSize(file string) int {
	size, _ := strconv.Atoi(strings.Split(file, " ")[0])
	return size
}

func isDir(line string) bool {
	return strings.Index(line, "dir") != -1
}

func isCmd(line string) bool {
	return strings.Index(line, "$ ") != -1
}
