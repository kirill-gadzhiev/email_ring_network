import React from "react";
import "./index.css"


const InformationRightColumn = (props) => {
    const { message } = props;
    return (
        <div className={"right-column__information"}>
            {message}
        </div>
    );
};

InformationRightColumn.defaultProps = {
    message: 'Недоступно',
};

export default InformationRightColumn;