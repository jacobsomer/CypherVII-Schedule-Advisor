import { react, useState } from "react";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import InputBase from "@material-ui/core/InputBase";

import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles((theme) => ({
    root: {
        "& .MuiFilledInput-root": {
            background: "rgb(232, 241, 250)",
        },
    },
}));

function Login(props) {
    const setValidAddress = props.setValidAddress;
    const validAddress = props.validAddress;
    const [submitEmailEnabled, setSubmitEmailEnabled] = useState(false);
    const [email, setEmail] = useState("");
    // types
    // valid, invalid, notYetServed

    const classes = useStyles();

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
        <div
            style={{
                position: "absolute",
                top: "50%",
                left: "50%",
                transform: "translate(-50%,-50%)",
            }}
        >
            <h3 style={{ width: "100%", textAlign: "center" }}>
                Major Advisor
            </h3>
            <h3 style={{ width: "100%", textAlign: "center" }}>
                DegreeWorks but better
            </h3>
            <TextField
                id="filled-basic"
                label="School email"
                variant="filled"
                className={classes.root}
                onChange={(e) => setEmail(e.target.value)}
            />
            <Button
                variant="contained"
                onClick={() => handleSubmit(email)}
                style={{ height: "55px", position: "relative" }}
            >
                Submit
            </Button>
        </div>
    );
}

export default Login;
