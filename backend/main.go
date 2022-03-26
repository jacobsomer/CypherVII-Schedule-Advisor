package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	/* 
		Returns Optimal Class Schedule given classes already taken and major.
	*/
	router.GET("/getOptimalSchedule/:schedule", func(c *gin.Context) {
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
	})

	router.GET("/getAvailableCourses/:courseToSwitch", func(c *gin.Context) {
		courseToSwitch := c.Param("courseToSwitch")
		c.String(http.StatusOK, courseToSwitch)
	})

	// TODO: upload opencourseList database locally such that it can be accessed from code. 
	// Possible solution: https://github.com/boltdb/bolt see README.md

	// TODO: upload Major/Graduation requirements database just as above

	router.Run(":8080")
}