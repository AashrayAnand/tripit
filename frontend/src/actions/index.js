import { ADD_TRIP } from './types';

export function addTrip(trip) {
  return {
    type: ADD_TRIP,
    payload: trip
  };
}