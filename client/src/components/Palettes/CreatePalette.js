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

  setPrimary = hex => {
    console.log(hex);
  };
  render() {
    if (this.state.colors) {
      return (
        <div className="container">
          <div className={`jumbotron ${style.jumbotron}`}>
            <div className="row" style={{ justifyContent: 'center' }}>
              <ColorContainer colors={this.state.colors} isPalette="true" />
            </div>
          </div>
          <CreateColor isPalette="true" setPrimary={this.setPrimary} />
        </div>
      );
    }
    return <div>Colors</div>;
  }
}

export default CreatePalette;
