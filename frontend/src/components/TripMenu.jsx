import React from 'react';

import TripList from './TripList';
import TripForm from '../containers/TripForm';

const TripMenu = () => {
  return (
    <div id='trip-menu' className='container-fluid m-0 d-flex row justify-content-center align-items-center'>
      <TripForm />
      <TripList />
    </div>
  );
}

export default TripMenu;