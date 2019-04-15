import React, { Component } from 'react';

import api from '../../apis/colors';
import ColorContainer from '../Colors/ColorContainer';
import PaletteContainer from '../Palettes/PaletteContainer';

class Profile extends Component {
  async componentDidMount() {
    // calls the api passing in the id parameter from the route
    const response = await api.get(
      `api/user?userId=${this.props.match.params.id}`
    );
    this.setState({ UserData: response.data });
  }

  renderProfile() {
    if (!this.state) {
      return <div>Retrieving User</div>;
    }
    return (
      <div>
        <h4>{this.state.UserData.userInfo.username.toUpperCase()}</h4>
        <h3>Created Colors</h3>
        <div className="row">
          {this.state.UserData.createdColors && (
            <ColorContainer colors={this.state.UserData.createdColors} />
          )}
        </div>
        <hr />
        <h3>Favorite Colors</h3>
        <div className="row">
          {this.state.UserData.favoriteColors && (
            <ColorContainer colors={this.state.UserData.favoriteColors} />
          )}
        </div>
        <hr />
        <h3>Palettes</h3>
        <div>
          {this.state.UserData.palettes && (
            <PaletteContainer palettes={this.state.UserData.palettes} />
          )}
        </div>
        <hr />
        <h3>Favorite Palettes</h3>
        <div>
          {this.state.UserData.favoritePalettes && (
            <PaletteContainer palettes={this.state.UserData.favoritePalettes} />
          )}
        </div>
      </div>
    );
  }
  render() {
    return <div className="container">{this.renderProfile()}</div>;
  }
}
export default Profile;
