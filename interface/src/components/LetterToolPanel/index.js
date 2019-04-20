import React from "react";
import ReactDOM from "react-dom";
import "./index.css"

class LetterToolPanel extends React.Component {
    render() {
        return (
            <div className="letter-reading__tool-panel">
                <div className="tool-panel__delete-button">Удалить</div>
            </div>
        );
    }
}

LetterToolPanel.defaultProps = {
};

export default LetterToolPanel;