import { combineReducers } from 'redux';
import { History } from 'history';
import { RouterState, connectRouter } from 'connected-react-router';
import { IState as todoState, reducer as todoReducer } from '~/reducers/todo';

export interface IState {
    todo: todoState;
    router: RouterState;
}

export default (history: History) =>
    combineReducers<IState>({
        router: connectRouter(history),
        todo: todoReducer,
    });
