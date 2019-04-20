import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import DialogList from "../DialogList";
import SettingsColumn from '../SettingsColumn';

import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import TabBar from "../TabBar";
import dialogs from "../../stories/DialogListSampleData.js";
import StatusBar from "../StatusBar";


class LeftColumn extends React.Component {
    render() {
        return (
            <div className={"left-column"}>
                <StatusBar connection={true}/>
                <Switch>
                    <Route path={"/settings"} render={() => <SettingsColumn />}/>
                    <Route render={() => <DialogList dialogs={[...dialogs, ...dialogs, ...dialogs]}/>}/>
                </Switch>
                <TabBar />
            </div>
        );
    }
}


LeftColumn.defaultProps = {};

export default LeftColumn;