import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import colorapi from '../../apis/colors';

import style from '../../style.css';

class Color extends Component {
  constructor() {
    super();
    this.state = {};
  }
  async componentDidMount() {
    const res = await colorapi.get(`/api/colors?color=${this.props.colorName}`);
    this.setState({ color: res.data });
  }

  renderColor = () => {
    if (this.state.color) {
      const color = this.state.color;
      const colorCode = `rgba(${color.r},${color.g},${color.b},${color.a})`;
      return (
        <Link to={`/color/${this.state.color.color}`}>
          <div className={style.colorDiv} style={{ borderColor: colorCode }}>
            <div
              className={style.colorBlock}
              style={{ background: colorCode }}
            />
            <p>
              Color: <span style={{ color: colorCode }}>{color.color}</span>
            </p>
          </div>
        </Link>
      );
    }
    return <div>Loading Color</div>;
  };
  render() {
    return <div>{this.renderColor()}</div>;
  }
}

export default Color;
