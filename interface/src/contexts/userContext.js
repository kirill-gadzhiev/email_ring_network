import React, {useEffect, useState} from 'react';

const defaultContext = {
    email: '',
    collision: false,
};

export const UserContext = React.createContext(defaultContext);

export const UserContextProvider = (props) => {
    const [state, setState] = useState(defaultContext);

    useEffect(() => {
        // do request in GO for a user in config -> setEmail
    }, []);

    return (
        <UserContext.Provider value={[state, setState]}>
            {props.children}
        </UserContext.Provider>
    );
};