import { combineReducers } from 'redux';
import { reducer as formReducer } from 'redux-form';

import colorReducer from './colorReducer';
import authReducer from './authReducer';
import favoritesReducer from './favoritesReducer';
import paletteReducer from './paletteReducer';

const rootReducer = combineReducers({
  form: formReducer,
  auth: authReducer,
  colors: colorReducer,
  favorites: favoritesReducer,
  palette: paletteReducer
});

export default rootReducer;
