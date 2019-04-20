import React from "react";
import ReactDOM from "react-dom";
import "./index.css"
import LetterToolPanel from "../LetterToolPanel";

class LetterReading extends React.Component {
    render() {
        return (
            <div className={"right-column__letter-reading"}>
                <div className="letter-reading__content">
                    <div className="letter-reading__info">
                        <div className={"letter-reading__author"}>
                            <div className="letter-reading__author__label">
                                От кого:
                            </div>
                            <div className="letter-reading__author__name">
                                {this.props.author}
                            </div>
                        </div>
                    </div>
                    <div className={"letter-reading__message"}>
                        {this.props.message}
                    </div>
                </div>
                <LetterToolPanel/>
            </div>
        );
    }
}

LetterReading.defaultProps = {
    author: 'Unknown',
    responder: 'Unknown',
    message: 'Empty',
};

export default LetterReading;