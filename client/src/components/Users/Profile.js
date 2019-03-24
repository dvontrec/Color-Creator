import React, { Component } from 'react';

import api from '../../apis/colors';
import ColorContainer from '../Colors/ColorContainer';

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
        <h4>Created Colors</h4>
        <div className="row">
          <ColorContainer colors={this.state.UserData.createdColors} />
        </div>
        <hr />
        <h4>Favorite Colors</h4>
        <div className="row">
          <ColorContainer colors={this.state.UserData.favoriteColors} />
        </div>
      </div>
    );
  }
  render() {
    return <div className="container">{this.renderProfile()}</div>;
  }
}
export default Profile;
