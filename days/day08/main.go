package main

import (
	"bufio"
	"fmt"
	"myAwesomeModule/utils"
	"os"
)

var input []string
var treeMap [][]int

// yeaaah... so these are not the prettiest spoons in the drawer
func read(filePath string) {
	file, err := os.Open(filePath)
	utils.CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	treeMap = make([][]int, len(input), len(input[0]))

	idx := 0
	for _, line := range input {
		for _, r := range line {
			treeMap[idx] = append(treeMap[idx], int(r-'0'))
		}

		idx++
	}

	utils.CheckError(scanner.Err())
}

func testLR(sideTrees []int, height int) int {
	for _, j := range sideTrees {
		if j >= height {
			return 1
		}
	}

	return 0
}

func testFB(min int, max int, row int, val int) int {
	for i := min; i < max; i++ {
		if treeMap[i][row] >= val {
			return 1
		}
	}

	return 0
}

func checkVisibility(col int, row int) bool {
	val := treeMap[col][row]

	rowCnt := 0
	rowCnt += testLR(treeMap[col][0:row], val)
	rowCnt += testLR(treeMap[col][row+1:], val)

	colCnt := 0
	colCnt += testFB(0, col, row, val)
	colCnt += testFB(col+1, len(treeMap), row, val)

	return rowCnt < 2 || colCnt < 2
}

func part1() int {
	utils.StartTimer()

	visibleCnt := len(treeMap[0])*2 + len(treeMap)*2 - 4
	for i := 1; i < len(treeMap)-1; i++ {
		for j := 1; j < len(treeMap[0])-1; j++ {
			if checkVisibility(i, j) {
				visibleCnt++
			}
		}
	}

	return visibleCnt
}

func countScenic(col int, row int) int {
	val := treeMap[col][row]
	lCnt, rCnt, tCnt, bCnt := 0, 0, 0, 0

	low := treeMap[col][0:row]
	for j := len(low) - 1; j >= 0; j-- {
		test := low[j]
		lCnt++

		if test >= val {
			break
		}
	}

	for _, j := range treeMap[col][row+1:] {
		rCnt++
		if j >= val {
			break
		}
	}

	for i := col - 1; i >= 0; i-- {
		tCnt++
		if treeMap[i][row] >= val {
			break
		}
	}

	for i := col + 1; i < len(treeMap); i++ {
		bCnt++
		if treeMap[i][row] >= val {
			break
		}
	}

	return lCnt * rCnt * tCnt * bCnt
}

func part2() int {
	utils.StartTimer()

	topScenery := 0
	for i := 1; i < len(treeMap)-1; i++ {
		for j := 1; j < len(treeMap[0])-1; j++ {
			top := countScenic(i, j)
			if top > topScenery {
				topScenery = top
			}
		}
	}

	return topScenery
}

func main() {
	utils.InitDay()

	read(utils.GetInputPath())

	utils.WriteSolutionInt(part1())
	utils.WriteSolutionInt(part2())

	fmt.Println("done.")
}
