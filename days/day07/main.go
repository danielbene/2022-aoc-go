package main

import (
	"bufio"
	"fmt"
	"myAwesomeModule/utils"
	"os"
	"strconv"
	"strings"
)

var folders map[string]int = make(map[string]int)

func read(filePath string) {
	file, err := os.Open(filePath)
	utils.CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	path := "/root"
	folders[path] = 0

	// this day was a nightmare
	// after countles errors angry me used this as an "inspiration"
	// https://github.com/jeff-frederic/AdventOfCode-2022/blob/main/day7.py
	for scanner.Scan() {
		line := scanner.Text()

		if string(line[0]) == "$" {
			if line[2:4] == "cd" {
				if string(line[5]) == "/" {
					path = "/root"
				} else if len(line) >= 7 && line[5:7] == ".." {
					path = path[0:strings.LastIndex(path, "/")]
				} else {
					dir := line[5:]
					path = path + "/" + dir
					folders[path] = 0
				}
			}
		} else if line[0:3] != "dir" {
			size, _ := strconv.Atoi(line[:strings.LastIndex(line, " ")])
			dir := path
			for i := 0; i < strings.Count(path, "/"); i++ {
				folders[dir] += size
				dir = dir[:strings.LastIndex(dir, "/")]
			}
		}
	}

	utils.CheckError(scanner.Err())
}

func part1() int {
	utils.StartTimer()

	sum := 0
	for _, v := range folders {
		if v <= 100_000 {
			sum += v
		}
	}

	return sum
}

func part2() int {
	utils.StartTimer()

	req := 30_000_000 - (70_000_000 - folders["/root"])
	var candidates []int

	for _, v := range folders {
		if v >= req {
			candidates = append(candidates, v)
		}
	}

	min, _ := utils.MinMax(candidates)

	return min
}

func main() {
	utils.InitDay()

	read(utils.GetInputPath())

	utils.WriteSolutionInt(part1())
	utils.WriteSolutionInt(part2())

	fmt.Println("done.")
}
