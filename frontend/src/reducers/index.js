import { combineReducers } from 'redux';
import tripsReducer from './trips';

export default combineReducers({
  trips: tripsReducer
});