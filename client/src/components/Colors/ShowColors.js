import React, { Component } from 'react';
import { connect } from 'react-redux';

import style from '../../style.css';

import { fetchColors } from '../../actions';
import ShowColor from '../Colors/ShowColor';

class ShowColors extends Component {
  componentDidMount() {
    this.props.fetchColors();
  }

  renderColors() {
    return this.props.colors.map(color => {
      return (
        <div key={color.color} className="col-md-4">
          <ShowColor colorName={color.color} />
        </div>
      );
    });
  }

  render() {
    return (
      <div className={`container ${style.colorContainer}`}>
        {this.renderColors()}
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
