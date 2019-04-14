import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';

import { getPalette } from '../../actions';
import api from '../../apis/colors';

class ShowPalette extends Component {
  state = {};
  async componentDidMount() {
    this.props.getPalette(this.props.match.params.id);
  }

  renderColor(color) {
    return (
      <Link to={`/color/${color.hex}`}>
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
        </div>
      </Link>
    );
  }

  renderPalette() {
    console.log(this.props);
    if (this.state.paletteData) {
      return (
        <div>
          <h1>{this.state.paletteData.PaletteName}</h1>
          <div>{this.renderColor(this.state.paletteData.Primary)}</div>
          <hr />
          <div>{this.renderColor(this.state.paletteData.Secondary)}</div>
          <hr />
          <div>{this.renderColor(this.state.paletteData.Tertiary)}</div>
          <Link className="btn btn-primary" to={`/profile/`}>
            View Creator Profile
          </Link>
          {/* {this.checkFavorite()} */}
        </div>
      );
    }
    return <div>Loading...</div>;
  }

  addFavorite = () => {
    const favs = this.props.favorites || this.state.favorites;
    favs.push(this.props.auth.userId);
    const favoritesQuery = `userId=${this.props.auth.userId}&userHash=${
      this.props.auth.userHash
    }&colorHex=${this.props.color.hex}`;
    this.props.addFavorites(favoritesQuery);
    this.setState({ favorites: favs });
  };

  removeFavorite = () => {
    const newFavs =
      this.props.favorites.filter(e => e != this.props.auth.userId) ||
      this.props.favorites ||
      this.state.favorites.filter(e => e != this.props.auth.userId);
    const favoritesQuery = `userId=${this.props.auth.userId}&userHash=${
      this.props.auth.userHash
    }&colorHex=${this.props.color.hex}`;
    this.props.removeFavorites(favoritesQuery);
    this.setState({ favorites: newFavs });
  };

  checkFavorite = () => {
    if (this.props.auth.userId) {
      const favArray = this.state.favorites || this.props.favorites;
      if (favArray && favArray.includes(this.props.auth.userId.toString())) {
        return (
          <button onClick={this.removeFavorite}>Remove From Favorites</button>
        );
      }
      return <button onClick={this.addFavorite}>Add To Favorite</button>;
    }
  };

  render() {
    return <div className="container">{this.renderPalette()}</div>;
  }
}

export default connect(
  null,
  { getPalette }
)(ShowPalette);
