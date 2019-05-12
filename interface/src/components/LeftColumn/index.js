import React, { useEffect } from "react";
import ReactDOM from "react-dom";
import "./index.css";
import DialogList from "../DialogList";
import SettingsColumn from '../SettingsColumn';

import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import { Redirect } from 'react-router';

import TabBar from "../TabBar";
import dialogs from "../../stories/DialogListSampleData.js";
import StatusBar from "../StatusBar";
import { useLettersContext } from "../../useContexts/useLettersContext.js";
import {useUserContext} from "../../useContexts/useUserContext.js";
import UsersList from "../UsersList";


const LeftColumn = () => {
    const { email } = useUserContext();
    if (!email) {
        return <Redirect to={'/auth'}/>;
    }

    const { setLetters } = useLettersContext();
    useEffect(() => {
        setLetters([...dialogs]);
    }, []);

    return (
        <div className={"left-column"}>
            <StatusBar />
            <Switch>
                <Route path={"/settings"} render={() => <SettingsColumn />}/>
                <Route path={"/users"} render={() => <UsersList />}/>
                <Route render={() => <DialogList/>}/>
            </Switch>
            <TabBar />
        </div>
    );
};


LeftColumn.defaultProps = {};

export default LeftColumn;