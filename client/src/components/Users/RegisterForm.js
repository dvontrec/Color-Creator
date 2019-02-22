import React, { Component } from 'react';
import { connect } from 'react-redux';
import { registerUser } from '../../actions';

import UserForm from './UserForm';

class RegisterForm extends Component {
  onSubmit = formValues => {
    console.log(formValues);
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
