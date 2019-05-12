import React, { useContext, useState } from 'react';
import { LettersContext } from "../contexts/lettersContext.js";
import { sendNewLetter } from "../CoreInteraction/InteractionService.js";

export const useLettersContext = () => {
    const [state, setState] = useContext(LettersContext);

    const [findParam, setFindParam] = useState({
        substring: '',
    });

    function sendLetter(letter) {
        const { letters } = state;
        letters.push(letter);
        setState(state => ({...state, letters}));
        findLetters(findParam.substring);

        sendNewLetter(letter);
    }

    const addLetter = (letter) => {
        const { letters } = state;
        letters.push(letter);
        setState(state => ({...state, letters}));
        findLetters(findParam.substring);
    };

    function deleteLetter(deleteId) {
        const { letters: oldLetters, searchedLetters: oldSearchedLetters } = state;
        const letters = oldLetters.filter(letter => letter.id !== deleteId);
        const searchedLetters = oldSearchedLetters.filter(letter => letter.id !== deleteId);

        setState(state => ({...state, letters, searchedLetters}));
    }

    function findLetters(substring) {
        setFindParam({substring});

        const { letters } = state;
        const searchedLetters = letters.filter( letter => {
            const { author, responder, message } = letter;
            const find = (str, sub) => str.toLowerCase().indexOf(sub.toLowerCase()) !== -1;

            return find(author, substring) || find(responder, substring) || find(message, substring);
        });
        setState(state => ({...state, searchedLetters}));
    }

    function setLetters(letters) {
        const searchedLetters = letters;
        setState(state => ({...state, searchedLetters, letters}));
    }

    function getLetterByID(id) {
        const { letters } = state;
        const letter = letters.find(letter => letter.id === id);
        const empty = { author: '', message: '', responder: '', id: '' };
        return letter ? letter : empty;
    }

    return {
        ...state,
        sendLetter,
        addLetter,
        deleteLetter,
        findLetters,
        setLetters,
        getLetterByID,
    }

};