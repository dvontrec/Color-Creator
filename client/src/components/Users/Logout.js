import React, { Component } from 'react';
import { connect } from 'react-redux';
import { signOut } from '../../actions';

class Logout extends Component {
  componentDidMount() {
    this.props.signOut();
  }

  render() {
    return (
      <div className="container">
        <h2>User Logged Out</h2>
      </div>
    );
  }
}

export default connect(
  null,
  { signOut }
)(Logout);
