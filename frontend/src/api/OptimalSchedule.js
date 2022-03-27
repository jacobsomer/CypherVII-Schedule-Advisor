import { react } from "react";

async function OptimalSchedule(scheduleParams) {
    const major = scheduleParams.major;

    const response = await fetch(
        "http://localhost:8080/getOptimalSchedule/" + major,
    );
    const json = await response.json();
    return json;
}

export default OptimalSchedule;
