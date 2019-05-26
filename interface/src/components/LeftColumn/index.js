import React from "react";
import "./index.css";
import DialogList from "../DialogList";
import SettingsColumn from '../SettingsColumn';

import { Route, Switch } from 'react-router-dom';
import { Redirect } from 'react-router';

import TabBar from "../TabBar";
import StatusBar from "../StatusBar";
import {useUserContext} from "../../useContexts/useUserContext.js";
import UsersList from "../UsersList";


const LeftColumn = () => {
    const { email } = useUserContext();
    if (!email) {
        return <Redirect to={'/auth'}/>;
    }

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