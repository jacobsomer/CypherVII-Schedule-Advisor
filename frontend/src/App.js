import { react, useState } from "react";
import "./App.css";
import Questions from "./components/Questions.js";
import Login from "./components/Login.js";
import LoginError from "./components/LoginError.js";
import Planner from "./components/Planner/Planner";
import OptimalSchedule from "./api/OptimalSchedule";

export function Render() {
    const [validAddress, setValidAddress] = useState("");
    const scheduleParams = {
        major: ""
        // TODO: add previous courses taken to params
    }
    const [majorPar, setMajorPar] = useState("")

    async function handleSubmit(major) {
        switch (major) {
            case "Computer Science":
                scheduleParams.major = "CSCI"
                break
            case "CompSci":
                scheduleParams.major = "CSCI"
                break
            case "International Relations":
                scheduleParams.major = "INRL"
                break
        }
        setMajorPar(scheduleParams.major)

        if (scheduleParams.major === "CSCI" || scheduleParams.major === "INRL") {
            // submit params used to generate possible schedule
            setValidAddress("planner");
        } else {
            setValidAddress("");
        }
    }

    switch (validAddress) {
        case "valid":
            return <Questions handleSubmit={handleSubmit} />;
        case "invalid":
            return <LoginError setValidAddress={setValidAddress} />; //<InvalidEmail/>;
        case "notYetServed":
            return (
                <Login
                    validAddress={validAddress}
                    setValidAddress={setValidAddress}
                />
            );

        case "":
            return (
                <Login
                    validAddress={validAddress}
                    setValidAddress={setValidAddress}
                />
            );
        case "planner":
            console.log("pre-planner", majorPar)
            return <Planner props={majorPar}/>;
        default:
            return (
                <Login
                    validAddress={validAddress}
                    setValidAddress={setValidAddress}
                />
            );
    }
}
function App() {
    return <div className="background">{Render()}</div>;
}

export default App;
