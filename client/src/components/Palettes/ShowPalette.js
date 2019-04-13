import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import api from '../../apis/colors';

class ShowPalette extends Component {
  state = {};
  async componentDidMount() {
    // calls the api passing in the id parameter from the route
    const response = await api.get(
      `/api/palettes?paletteID=${this.props.match.params.id}`
    );
    this.setState({ paletteData: response.data });
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
    if (this.state.paletteData) {
      return (
        <div>
          <h1>{this.state.paletteData.PaletteName}</h1>
          <div>{this.renderColor(this.state.paletteData.Primary)}</div>
          <hr />
          <div>{this.renderColor(this.state.paletteData.Secondary)}</div>
          <hr />
          <div>{this.renderColor(this.state.paletteData.Tertiary)}</div>
        </div>
      );
    }
    return <div>Loading...</div>;
  }

  render() {
    console.log(this.state);
    return <div className="container">{this.renderPalette()}</div>;
  }
}

export default ShowPalette;
