import React, { Component } from 'react';
import { Link } from 'react-router-dom';

import style from '../../style.css';
import api from '../../apis/colors';

class Profile extends Component {
  async componentDidMount() {
    // calls the api passing in the id parameter from the route
    const response = await api.get(
      `api/user?userId=${this.props.match.params.id}`
    );
    this.setState({ UserData: response.data });
  }
  renderColor = color => {
    if (color) {
      const colorCode = `rgba(${color.r},${color.g},${color.b},${color.a})`;
      return (
        <div key={color.hex} className="col-md-1">
          <Link to={`/color/${color.hex}`}>
            <div
              className={style.colorBlock}
              style={{ background: colorCode }}
            />
          </Link>
        </div>
      );
    }
    return <div>Loading Color</div>;
  };

  renderProfile() {
    if (!this.state) {
      return <div>Retrieving User</div>;
    }
    const userInfo = this.state.UserData.userInfo;
    const favoriteColors = this.state.UserData.favoriteColors;
    const createdColors = this.state.UserData.createdColors;
    return (
      <div>
        <h4>{userInfo.username.toUpperCase()}</h4>
        <h4>Created Colors</h4>
        <div className="row">
          {createdColors && createdColors.map(color => this.renderColor(color))}
        </div>
        <hr />
        <h4>Favorite Colors</h4>
        <div className="row">
          {favoriteColors &&
            favoriteColors.map(color => this.renderColor(color))}
        </div>
      </div>
    );
  }
  render() {
    return <div className="container">{this.renderProfile()}</div>;
  }
}
export default Profile;
