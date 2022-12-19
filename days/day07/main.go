package main

import (
	"bufio"
	"fmt"
	"myAwesomeModule/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var folders map[string]Dir

type Dir struct {
	Size     uint64
	Children []string
}

func read(filePath string) {
	file, err := os.Open(filePath)
	utils.CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	folders = make(map[string]Dir)

	reUp := regexp.MustCompile(`\$ cd \.\.`)
	reFolder := regexp.MustCompile(`\$ cd \/|\$ cd \w+`)
	reFile := regexp.MustCompile(`[0-9]+ [a-z.]+`)
	reChild := regexp.MustCompile(`dir [a-z]+`)

	f := "/"
	for scanner.Scan() {
		line := scanner.Text()

		if reUp.FindString(line) != "" {
			continue
		} else if reFolder.FindString(line) != "" {
			f = strings.ReplaceAll(line, "$ cd ", "")
			folders[f] = Dir{}
		} else if reFile.FindString(line) != "" {
			val, _ := strconv.Atoi(strings.Split(line, " ")[0])

			// https://stackoverflow.com/a/69006398
			if entry, ok := folders[f]; ok {
				entry.Size += uint64(val)
				folders[f] = entry
			}
		} else if reChild.FindString(line) != "" {
			// https://stackoverflow.com/a/69006398
			if entry, ok := folders[f]; ok {
				entry.Children = append(entry.Children, strings.Split(line, " ")[1])
				folders[f] = entry
			}
		} else {
			if line != "$ ls" {
				panic("Unhandled line type found.")
			}
		}
	}

	//fmt.Println(folders)

	utils.CheckError(scanner.Err())
}

func valueRecursion(d Dir) uint64 {
	val := d.Size
	for _, ch := range d.Children {
		val += valueRecursion(folders[ch])
	}

	return val
}

func part1() uint64 {
	utils.StartTimer()

	valueMap := make(map[string]uint64)
	for key, value := range folders {
		// it seems faster going through everything than checking if key already calculated
		valueMap[key] = valueRecursion(value)

		/*if len(folders) == len(valueMap) {
			break
		}*/
	}

	var sum uint64
	for _, val := range valueMap {
		if val <= 100000 {
			sum += val
		}
	}

	/*for key, value := range valueMap {
		fmt.Print(key)
		fmt.Print(" ")
		fmt.Print(value)
		fmt.Println()
	}*/

	//fmt.Println(valueMap)

	return sum
}

func part2() uint64 {
	utils.StartTimer()

	// part2 solution

	return 0
}

func main() {
	utils.InitDay()

	read(utils.GetInputPath())

	utils.WriteSolutionUint64(part1())
	utils.WriteSolutionUint64(part2())

	fmt.Println("done.")
}
