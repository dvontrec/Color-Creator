import React, { Component } from 'react';
import { connect } from 'react-redux';
import { signIn } from '../../actions';

import UserForm from './UserForm';

class LoginForm extends Component {
  onSubmit = formValues => {
    const username = formValues.username;
    const password = formValues.password;
    const userQuery = `auth?username=${username}&password=${password}`;
    this.props.signIn(userQuery);
  };

  render() {
    return (
      <div className="container">
        <h3>Login</h3>
        <UserForm submitButton="Login" onSubmit={this.onSubmit} />
      </div>
    );
  }
}

export default connect(
  null,
  { signIn }
)(LoginForm);
