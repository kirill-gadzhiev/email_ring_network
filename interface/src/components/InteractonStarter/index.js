import React, {useEffect} from 'react';
import {useNetworkContext} from "../../useContexts/useNetworkContext.js";
import bus from '../../CoreInteraction/InteractionService.js';
import { EVENT_TYPES } from "../../CoreInteraction/ws.js";
import {useLettersContext} from "../../useContexts/useLettersContext.js";
import {useComPortsContext} from "../../useContexts/useComPortsContext.js";

const InteractionStarter = () => {
    const { setAvailableUsers } = useNetworkContext();
    const { letters, addLetter } = useLettersContext();
    const { inCom, outCom, ports, mergeState } = useComPortsContext();

    useEffect( () => {
        bus.on(EVENT_TYPES.NETWORK_STATUS, data => {
            setAvailableUsers(data.networkStatus.availableUsers);
        });
    }, []);

    useEffect( () => {
        const event = EVENT_TYPES.MESSAGE_RECEIVED;
        bus.unsubscribeAll(event);
        bus.on(event, data => {
            addLetter(data.letter);
        });
    }, [letters]);

    useEffect( () => {
        const event = EVENT_TYPES.COM_PORTS_SETTINGS;
        bus.unsubscribeAll(event);
        bus.on(event, data => {
            console.log("RECEIVED COM-ports UPDATE: ", data.comPortsSettings);
            mergeState(data.comPortsSettings);
        });
    }, [inCom, outCom, ports]);

    return null;
};

export default InteractionStarter;