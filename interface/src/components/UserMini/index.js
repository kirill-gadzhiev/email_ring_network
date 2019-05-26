import React from "react";
import "./index.css";


const UserMini = (props) => {
    const { email } = props;
    return (
        <div className={"user-list__user-mini"}>
            <div className="user-mini__email">
                { email }
            </div>
        </div>
    );
};

UserMini.defaultProps = {
    email: 'Unknown',
};

export default UserMini;