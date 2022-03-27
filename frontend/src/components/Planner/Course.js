import { react, useState } from "react";

function getColor(subject) {
    if (subject == "CSCI") {
        return "#7777FF";
    }
    if (subject == "JAPN" || subject == "GRMN") {
        return "#FF66AA";
    }
    return "#BB88FF";
}

export default function Course(props) {
    const subject = props.Subj;
    const id = props.Id;
    return (
        <div
            style={{
                backgroundColor: "black",
                width: "90%",
                color: getColor(subject),
                position: "relative",
                margin: "3px",
                borderRadius: "px",
                textAlign: "center",
                padding: "0px",
                borderColor: getColor(subject),
                borderStyle: "solid",
                borderWidth: "1px",
            }}
        >
            <div style={{ padding: "5px", fontWeight: "300" }}>
                {subject} {id}
            </div>
        </div>
    );
}
