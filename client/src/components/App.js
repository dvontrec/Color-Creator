import React, { Component } from 'react';
import { Route, HashRouter, IndexRoute } from 'react-router-dom';

import ShowColors from './Colors/ShowColors';
import ShowColor from './Colors/ShowColor';

class App extends Component {
  render() {
    return (
      <HashRouter>
        <div>
          <Route exact path="/" component={ShowColors} />
          <Route path="/color" component={ShowColor} />
        </div>
      </HashRouter>
    );
  }
}

export default App;
