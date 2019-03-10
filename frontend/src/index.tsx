import * as React from 'react';
import { render } from 'react-dom';
import App from '~/containers/app';
import '../assets/css/main.scss';

render(
    <App />,
    document.querySelector('.app'),
);
