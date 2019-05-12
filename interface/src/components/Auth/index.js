import React, { useState, useEffect } from "react";
import "./index.css";
import { useUserContext } from "../../useContexts/useUserContext.js";
import { Redirect } from "react-router";
import {EVENT_TYPES} from "../../CoreInteraction/ws.js";
import bus from "../../CoreInteraction/InteractionService.js"

const Auth = () => {
    const { email, setEmail, collision, setCollision } = useUserContext();

    const [inputValue, setInputValue] = useState(email);

    const [redirect, setRedirect] = useState(false);

    useEffect(() => {
        bus.on(EVENT_TYPES.SET_USER_RESPONSE, (data) => {
            if (data.status) {
                setRedirect(true);
            } else {
                setCollision(true);
            }
        });
    }, []);

    const onApply = () => {
        if (inputValue) {
            setEmail(inputValue);
        }
    };

    const onChange = (event) => {
        const { value } = event.currentTarget;
        setInputValue(value);
    };

    if (redirect) {
        return <Redirect to={"/"} />;
    }

    const headerText = collision ? "Такой пользователь уже находится в сети. Введите другой email" : "Авторизация";

    return (
        <div className="auth">
            <h2 className="auth__header">{headerText}</h2>
            <input onChange={onChange} className="auth__email-input" type="text" placeholder="Email" value={inputValue}/>
            <a onClick={onApply} className="auth__submit-button">Применить</a>
        </div>
    );
};


Auth.defaultProps = {};

export default Auth;