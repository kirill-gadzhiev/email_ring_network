import React from "react";
import ReactDOM from "react-dom";
import "./index.css"

class DialogMini extends React.Component {
    render() {
        return (
            <div className={"dialog-list__dialog-mini"}>
                <div className="dialog-mini__author">
                    {this.props.author}
                </div>
                <div className="dialog-mini__message">
                    {this.props.message}
                </div>
            </div>
        );
    }
}

DialogMini.defaultProps = {
    author: 'Unknown',
    message: 'Empty',
};

export default DialogMini;