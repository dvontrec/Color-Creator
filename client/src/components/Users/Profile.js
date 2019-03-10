import React, { Component } from 'react';
import api from '../../apis/colors';

class Profile extends Component {
  async componentDidMount() {
    // saves the response from the api call to be the response data
    const response = await api.get(`api/user?userId=2`);
    console.log(response.data);
  }
  render() {
    return <div className="container">Profile</div>;
  }
}
export default Profile;
