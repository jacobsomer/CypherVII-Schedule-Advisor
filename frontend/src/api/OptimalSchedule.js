import { react } from "react";

async function OptimalSchedule(props) {
    const ma = {
        major1: props.major1,
        major2: props.major2,
        minor1: props.minor1,
        minor2: props.minor2,
    };
    const response = await fetch(
        "http://localhost:8080/getOptimalSchedule/" + major,
    );
    const json = await response.json();
    return json;
}

export default OptimalSchedule;
