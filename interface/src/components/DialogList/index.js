import React from "react";
import ReactDOM from "react-dom";
import "./index.css";

import DialogMini from '../DialogMini/index.js'

class DialogList extends React.Component {
    render() {
        return (
            <div className={"dialog-list"}>
                {this.props.dialogs.map( dialog => <DialogMini author={dialog.author} message={dialog.message}/>)}
            </div>
        );
    }
}

DialogList.defaultProps = {
    dialogs: [],
};

export default DialogList;