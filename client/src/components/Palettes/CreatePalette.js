import React, { Component } from 'react';
import colorapi from '../../apis/colors';

import style from '../../style.css';
import ColorContainer from '../Colors/ColorContainer';

class CreatePalette extends Component {
  state = {};
  async componentDidMount() {
    const res = await colorapi.get('/api/colors?color=ff0000');
    const colorArray = [res.data, res.data, res.data];
    console.log(colorArray);
    this.setState({ colors: colorArray });
  }
  render() {
    if (this.state.colors) {
      return (
        <div className="container">
          <div className={`jumbotron ${style.jumbotron}`}>
            <div className="row" style={{ justifyContent: 'center' }}>
              <ColorContainer colors={this.state.colors} isPalette="true" />
            </div>
          </div>
        </div>
      );
    }
    return <div>Colors</div>;
  }
}

export default CreatePalette;
