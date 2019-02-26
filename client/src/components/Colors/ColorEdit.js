import React, { Component } from 'react';
import { connect } from 'react-redux';
import _ from 'lodash';
import { fetchColor, editColor } from '../../actions';

import ColorForm from './ColorForm';

class ColorEdit extends Component {
  componentDidMount() {
    this.props.fetchColor(this.props.match.params.color);
  }

  onSubmit = formValues => {
    if (this.props.color.creatorHash != this.props.auth.userHash) {
      return;
    }
    console.log(formValues);
    const colorQuery = `color=${this.props.match.params.color}&name=${
      formValues.color
    }`;
    this.props.editColor(colorQuery);
  };

  renderColor() {
    if (!this.props.color) {
      return <div>Getting Color Data</div>;
    }
    const color = this.props.color;
    return (
      <div className="row">
        <div className="col-md-4">
          <ColorForm
            initialValues={_.pick(this.props.color, 'color')}
            onSubmit={this.onSubmit}
          />
          <div>R: {color.r}</div>
          <div>G: {color.g}</div>
          <div>B: {color.b}</div>
          <div>Hex: {color.hex}</div>
        </div>
        <div
          className="col-md-8"
          style={{
            backgroundColor: color.hex,
            border: '3px solid black',
            minHeight: '100px'
          }}
        />
      </div>
    );
  }

  render() {
    const color = this.props.color;
    return <div className="container">{this.renderColor()}</div>;
  }
}

const mapStateToProps = (state, ownProps) => {
  return { color: state.colors[ownProps.match.params.color], auth: state.auth };
};

export default connect(
  mapStateToProps,
  { fetchColor, editColor }
)(ColorEdit);
