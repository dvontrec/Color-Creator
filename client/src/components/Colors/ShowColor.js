import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';

import colorapi from '../../apis/colors';

import style from '../../style.css';

class Color extends Component {
  constructor() {
    super();
    this.state = {};
  }
  async componentDidMount() {
    const res = await colorapi.get(`/api/colors?color=${this.props.colorHex}`);
    this.setState({ color: res.data });
  }

  renderColor = () => {
    if (this.state.color) {
      const color = this.state.color;
      const colorCode = `rgba(${color.r},${color.g},${color.b},${color.a})`;
      if (
        this.state.color.creatorId == this.props.auth.userId &&
        this.state.color.creatorHash == this.props.auth.userHash
      ) {
        return (
          <div className={style.colorDiv} style={{ borderColor: colorCode }}>
            <div
              className={style.colorBlock}
              style={{ background: colorCode }}
            />
            <Link to={`/edit/color/${this.state.color.hex}`}>
              <p>
                Color: <span style={{ color: colorCode }}>{color.color}</span>
              </p>
            </Link>
          </div>
        );
      }
      return (
        <div className={style.colorDiv} style={{ borderColor: colorCode }}>
          <div className={style.colorBlock} style={{ background: colorCode }} />
          <Link to={`/color/${this.state.color.hex}`}>
            <p>
              Color: <span style={{ color: colorCode }}>{color.color}</span>
            </p>
          </Link>
        </div>
      );
    }
    return <div>Loading Color</div>;
  };
  render() {
    return <div>{this.renderColor()}</div>;
  }
}
const mapStateToProps = state => {
  return { auth: state.auth };
};
export default connect(mapStateToProps)(Color);
