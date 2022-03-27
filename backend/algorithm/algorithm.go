package algorithm

import (
	"backend/database"
	"fmt"
)

const debug = false

// Schedule is the struct format in which the scheduled classes are stored
// Each semester is a row in the struct
type Schedule [][]database.Class

// MakeSchedule returns a Schedule struct given a major and degree
func MakeSchedule(major, degree string) Schedule {
	fulfilled := make(map[database.Class]int)
	schedule := Schedule{}

	majorRequirements := database.GetMajorRequirements(major)

	generalRequirements := database.GetGeneralRequirements(degree)

	// print the course requirements when debug is true
	if debug {
		for _, class := range majorRequirements {
			fmt.Printf("Subj: %s, ID: %s \n", class.Subj, class.Id)
		}
		for _, class := range generalRequirements {
			fmt.Printf("Subj: %s, ID: %s \n", class.Subj, class.Id)
		}
	}

	semesters := 0

	// while the requirements have entries try to place them in the schedule
	for len(majorRequirements)+len(generalRequirements) != 0 {
		semester := make([]database.Class, 0)

		// loop over the major requirements first and try to fill as much as possible
		credits := loopOverRequirements(&majorRequirements, &fulfilled, &semester, 0)

		// loop over the general requirements and try to fill what can be taken
		loopOverRequirements(&generalRequirements, &fulfilled, &semester, credits)

		// append the semester to the schedule
		schedule = append(schedule, semester)
		semesters++
		if semesters > 8 {
			break
		}
	}
	return schedule
}

// loopOverRequirements takes pointers to required classes, fulfilled classes, a slice of semester classes,
// and semester credits, and returns the semester credits
func loopOverRequirements(r *map[database.Class]database.Class, f *map[database.Class]int, s *[]database.Class, c int) int {
	semesterCredits := c

	alreadyFulfilled := make(map[database.Class]int)

	for class, _ := range *f {
		alreadyFulfilled[class] = 1
	}

	// current class. each iteration is a class in the major requirements
currentClass:
	for i, class := range *r {
		classRequirements := database.CheckPrerequisite(class)

		// meeting requirements. each iteration is over one AND requirement block
	meetingRequirements:
		for _, reqs := range classRequirements {
			if classRequirements == nil {
				break
			}

			var temp database.Class
			for _, req := range reqs {
				if debug {
					fmt.Printf("class: %s %s and req: %s %s \n", class.Subj, class.Id, req.Subj, req.Id)
				}

				_, ok := (alreadyFulfilled)[req]
				if ok {

					// if a requirement is found to be fulfilled in the block the meeting requirement loop is continued
					continue meetingRequirements
				}

				temp = req
			}

			if _, ok := (*r)[temp]; !ok {
				(*r)[temp] = temp
			}

			// if one of the AND blocks did not have any fulfilled requirements, the meetingRequirements loop is not continued
			// that means this code is reached
			continue currentClass
		}

		//fmt.Printf("all requirements met for %s %s \n", class.Subj, class.Id)

		credits := database.GetCredits(class)
		if credits == 0 {
			credits = 3
		}

		// if the credit limit is exceeded, continue the loop
		if semesterCredits+credits > 15 {
			continue
		}

		if _, ok := (*f)[class]; !ok {
			// if the meetingRequirements loop finished with a breakout this code is reached
			semesterCredits += credits // update the credits for the semester
			*s = append(*s, class)     // append the semester with the class
			(*f)[class] = 1            // make an entry for the class in the fulfilled map
			delete(*r, i)              // delete the majorRequirement from the required map
		}

	}
	return semesterCredits
}
