import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import {
  fetchColor,
  getFavoritesByColor,
  addFavorites,
  removeFavorites
} from '../../actions';

class ColorDisplay extends Component {
  state = {
    favorites: this.props.favorites || []
  };
  componentDidMount() {
    this.props.fetchColor(this.props.match.params.color);
    this.props.getFavoritesByColor(this.props.match.params.color);
  }

  // function called when props update
  componentDidUpdate() {
    // this.props.getFavoritesByColor(this.props.match.params.color);
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
        <Link
          className="btn btn-primary"
          to={`/profile/${this.props.color.creatorId}`}
        >
          View Creator Profile
        </Link>
        {this.checkFavorite()}
      </div>
    );
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
          <button className="btn btn-secondary" onClick={this.removeFavorite}>
            Remove From Favorites
          </button>
        );
      }
      return (
        <button className="btn btn-success" onClick={this.addFavorite}>
          Add To Favorite
        </button>
      );
    }
  };

  render() {
    return <div className="container">{this.renderColor()}</div>;
  }
}

const mapStateToProps = (state, ownProps) => {
  return {
    color: state.colors[ownProps.match.params.color],
    auth: state.auth,
    favorites: Object.values(state.favorites)[0]
  };
};

export default connect(
  mapStateToProps,
  { fetchColor, getFavoritesByColor, addFavorites, removeFavorites }
)(ColorDisplay);
