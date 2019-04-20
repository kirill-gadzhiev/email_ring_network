import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import { withRouter, Link } from 'react-router-dom'

class SettingsItem extends React.Component {
    render() {
        return (
            <Link to={this.props.linkTo}>
                <div className="settings-column__settings-item">
                    {this.props.name}
                </div>
            </Link>
        );
    }
}


SettingsItem.defaultProps = {
    name: 'Unknown',
    linkTo: '/',
};

export default withRouter(props => <SettingsItem {...props}/>);