import { combineReducers } from 'redux';

const rootReducer = combineReducers({
  // remove this when there are reducers to alter state
  state: (state = {}) => state
});

export default rootReducer;
