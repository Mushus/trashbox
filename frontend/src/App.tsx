import * as React from 'react';
import { createBrowserHistory } from 'history';
import { applyMiddleware, compose, createStore } from 'redux';
import { Provider } from 'react-redux';
import { Route, Switch } from 'react-router';
import { routerMiddleware, ConnectedRouter } from 'connected-react-router';
import App from '~/containers/app';
import createRootReducer from '~/reducers';

const history = createBrowserHistory();
const store = createStore(
    createRootReducer(history),
    compose(
        applyMiddleware(
            routerMiddleware(history),
        )
    )
);

const component: React.SFC = () => (
    <Provider store={store}>
        <ConnectedRouter history={history}>
            <Switch>
                <Route exact path="/" component={App} />
            </Switch>
        </ConnectedRouter>
    </Provider>
);

export default component;
