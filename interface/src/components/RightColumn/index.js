import React from "react";
import "./index.css";

import { Route } from 'react-router-dom';
import LetterReading from "../LetterReading";
import DetailedSettings from "../DetailedSettings";
import ComPortSettings from "../ComPortSettings";
import LetterWriting from "../LetterWriting";

const RightColumn = () => {
        return (
            <div className={"right-column"}>
                <Route path={"/letter/new"} render={() => <LetterWriting />} />
                <Route path={"/letter/:id"} component={LetterReading}/>
                <Route exact path={"/settings"} render={() => <DetailedSettings />}/>
                <Route exact path={"/settings/ports"}
                       render={() => <ComPortSettings />}/>
            </div>
        );
};


RightColumn.defaultProps = {};

export default RightColumn;