package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"regexp"
)

const debug = false

type Class struct {
	Subj   string
	Id     string
	Cred   int
	Start  int
	End    int
	PreReq string
}

func FilterBySubject(subject string) []Class {
	db, err := sql.Open("sqlite3", "./database/courses.db")
	checkErr(err)

	defer db.Close()

	classes := make([]Class, 0)

	lim := 1000
	if debug {
		lim = 10
	}

	rows, err := db.Query("SELECT Subject, ID, Credits, Start, End, Prerequisites FROM courses WHERE Subject = ? LIMIT ?", subject, lim)
	checkErr(err)
	for rows.Next() {
		temp := Class{}
		err = rows.Scan(&temp.Subj, &temp.Id, &temp.Cred, &temp.Start, &temp.End, &temp.PreReq)
		if debug {
			fmt.Printf("%s, %s, Credits: %v From: %v-%v, PreRequs: %s\n", temp.Subj, temp.Id, temp.Cred, temp.Start, temp.End, temp.PreReq)
		}
		classes = append(classes, temp)
	}

	return classes

}

func CheckPrerequisite(class Class) [][]Class {
	db, err := sql.Open("sqlite3", "./database/courses.db")
	checkErr(err)

	defer db.Close()

	wantedClasses := make([]Class, 0)
	neededClasses := make([][]Class, 0)

	lim := 1000
	if debug {
		lim = 10
	}

	// get the prerequisites for the class from the db
	rows, err := db.Query("SELECT Prerequisites FROM courses WHERE Subject = ? AND ID = ? LIMIT 1", class.Subj, class.Id, lim)
	checkErr(err)

	//var preRequ string
	for rows.Next() {
		temp := Class{}
		err = rows.Scan(&temp.PreReq)
		if debug {
			fmt.Printf("%s, %s, Credits: %v From: %v-%v, PreRequs: %s\n", temp.Subj, temp.Id, temp.Cred, temp.Start, temp.End, temp.PreReq)
		}
		wantedClasses = append(wantedClasses, temp)
	}

	preRequString := wantedClasses[0].PreReq

	AND := regexp.MustCompile(`AND`)
	OR := regexp.MustCompile(`OR`)
	Space := regexp.MustCompile(` `)

	orLists := AND.Split(preRequString, -1)

	for _, orList := range orLists {

		// trim of parentheses
		if string(orList[1]) == "(" {
			orList = orList[2 : len(orList)-2]
		}

		orClassList := make([]Class, 0)
		orClasses := OR.Split(orList, -1)

		for _, orClass := range orClasses {

			// trim of spaces
			if string(orClass[0]) == " " {
				orClass = orClass[1:]
			}

			arr := Space.Split(orClass, -1)

			tempClass := Class{
				Subj: arr[0],
				Id:   arr[1],
			}

			orClassList = append(orClassList, tempClass)
		}
		neededClasses = append(neededClasses, orClassList)
	}

	return neededClasses

}

//func FilterByPrerequisite() []Class {
//
//}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
