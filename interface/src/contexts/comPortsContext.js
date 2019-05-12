import React, {useState} from 'react';


const defaultSpeed = 115200;
const defaultContext = {
    inCom: {
        name: 'COM1',
        speed: defaultSpeed,
    },
    outCom: {
        name: 'COM1',
        speed: defaultSpeed,
    },
    ports: ['COM1', 'COM2', 'COM3', 'COM4', 'COM5'],
};

export const ComPortsContext = React.createContext(defaultContext);

export const ComPortsContextProvider = (props) => {
    const [state, setState] = useState(defaultContext);

    return (
        <ComPortsContext.Provider value={[state, setState]}>
            {props.children}
        </ComPortsContext.Provider>
    );
};



