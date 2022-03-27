package algorithm

import (
	db "backend/database"
	"fmt"
)

// Schedule is the struct format in which the scheduled classes are stored
// Each semester is a row in the struct
type Schedule [][]db.Class

// MakeSchedule returns a Schedule struct given a major and degree
func MakeSchedule(major, degree string) Schedule {
	fulfilled := make(map[db.Class]int)
	schedule := Schedule{}

	majorRequirements := db.GetMajorRequirements(major)

	generalRequirements := db.GetGeneralRequirements(degree)

	for _, class := range majorRequirements {
		fmt.Printf("Subj: %s, ID: %s \n", class.Subj, class.Id)
	}
	for _, class := range generalRequirements {
		fmt.Printf("Subj: %s, ID: %s \n", class.Subj, class.Id)
	}

	semesters := 0

	// while the requirements have entries try to place them in the schedule
	for len(majorRequirements)+len(generalRequirements) != 0 {
		semester := make([]db.Class, 0)

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
func loopOverRequirements(r *map[db.Class]db.Class, f *map[db.Class]int, s *[]db.Class, c int) int {
	semesterCredits := c

	alreadyFulfilled := make(map[db.Class]int)

	for class, _ := range *f {
		alreadyFulfilled[class] = 1
	}

	// current class. each iteration is a class in the major requirements
currentClass:
	for i, class := range *r {
		classRequirements := db.CheckPrerequisite(class)

		// meeting requirements. each iteration is over one AND requirement block
	meetingRequirements:
		for _, reqs := range classRequirements {
			if classRequirements == nil {
				break
			}

			var temp db.Class
			for _, req := range reqs {
				fmt.Printf("class: %s %s and req: %s %s \n", class.Subj, class.Id, req.Subj, req.Id)

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

		credits := db.GetCredits(class)
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
