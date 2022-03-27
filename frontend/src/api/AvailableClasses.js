import {react} from "react"

function AvailableCourses(props){ 
    const params = {
       courseToSwitch: props.courseToSwitch
    }
    const response = await fetch("https://vast-sierra-57366.herokuapp.com/getAvailableCourses/"+JSON.stringify(params))
    const json = await response.json();
    return json
}

export default AvailableCourses;