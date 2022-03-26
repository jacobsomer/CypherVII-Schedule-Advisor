import {react} from "react"

function AvailableCourses(props){ 
    const params = {
       courseToSwitch: props.courseToSwitch
    }
    const response = await fetch("http://localhost:8080/getAvailableCourses/"+JSON.stringify(params))
    const json = await response.json();
    return json
}

export default AvailableCourses;