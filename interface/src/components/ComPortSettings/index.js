import React from "react";
import ReactDOM from "react-dom";
import "./index.css";


import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

class ComPortSettings extends React.Component {
    portsSelect() {
        return this.props.portsList.map( port => <option value={port}>{port}</option>);
    }

    portsSpeeds() {
        return this.props.speeds.map( speed => <option value={speed}>{speed} бит/с</option>)
    }

    render() {
        return (
            <div className={"right-column__detailed-settings"}>
                <div className="detailed-settings__com-ports">
                    <form id="com-ports-form">

                        <h2 className="com-ports__title">Выберите COM-порты</h2>

                        <div>
                            <label className="com-ports__label" htmlFor="in-com">Входящий порт</label>
                            <select id="in-com">
                                {this.portsSelect()}
                            </select>
                            <select className="com-ports__speed" id="in-com-speed">
                                {this.portsSpeeds()}
                            </select>
                        </div>

                        <div>
                            <label className="com-ports__label" htmlFor="out-com">Исходящий порт</label>
                            <select id="out-com">
                                {this.portsSelect()}
                            </select>
                            <select className="com-ports__speed" id="out-com-speed">
                                {this.portsSpeeds()}
                            </select>
                        </div>

                        <div className="com-ports__apply-button">Применить</div>

                    </form>
                </div>
            </div>
        );
    }
}


ComPortSettings.defaultProps = {
    portsList: [],
    speeds: [50, 75, 110, 150, 300, 600, 1200, 2400, 4800, 9600, 19200, 38400, 57600, 115200].reverse(),
};

export default ComPortSettings;