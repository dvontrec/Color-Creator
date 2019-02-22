export default (state = {}, action) => {
  switch (action.type) {
    case 'REGISTER_USER':
      return { ...state, isSignedIn: false, userId: null };
    case 'SIGN_IN':
      return { ...state, isSignedIn: true, userId: action.payload };
    default:
      return state;
  }
};
