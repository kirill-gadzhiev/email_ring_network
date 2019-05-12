import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import { Link } from 'react-router-dom';


const LetterWritingToolPanel = (props) => {
    return (
        <div className="letter-writing__tool-panel">
            <a className="tool-panel__send-button" onClick={props.onSendPressed}>Отправить</a>
            <a className="tool-panel__cancel-button">
                <Link to={'/'}>Отмена</Link>
            </a>
        </div>
    );
};

LetterWritingToolPanel.defaultProps = {
    onSendPressed: () => null,
};

export default LetterWritingToolPanel;