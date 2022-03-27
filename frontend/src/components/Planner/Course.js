import { react, useState } from "react";

export default function Semester(props) {
    const classes = props.classes;

    const Schedule = ({ classes }) => (
        <>
            {courses.map((courseData) => (
                <Course courseData={courseData} />
            ))}
        </>
    );

    return <Schedule />;
}
