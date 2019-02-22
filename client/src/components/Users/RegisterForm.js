import React, { Component } from 'react';
import { connect } from 'react-redux';
import { registerUser } from '../../actions';

import UserForm from './UserForm';

class RegisterForm extends Component {
  onSubmit = formValues => {
    const username = formValues.username;
    const password = formValues.password;
    const userQuery = `auth?username=${username}&password=${password}`;
    this.props.registerUser(userQuery);
  };

  render() {
    return (
      <div className="container">
        <h3>Register User</h3>
        <UserForm submitButton="Register" onSubmit={this.onSubmit} />
      </div>
    );
  }
}

export default connect(
  null,
  { registerUser }
)(RegisterForm);
