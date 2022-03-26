package algorithm

import (
	db "backend/database"
	"fmt"
)

type Schedule [][]db.Class

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
		loopOverRequirements(&majorRequirements, &fulfilled, &semester)

		// loop over the general requirements and try to fill what can be taken
		loopOverRequirements(&generalRequirements, &fulfilled, &semester)

		// append the semester to the schedule
		schedule = append(schedule, semester)
		semesters++
		if semesters > 8 {
			break
		}
	}
	return schedule
}

func loopOverRequirements(r *map[int]db.Class, f *map[db.Class]int, s *[]db.Class) {
	currentlyFulfilled := *f
	semesterCredits := 0

	// current class. each iteration is a class in the major requirements
currentClass:
	for i, class := range *r {
		classRequirements := db.CheckPrerequisite(class)

		// meeting requirements. each iteration is over one AND requirement block
	meetingRequirements:
		for _, reqs := range classRequirements {

			for _, req := range reqs {
				fmt.Printf("class: %s %s and req: %s %s \n", class.Subj, class.Id, req.Subj, req.Id)

				if _, ok := (currentlyFulfilled)[req]; ok {

					// if a requirement is found to be fulfilled in the block the meeting requirement loop is continued
					continue meetingRequirements
				}

			}
			// if one of the AND blocks did not have any fulfilled requirements, the meetingRequirements loop is not continued
			// that means this code is reached
			continue currentClass
		}

		// if the credit limit is exceeded, continue the loop
		if semesterCredits+db.GetCredits(class) > 18 {
			continue
		}

		// if the meetingRequirements loop finished with a breakout this code is reached
		*s = append(*s, class) // append the semester with the class
		(*f)[class] = 1        // make an entry for the class in the fulfilled map
		delete(*r, i)          // delete the majorRequirement from the required map

	}
}
