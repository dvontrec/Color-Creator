import colors from '../apis/colors';

import history from '../history';

export const fetchColors = () => async dispatch => {
  const response = await colors.get('/api/colors');
  dispatch({ type: 'FETCH_COLORS', payload: response.data });
};

export const fetchColor = color => async dispatch => {
  const response = await colors.get(`/api/colors?color=${color}`);
  dispatch({ type: 'FETCH_COLOR', payload: response.data });
};

export const createColor = colorQuery => async dispatch => {
  const response = await colors.post(`api/colors?${colorQuery}`);
  //  dispatches an action with the response as the payload
  dispatch({ type: 'CREATE_COLOR', payload: response.data });
  //  Do programmatic navigation to get user back to root route
  history.push('/');
};
