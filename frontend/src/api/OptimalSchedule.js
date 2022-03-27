async function OptimalSchedule(props) {

    // turn the parameters into the right format for the JSON
    const scheduleParams = {
        major: props.props
    }

    const response = await fetch(
        "http://localhost:8080/getOptimalSchedule/" + JSON.stringify(scheduleParams),
    );
    return await response.json();
}

export default OptimalSchedule;
