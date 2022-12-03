package main

import (
	"fmt"
	d "myAwesomeModule/days"
	"os"
)

// run with `go run .`, or with `go run ./days/.` from project root
func main() {
	switch os.Args[1] {
	case "1":
		d.Day1()
	case "2":
		d.Day2()
	default:
		fmt.Println("Ain't nobody got day for that.")
	}
}
