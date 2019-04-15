export default (state = {}, action) => {
  switch (action.type) {
    case 'GET_COLOR_FAVORITES':
      return { ...state, favorites: action.payload.favorites };
    case 'GET_PALETTE_FAVORITES':
      return { ...state, favorites: action.payload.favorites };
    default:
      return state;
  }
};
