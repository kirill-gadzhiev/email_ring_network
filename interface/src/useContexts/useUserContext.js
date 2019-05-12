import React, { useContext, useEffect } from 'react';
import { UserContext } from "../contexts/userContext.js";
import { setUser } from "../CoreInteraction/InteractionService.js";

export const useUserContext = () => {
    const [state, setState] = useContext(UserContext);

    function setEmail(email) {
        setState(state => ({...state, email}));
        setUser({email});
    }

    function setCollision(collision) {
        setState(state => ({...state, collision}));
    }

    return {
        ...state,
        setEmail,
        setCollision,
    }
};