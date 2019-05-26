import React from 'react';
import '../index.css';

import { storiesOf } from '@storybook/react';

import DialogMini from '../components/DialogMini';
import { BrowserRouter as Router } from 'react-router-dom';


    storiesOf('MessageWriting')
 // storiesOf('DialogMini', module)
 //    .add('default', () => <Router><DialogMini /></Router>)
 //    .add('with data', () => <Router><DialogMini author="arkady1995@mail.ru" message="Привет, как дела?"/></Router>)
 //    .add('with long data', () => <Router><DialogMini
 //        author="arkady1995198273981273981723981@mail.ru"
 //        responder="kek@mail.ru"
 //        message="Привет, как дела? Привет, как дела? Привет, как дела? Привет, как дела?"/></Router>)
 //    .add('from me', () => (
 //        <Router>
 //            <DialogMini id="2"
 //                        message="from me blabla"
 //                        author="kek@mail.ru"
 //                        responder="another@mail.ru"
 //                        checked={true}
 //            />
 //        </Router>))
 //    .add('to me', () => (
 //        <Router>
 //            <DialogMini id="3"
 //                        message="to me blabla"
 //                        responder="kek@mail.ru"
 //                        author="another@mail.ru"
 //                        checked={false}
 //            />
 //        </Router>));

// storiesOf('DialogList', module)
//     .add('default', () => <DialogList />)
//     .add('with data', () => <DialogList dialogs={[...dialogs, ...dialogs, ...dialogs]}/>);
//
// storiesOf('TabBar', module)
//     .add('default', () => <Router><TabBar /></Router>);
//
// storiesOf('SettingsColumn', module)
//     .add('default', () => <SettingsColumn/>)
//     .add('with data', () => <SettingsColumn settingsItems={settings}/>);
//
// storiesOf('LeftColumn', module)
//     .add('default', () => <LeftColumn /> );
//
// storiesOf('LetterReading', module)
//     .add('with data', () => <LetterReading {...letter} /> );
//
// storiesOf('DetailedSettings', module)
//     .add('default', () => <DetailedSettings /> );
//
// storiesOf('RightColumn', module)
//     .add('default', () => <RightColumn /> );
//
// storiesOf('LetterWriting', module)
//     .add('default', () => <LetterWriting/>);
//
// storiesOf('App', module)
//     .add('default', () => <App />);
