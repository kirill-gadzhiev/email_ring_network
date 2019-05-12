import React, { useContext } from 'react';
import { ComPortsContext } from "../contexts/comPortsContext.js";

export const useComPortsContext = () => {
    const [state, setState] = useContext(ComPortsContext);

    function setInCom(inCom) {
        setState(state => ({...state, inCom}));
    }

    function setOutCom(outCom) {
        setState(state => ({...state, outCom}));
    }

    function setPorts(ports) {
        setState(state => ({...state, ports}));
    }

    function mergeState(newState) {
        setState(state => ({...state, ...newState}));
    }

    return {
        ...state,
        setInCom,
        setOutCom,
        setPorts,
        mergeState,
    }

};