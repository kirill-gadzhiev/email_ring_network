import React, { useState } from "react";
import "./index.css"
import LetterWritingToolPanel from "../LetterWritingToolPanel";
import { useLettersContext } from "../../useContexts/useLettersContext.js";
import { Redirect } from 'react-router';
import { useUserContext } from "../../useContexts/useUserContext.js";
import { useNetworkContext } from "../../useContexts/useNetworkContext.js";
import InformationRightColumn from "../InformationRightColumn";

const generateID = responder => {
    const currentTimestamp = Date.now();
    const rand = Math.random() * 1000;

    const id = `${currentTimestamp}${rand}${responder}`;
    return id;
};

const LetterWriting = (props) => {
    const [state, setState] = useState({
        responder: '',
        message: '',
        redirect: false,
    });

    const { sendLetter } = useLettersContext();
    const { email: userEmail } = useUserContext();
    const { connection } = useNetworkContext();

    const onSendPressed = () => {
        const { responder, message } = state;
        if (!props.availableUsers.includes(responder)) {
            console.log('Такого пользователя нет в сети!');
            return;
        }

        const id = generateID(responder);
        const author = userEmail;
        const date = Date.now();
        const letter = {id, responder, message, author, date};

        sendLetter(letter);
        setState({
            ...state,
            redirect: true,
        });
    };

    const onInputChange = (event) => {
        const input = event.currentTarget;
        const { name } = input.dataset;
        const value = name === 'message' ? input.innerText : input.value;
        setState({
            ...state,
            [name]: value,
        });
    };

    if (state.redirect) {
        return <Redirect to={'/'}/>;
    }

    if (!connection) {
        return <InformationRightColumn message={"Отсутствует соединение с сетью"}/>;
    }

    return (
        <div className={"right-column__letter-writing"}>
            <div className="letter-writing__content">
                <label htmlFor={"new-letter-to-input"} className="letter-writing__info">
                    <div className={"letter-writing__to"}>
                        <div className="letter-writing__to__label">
                            Кому:
                        </div>
                        <input id={"new-letter-to-input"}
                               onChange={onInputChange}
                               data-name="responder"
                               className={"letter-writing__to__input"}
                               type="text" />
                    </div>
                </label>
                <div onInput={onInputChange}
                     data-name="message"
                     className={"letter-writing__message"}
                     contentEditable={true} />
            </div>
            <LetterWritingToolPanel onSendPressed={onSendPressed} />
        </div>
    );
};

LetterWriting.defaultProps = {
    availableUsers: [],
};

export default LetterWriting;