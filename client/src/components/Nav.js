import React from 'react';
import { Link } from 'react-router-dom';

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
      </ul>
    </div>
  );
};

export default Nav;
