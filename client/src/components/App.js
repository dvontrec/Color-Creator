import React, { Component } from 'react';
import { Route, HashRouter, IndexRoute } from 'react-router-dom';

import ShowColors from './Colors/ShowColors';
import CreateColor from './Colors/CreateColor';
import Nav from './Nav';

class App extends Component {
  render() {
    return (
      <HashRouter>
        <div>
          <div>
            <Nav />
          </div>
          <Route exact path="/" component={ShowColors} />
          <Route path="/newcolor/" component={CreateColor} />
        </div>
      </HashRouter>
    );
  }
}

export default App;
