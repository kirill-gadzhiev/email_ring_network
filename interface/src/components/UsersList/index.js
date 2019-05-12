import React, { useEffect } from "react";
import "./index.css";

import UserMini from '../UserMini';
import { useNetworkContext } from "../../useContexts/useNetworkContext.js";

const UsersList = (props) => {

    const { availableUsers } = useNetworkContext();

    return (
        <div className={"user-list"}>
            {availableUsers.map( user => <UserMini key={user.email} {...user}/>)}
        </div>
    );
};

UsersList.defaultProps = {
};

export default UsersList;