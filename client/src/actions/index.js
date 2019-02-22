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

export const registerUser = userQuery => async dispatch => {
  const response = await api.post(`api/auth?${userQuery}`);
  dispatch({ type: 'REGISTER_USER', payload: response.data });
  history.push('/');
};

export const signIn = userQuery => async dispatch => {
  const response = await api.get(`api/auth?${userQuery}`);
  dispatch({ type: 'SIGN_IN', payload: response.data });
  history.push('/');
};
