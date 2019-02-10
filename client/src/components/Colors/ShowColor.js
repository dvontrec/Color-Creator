import React, { Component } from 'react';
import colorapi from '../../apis/colors';

import { fetchColor } from '../../actions';

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
        <div
          style={{ color: `rgba(${color.r},${color.g},${color.b},${color.a})` }}
        >
          <p>Color:{color.color}</p>
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
