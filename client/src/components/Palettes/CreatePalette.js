import React, { Component } from 'react';
import { connect } from 'react-redux';
import colorapi from '../../apis/colors';
import { createPalette } from '../../actions';

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

  // Sets the color to the position in state
  async setPrimary(colorState) {
    console.log('primary: ', colorState);
    // tries to find the color by hex
    let res = await colorapi.get(`/api/colors?color=${colorState.hex}`);
    // if color exists set it
    if (res.status === 200) {
      console.log('response', res);
      this.state.colors[0] = res.data;
      this.forceUpdate();
    } else {
      // if color does not exist prompt for color if nothing is input
      const colName =
        colorState.color ||
        prompt("You've Discovered a new Color!! Give ia a name");
      const queryString = `color=${colName}&r=${colorState.r}&g=${
        colorState.g
      }&b=${colorState.b}&a=${colorState.a}&hex=${colorState.hex}&creatorId=${
        this.props.auth.userId
      }&creatorHash=${this.props.auth.userHash}`;
      console.log(queryString);
      // sets the color in the palette container
      res = await colorapi.post(`/api/colors?${queryString}`);
      this.state.colors[0] = res.data;
      console.log('state', this.state.colors);
      this.forceUpdate();
    }
  }

  async setSecondary(colorState) {
    console.log('primary: ', colorState);
    let res = await colorapi.get(`/api/colors?color=${colorState.hex}`);
    if (res.status === 200) {
      console.log('response', res);
      this.state.colors[1] = res.data;
      this.forceUpdate();
    } else {
      const colName =
        colorState.color ||
        prompt("You've Discovered a new Color!! Give ia a name");
      const queryString = `color=${colName}&r=${colorState.r}&g=${
        colorState.g
      }&b=${colorState.b}&a=${colorState.a}&hex=${colorState.hex}&creatorId=${
        this.props.auth.userId
      }&creatorHash=${this.props.auth.userHash}`;
      console.log(queryString);
      res = await colorapi.post(`/api/colors?${queryString}`);
      this.state.colors[1] = res.data;
      console.log('state', this.state.colors);
      this.forceUpdate();
    }
  }

  async setTertiary(colorState) {
    console.log('primary: ', colorState);
    let res = await colorapi.get(`/api/colors?color=${colorState.hex}`);
    if (res.status === 200) {
      console.log('response', res);
      this.state.colors[2] = res.data;
      this.forceUpdate();
    } else {
      const colName =
        colorState.color ||
        prompt("You've Discovered a new Color!! Give ia a name");
      const queryString = `color=${colName}&r=${colorState.r}&g=${
        colorState.g
      }&b=${colorState.b}&a=${colorState.a}&hex=${colorState.hex}&creatorId=${
        this.props.auth.userId
      }&creatorHash=${this.props.auth.userHash}`;
      console.log(queryString);
      res = await colorapi.post(`/api/colors?${queryString}`);
      this.state.colors[2] = res.data;
      console.log('state', this.state.colors);
      this.forceUpdate();
    }
  }
  // function for saving pallet by hex to the api
  submitPalette = () => {
    const paletteName = prompt('Please enter a name for the palette: ');
    const primaryHex = this.state.colors[0].hex;
    const secondaryHex = this.state.colors[1].hex;
    const tertiaryHex = this.state.colors[2].hex;
    const paletteData = `userID=${
      this.props.auth.userId
    }&paletteName=${paletteName}&primaryHex=${primaryHex}&secondaryHex=${secondaryHex}&tertiaryHex=${tertiaryHex}`;
    this.props.createPalette(paletteData);
  };

  renderPalette = () => {
    return (
      <div className={`jumbotron ${style.jumbotron}`}>
        <div className="row" style={{ justifyContent: 'center' }}>
          <ColorContainer colors={this.state.colors} isPalette="true" />
          <button className="btn btn-primary" onClick={this.submitPalette}>
            Submit
          </button>
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
            setTertiary={this.setTertiary.bind(this)}
          />
        </div>
      );
    }
    return <div>Colors</div>;
  }
}

const mapStateToProps = ({ auth }) => {
  return { auth };
};

export default connect(
  mapStateToProps,
  { createPalette }
)(CreatePalette);
