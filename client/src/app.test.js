import React from 'react';
import ReactDom from 'react-dom';
import App from './components/App';

// Test to see if app renders without crashing
it('renders without crashing', () => {
  const div = document.createElement('div');
  ReactDom.render(<App />, div);
  ReactDom.unmountComponentAtNode(div);
});
