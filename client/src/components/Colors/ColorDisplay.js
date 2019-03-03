import React, { Component } from 'react';
import { connect } from 'react-redux';
import { fetchColor, getFavoritesByColor } from '../../actions';

class ColorDisplay extends Component {
  componentDidMount() {
    this.props.fetchColor(this.props.match.params.color);
    this.props.getFavoritesByColor(this.props.match.params.color);
  }

  renderColor() {
    if (!this.props.color) {
      return <div>Getting Color Data</div>;
    }
    const color = this.props.color;
    return (
      <div className="row">
        <div className="col-md-4">
          <div>Color Name: {color.color}</div>
          <div>R: {color.r}</div>
          <div>G: {color.g}</div>
          <div>B: {color.b}</div>
          <div>Hex: {color.hex}</div>
        </div>
        <div
          className="col-md-8"
          style={{
            backgroundColor: `#${color.hex}`,
            border: '3px solid black',
            minHeight: '100px'
          }}
        />
        {this.checkFavorite()}
      </div>
    );
  }

  checkFavorite() {
    const favArray = Array(this.props.favorites);
    console.log(favArray);
    if (favArray.includes(this.props.loggedUser)) {
      return <p>Logged In</p>;
    }
    return <p>Logged Out</p>;
  }

  render() {
    console.log(this.props);
    return <div className="container">{this.renderColor()}</div>;
  }
}

const mapStateToProps = (state, ownProps) => {
  return {
    color: state.colors[ownProps.match.params.color],
    loggedUser: state.auth.userId,
    favorites: state.favorites.favorites
  };
};

export default connect(
  mapStateToProps,
  { fetchColor, getFavoritesByColor }
)(ColorDisplay);
