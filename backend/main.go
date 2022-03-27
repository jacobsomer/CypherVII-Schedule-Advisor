package main

import (
	"backend/algorithm"
	"backend/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const debug = false

// Major is the incoming JSON format for GetSchedule
type Major struct {
	Major string
}

// Course is the incoming JSON format for GetPrerequisites
type Course struct {
	Subject string
	Id      string
}

func main() {
	router := gin.Default()

	// Returns a schedule given a major class
	router.GET("/getOptimalSchedule/:schedule", GetSchedule)

	// Returns the Available Classes.
	router.GET("/getPrerequisites/:courseToSwitch", GetPrerequisites)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

// GetSchedule returns a schedule based on the input major
func GetSchedule(c *gin.Context) {
	// Add CORS headers
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000");
	c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS");

	sReq := c.Param("schedule")

	var major Major

	_ = json.Unmarshal([]byte(sReq), &major)

	schedule := algorithm.MakeSchedule(major.Major, "B.S.")

	scheduleJSON := marshalJSON(schedule)

	c.String(http.StatusOK, scheduleJSON)
}

// GetPrerequisites returns a list of course requirements
func GetPrerequisites(c *gin.Context) {
	crs := c.Param("courseToSwitch")

	var course Course // create a new Course object

	_ = json.Unmarshal([]byte(crs), &course) // unmarshal the json into the course object

	class := database.Class{
		Subj: course.Subject,
		Id:   course.Id,
	}

	classes := database.CheckPrerequisite(class) // gets the prerequisites for the class

	// print the classes to the console
	if debug {
		output := ""
		for _, cl := range classes {
			for _, cl := range cl {
				fmt.Printf("%s, %s\n", cl.Subj, cl.Id)
				output += " " + fmt.Sprintf("%s, %s\n", cl.Subj, cl.Id)
			}
		}
	}

	classesJSON := marshalJSON(classes) // marshals the object

	c.String(http.StatusOK, classesJSON) // writes the json to the page
}

// marshalJSON converts a struct into a JSON formatted string
func marshalJSON(rawStruct interface{}) string {
	marshalled, err := json.Marshal(rawStruct)
	checkErr(err)

	return string(marshalled)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
