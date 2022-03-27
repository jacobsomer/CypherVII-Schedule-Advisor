package main

import (
	"backend/algorithm"
	"fmt"
)

func main() {
	schedule := algorithm.MakeSchedule("INRL", "B.S.")
	println(schedule)
	for _, semester := range schedule {
		for _, class := range semester {
			fmt.Printf("Class: %s %s, ", class.Subj, class.Id)
		}
		fmt.Println()
		fmt.Println("end of semester")
	}
}
