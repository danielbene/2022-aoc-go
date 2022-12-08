package utils

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var start time.Time
var root string
var solutionPath string
var inputPath string

func InitDay() {
	root, _ = os.Getwd()
	solutionPath = root + "/solution.md"
	inputPath = root + "/input.txt"

	if err := os.Truncate(solutionPath, 0); err != nil {
		fmt.Printf("Failed to truncate solution file: %v", err)
	}

	// TODO: add input parse
}

func StartTimer() {
	start = time.Now()
}

func WriteSolution(solution string) {
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
