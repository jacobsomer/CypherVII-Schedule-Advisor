import { react, useState } from "react";
import OptimalSchedule from "../api/OptimalSchedule";

function Questions(props) {
    const [major1, setMajor1] = useState(0);
    const [major2, setMajor2] = useState(0);
    const [minor1, setMinor1] = useState(0);
    const [minor2, setMinor2] = useState(0);

    function handleSubmit() {
        // submit params used to generate possible schedule
        const scheduleParams = {
            major1: major1,
            major2: major2,
            minor1: minor1,
            minor2: minor2,
            // TODO: add previous courses taken to params
        };
        OptimalSchedule(scheduleParams);
    }
    return (
        <div>
            <form onSubmit={() => handleSubmit()}>
                <label>
                    Major1:
                    <input
                        type="text"
                        value={major1}
                        onChange={(e) => setMajor1(e.target.value)}
                    />
                </label>
                <label>
                    Major2:
                    <input
                        type="text"
                        value={major2}
                        onChange={(e) => setMajor2(e.target.value)}
                    />
                </label>
                <label>
                    Minor1:
                    <input
                        type="text"
                        value={minor1}
                        onChange={(e) => setMinor1(e.target.value)}
                    />
                </label>
                <label>
                    Minor2:
                    <input
                        type="text"
                        value={minor2}
                        onChange={(e) => setMinor2(e.target.value)}
                    />
                </label>
                <input type="submit" value="Submit" />
            </form>
        </div>
    );
}

export default Questions;
