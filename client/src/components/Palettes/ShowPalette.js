import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';

import {
  getPalette,
  getFavoritesByPalette,
  addPaletteFavorites,
  removePaletteFavorites
} from '../../actions';
import api from '../../apis/colors';

class ShowPalette extends Component {
  state = {
    favorites: this.props.favorites || []
  };
  async componentDidMount() {
    this.props.getPalette(this.props.match.params.id);
    this.props.getFavoritesByPalette(this.props.match.params.id);
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
    if (this.props.palette.palette) {
      return (
        <div>
          <h1>{this.props.palette.palette.PaletteName}</h1>
          <div>{this.renderColor(this.props.palette.palette.Primary)}</div>
          <hr />
          <div>{this.renderColor(this.props.palette.palette.Secondary)}</div>
          <hr />
          <div>{this.renderColor(this.props.palette.palette.Tertiary)}</div>
          {this.checkFavorite()}
        </div>
      );
    }
    return <div>Loading...</div>;
  }

  addFavorite = () => {
    const favs = this.props.favorites || this.state.favorites;
    favs.push(this.props.auth.userId);
    const favoritesQuery = `userId=${this.props.auth.userId}&paletteId=${
      this.props.match.params.id
    }`;
    this.props.addPaletteFavorites(favoritesQuery);
    this.setState({ favorites: favs });
  };

  removeFavorite = () => {
    const newFavs =
      this.props.favorites.filter(e => e != this.props.auth.userId) ||
      this.props.favorites ||
      this.state.favorites.filter(e => e != this.props.auth.userId);
    const favoritesQuery = `userId=${this.props.auth.userId}&paletteId=${
      this.props.match.params.id
    }`;
    this.props.removePaletteFavorites(favoritesQuery);
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

const matchStateToProps = state => {
  return {
    palette: state.palette,
    auth: state.auth,
    favorites: Object.values(state.favorites)[0]
  };
};

export default connect(
  matchStateToProps,
  {
    getPalette,
    getFavoritesByPalette,
    addPaletteFavorites,
    removePaletteFavorites
  }
)(ShowPalette);
