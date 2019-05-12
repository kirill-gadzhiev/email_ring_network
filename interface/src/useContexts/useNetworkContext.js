import React, { useContext, useEffect } from 'react';
import { NetworkContext } from "../contexts/networkContext.js";


export const useNetworkContext = () => {
    const [state, setState] = useContext(NetworkContext);

    function setConnection(connection) {
        setState(state => ({...state, connection}));
    }

    function setAvailableUsers(availableUsers) {
        setState(state => ({...state, availableUsers}));
    }

    return {
        ...state,
        setConnection,
        setAvailableUsers,
    }
};