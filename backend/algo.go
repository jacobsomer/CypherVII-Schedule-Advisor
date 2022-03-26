package main

import (
	"backend/algorithm"
	"fmt"
)

func main() {
	schedule := algorithm.MakeSchedule("CSCI", "B.S.")
	println(schedule)
	for _, semester := range schedule {
		for _, class := range semester {
			fmt.Printf("Class: %s, %s, ", class.Subj, class.Id)
		}
		fmt.Println()
	}
}
