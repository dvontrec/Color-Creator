import React from 'react';
import { Link } from 'react-router-dom';

const Nav = props => {
  return (
    <div>
      <Link to="/">Colors</Link>
      <Link to="/newcolor">Create Color</Link>
    </div>
  );
};

export default Nav;
