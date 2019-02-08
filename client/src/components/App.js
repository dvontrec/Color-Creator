import React, { Component } from 'react';
import { Route, HashRouter, IndexRoute } from 'react-router-dom';

import Color from './Color';

class App extends Component {
  render() {
    return (
      <HashRouter>
        <div>
          <Route exact path="/" component={Color} />
        </div>
      </HashRouter>
    );
  }
}

export default App;
