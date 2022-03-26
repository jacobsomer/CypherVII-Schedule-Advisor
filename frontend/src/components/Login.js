import { react, useState } from "react";

function Login(props) {
    const setValidAddress = props.setValidAddress;
    const validAddress = props.validAddress;
    const [submitEmailEnabled, setSubmitEmailEnabled] = useState(false);
    const [email, setEmail] = useState("");
    // types
    // valid, invalid, notYetServed

    function handleSubmit(email) {
        // regex to test if email is valid
        setSubmitEmailEnabled(true);

        let re =
            /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
        var inValid = true;
        if (re.test(email)) {
            if (email.endsWith("wm.edu")) {
                setValidAddress("valid");
                inValid = false;
            } else if (email.endsWith(".edu")) {
                setValidAddress("notYetServed");
                inValid = false;
            }
        }
        if (inValid) {
            setValidAddress("invalid");
        }
    }
    return (
        <div>
            <form onSubmit={() => handleSubmit(email)}>
                <label>
                    What is Your School Email?
                    <input
                        type="text"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        disabled={submitEmailEnabled}
                    />
                </label>
                <input type="submit" value="Submit" />
            </form>
        </div>
    );
}

export default Login;
