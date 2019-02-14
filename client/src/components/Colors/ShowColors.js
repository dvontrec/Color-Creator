import React, { Component } from 'react';
import { connect } from 'react-redux';

import style from '../../style.css';

import { fetchColors } from '../../actions';
import ShowColor from '../Colors/ShowColor';
import ColorSearch from './ColorSearch';

class ShowColors extends Component {
  state = {
    term: ''
  };
  componentDidMount() {
    this.props.fetchColors();
  }

  searchColors = () => {
    if (this.state.termterm === '') {
      return this.state.color;
    } else {
      return this.props.colors.filter(c =>
        c.color.toLowerCase().includes(this.state.term.toLowerCase())
      );
    }
  };

  inputChange = e => {
    this.setState({ term: e.target.value }, () => this.searchColors());
  };

  renderColors(colors = this.searchColors()) {
    return colors.map(color => {
      return (
        <div key={color.color} className="col-md-4">
          <ShowColor colorName={color.color} />
        </div>
      );
    });
  }

  render() {
    return (
      <div>
        <div className="container">
          <ColorSearch value={this.state.term} onChange={this.inputChange} />
          <div className={style.colorContainer}>{this.renderColors()}</div>
        </div>
      </div>
    );
  }
}

const mapStateToProps = state => {
  return { colors: Object.values(state.colors) };
};

export default connect(
  mapStateToProps,
  { fetchColors }
)(ShowColors);
