import React, { Component } from 'react';
import colorapi from '../../apis/colors';

import style from '../../style.css';
import ColorContainer from '../Colors/ColorContainer';
import CreateColor from '../Colors/CreateColor';

class CreatePalette extends Component {
  state = {};
  async componentDidMount() {
    const res = await colorapi.get('/api/colors?color=ffffff');
    const colorArray = [res.data, res.data, res.data];
    console.log(colorArray);
    this.setState({ colors: colorArray });
  }

  async setPrimary(hex) {
    console.log(this.state.colors);
    const res = await colorapi.get(`/api/colors?color=${hex}`);
    this.state.colors[0] = res.data;
    this.forceUpdate();
  }
  async setSecondary(hex) {
    console.log(this.state.colors);
    const res = await colorapi.get(`/api/colors?color=${hex}`);
    this.state.colors[1] = res.data;
    this.forceUpdate();
  }

  renderPalette = () => {
    return (
      <div className={`jumbotron ${style.jumbotron}`}>
        <div className="row" style={{ justifyContent: 'center' }}>
          <ColorContainer colors={this.state.colors} isPalette="true" />
        </div>
      </div>
    );
  };
  render() {
    if (this.state.colors) {
      return (
        <div className="container">
          {this.renderPalette()}
          <CreateColor
            isPalette="true"
            setPrimary={this.setPrimary.bind(this)}
            setSecondary={this.setSecondary.bind(this)}
          />
        </div>
      );
    }
    return <div>Colors</div>;
  }
}

export default CreatePalette;
