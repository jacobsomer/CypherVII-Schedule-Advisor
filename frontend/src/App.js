import { react, useState } from "react";
import "./App.css";
import Questions from "./components/Questions.js";
import Login from "./components/Login.js";
import LoginError from "./components/LoginError.js";
import Planner from "./components/Planner/Planner";
import OptimalSchedule from "./api/OptimalSchedule";

export function Render() {
    const [validAddress, setValidAddress] = useState("");
    async function handleSubmit(major) {
        if (major == "Computer Science") {
            // submit params used to generate possible schedule
            const scheduleParams = {
                major: "CompSci",
                // TODO: add previous courses taken to params
            };
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
            return <Planner />;
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
