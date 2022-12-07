package day7

import (
	"reflect"
	"strconv"
	"strings"

	inpututils "github.com/OlegGulevskyy/aoc2022/input-utils"
)

var TotalSizeThreshold = 100000

type Dir struct {
	Name     string
	Size     int
	Files    []string
	Children []*Dir
	Parent   *Dir
}

var ans int

func p1(input string) int {
	lines := inpututils.Lines(input)
	fs := make(map[string]*Dir)
	fs["/"] = &Dir{Name: "/", Size: 0, Children: []*Dir{}, Parent: nil}

	currDir := fs["/"]

	for l := 2; l < len(lines); l++ {
		line := lines[l]

		if isCmd(line) {
			cmd := strings.Split(line, " ")[1]
			if cmd == "ls" {
				continue
			}

			toDir := strings.Split(line, " ")[2]
			if toDir == ".." {
				currDir = currDir.Parent
			} else {
				for _, d := range currDir.Children {
					if d.Name == toDir {
						currDir = d
						break
					}
				}
			}

		} else if isDir(line) {
			dirName := strings.Split(line, " ")[1]
			dir := &Dir{
				Name:   dirName,
				Parent: currDir,
			}

			currDir.Children = append(currDir.Children, dir)
		} else {
			fileName := strings.Split(line, " ")[1]
			fileSize := fileSize(strings.Split(line, " ")[0])

			currDir.Size += fileSize
			currDir.Files = append(currDir.Files, fileName)

			incrementParentSize(currDir, fileSize)
		}
	}

	handleFs(fs["/"])
	return ans
}

func incrementParentSize(p *Dir, size int) {
	if p.Parent != nil {
		p.Parent.Size += size
		incrementParentSize(p.Parent, size)
	}
}

func handleFs(dir *Dir) {
	v := reflect.ValueOf(dir)
	indV := reflect.Indirect(v)

	if dir.Size <= TotalSizeThreshold {
		ans += dir.Size
	}

	for i := 0; i < indV.NumField(); i++ {
		fieldName := indV.Type().Field(i).Name
		fieldValue := indV.Field(i)

		if fieldName == "Children" {
			for _, child := range fieldValue.Interface().([]*Dir) {
				handleFs(child)
			}
		}

	}
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
