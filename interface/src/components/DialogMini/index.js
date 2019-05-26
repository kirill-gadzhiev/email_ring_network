import React, { useState } from "react";
import "./index.css";
import { withRouter, Link } from 'react-router-dom';
import { Redirect } from 'react-router';
import { down, up, doubleArrow } from './svg.js';
import { useUserContext } from "../../useContexts/useUserContext";
import {useLettersContext} from "../../useContexts/useLettersContext";


const DialogMini = (props) => {
    const { id, author, message, responder, checked } = props;
    const { email: currentUser } = useUserContext();
    const { setLetterChecked, sendCheckedSubEvent } = useLettersContext();


    const toMe = responder === currentUser;
    const toIconClassName = toMe ? 'dialog-mini__icon-in' : 'dialog-mini__icon-out';
    const checkedIconClassName = checked ? 'dialog-mini__icon--checked' : 'dialog-mini__icon--unchecked';

    const onClick = () => {
        if (toMe) {
            sendCheckedSubEvent(id);
            setLetterChecked(id);
        }
        props.history.push(`/letter/${id}`)
    };

    return (
        <div onClick={onClick} className={"dialog-list__dialog-mini"}>
            <div className={toIconClassName}>
                {toMe ? down : up}
            </div>
            <div className="dialog-mini__main">
                <div className="dialog-mini__author">
                    {author}
                </div>
                <div className="dialog-mini__message">
                    {message}
                </div>
            </div>
            <div className={checkedIconClassName}>
                {doubleArrow}
            </div>
        </div>

    );
};

DialogMini.defaultProps = {
    author: 'Unknown',
    responder: 'Unknown',
    message: 'Empty',
    id: null,
};

export default withRouter(props => <DialogMini {...props}/>);