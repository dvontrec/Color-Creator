import React, { Component } from 'react';
import colorapi from '../../apis/colors';

import { fetchColor } from '../../actions';
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
      return (
        <div className={style.colorDiv}>
          <div
            className={style.colorBlock}
            style={{
              background: `rgba(${color.r},${color.g},${color.b},${color.a})`
            }}
          />
          <p>
            Color:{' '}
            <span
              style={{
                color: `rgba(${color.r},${color.g},${color.b},${color.a})`
              }}
            >
              {color.color}
            </span>
          </p>
        </div>
      );
    }
    return <div>nope</div>;
  };
  render() {
    return <div>{this.renderColor()}</div>;
  }
}

export default Color;
