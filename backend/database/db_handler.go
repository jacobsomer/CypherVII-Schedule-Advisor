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

// FilterBySubject returns all courses of a given subject
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

// CheckPrerequisite returns a list of lists of prerequisite classes
func CheckPrerequisite(class Class) [][]Class {

	// open a new db connection
	db, err := sql.Open("sqlite3", "./database/courses.db")
	checkErr(err)

	defer db.Close() // defer the closure of the db connection

	// make two variables to
	wantedClasses := make([]Class, 0)
	neededClasses := make([][]Class, 0)

	// get the prerequisites for the class from the db
	rows, err := db.Query("SELECT Prerequisites FROM courses WHERE Subject = ? AND ID = ?", class.Subj, class.Id)
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

	// if the class does not exist in the db return nil
	if len(wantedClasses) == 0 {
		return nil
	}
	preRequString := wantedClasses[0].PreReq

	// create RegExpressions to clean up the strings
	AND := regexp.MustCompile(`AND`)
	OR := regexp.MustCompile(`OR`)
	Space := regexp.MustCompile(` `)
	Parentheses := regexp.MustCompile(`([()])`)

	orLists := AND.Split(preRequString, -1)

	for _, orList := range orLists {

		orList = Parentheses.ReplaceAllString(orList, "")

		if len(orList) == 0 {
			continue
		}

		orClassList := make([]Class, 0)
		orClasses := OR.Split(orList, -1)

		for _, orClass := range orClasses {

			for string(orClass[0]) == " " {
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

// GetMajorRequirements returns a list of required courses given a major
func GetMajorRequirements(major string) map[Class]Class {

	db, err := sql.Open("sqlite3", "./database/courses.db")
	checkErr(err)

	defer db.Close()

	classes := make(map[Class]Class)

	var rows *sql.Rows

	if major == "CSCI" {
		rows, err = db.Query("SELECT * FROM csci_requirements")
		checkErr(err)
	} else if major == "INRL" {
		rows, err = db.Query("SELECT * FROM inrl_requirement")
		checkErr(err)
	} else {
		rows, err = db.Query("SELECT * FROM csci_requirements")
		checkErr(err)
	}

	for rows != nil && rows.Next() {
		temp := Class{}
		err = rows.Scan(&temp.Subj, &temp.Id)
		classes[temp] = temp
	}

	return classes

}

// GetGeneralRequirements returns the general education required classes
func GetGeneralRequirements(degree string) map[Class]Class {
	if degree == "B.S." {
		degree = "bs_requirements"
	}

	db, err := sql.Open("sqlite3", "./database/courses.db")
	checkErr(err)

	defer db.Close()

	classes := make(map[Class]Class)

	rows, err := db.Query("SELECT * FROM bs_requirements")
	checkErr(err)

	i := 0
	for rows.Next() {
		temp := Class{}
		err = rows.Scan(&temp.Subj, &temp.Id)
		classes[temp] = temp
		i++
	}

	return classes

}

// GetCredits returns the number of credits of a given class
func GetCredits(class Class) int {
	db, err := sql.Open("sqlite3", "./database/courses.db")
	checkErr(err)

	defer db.Close()

	credit := 0

	rows, err := db.Query("SELECT Subject, ID, Credits, Start, End, Prerequisites FROM courses WHERE Subject = ? AND ID = ?", class.Subj, class.Id)
	checkErr(err)
	for rows.Next() {
		temp := Class{}
		err = rows.Scan(&temp.Subj, &temp.Id, &temp.Cred, &temp.Start, &temp.End, &temp.PreReq)
		if debug {
			fmt.Printf("%s, %s, Credits: %v From: %v-%v, PreRequs: %s\n", temp.Subj, temp.Id, temp.Cred, temp.Start, temp.End, temp.PreReq)
		}
		credit = temp.Cred
	}

	return credit
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
