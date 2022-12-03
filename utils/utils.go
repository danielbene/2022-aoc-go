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
	getDayID()
	start = time.Now()
}

func WriteSolution(solution string) {
	duration := time.Since(start)

	root, _ := os.Getwd()
	f, err := os.OpenFile(root+"/solutions/"+dayId+"_solution", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	checkError(err)

	defer f.Close()

	var sb strings.Builder
	sb.WriteString("Solution: " + solution + "\n")
	sb.WriteString("Duration: " + duration.String() + "\n")
	sb.WriteString("--------------------------\n")

	_, err2 := f.WriteString(sb.String())
	checkError(err2)
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func getDayID() {
	_, file, _, _ := runtime.Caller(2)
	dayId = strings.Split(filepath.Base(file), ".")[0]
}
