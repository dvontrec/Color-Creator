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
    const res = await fetch('/api/yellow');
    const json = await res.json();
    this.setState({ color: json.color });
  }

  render() {
    return (
      <div>
        <p>Color:{this.state.color}</p>
      </div>
    );
  }
}

export default Color;
