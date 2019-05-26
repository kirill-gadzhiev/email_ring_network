import React, {useEffect, useState} from 'react';

const defaultContext = {
    availableUsers: [],
    connection: false,
};

export const NetworkContext = React.createContext(defaultContext);

export const NetworkContextProvider = (props) => {
    const [state, setState] = useState(defaultContext);

    // useEffect(() => {
    //     // do request in GO for a user in config -> setEmail
    // }, []);

    return (
        <NetworkContext.Provider value={[state, setState]}>
            {props.children}
        </NetworkContext.Provider>
    );
};