import React, { Component } from 'react';
import { Route, Router, IndexRoute } from 'react-router-dom';

import Nav from './Nav';
import history from '../history';
import ShowColors from './Colors/ShowColors';
import CreateColor from './Colors/CreateColor';

import ColorDisplay from './Colors/ColorDisplay';

class App extends Component {
  render() {
    return (
      <Router history={history}>
        <div>
          <div>
            <Nav />
          </div>
          <Route exact path="/" component={ShowColors} />
          <Route path="/newcolor/" component={CreateColor} />
          <Route path="/color/:color" component={ColorDisplay} />
        </div>
      </Router>
    );
  }
}

export default App;
