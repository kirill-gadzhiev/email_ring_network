import React from 'react';
import '../index.css';

import { storiesOf } from '@storybook/react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import DialogMini from '../components/DialogMini';
import DialogList from '../components/DialogList';
import TabBar from '../components/TabBar'
import SettingsColumn from "../components/SettingsColumn";
import LeftColumn from "../components/LeftColumn";
import LetterReading from "../components/LetterReading";

import letter from './LetterSampleData.js';
import dialogs from './DialogListSampleData.js';
import settings from './SettingsSampleData.js'
import DetailedSettings from "../components/DetailedSettings";
import RightColumn from "../components/RightColumn";
import App from "../App.js";


storiesOf('DialogMini', module)
    .add('default', () => <DialogMini />)
    .add('with data', () => <DialogMini author="arkady1995@mail.ru" message="Привет, как дела?"/>)
    .add('with long data', () => <DialogMini
        author="arkady1995198273981273981723981@mail.ru"
        message="Привет, как дела? Привет, как дела? Привет, как дела? Привет, как дела?"/>);

storiesOf('DialogList', module)
    .add('default', () => <DialogList />)
    .add('with data', () => <DialogList dialogs={[...dialogs, ...dialogs, ...dialogs]}/>);

storiesOf('TabBar', module)
    .add('default', () => <Router><TabBar /></Router>);

storiesOf('SettingsColumn', module)
    .add('default', () => <SettingsColumn/>)
    .add('with data', () => <SettingsColumn settingsItems={settings}/>);

storiesOf('LeftColumn', module)
    .add('default', () => <LeftColumn /> );

storiesOf('LetterReading', module)
    .add('with data', () => <LetterReading {...letter} /> );

storiesOf('DetailedSettings', module)
    .add('default', () => <DetailedSettings /> );

storiesOf('RightColumn', module)
    .add('default', () => <RightColumn /> );

storiesOf('App', module)
    .add('default', () => <App />);
