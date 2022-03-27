import { react } from "react";

async function OptimalSchedule(props) {
  const params = {
    major1: props.major1,
    major2: props.major2,
    minor1: props.minor1,
    minor2: props.minor2,
  };
  const response = await fetch(
    "http://localhost:8080/getOptimalSchedule/" + JSON.stringify(params)
  );
  const json = await response.json();
  console.log("json");
  return json;
}

export default OptimalSchedule;
