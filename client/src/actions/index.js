import api from '../apis/colors';

import history from '../history';

export const fetchColors = () => async dispatch => {
  const response = await api.get('/api/colors');
  dispatch({ type: 'FETCH_COLORS', payload: response.data });
};

export const fetchColor = color => async dispatch => {
  const response = await api.get(`/api/colors?color=${color}`);
  dispatch({ type: 'FETCH_COLOR', payload: response.data });
};

export const createColor = colorQuery => async dispatch => {
  const response = await api.post(`api/colors?${colorQuery}`);
  //  dispatches an action with the response as the payload
  dispatch({ type: 'CREATE_COLOR', payload: response.data });
  //  Do programmatic navigation to get user back to root route
  history.push('/');
};

export const editColor = (color, colorQuery) => async dispatch => {
  const response = await api.patch(`api/colors`);
  dispatch({ type: 'EDIT_COLOR', payload: response.data });
};

export const registerUser = userQuery => async dispatch => {
  const response = await api.post(`api/${userQuery}`);
  dispatch({ type: 'REGISTER_USER', payload: response.data });
  history.push('/');
};

export const signIn = userQuery => async dispatch => {
  const response = await api.get(`api/${userQuery}`);
  dispatch({ type: 'SIGN_IN', payload: response.data });
  localStorage.setItem('id', response.data.id);
  localStorage.setItem('hash', response.data.hash);
  history.push('/');
};

export const signOut = () => {
  localStorage.removeItem('id');
  localStorage.removeItem('hash');
  return {
    type: 'SIGN_OUT'
  };
};
