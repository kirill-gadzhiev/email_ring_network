import React from "react";
import ReactDOM from "react-dom";
import "./index.css";

import {BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import LetterReading from "../LetterReading";
import DetailedSettings from "../DetailedSettings";
import ComPortSettings from "../ComPortSettings";

class RightColumn extends React.Component {
    render() {
        return (
            <div className={"right-column"}>
                <Route exact path={"/"} render={() => <LetterReading message={"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."} />}/>
                <Route exact path={"/settings"} render={() => <DetailedSettings />}/>
                <Route exact path={"/settings/ports"}
                       render={() => <ComPortSettings portsList={['COM1', 'COM2', 'COM3', 'COM4']} />}/>
            </div>
        );
    }
}


RightColumn.defaultProps = {};

export default RightColumn;