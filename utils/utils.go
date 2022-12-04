package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var start time.Time
var dayId string

func Hey() {
	name := "util"
	fmt.Println("Hello", name)
}

func StartTimer() {
	initDayID()
	start = time.Now()
}

func WriteSolution(solution string) {
	duration := time.Since(start)

	root, _ := os.Getwd()
	f, err := os.OpenFile(root+"/solutions/"+dayId+"_solution", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	CheckError(err)

	defer f.Close()

	var sb strings.Builder
	sb.WriteString("Solution: " + solution + "\n")
	sb.WriteString("Duration: " + duration.String() + "\n")
	sb.WriteString("--------------------------\n")

	_, err2 := f.WriteString(sb.String())
	CheckError(err2)
}

func GetDayID() string {
	return dayId
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func initDayID() {
	_, file, _, _ := runtime.Caller(2)
	dayId = strings.Split(filepath.Base(file), ".")[0]
}
