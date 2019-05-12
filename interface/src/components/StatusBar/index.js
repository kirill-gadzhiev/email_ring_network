import React from "react";
import ReactDOM from "react-dom";
import "./index.css"
import connectionSvg from './connection-icon.js'
import newLetterSvg from './new-letter-icon.js'

import { withRouter, Link } from 'react-router-dom'
import { useLettersContext } from "../../useContexts/useLettersContext.js";
import {useNetworkContext} from "../../useContexts/useNetworkContext.js";

var classNames = require('classnames');

const StatusBar = (props) => {
    const { connection } = useNetworkContext();
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

    const { findLetters } = useLettersContext();

    return (
        <div className={"left-column__status-bar"}>
            <input type="text"
                   className="status-bar__search"
                   placeholder="Поиск"
                   onChange={(e) => findLetters(e.currentTarget.value)}/>
            <div className={connectionClassNames}>
                {connectionSvg}
            </div>
            <Link to={'/letter/new'}>
                <div className={newLetterClassNames}>
                    {newLetterSvg}
                </div>
            </Link>
        </div>
    );
};

StatusBar.defaultProps = {
};

export default StatusBar;
