import { react, useState } from "react";
import Course from "./Course";
import Grid from "@material-ui/core/Grid";

export default function Semester(props) {
    const list = props.courses;
    const half = Math.ceil(list.length / 2);

    const firstHalf = list.slice(0, half);
    const secondHalf = list.slice(-half);

    return (
        <div
            style={{
                display: "flex",
                flexDirection: "row",
                maxWidth: "100%",
                position: "relative",
                // backgroundColor: "red",
                left: "50%",
                transform: "translate(-50%)",
            }}
        >
            <div
                // spacing={3}
                style={{
                    width: "100%",
                    display: "flex",
                    flexDirection: "column",
                    // backgroundColor: "#e25969",
                    position: "relative",
                    alignItems: "center",
                    paddingTop: "10px",
                    paddingBottom: "30px",
                    justifyContent: "left",
                }}
            >
                {firstHalf.map((courseData) => (
                    <Course Subj={courseData.Subj} Id={courseData.Id} />
                ))}
            </div>
            <div
                // spacing={3}
                style={{
                    width: "100%",
                    display: "flex",
                    flexDirection: "column",
                    // backgroundColor: "#e25969",
                    position: "relative",
                    alignItems: "center",
                    paddingTop: "10px",
                    justifyContent: "left",
                }}
            >
                {secondHalf.map((courseData) => (
                    <Course Subj={courseData.Subj} Id={courseData.Id} />
                ))}
            </div>
        </div>
    );
}
