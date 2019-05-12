import React from 'react';
import LeftColumn from './components/LeftColumn';
import RightColumn from "./components/RightColumn";
import Auth from "./components/Auth";
import "./App.css";
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import {AppContextProvider} from "./AppContextProvider.js";
import InteractionStarter from "./components/InteractonStarter";

const App = () => {
    return (
        <AppContextProvider>
            <InteractionStarter />
            <div className="App">
                <Router>
                    <Switch>
                        <Route path={"/auth"} render={() => <Auth />}/>
                        <Route render={() => (
                            <>
                                <LeftColumn />
                                <RightColumn />
                            </>
                        )}/>
                    </Switch>
                </Router>
            </div>
        </AppContextProvider>
    );
};

export default App;
