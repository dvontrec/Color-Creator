import React, { Component } from 'react';
import { connect } from 'react-redux';
import { fetchColor, getFavoritesByColor, addFavorites } from '../../actions';

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

  addFavorite = () => {
    const favoritesQuery = `userId=${this.props.auth.userId}&userHash=${
      this.props.auth.userHash
    }&colorHex=${this.props.color.hex}`;
    this.props.addFavorites(favoritesQuery);
  };

  removeFavorite = () => {
    console.log('here');
  };

  checkFavorite = () => {
    const favArray = Array(Object.values(this.props.favorites)[0]);
    console.log(favArray[0]);
    console.log(this.props.auth.userId.toString());
    if (
      favArray[0] &&
      favArray[0].includes(this.props.auth.userId.toString())
    ) {
      return (
        <button onClick={this.removeFavorite}>Remove From Favorites</button>
      );
    }
    return <button onClick={this.addFavorite}>Favorite</button>;
  };

  render() {
    return <div className="container">{this.renderColor()}</div>;
  }
}

const mapStateToProps = (state, ownProps) => {
  return {
    color: state.colors[ownProps.match.params.color],
    auth: state.auth,
    favorites: state.favorites
  };
};

export default connect(
  mapStateToProps,
  { fetchColor, getFavoritesByColor, addFavorites }
)(ColorDisplay);
