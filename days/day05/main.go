package main

import (
	"bufio"
	"fmt"
	"myAwesomeModule/utils"
	"os"
	"strings"
)

var mapLines []string
var moveLines []string
var colCount int
var containers map[int][]string

func read(filePath string) {
	file, err := os.Open(filePath)
	utils.CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	separatorFound := false
	for scanner.Scan() {
		if scanner.Text() == "" {
			separatorFound = true
			continue
		}

		if !separatorFound {
			mapLines = append(mapLines, scanner.Text())
		} else {
			moveLines = append(moveLines, scanner.Text())
		}
	}

	colCount = len(strings.ReplaceAll(mapLines[len(mapLines)-1], " ", ""))
	containers = make(map[int][]string)
	for j := len(mapLines) - 2; j >= 0; j-- {
		for k := 1; k <= colCount; k++ {
			col := k*4 - 3
			if len(mapLines[j]) < col {
				continue
			}

			char := string(mapLines[j][col])
			if char != " " {
				containers[k] = append(containers[k], char)
			}
		}
	}

	fmt.Println(containers)

	utils.CheckError(scanner.Err())
}

func part1() string {
	utils.StartTimer()

	// part1 solution

	return "asd"
}

func part2() string {
	utils.StartTimer()

	// part2 solution

	return "asd"
}

func main() {
	utils.InitDay()

	read(utils.GetInputPath())

	utils.WriteSolutionStr(part1())
	utils.WriteSolutionStr(part2())

	fmt.Println("done.")
}
