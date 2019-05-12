import React, { useEffect } from "react";
import ReactDOM from "react-dom";
import "./index.css";

import DialogMini from '../DialogMini/index.js'
import { useLettersContext } from "../../useContexts/useLettersContext.js";

const DialogList = () => {

    const { searchedLetters } = useLettersContext();

    const cmp = (a,b) => b.date - a.date; // обратный порядок писем (новее => выше)

    return (
        <div className={"dialog-list"}>
            {searchedLetters.sort(cmp).map( letter => <DialogMini key={letter.id} {...letter}/>)}
        </div>
    );
};

DialogList.defaultProps = {
};

export default DialogList;