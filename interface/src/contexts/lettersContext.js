import React, {useState} from 'react';
import dialogs from "../stories/DialogListSampleData.js";


const defaultContext = {
    letters: [],
    searchedLetters: [],
};

export const LettersContext = React.createContext(defaultContext);

export const LettersContextProvider = (props) => {
    const [state, setState] = useState(defaultContext);

    return (
        <LettersContext.Provider value={[state, setState]}>
            {props.children}
        </LettersContext.Provider>
    );
};