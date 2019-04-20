import React from "react";
import ReactDOM from "react-dom";
import "./index.css"
import connectionSvg from './connection-icon.js'
import newLetterSvg from './new-letter-icon.js'

import { withRouter, Link } from 'react-router-dom'

var classNames = require('classnames');

class StatusBar extends React.Component {
    render() {
        const { connection } = this.props;
        const connectionClassNames = classNames({
            "status-bar__button": true,
            "status-bar__connection--connected": connection,
            "status-bar__connection--disconnected": !connection,
        });
        const newLetterClassNames = classNames({
            "status-bar__button": true,
            "status-bar__new-letter--connected": connection,
            "status-bar__new-letter--disconnected": !connection,
        });

        return (
            <div className={"left-column__status-bar"}>
                <input type="text" className="status-bar__search" placeholder="Поиск" />
                <div className={connectionClassNames}>
                    {connectionSvg}
                </div>
                <div className={newLetterClassNames}>
                    {newLetterSvg}
                </div>
            </div>
        );
    }
}

StatusBar.defaultProps = {
    connection: false,
};

export default StatusBar;
