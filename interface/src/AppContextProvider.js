import React from "react";
import { UserContextProvider } from "./contexts/userContext.js";
import { LettersContextProvider } from "./contexts/lettersContext.js";
import { ComPortsContextProvider } from "./contexts/comPortsContext.js";
import { NetworkContextProvider } from "./contexts/networkContext.js";

export const AppContextProvider = (props) => {
    return (
        <NetworkContextProvider>
            <UserContextProvider>
                <LettersContextProvider>
                    <ComPortsContextProvider>
                        {props.children}
                    </ComPortsContextProvider>
                </LettersContextProvider>
            </UserContextProvider>
        </NetworkContextProvider>
    );
};