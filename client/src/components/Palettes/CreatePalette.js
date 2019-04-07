import React, { Component } from 'react';
import { connect } from 'react-redux';
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

  async setPrimary(colorState) {
    console.log('primary: ', colorState);
    let res = await colorapi.get(`/api/colors?color=${colorState.hex}`);
    if (res.status === 200) {
      console.log('response', res);
      this.state.colors[0] = res.data;
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

export default connect(mapStateToProps)(CreatePalette);
