import React, { Component } from 'react';
import { Field, reduxForm } from 'redux-form';

class ColorForm extends Component {
  render() {
    return <div className="container">Color Form</div>;
  }
}

export default reduxForm({
  form: 'colorForm'
})(ColorForm);
