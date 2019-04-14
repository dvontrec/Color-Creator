import _ from 'lodash';

export default (state = {}, action) => {
  switch (action.type) {
    case 'FETCH_COLORS':
      return { ...state, ..._.mapKeys(action.payload, 'hex') };
    case 'FETCH_COLOR':
      // Returns the state and the color with a key of the id
      return { ...state, [action.payload.hex]: action.payload };
    case 'EDIT_COLOR':
      return { ...state, [action.payload.hex]: action.payload };
    default:
      return state;
  }
};
