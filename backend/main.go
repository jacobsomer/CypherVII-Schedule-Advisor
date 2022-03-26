package main

import (
	db "backend/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Returns Optimal Class Schedule given classes already taken and major.
	router.GET("/getOptimalSchedule/:schedule", GetOptimalScheduleHandler)

	// Returns the Available Classes.
	router.GET("/getAvailableCourses/:courseToSwitch", GetAvailableCourses)

	// TODO: upload opencourseList database locally such that it can be accessed from code.
	// Possible solution: https://github.com/boltdb/bolt see README.md

	// TODO: upload Major/Graduation requirements database just as above

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func GetOptimalScheduleHandler(c *gin.Context) {
	schedule := c.Param("schedule")
	/* Schedule is an object given by the frontend we will define more precisely
	Struct Params:
		- classesTaken: List<int>
		- Majors/minors Desired: String
		- Potentially other preferences (Stretch goal): String
	*/
	// TODO: write function that takes in above params and return class schedule
	// Ex. func getOptimalSchedule(classesTaken, majors, minor, ....)
	// See https://drive.google.com/file/d/1SDH2SaaErjDdW9vOl25BFjtamK0eBbC4/view?usp=sharing for pseudocode
	c.String(http.StatusOK, "Hello %s", schedule)
}

func GetAvailableCourses(c *gin.Context) {
	//class := c.Param("courseToSwitch")

	class := db.Class{
		Subj: "CSCI",
		Id:   "303",
	}

	//classes := db.FilterBySubject("CSCI")

	classes := db.CheckPrerequisite(class)

	output := ""
	for _, cl := range classes {
		for _, cl := range cl {
			fmt.Printf("%s, %s\n", cl.Subj, cl.Id)
			output += " " + fmt.Sprintf("%s, %s\n", cl.Subj, cl.Id)
		}
	}

	//cl := classes[0][0]
	//singleClass := fmt.Sprintf("%s, %s, Credits: %v From: %v-%v, PreRequs: %s\n", cl.Subj, cl.Id, cl.Cred, cl.Start, cl.End, cl.PreReq)

	//courseToSwitch := c.Param("courseToSwitch")
	c.String(http.StatusOK, output)
}
