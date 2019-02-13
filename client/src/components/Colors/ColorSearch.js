import React from 'react';
import style from '../../style.css';

const ColorSearch = props => {
  return (
    <div className={style.searchBar}>
      <input
        type="text"
        value={props.value}
        onChange={props.onChange}
        className="col-sm-12"
        placeholder="Search For Color"
      />
    </div>
  );
};

export default ColorSearch;
