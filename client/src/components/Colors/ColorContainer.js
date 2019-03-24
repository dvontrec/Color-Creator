import React from 'react';
import { Link } from 'react-router-dom';
import style from '../../style.css';

const ColorContainer = props => {
  return props.color.map(color => {
    const colorCode = `rgba(${color.r},${color.g},${color.b},${color.a})`;
    return (
      <div key={color.hex} className="col-md-1">
        <Link to={`/color/${color.hex}`}>
          <div className={style.colorBlock} style={{ background: colorCode }} />
        </Link>
      </div>
    );
  });
};

export default ColorContainer;
