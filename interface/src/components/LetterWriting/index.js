import React, { useState, useEffect } from "react";
import "./index.css"
import LetterWritingToolPanel from "../LetterWritingToolPanel";
import { useLettersContext } from "../../useContexts/useLettersContext.js";
import { Redirect } from 'react-router';
import { useUserContext } from "../../useContexts/useUserContext.js";
import { useNetworkContext } from "../../useContexts/useNetworkContext.js";
import InformationRightColumn from "../InformationRightColumn";
import queryString from "query-string";
import { withRouter, Link } from 'react-router-dom'


const generateID = responder => {
    const currentTimestamp = Date.now();
    const rand = Math.random() * 1000;

    const id = `${currentTimestamp}${rand}${responder}`;
    return id;
};

const LetterWriting = (props) => {
    const { sendLetter, getLetterByID } = useLettersContext();

    const defaultState = {
        responder: '',
        message: '',
        redirect: null,
    };

    const [state, setState] = useState(defaultState);

    useEffect( () => {
        const { search } = props.location;
        const { action, id } = queryString.parse(search);
        if (action && id) {
            const originFieldName = action === 'reply' ? 'author' : 'message';
            const newFieldName = action === 'reply' ? 'responder' : 'message';
            const originLetter = getLetterByID(id);
            console.log(originLetter[originFieldName]);
            setState(state => ({
                ...state,
                [newFieldName]: originLetter[originFieldName],
            }));
        }
    }, []);

    const { email: userEmail } = useUserContext();
    const { connection, availableUsers } = useNetworkContext();

    const onSendPressed = () => {
        const { responder, message } = state;
        const availableUserEmails = availableUsers.map( user => user.email);
        if (!availableUserEmails.includes(responder)) {
            console.log('Такого пользователя нет в сети!', availableUsers, responder);
            return;
        }

        const id = generateID(responder);
        const author = userEmail;
        const date = Date.now();
        const checkedSubEvent = false;
        const letter = {id, responder, message, author, date, checkedSubEvent};

        sendLetter(letter);
        setState({
            ...state,
            redirect: true,
        });
    };

    const onInputChange = (event) => {
        const input = event.currentTarget;
        const { name } = input.dataset;
        const value = input.value;
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
                               value={state.responder}
                               data-name="responder"
                               className={"letter-writing__to__input"}
                               type="text"
                        />
                    </div>
                </label>
                <textarea onChange={onInputChange}
                     data-name="message"
                     className={"letter-writing__message"}
                     value={state.message}
                     type="text"
                />
            </div>
            <LetterWritingToolPanel onSendPressed={onSendPressed} />
        </div>
    );
};

LetterWriting.defaultProps = {
};

export default withRouter( props => <LetterWriting {...props} />);