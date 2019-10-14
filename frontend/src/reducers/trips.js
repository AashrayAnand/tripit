import { ADD_TRIP } from '../actions/types';

export default function(state = [], action) {
  switch (action.type) {
    case ADD_TRIP:
      return [...state, action.payload];
    default:
      return state;
  }
}