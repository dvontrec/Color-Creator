import React, { Component } from 'react';
import { Field, reduxForm } from 'redux-form';

class ColorForm extends Component {
  renderInput = ({ input, label }) => {
    return (
      <div>
        <label>{label}</label>
        <input {...input} />
      </div>
    );
  };

  onSubmit = formValues => {
    this.props.onSubmit(formValues);
  };

  render() {
    return (
      <div>
        <form className="row" onSubmit={this.props.handleSubmit(this.onSubmit)}>
          <Field name="color" component={this.renderInput} label="Color Name" />
          <button className="btn-primary btn-sm">Update</button>
        </form>
      </div>
    );
  }
}

export default reduxForm({
  form: 'colorForm'
})(ColorForm);
