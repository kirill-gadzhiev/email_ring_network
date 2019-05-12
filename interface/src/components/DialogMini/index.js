import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import { withRouter, Link } from 'react-router-dom'


const DialogMini = (props) => {
    const { id, author, message} = props;
    return (
        <Link to={`/letter/${id}`}>
            <div className={"dialog-list__dialog-mini"}>
                <div className="dialog-mini__author">
                    {author}
                </div>
                <div className="dialog-mini__message">
                    {message}
                </div>
            </div>
        </Link>
    );
};

DialogMini.defaultProps = {
    author: 'Unknown',
    message: 'Empty',
    id: null,
};

export default withRouter(props => <DialogMini {...props}/>);