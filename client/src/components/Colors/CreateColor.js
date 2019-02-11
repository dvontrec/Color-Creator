import React, { Component } from 'react';

class CreateColor extends Component {
  constructor() {
    super();
    this.state = {
      r: this.getRandomNum(255),
      g: this.getRandomNum(255),
      b: this.getRandomNum(255),
      a: 100
    };
  }

  getRandomNum(max) {
    const num = Math.floor(Math.random() * max);
    return num;
  }

  updateSlider = e => {
    const Skey = String(e.target.name);
    this.setState({ [Skey]: e.target.value });
  };

  render() {
    const colorCode = `rgba(${this.state.r},${this.state.g},${
      this.state.b
    },${this.state.a / 100})`;
    // console.log(colorCode);
    return (
      <div
        className="colorCreator"
        style={{ backgroundColor: colorCode, height: '100vh' }}
      >
        <div className="color-group">
          <div className="input-group">
            <label>Red </label>
            <input
              type="number"
              onChange={this.updateSlider}
              value={this.state.r}
              name="r"
            />
            <input
              type="range"
              min="0"
              max="255"
              value={this.state.r}
              onChange={this.updateSlider}
              name="r"
            />
          </div>

          <div className="input-group">
            <label>Green </label>
            <input
              type="number"
              onChange={this.updateSlider}
              value={this.state.g}
              name="g"
            />
            <input
              type="range"
              min="0"
              max="255"
              value={this.state.g}
              onChange={this.updateSlider}
              name="g"
            />
          </div>
          <div className="input-group">
            <label>Blue </label>
            <input
              type="number"
              onChange={this.updateSlider}
              value={this.state.b}
              name="b"
            />
            <input
              type="range"
              min="0"
              max="255"
              value={this.state.b}
              onChange={this.updateSlider}
              name="b"
            />
          </div>
          <div className="input-group">
            <label>Alpha/Opacity </label>
            <input
              type="number"
              onChange={this.updateSlider}
              value={this.state.a}
              name="a"
            />
            <input
              type="range"
              min="0"
              max="100"
              value={this.state.a}
              onChange={this.updateSlider}
              name="a"
            />
          </div>
        </div>
      </div>
    );
  }
}

export default CreateColor;
