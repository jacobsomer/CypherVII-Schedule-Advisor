import { react, useState } from "react";
import "./App.css";
import Questions from "./components/Questions.js";
import Login from "./components/Login.js";

export function Render() {
  const [validAddress, setValidAddress] = useState("");
  switch (validAddress) {
    case "valid":
      return <Questions />;
    case "invalid":
      return <></>; //<InvalidEmail/>;
    case "notYetServed":
      return <></>; //<NotYetServed />;
    default:
      console.log("hello1");
      return (
        <Login validAddress={validAddress} setValidAddress={setValidAddress} />
      );
  }
}
function App() {
  return <div className="background">{Render()}</div>;
}

export default App;
