import { Link } from "react-router-dom";

function LoginError(props) {
    return (
        <div>
            <h1>You Did Not Use An Educational Email</h1>
            <p>Please use an educational email, one that ends with .edu.</p>

            <button onClick={() => props.setValidAddress("")}>Go Back</button>
        </div>
    );
}

export default LoginError;
