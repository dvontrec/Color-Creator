import React, { Component } from 'react';
import axios from 'axios';

class Color extends Component {
  constructor() {
    super();
    this.state = {
      color: 'green'
    };
  }

  async componentDidMount() {
    res = await fetch('http://localhost:8083/blue');
    json = await res.json;
    this.setState({ color: res.color });
  }

  render() {
    return (
      <div>
        <p>{this.state.color}</p>
      </div>
    );
  }
}

export default Color;
