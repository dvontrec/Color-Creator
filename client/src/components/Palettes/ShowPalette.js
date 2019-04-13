import React, { Component } from 'react';
import { connect } from 'react-redux';

class ShowPalette extends Component {
  render() {
    return <div>This will show the palette</div>;
  }
}

export default connect(null)(ShowPalette);
