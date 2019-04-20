import React from "react";
import ReactDOM from "react-dom";
import "./index.css";


import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

class DetailedSettings extends React.Component {
    render() {
        return (
            <div className={"right-column__detailed-settings"}>
                <div>Здесь пока ничего нет :(</div>
            </div>
        );
    }
}


DetailedSettings.defaultProps = {};

export default DetailedSettings;