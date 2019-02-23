import React, { Component } from 'react';
import { Field, reduxForm } from 'redux-form';

class UserForm extends Component {
  renderInput = ({ input, label }) => {
    return (
      <div>
        <label>{label}: </label>
        <input {...input} autoComplete="off" />
      </div>
    );
  };

  renderPasswordInput = ({ input, label }) => {
    return (
      <div>
        <label>{label}: </label>
        <input {...input} type="password" autoComplete="off" />
      </div>
    );
  };

  onSubmit = formValues => {
    this.props.onSubmit(formValues);
  };

  render() {
    return (
      <form onSubmit={this.props.handleSubmit(this.onSubmit)}>
        <Field name="username" component={this.renderInput} label="Username" />
        <Field
          name="password"
          component={this.renderPasswordInput}
          label="Password"
        />
        {this.props.isRegister && (
          <Field
            name="password2"
            component={this.renderPasswordInput}
            label="Repeat Password"
            type="password"
          />
        )}
        <button>{this.props.submitButton}</button>
      </form>
    );
  }
}

const validate = formValues => {
  const errors = {};

  if (!formValues.username) {
    errors.username = 'You must enter a valid username';
  }
  if (!formValues.password) {
    errors.password = 'You must enter a valid password';
  }
  return errors;
};
export default reduxForm({
  form: 'userForm',
  validate
})(UserForm);
