import React, { Component } from 'react';
import { connect } from 'react-redux';
import { registerUser } from '../../actions';

import UserForm from './UserForm';

class RegisterForm extends Component {
  onSubmit = formValues => {
    const username = formValues.username;
    const password = formValues.password;
    const password2 = formValues.password2;
    if (password === password2) {
      const userQuery = `auth?username=${username}&password=${password}`;
      this.props.registerUser(userQuery);
    } else {
      alert('passwords must match');
    }
  };

  render() {
    return (
      <div className="container">
        <h3>Register User</h3>
        <UserForm
          submitButton="Register"
          onSubmit={this.onSubmit}
          isRegister={true}
        />
      </div>
    );
  }
}

export default connect(
  null,
  { registerUser }
)(RegisterForm);
