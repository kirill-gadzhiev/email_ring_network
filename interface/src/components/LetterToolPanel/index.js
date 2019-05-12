import React from "react";
import "./index.css";
import { Redirect } from 'react-router';
import { useState } from 'react';

import { useLettersContext } from "../../useContexts/useLettersContext.js";


const LetterToolPanel = (props) => {
    const { deleteLetter } = useLettersContext();

    const [state, setState] = useState({
        redirect: false,
    });

    const onDeletePressed = () => {
        deleteLetter(props.id);
        setState({ redirect: true });
    };

    if (state.redirect) {
        return <Redirect to={'/'}/>;
    }

    return (
        <div className="letter-reading__tool-panel">
            <a onClick={() => onDeletePressed()} className="tool-panel__delete-button">Удалить</a>
        </div>
    );
};

LetterToolPanel.defaultProps = {
    id: null,
};

export default LetterToolPanel;