import React, { Component } from 'react';
import axios from 'axios';

class Color extends Component {
  constructor() {
    super();
    this.state = {
      color: {
        color: 'teal'
      }
    };
  }

  async componentDidMount() {
    const res = await fetch('/api/colors?color=green');
    const json = await res.json();
    this.setState({ color: json });
  }

  render() {
    const color = this.state.color;
    return (
      <div
        style={{ color: `rgba(${color.r},${color.g},${color.b},${color.a})` }}
      >
        <p>Color:{color.color}</p>
      </div>
    );
  }
}

export default Color;
