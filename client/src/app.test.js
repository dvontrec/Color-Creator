import React from 'react';
import ReactDom from 'react-dom';
import { Provider } from 'react-redux';

import App from './components/App';
import { createStore, applyMiddleware, compose } from 'redux';
import reduxThunk from 'redux-thunk';
import reducers from './reducers';

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
const createStoreWithMiddleware = composeEnhancers(applyMiddleware(reduxThunk))(
  createStore
);

// Test to see if app renders without crashing
it('renders without crashing', () => {
  const div = document.createElement('div');
  ReactDom.render(
    <Provider store={createStoreWithMiddleware(reducers)}>
      <App />
    </Provider>,
    div
  );
  ReactDom.unmountComponentAtNode(div);
});
