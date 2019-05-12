import React, { useState } from "react";
import "./index.css";
import {useComPortsContext} from "../../useContexts/useComPortsContext.js";
import {setComPortsSettings} from "../../CoreInteraction/InteractionService.js";

const renderSpeed = speed => `${speed}бит/с`;
const compareSpeed = (x, y) => Number(x) === Number(y);
const compareName = (a, b) => String(a) === String(b);

const renderItems = (items, currentItem, compareItem, renderItem = item => item) => {
    return items.map( item => {
            return <option selected={compareItem(item, currentItem)}
                    value={item}>
                {renderItem(item)}
            </option>
        }
    );
};

const ComPortSettings = (props) => {

    const { speeds } = props;

    const { ports, inCom, outCom, setInCom, setOutCom } = useComPortsContext();

    const [ state, setState ] = useState({
        inCom,
        outCom,
    });

    const handleInputChange = (event) => {
        const { dataset, value } = event.currentTarget;
        const { field, port } = dataset;

        const changedPort = {
            ...state[port],
            [field]: value,
        };

        setState({
            ...state,
            [port]: changedPort,
        });
    };

    const apply = () => {
        const { inCom, outCom } = state;
        setInCom(inCom);
        setOutCom(outCom);

        inCom.speed = +inCom.speed;
        outCom.speed = +outCom.speed;
        setComPortsSettings({ports, inCom, outCom});
    };

    return (
        <div className={"right-column__detailed-settings"}>
            <div className="detailed-settings__com-ports">
                <form id="com-ports-form">

                    <h2 className="com-ports__title">Выберите COM-порты</h2>

                    <div>
                        <label className="com-ports__label" htmlFor="InComName">Входящий порт</label>
                        <select id={"InComName"}
                                data-port="inCom" data-field="name"
                                onChange={handleInputChange}>
                            {renderItems(ports, state.inCom.name, compareName)}
                        </select>
                        <select className="com-ports__speed"
                                data-port="inCom" data-field="speed"
                                onChange={handleInputChange}>
                            {renderItems(speeds, state.inCom.speed, compareSpeed, renderSpeed)}
                        </select>
                    </div>

                    <div>
                        <label className="com-ports__label" htmlFor="OutComName">Исходящий порт</label>
                        <select id={"OutComName"}
                                data-port="outCom" data-field="name"
                                onChange={handleInputChange}>
                            {renderItems(ports, state.outCom.name, compareName)}
                        </select>
                        <select className="com-ports__speed"
                                data-port="outCom" data-field="speed"
                                onChange={handleInputChange}>
                            {renderItems(speeds, state.outCom.speed, compareSpeed, renderSpeed)}
                        </select>
                    </div>
                    
                    <a className="com-ports__apply-button" onClick={apply}>Применить</a>
                </form>
            </div>
        </div>
    );
};


ComPortSettings.defaultProps = {
    speeds: [50, 75, 110, 150, 300, 600, 1200, 2400, 4800, 9600, 19200, 38400, 57600, 115200].reverse(),
};

export default ComPortSettings;