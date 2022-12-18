package main

import (
	"bufio"
	"fmt"
	"myAwesomeModule/utils"
	"os"
	"strconv"
	"strings"
)

var mapLines []string
var moveLines []string
var colCount int
var containers map[int][]string
var containersCopy map[int][]string

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

	containersCopy = make(map[int][]string, len(containers))
	for k, v := range containers {
		// slice map values needs to be explicity copied or values will "corrupt" or smthg
		c := make([]string, len(v))
		copy(c, v)
		containersCopy[k] = c
	}

	utils.CheckError(scanner.Err())
}

func part1() string {
	utils.StartTimer()

	for _, line := range moveLines {
		parts := strings.Split(line, " ")

		moveCnt, _ := strconv.Atoi(parts[1])
		fromIdx, _ := strconv.Atoi(parts[3])
		toIdx, _ := strconv.Atoi(parts[5])

		for i := 0; i < moveCnt; i++ {
			lastIdx := len(containers[fromIdx]) - 1
			char := containers[fromIdx][lastIdx]
			containers[fromIdx] = containers[fromIdx][:lastIdx]
			containers[toIdx] = append(containers[toIdx], char)
		}
	}

	var solution []string
	for i := 1; i <= colCount; i++ {
		solution = append(solution, containers[i][len(containers[i])-1])
	}

	return strings.Join(solution, "")
}

func part2() string {
	utils.StartTimer()

	for _, line := range moveLines {
		parts := strings.Split(line, " ")

		moveCnt, _ := strconv.Atoi(parts[1])
		fromIdx, _ := strconv.Atoi(parts[3])
		toIdx, _ := strconv.Atoi(parts[5])

		containersCopy[toIdx] = append(containersCopy[toIdx], containersCopy[fromIdx][len(containersCopy[fromIdx])-moveCnt:]...)
		containersCopy[fromIdx] = containersCopy[fromIdx][:len(containersCopy[fromIdx])-moveCnt]
	}

	var solution []string
	for i := 1; i <= colCount; i++ {
		solution = append(solution, containersCopy[i][len(containersCopy[i])-1])
	}

	return strings.Join(solution, "")
}

func main() {
	utils.InitDay()

	read(utils.GetInputPath())

	utils.WriteSolutionStr(part1())
	utils.WriteSolutionStr(part2())

	fmt.Println("done.")
}
