import React from "react";
import "./index.css"
import  {settingsSvg, dialogsSvg, usersSvg} from './svg.js'

import { withRouter, Link } from 'react-router-dom'

var classNames = require('classnames');

class TabBar extends React.Component {
    render() {
        const { pathname } = this.props.location;
        const isUsers = pathname.match(/\/users/gi);
        const isSettings = pathname.match(/\/settings/gi);
        const isDialogs = !(isSettings || isUsers);  // mr de morgan dobriy den

        const usersClassNames = classNames({
            "tab-bar__users-tab": true,
            "tab-bar__tab--checked": isUsers,
        });
        const dialogsClassNames = classNames({
            "tab-bar__dialogs-tab": true,
            "tab-bar__tab--checked": isDialogs,
        });
        const settingsClassNames = classNames({
            "tab-bar__settings-tab": true,
            "tab-bar__tab--checked": isSettings,
        });

        return (
            <div className={"left-column__tab-bar"}>
                <Link to={'/users'}>
                    <div className={usersClassNames}>
                        {usersSvg}
                    </div>
                </Link>
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
