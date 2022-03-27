import { React, useState, useEffect } from "react";
import { makeStyles, createStyles, Theme } from "@material-ui/core/styles";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Button from "@material-ui/core/Button";
import OptimalSchedule from "../../api/OptimalSchedule";
import Semester from "./Semester";

const useStyles = makeStyles((theme) =>
    createStyles({
        Semester: {
            // paddingTop: "-10px",
            backgroundColor: "white",
            padding: theme.spacing(2),
            textAlign: "center",
            color: "black",
        },
        root: {
            flexGrow: 1,
        },
    }),
);

function RenderSemester(props) {
    const schedule = props.schedule;
    const index = props.index;

    return (
        <>
            {schedule == "" ? (
                <></>
            ) : (
                <>
                    <Semester courses={schedule[index]} />
                </>
            )}
        </>
    );
}

export default function Planner(scheduleParams) {
    const classes = useStyles();

    // const [major, setMajor] = useState("Computer Science");
    const [schedule, setSchedule] = useState("");

    useEffect(() => {
        // setMajor(scheduleParams.major)
        handleClick();
    }, []);

    async function handleClick() {
        // submit params used to generate possible schedule
        // const params = {
        //     major: major,
        //     // TODO: add previous courses taken to params
        // };
        console.log("schedule params: ", scheduleParams)
        const schedule = await OptimalSchedule(scheduleParams);
        setSchedule(schedule);
    }

    return (
        <div
            style={{
                position: "absolute",
                width: "90%",
                backgroundColor: "rgba(0,0,0,.5)",
                padding: "10px",
                transform: "translate(-50%)",
                left: "50%",
                borderRadius: "10px",
                height: "fitContent",
                blockSize: "fitContent",
            }}
        >
            <div
                style={{
                    position: "relative",
                    width: "90%",
                    backgroundColor: "transparent",
                    padding: "10px",
                    transform: "translate(-50%)",
                    left: "50%",
                    borderRadius: "10px",
                }}
            >
                <Button
                    style={{ left: "50%", transform: "translate(-50%)" }}
                    variant="contained"
                    onClick={() => handleClick()}
                >
                    Generate New Schedule
                </Button>
                <Grid container spacing={3}>
                    <div
                        style={{
                            position: "relative",
                            width: "80%",
                            backgroundColor: "transparent",
                            height: "0px",
                            padding: "10px",
                        }}
                    ></div>
                    <Grid item xs={12} sm={6}>
                        <Paper className={classes.Semester}>Fall 2022</Paper>
                        <RenderSemester schedule={schedule} index={0} />
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <Paper className={classes.Semester}>Spring 2023</Paper>
                        {<RenderSemester schedule={schedule} index={1} />}
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <Paper className={classes.Semester}>Fall 2023</Paper>
                        {<RenderSemester schedule={schedule} index={2} />}
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <Paper className={classes.Semester}>Spring 2024</Paper>
                        {<RenderSemester schedule={schedule} index={3} />}
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <Paper className={classes.Semester}>Fall 2024</Paper>
                        {<RenderSemester schedule={schedule} index={4} />}
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <Paper className={classes.Semester}>Spring 2025</Paper>
                        {<RenderSemester schedule={schedule} index={5} />}
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <Paper className={classes.Semester}>Fall 2025</Paper>
                        {<RenderSemester schedule={schedule} index={6} />}
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <Paper className={classes.Semester}>Spring 2026</Paper>
                        {<RenderSemester schedule={schedule} index={7} />}
                    </Grid>
                </Grid>
            </div>
        </div>
    );
}
