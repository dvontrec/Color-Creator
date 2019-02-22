import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';

import style from '../style.css';

class Nav extends Component {
  renderNav = () => {
    if (this.props.isSignedIn) {
      return (
        <li>
          <Link to="/logout">Log Out</Link>
        </li>
      );
    }
    return (
      <li>
        <Link to="/register">Register</Link>/<Link to="/login">Login</Link>
      </li>
    );
  };

  render() {
    this.renderNav();
    return (
      <div className={style.nav}>
        <ul>
          <li>
            <Link to="/">Colors</Link>
          </li>
          <li>
            <Link to="/newcolor">Create Color</Link>
          </li>
          {this.renderNav()}
        </ul>
      </div>
    );
  }
}

const mapStateToProps = state => {
  return {
    isSignedIn: state.auth.isSignedIn
  };
};

export default connect(mapStateToProps)(Nav);
