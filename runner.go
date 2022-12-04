package main

import (
	"bytes"
	"fmt"
	d "myAwesomeModule/days"
	u "myAwesomeModule/utils"
	"os"
)

func main() {
	dayNum := os.Args[1]
	if len(os.Args) > 2 && os.Args[2] == "init" {
		templating(dayNum)
		fmt.Println("Day creation done.")
	} else {
		switch dayNum {
		case "1":
			d.Day1()
		case "2":
			//d.Day2()
		default:
			fmt.Println("Ain't nobody got day for that.")
		}
	}
}

func templating(dayNum string) {
	os.Create("inputs/day" + dayNum + "_input")
	os.Create("solutions/day" + dayNum + "_solution")

	data, err := os.ReadFile("template")
	u.CheckError(err)

	data = bytes.Replace(data, []byte("{DAYNUM}"), []byte(dayNum), -1)

	err = os.WriteFile("days/day"+dayNum+".go", data, 0644)
	u.CheckError(err)
}
