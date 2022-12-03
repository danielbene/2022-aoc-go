package main

import (
	"fmt"
	d "myAwesomeModule/days"
	"os"
)

func main() {
	if len(os.Args) > 2 && os.Args[2] == "init" {
		fmt.Println("TODO: create new files")
	} else {
		switch os.Args[1] {
		case "1":
			d.Day1()
		case "2":
			//d.Day2()
		default:
			fmt.Println("Ain't nobody got day for that.")
		}
	}
}
