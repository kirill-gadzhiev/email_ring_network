import { ws, EVENT_TYPES } from "./ws.js";

export const sendNewLetter = letter => {
    ws.send(JSON.stringify({
        "event": EVENT_TYPES.MESSAGE_SEND,
        "letter": letter,
    }));
};

export const setUser = user => {
    ws.send(JSON.stringify({
        "event": EVENT_TYPES.SET_USER,
        "user": user,
    }));
};

export const sendInterfaceReady = () => {
    ws.send(JSON.stringify({
        "event": EVENT_TYPES.INTERFACE_READY,
    }));
};

export const setComPortsSettings = comPortsSettings => {
    ws.send(JSON.stringify({
        "event": EVENT_TYPES.COM_PORTS_SETTINGS_CHANGE,
        "comPortsSettings": comPortsSettings,
    }));
};

export const sendCloseWindow = () => {
    ws.send(JSON.stringify({
        "event": EVENT_TYPES.CLOSE_WINDOW,
    }));
}


class EventBus {
    constructor() {
        this.listeners = {};
    }

    on(event, callback) {
        if (!this.listeners[event]) {
            this.listeners[event] = [];
        }
        this.listeners[event].push(callback);
    }

    emit(event, data) {
        if (this.listeners[event]) {
            this.listeners[event].forEach(listener => {
                listener(data);
            });
        }
    }

    unsubscribeAll(event) {
        this.listeners[event] = [];
    }
}

const bus = new EventBus();

ws.onmessage = (message) => {
    const data = JSON.parse(message.data);
    if (!data.event) return;

    console.log('RECEIVED: ', data.event, "data: ", data);
    bus.emit(data.event, data);
};
ws.onopen = () => sendInterfaceReady();


export default bus;

