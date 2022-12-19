package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var start time.Time
var root string
var solutionPath string
var inputPath string
var testInputPath string

func InitDay() {
	root, _ = os.Getwd()
	solutionPath = root + "/solution.md"
	inputPath = root + "/input.txt"
	testInputPath = root + "/input_test.txt"

	if err := os.Truncate(solutionPath, 0); err != nil {
		fmt.Printf("Failed to truncate solution file: %v", err)
	}
}

func StartTimer() {
	start = time.Now()
}

func WriteSolutionUint64(solution uint64) {
	WriteSolutionStr(strconv.FormatUint(solution, 10))
}

func WriteSolutionInt(solution int) {
	WriteSolutionStr(strconv.Itoa(solution))
}

func WriteSolutionStr(solution string) {
	duration := time.Since(start)

	f, err := os.OpenFile(solutionPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	CheckError(err)

	defer f.Close()

	var sb strings.Builder
	sb.WriteString("Solution: " + solution + "\n")
	sb.WriteString("Duration: " + duration.String() + "\n")
	sb.WriteString("--------------------------\n")

	_, err2 := f.WriteString(sb.String())
	CheckError(err2)
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func GetInputPath() string {
	return inputPath
}

func GetTestInputPath() string {
	return testInputPath
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

/*func GetDayID() string {
	return dayId
}*/

/*func initDayID() {
	_, file, _, _ := runtime.Caller(2)
	dayId = strings.Split(filepath.Base(file), ".")[0]
}*/

/*func templating(dayNum string) {
	os.Create("inputs/day" + dayNum + "_input")
	os.Create("solutions/day" + dayNum + "_solution")

	data, err := os.ReadFile("template")
	u.CheckError(err)

	data = bytes.Replace(data, []byte("{DAYNUM}"), []byte(dayNum), -1)

	err = os.WriteFile("days/day"+dayNum+".go", data, 0644)
	u.CheckError(err)
}*/
