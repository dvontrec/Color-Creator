import React, { Component } from 'react';
import { Route, Router, IndexRoute } from 'react-router-dom';

import Nav from './Nav';
import history from '../history';
import ShowColors from './Colors/ShowColors';
import CreateColor from './Colors/CreateColor';
import RegisterForm from './Users/RegisterForm';
import LoginForm from './Users/LoginForm';
import Logout from './Users/Logout';
import ColorDisplay from './Colors/ColorDisplay';
import ColorEdit from './Colors/ColorEdit';
import Profile from './Users/Profile';

class App extends Component {
  render() {
    return (
      <Router history={history}>
        <div>
          <div>
            <Nav />
          </div>
          <Route exact path="/" component={ShowColors} />
          <Route path="/new/color/" component={CreateColor} />
          <Route path="/edit/color/:color" component={ColorEdit} />
          <Route path="/color/:color" component={ColorDisplay} />
          <Route path="/register" component={RegisterForm} />
          <Route path="/login" component={LoginForm} />
          <Route path="/logout" component={Logout} />
          <Route path="/profile/:id" component={Profile} />
        </div>
      </Router>
    );
  }
}

export default App;
