import { react, useState } from "react";
import "./App.css";
import Questions from "./components/Questions.js";
import Login from "./components/Login.js";
import LoginError from "./components/LoginError.js";
import Planner from "./components/Planner/Planner";

export function Render() {
    const [validAddress, setValidAddress] = useState("");
    switch (validAddress) {
        case "valid":
            return <Questions />;
        case "invalid":
            return <LoginError setValidAddress={setValidAddress} />; //<InvalidEmail/>;
        case "notYetServed":
            return <></>; //<NotYetServed />;

        case "":
            return (
                <Login
                    validAddress={validAddress}
                    setValidAddress={setValidAddress}
                />
            );
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
    return (
        <div className="background">
            {/* {Render()} */}
            <Planner />
        </div>
    );
}

export default App;
