import React from "react";
import "./index.css";


const DialogMini = (props) => {
    const { email } = props;
    return (
        <div className={"user-list__user-mini"}>
            <div className="user-mini__email">
                { email }
            </div>
        </div>
    );
};

DialogMini.defaultProps = {
    email: 'Unknown',
};

export default DialogMini;