import React from "react";
import { makeStyles, createStyles, Theme } from "@material-ui/core/styles";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";

const useStyles = makeStyles((theme) =>
    createStyles({
        Semester: {
            // paddingTop: "-10px",
            padding: theme.spacing(2),
            textAlign: "center",
            color: theme.palette.text.secondary,
        },
        root: {
            flexGrow: 1,
        },
    }),
);

export default function Planner() {
    const classes = useStyles();

    return (
        <div
            style={{
                position: "absolute",
                width: "90%",
                backgroundColor: "rgba(0,0,0,.7)",
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
                <Grid container spacing={3}>
                    <div
                        style={{
                            position: "relative",
                            width: "80%",
                            backgroundColor: "white",
                            height: "23px",
                            padding: "40px",
                            transform: "translate(-50%)",
                            left: "50%",
                            top: "20px",
                            borderRadius: "5px",
                        }}
                    >
                        <h2>Schedule</h2>
                    </div>
                    <div
                        style={{
                            position: "relative",
                            width: "80%",
                            backgroundColor: "transparent",
                            height: "0px",
                            padding: "20px",
                        }}
                    ></div>
                    <Grid item xs={12} sm={6}>
                        <Paper className={classes.Semester}>Fall 2022</Paper>
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <Paper className={classes.Semester}>Spring 2023</Paper>
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <Paper className={classes.Semester}>Fall 2023</Paper>
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <Paper className={classes.Semester}>Spring 2024</Paper>
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <Paper className={classes.Semester}>Fall 2024</Paper>
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <Paper className={classes.Semester}>Spring 2025</Paper>
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <Paper className={classes.Semester}>Fall 2025</Paper>
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <Paper className={classes.Semester}>Spring 2026</Paper>
                    </Grid>
                </Grid>
            </div>
        </div>
    );
}
