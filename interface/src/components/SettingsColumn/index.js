import React from "react";
import ReactDOM from "react-dom";
import "./index.css";

import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import SettingsItem from "../SettingsItem";

import settings from './SettingsData.js';

class SettingsColumn extends React.Component {
    render() {
        return (
            <div class="left-column__settings-column">
                {this.props.settingsItems.map(item => <SettingsItem {...item}/>)}
            </div>
        );
    }
}


SettingsColumn.defaultProps = {
    settingsItems: settings,
};

export default SettingsColumn;