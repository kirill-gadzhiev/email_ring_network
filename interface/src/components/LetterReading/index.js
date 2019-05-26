import React from "react";
import "./index.css"
import LetterToolPanel from "../LetterToolPanel";
import { useLettersContext } from "../../useContexts/useLettersContext.js";
import { useUserContext } from "../../useContexts/useUserContext.js";

const LetterReading = (props) => {
    const { getLetterByID } = useLettersContext();

    const id = props.match.params.id;
    const { author, message, responder } = getLetterByID(id);

    const { email: userEmail } = useUserContext();
    const isCurrentUserAuthor = userEmail === author;
    const authorLabel = isCurrentUserAuthor ? "Кому:" : "От кого:";
    const authorValue = isCurrentUserAuthor ? responder : author;

    return (
        <div className={"right-column__letter-reading"}>
            <div className="letter-reading__content">
                <div className="letter-reading__info">
                    <div className={"letter-reading__author"}>
                        <div className="letter-reading__author__label">
                            {authorLabel}
                        </div>
                        <div className="letter-reading__author__name">
                            {authorValue}
                        </div>
                    </div>
                </div>
                <div className={"letter-reading__message"}>
                    {message}
                </div>
            </div>
            <LetterToolPanel id={id}/>
        </div>
    );
};

LetterReading.defaultProps = {
};

export default LetterReading;