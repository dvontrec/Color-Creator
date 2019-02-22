import { combineReducers } from 'redux';
import { reducer as formReducer } from 'redux-form';

import colorReducer from './colorReducer';
import authReducer from './authReducer';

const rootReducer = combineReducers({
  form: formReducer,
  auth: authReducer,
  colors: colorReducer
});

export default rootReducer;
