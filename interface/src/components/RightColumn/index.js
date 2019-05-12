import React from "react";
import ReactDOM from "react-dom";
import "./index.css";

import {BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import LetterReading from "../LetterReading";
import DetailedSettings from "../DetailedSettings";
import ComPortSettings from "../ComPortSettings";
import LetterWriting from "../LetterWriting";

const RightColumn = () => {
        return (
            <div className={"right-column"}>
                <Route exact path={"/letter/new"} render={() => <LetterWriting availableUsers={['hello55@mail.ru']}/>} />
                <Route path={"/letter/:id"} component={LetterReading}/>
                <Route exact path={"/settings"} render={() => <DetailedSettings />}/>
                <Route exact path={"/settings/ports"}
                       render={() => <ComPortSettings />}/>
            </div>
        );
};


RightColumn.defaultProps = {};

export default RightColumn;