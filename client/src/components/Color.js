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
    const res = await fetch('http://localhost:8001/blue');
    const json = await res.json();
    this.setState({ color: json.color });
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
