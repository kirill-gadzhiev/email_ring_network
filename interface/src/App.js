import React, { Component } from 'react';
import LeftColumn from './components/LeftColumn';
import RightColumn from "./components/RightColumn";
import "./App.css";
import DialogMini from './components/DialogMini';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';


class App extends Component {
  render() {
    return (
      <div className="App">
          <Router>
            <LeftColumn />
            <RightColumn />
          </Router>
      </div>
    );
  }
}

export default App;
