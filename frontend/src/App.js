import { react, useState } from "react";
import "./App.css";
import Questions from "./components/Questions.js";
import Login from "./components/Login.js";

export function Render() {
    const [validAddress, setValidAddress] = useState("");
    switch (validAddress) {
        case "valid":
            console.log("hello4");
            return <Questions />;
        case "invalid":
            console.log("hello3");
            return <></>; //<InvalidEmail/>;
        case "notYetServed":
            console.log("hello2");
            return <></>; //<NotYetServed />;
        default:
            console.log("hello1");
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
