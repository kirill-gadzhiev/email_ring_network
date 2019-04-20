import React from "react";
import ReactDOM from "react-dom";
import "./index.css"
import  {settingsSvg, dialogsSvg} from './svg.js'

import { withRouter, Link } from 'react-router-dom'

var classNames = require('classnames');

class TabBar extends React.Component {
    render() {
        const { pathname } = this.props.location;
        const isSettings = pathname.match(/\/settings/gi);
        const dialogsClassNames = classNames({
            "tab-bar__dialogs-tab": true,
            "tab-bar__tab--checked": !isSettings,
        });
        const settingsClassNames = classNames({
            "tab-bar__settings-tab": true,
            "tab-bar__tab--checked": isSettings,
        });

        return (
            <div className={"left-column__tab-bar"}>
                <Link to={'/'}>
                    <div className={dialogsClassNames}>
                        {dialogsSvg}
                    </div>
                </Link>
                <Link to={'/settings'}>
                    <div className={settingsClassNames}>
                        {settingsSvg}
                    </div>
                </Link>
            </div>
        );
    }
}

TabBar.defaultProps = {
};

export default withRouter(props => <TabBar {...props}/>);
// export default TabBar;
