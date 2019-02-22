import React from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';

import style from '../style.css';

const Nav = props => {
  return (
    <div className={style.nav}>
      <ul>
        <li>
          <Link to="/">Colors</Link>
        </li>
        <li>
          <Link to="/newcolor">Create Color</Link>
        </li>
        <li>
          <Link to="/register">Register</Link>/<Link to="/login">Login</Link>
        </li>
      </ul>
    </div>
  );
};

const mapStateToProps(state){

}

export default connect()(Nav);
