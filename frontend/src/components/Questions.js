import { useState } from "react";
import OptimalSchedule from "../api/OptimalSchedule";

function Questions(props) {
  const [major, setMajor] = useState("");

  function handleSubmit() {
    // submit params used to generate possible schedule
    const scheduleParams = {
      major: major,
      // TODO: add previous courses taken to params
    };
    OptimalSchedule(scheduleParams);
  }
  return (
    <div>
      <form onSubmit={() => handleSubmit()}>
        <label>
          What is Your Major?
          <input
            type="text"
            value={major}
            onChange={(e) => setMajor(e.target.value)}
          />
        </label>
        <input type="submit" value="Submit" />
      </form>
    </div>
  );
}

export default Questions;
