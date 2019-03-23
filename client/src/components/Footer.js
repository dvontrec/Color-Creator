import React, { Component } from 'react';
import style from '../style.css';

const Footer = () => {
  return (
    <footer className={style.footer}>
      <h6 className="text-center copyright">
        ©{' '}
        <script type="text/javascript">
          var d = new Date();document.write(d.getFullYear());
        </script>{' '}
        <span>
          <a href="https://dvontrecoleman.com/" target="blank">
            Dvontre Coleman.com
          </a>
        </span>{' '}
      </h6>
    </footer>
  );
};

export default Footer;
