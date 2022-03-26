package main

import (
	db "backend/database"
	"fmt"
)

func main() {
	classes := db.FilterBySubject("CSCI")
	println(classes)
	for _, cl := range classes {
		fmt.Printf("%s, %s, Credits: %v From: %v-%v, PreRequs: %s\n", cl.Subj, cl.Id, cl.Cred, cl.Start, cl.End, cl.PreReq)
	}
}
