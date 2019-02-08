import React, { Component } from 'react';
import { Route, HashRouter, IndexRoute } from 'react-router-dom';

import Home from './Home';
import Away from './Away';
import Nav from './Nav';

class App extends Component {
  render() {
    return (
      <HashRouter>
        <div>
          <Nav />
          <Route exact path="/" component={Home} />
          <Route path="/away" component={Away} />
        </div>
      </HashRouter>
    );
  }
}

export default App;
