import { useState } from "react";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles((theme) => ({
    root: {
        "& .MuiFilledInput-root": {
            background: "rgb(232, 241, 250)",
        },
    },
}));

function Questions(props) {
    const [major, setMajor] = useState("");
    const classes = useStyles();
    const handleSubmit = props.handleSubmit;

    return (
        <div
            style={{
                position: "absolute",
                top: "50%",
                left: "50%",
                transform: "translate(-50%,-50%)",
            }}
        >
            <h5 style={{ width: "100%", textAlign: "center" }}>
                try "Computer Science"
            </h5>
            <TextField
                id="filled-basic"
                label="Major"
                variant="filled"
                className={classes.root}
                onChange={(e) => setMajor(e.target.value)}
            />
            <Button
                variant="contained"
                onClick={() => handleSubmit(major)}
                style={{ height: "55px", position: "relative" }}
            >
                Submit
            </Button>
        </div>
    );
}

export default Questions;
