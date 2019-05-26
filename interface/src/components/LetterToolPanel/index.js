import React from "react";
import "./index.css";
import { Redirect } from 'react-router';
import { useState } from 'react';

import { useLettersContext } from "../../useContexts/useLettersContext.js";
import { Link } from 'react-router-dom';

const LetterToolPanel = (props) => {
    const { deleteLetter } = useLettersContext();

    const [state, setState] = useState({
        redirect: null,
    });

    const { id: letterID } = props;

    const onDeletePressed = () => {
        deleteLetter(letterID);
        setState({ redirect: "/" });
    };

    if (state.redirect) {
        return <Redirect to={state.redirect}/>;
    }

    return (
        <div className="letter-reading__tool-panel">
            <a onClick={onDeletePressed} className="tool-panel__delete-button">Удалить</a>
            <Link to={`/letter/new?action=forward&id=${letterID}`}>
                <a className="tool-panel__forward-button">Переслать</a>
            </Link>
            <Link to={`/letter/new?action=reply&id=${letterID}`}>
                <a className="tool-panel__reply-button">Ответить</a>
            </Link>
        </div>
    );
};

LetterToolPanel.defaultProps = {
    id: null,
};

export default LetterToolPanel;